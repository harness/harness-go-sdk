# Mesh Identity (SPIRE)

Go SDK for Harness SPIRE-based service-to-service identity. Parity with Java
`io.harness.security.mesh`: same env vars, JWT shape (`ES256` + `x5c`), Prometheus
metric names/labels, inbound dispatch matrix, and outbound `X-Harness-Identity` header.

**HMAC is not included.** Plug your existing HMAC (or other) auth through `FallbackAuth`.
When inbound is enabled and `FallbackEnabled` is true (default), always supply a
`FallbackAuth` — a nil fallback rejects invalid mesh tokens with 401 rather than
failing open (Java always has an HMAC filter behind mesh).

Epic: [PL-72879](https://harness.atlassian.net/browse/PL-72879) · Story: [PL-73568](https://harness.atlassian.net/browse/PL-73568)

## Install

Requires **Go 1.25+**. The module `go` directive was raised so the transitive
`google.golang.org/grpc` dependency can stay on a patched release
(`>= v1.82.1`, [GHSA-hrxh-6v49-42gf](https://github.com/advisories/GHSA-hrxh-6v49-42gf)).
`go-spiffe/v2` itself only needs Go 1.24; the higher floor comes from the gRPC security fix.

```bash
go get github.com/harness/harness-go-sdk@latest
```

```go
import "github.com/harness/harness-go-sdk/harness/mesh"
```

## Configuration (env)

| Env | Default | Meaning |
|-----|---------|---------|
| `MESH_IDENTITY_INBOUND_ENABLED` | `false` | Validate inbound mesh JWT |
| `MESH_IDENTITY_OUTBOUND_ENABLED` | `false` | Stamp outbound `X-Harness-Identity` |
| `MESH_IDENTITY_OUTBOUND_ONLY` | `false` | Mesh-only stage (omit HMAC in *your* transport) |
| `MESH_IDENTITY_FALLBACK_ENABLED` | `true` | On mesh fail → call `FallbackAuth` |
| `MESH_IDENTITY_REJECT_AUTH_WITHOUT_MESH_HEADER` | `false` | Stage 5: require mesh header |
| `MESH_IDENTITY_AUDIENCE` | _(required if inbound)_ | Primary inbound `aud` (`AuthorizationServiceHeader` value) |
| `MESH_IDENTITY_ALLOWED_AUDIENCES` | `gateway-internal-service` | Comma-separated extra audiences |
| `SPIFFE_ENDPOINT_SOCKET` | `unix:///run/spire/agent-sockets/spire-agent.sock` | Workload API UDS |

## Rollout stages

| Stage | Inbound | Outbound | OutboundOnly | Fallback | RejectWithoutHeader |
|-------|---------|----------|--------------|----------|---------------------|
| 0 | false | false | — | — | — |
| 1 | true | false | — | true | false |
| 2 | true | true | false | true | false |
| 3 | true | true | true | true | false |
| 4 | true | true | true | false | false |
| 5 | true | true | true | false | true |

## Inbound (stdlib middleware)

```go
cfg := mesh.ConfigFromEnv()
holder, err := mesh.Bootstrap(ctx, cfg, nil)
if err != nil { log.Fatal(err) }
defer holder.Close()

hmacFallback := mesh.FallbackAuthFunc(func(r *http.Request) (mesh.Principal, error) {
    // existing HMAC / service auth
    return myHMACAuthenticate(r)
})

mux := http.NewServeMux()
mux.HandleFunc("/api/health", health)
handler := mesh.Middleware(holder, hmacFallback)(mux)
http.ListenAndServe(":8080", handler)
```

Principal is available via `mesh.PrincipalFromContext(r.Context())`.

### Gin

```go
r.Use(func(c *gin.Context) {
    // Run mesh middleware against the underlying handler, or:
    // wrap once at the http.Server level with mesh.Middleware.
    c.Next()
})
// Prefer wrapping the gin engine as http.Handler:
http.ListenAndServe(":8080", mesh.Middleware(holder, hmacFallback)(r))
```

## Outbound (`http.RoundTripper`)

Wrap **outside** your HMAC transport so dual-header mode keeps HMAC underneath:

```go
hmacTransport := myHMACRoundTripper(http.DefaultTransport)
meshTransport := mesh.NewRoundTripper(holder, mesh.OutboundConfig{
    TargetServiceID: mesh.ServiceAccessControlService, // callee aud
}, hmacTransport)

client := &http.Client{Transport: meshTransport}
```

When `MESH_IDENTITY_OUTBOUND_ONLY=true`, omit the HMAC transport yourself and pass
`http.DefaultTransport` (or your TLS transport) as `next`.

## Metrics

Registered on the default Prometheus registry (same names as Java):

- `harness_mesh_jwt_mint_total` / `harness_mesh_jwt_mint_duration_seconds`
- `harness_mesh_jwt_validate_total` / `harness_mesh_jwt_validate_duration_seconds`
- `harness_mesh_hmac_fallback_total`
- `harness_mesh_bundle_refresh_total`
- `harness_mesh_auth_transport_total`
- `harness_mesh_outbound_config_error_total`

## Service IDs

Outbound `TargetServiceID` and SPIFFE path last segment must match Java
`AuthorizationServiceHeader` values (e.g. `NextGenManager`, `accessControlService`,
`PipelineService`). See `serviceid.go`.

## Testing without SPIRE

Inject a `StaticSource` via `BootstrapOptions.Source` (see unit tests).
