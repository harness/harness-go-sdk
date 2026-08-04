package mesh

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// FallbackAuth is the service's existing HMAC (or other) auth path.
// Called when mesh header is absent (reject off) or mesh validation fails with fallback on.
type FallbackAuth interface {
	Authenticate(r *http.Request) (Principal, error)
}

// FallbackAuthFunc adapts a function to FallbackAuth.
type FallbackAuthFunc func(r *http.Request) (Principal, error)

// Authenticate implements FallbackAuth.
func (f FallbackAuthFunc) Authenticate(r *http.Request) (Principal, error) {
	return f(r)
}

// Middleware returns stdlib middleware implementing the Java MeshJWTAuthenticationFilter dispatch matrix.
// When holder is noop / inbound disabled, it only invokes fallback (or passes through if fallback is nil).
func Middleware(h *Holder, fallback FallbackAuth) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if h == nil || h.IsNoop() || !h.InboundEnabled() || h.Validator() == nil {
				if err := runFallback(w, r, fallback, next); err != nil {
					writeUnauthorized(w, err.Error())
				}
				return
			}

			cfg := h.Config()
			metrics := h.Metrics()
			meshToken := strings.TrimSpace(r.Header.Get(IdentityHeader))

			if meshToken != "" {
				p, err := validateInbound(h, meshToken)
				if err == nil {
					metrics.RecordAuthTransport(TransportSPIRE)
					ctx := ContextWithPrincipal(r.Context(), p)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
				if cfg.FallbackEnabled {
					log.Printf("mesh: JWT validation failed; falling back to HMAC: %v", err)
					metrics.RecordHMACFallback(ReasonFrom(err))
					// fall through to HMAC
				} else {
					log.Printf("mesh: JWT rejected (fallback disabled): %v", err)
					writeUnauthorized(w, ErrInvalidMeshToken.Error())
					return
				}
			} else if cfg.RejectWithoutMeshHeader {
				log.Printf("mesh: rejecting request without %s header", IdentityHeader)
				writeUnauthorized(w, ErrMeshHeaderRequired.Error())
				return
			}

			metrics.RecordAuthTransport(TransportHMAC)
			if err := runFallback(w, r, fallback, next); err != nil {
				writeUnauthorized(w, err.Error())
			}
		})
	}
}

func runFallback(w http.ResponseWriter, r *http.Request, fallback FallbackAuth, next http.Handler) error {
	if fallback == nil {
		next.ServeHTTP(w, r)
		return nil
	}
	p, err := fallback.Authenticate(r)
	if err != nil {
		return err
	}
	ctx := ContextWithPrincipal(r.Context(), p)
	next.ServeHTTP(w, r.WithContext(ctx))
	return nil
}

func validateInbound(h *Holder, token string) (Principal, error) {
	metrics := h.Metrics()
	cfg := h.Config()
	done := metrics.StartValidate()
	defer done()

	caller := LabelUnknown
	trustDomain := LabelUnknown
	v, err := h.Validator().Validate(token, cfg.Audience, cfg.AllowedAudiences)
	if err != nil {
		metrics.RecordValidate(OutcomeFailure, caller, cfg.Audience, trustDomain)
		return Principal{}, err
	}
	trustDomain = v.TrustDomain

	serviceID := ResolveServiceIDFromSPIFFE(v.SpiffeID)
	if serviceID == "" {
		err := &ValidationError{
			Reason: ReasonMalformedToken,
			Err:    fmt.Errorf("SPIFFE ID %s has no configured principalMapping entry", v.SpiffeID),
		}
		metrics.RecordValidate(OutcomeFailure, caller, cfg.Audience, trustDomain)
		return Principal{}, err
	}
	caller = serviceID

	typeClaim, _ := v.Claims["type"].(string)
	isDelegated := typeClaim != "" && typeClaim != string(PrincipalTypeService)

	var primary Principal
	if isDelegated {
		if err := verifyActMatchesCaller(v.Claims, v.SpiffeID); err != nil {
			metrics.RecordValidate(OutcomeFailure, caller, cfg.Audience, trustDomain)
			return Principal{}, err
		}
		primary = principalFromClaims(v.Claims)
		if primary.Name == "" {
			err := &ValidationError{Reason: ReasonMalformedToken, Err: fmt.Errorf("delegated token missing principal claims")}
			metrics.RecordValidate(OutcomeFailure, caller, cfg.Audience, trustDomain)
			return Principal{}, err
		}
	} else {
		primary = Principal{Type: PrincipalTypeService, Name: serviceID, SpiffeID: v.SpiffeID}
	}

	metrics.RecordValidate(OutcomeSuccess, caller, cfg.Audience, trustDomain)
	return primary, nil
}

func verifyActMatchesCaller(claims map[string]any, expectedSpiffeID string) error {
	actRaw, ok := claims["act"]
	if !ok {
		return &ValidationError{Reason: ReasonMalformedToken, Err: fmt.Errorf("delegated token missing act claim")}
	}
	act, ok := actRaw.(map[string]any)
	if !ok {
		return &ValidationError{Reason: ReasonMalformedToken, Err: fmt.Errorf("delegated token missing act claim")}
	}
	actSub, _ := act["sub"].(string)
	if actSub != expectedSpiffeID {
		return &ValidationError{
			Reason: ReasonMalformedToken,
			Err:    fmt.Errorf("act.sub mismatch: expected %s got %v", expectedSpiffeID, act["sub"]),
		}
	}
	return nil
}

func principalFromClaims(claims map[string]any) Principal {
	p := Principal{
		Type:      PrincipalType(claimString(claims, "type")),
		Name:      claimString(claims, "name"),
		Email:     claimString(claims, "email"),
		Username:  claimString(claims, "username"),
		AccountID: claimString(claims, "accountId"),
		Role:      claimString(claims, "role"),
	}
	if st := claimString(claims, "scopedTokenId"); st != "" {
		p.ScopedTokenID = st
	}
	return p
}

func claimString(claims map[string]any, key string) string {
	v, ok := claims[key]
	if !ok || v == nil {
		return ""
	}
	s, _ := v.(string)
	return s
}

func writeUnauthorized(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(map[string]string{"message": msg})
}
