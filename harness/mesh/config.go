package mesh

import (
	"os"
	"strconv"
	"strings"
)

// Config mirrors Java MeshIdentityConfig / meshIdentity YAML block.
type Config struct {
	InboundEnabled               bool
	OutboundEnabled              bool
	OutboundOnly                 bool
	FallbackEnabled              bool
	RejectWithoutMeshHeader      bool
	Audience                     string
	AllowedAudiences             []string
	SPIFFEEndpointSocket         string
}

// ConfigFromEnv loads config from MESH_IDENTITY_* and SPIFFE_ENDPOINT_SOCKET env vars.
func ConfigFromEnv() Config {
	cfg := Config{
		InboundEnabled:          envBool("MESH_IDENTITY_INBOUND_ENABLED", false),
		OutboundEnabled:         envBool("MESH_IDENTITY_OUTBOUND_ENABLED", false),
		OutboundOnly:            envBool("MESH_IDENTITY_OUTBOUND_ONLY", false),
		FallbackEnabled:         envBool("MESH_IDENTITY_FALLBACK_ENABLED", true),
		RejectWithoutMeshHeader: envBool("MESH_IDENTITY_REJECT_AUTH_WITHOUT_MESH_HEADER", false),
		Audience:                os.Getenv("MESH_IDENTITY_AUDIENCE"),
		SPIFFEEndpointSocket:    envOr("SPIFFE_ENDPOINT_SOCKET", DefaultSPIFFESocket),
	}
	if raw, ok := os.LookupEnv("MESH_IDENTITY_ALLOWED_AUDIENCES"); ok {
		cfg.AllowedAudiences = splitCSV(raw)
		if len(cfg.AllowedAudiences) == 0 {
			cfg.AllowedAudiences = []string{DefaultAllowedAudience}
		}
	} else {
		cfg.AllowedAudiences = []string{DefaultAllowedAudience}
	}
	return cfg
}

// Validate checks bootstrap-time rules (match Java MeshIdentityBootstrap).
func (c Config) Validate() error {
	if c.InboundEnabled && strings.TrimSpace(c.Audience) == "" {
		return ErrAudienceRequired
	}
	return nil
}

// MeshActive reports whether either inbound or outbound is enabled.
func (c Config) MeshActive() bool {
	return c.InboundEnabled || c.OutboundEnabled
}

func envBool(key string, def bool) bool {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return def
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return def
	}
	return b
}

func envOr(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

func splitCSV(raw string) []string {
	parts := strings.Split(raw, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}
