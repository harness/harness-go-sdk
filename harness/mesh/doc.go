// Package mesh provides SPIRE-based service-to-service mesh identity for Harness Go services.
//
// It mirrors the Java package io.harness.security.mesh: same config env vars, JWT shape
// (ES256 + x5c), Prometheus metric names/labels, inbound dispatch matrix, and outbound
// X-Harness-Identity stamping.
//
// HMAC is not reimplemented. Services plug existing HMAC (or other) auth through
// FallbackAuth. The mesh middleware calls it when the mesh header is absent (and
// reject-without-header is off) or when mesh validation fails and FallbackEnabled is true.
//
// Getting started:
//
//	cfg, err := mesh.ConfigFromEnv()
//	holder, err := mesh.Bootstrap(ctx, cfg)
//	defer holder.Close()
//
//	mux := http.NewServeMux()
//	handler := mesh.Middleware(holder, myHMACFallback)(mux)
//
//	client := &http.Client{
//	  Transport: mesh.NewRoundTripper(holder, mesh.OutboundConfig{
//	    TargetServiceID: mesh.ServiceNextGenManager,
//	  }, existingHMACTransport),
//	}
//
// See README.md for rollout stages and env var details.
package mesh
