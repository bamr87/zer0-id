package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

// Build metadata, set at build time with
//
//	go build -ldflags "-X main.version=v0.1.0 -X main.commit=$(git rev-parse HEAD) -X main.builtAt=$(date -u +%FT%TZ)"
//
// When unset, buildInfo falls back to the VCS stamp Go embeds in the binary.
var (
	version = ""
	commit  = ""
	builtAt = ""
)

const serviceName = "zer0-id"

// Server timeouts (UPS-BE hardening; a zero-value http.Server never times out — gosec G114).
const (
	readHeaderTimeout = 5 * time.Second
	readTimeout       = 10 * time.Second
	writeTimeout      = 15 * time.Second
	idleTimeout       = 60 * time.Second
	shutdownTimeout   = 10 * time.Second
)

func issuer() string {
	if v := os.Getenv("ZID_ISSUER"); v != "" {
		return v
	}
	return "http://localhost:8080"
}

// addr is the listen address: ZID_ADDR, or :8080 when unset.
func addr() string {
	if v := os.Getenv("ZID_ADDR"); v != "" {
		return v
	}
	return ":8080"
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func discovery(w http.ResponseWriter, _ *http.Request) {
	iss := issuer()
	writeJSON(w, http.StatusOK, map[string]any{
		"issuer":                   iss,
		"authorization_endpoint":   iss + "/authorize",
		"token_endpoint":           iss + "/token",
		"jwks_uri":                 iss + "/jwks.json",
		"response_types_supported": []string{"code"},
	})
}

// healthz is liveness (UPS-BE-10): the process is up and serving.
func healthz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// readyz is readiness (UPS-BE-10). The service has no dependencies yet, so
// checks is empty; a dependency added later reports here and turns the
// status to "degraded" (503) when it fails.
func readyz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"status": "ok", "checks": map[string]string{}})
}

// buildInfo returns version, commit and build time: ldflags first, then the
// module version and VCS stamp from debug.ReadBuildInfo.
func buildInfo() (v, c, at string) {
	v, c, at = version, commit, builtAt
	if bi, ok := debug.ReadBuildInfo(); ok {
		if v == "" && bi.Main.Version != "" {
			v = bi.Main.Version
		}
		for _, s := range bi.Settings {
			switch {
			case s.Key == "vcs.revision" && c == "":
				c = s.Value
			case s.Key == "vcs.time" && at == "":
				at = s.Value
			}
		}
	}
	if v == "" {
		v = "dev"
	}
	return v, c, at
}

// versionHandler is build metadata (UPS-BE-11).
func versionHandler(w http.ResponseWriter, _ *http.Request) {
	v, c, at := buildInfo()
	writeJSON(w, http.StatusOK, map[string]string{
		"name":     serviceName,
		"version":  v,
		"commit":   c,
		"built_at": at,
	})
}

func mux() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/healthz", healthz)
	m.HandleFunc("/readyz", readyz)
	m.HandleFunc("/version", versionHandler)
	m.HandleFunc("/.well-known/openid-configuration", discovery)
	return m
}

func newServer(addr string, h http.Handler) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           h,
		ReadHeaderTimeout: readHeaderTimeout,
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
	}
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := newServer(addr(), mux())
	errc := make(chan error, 1)
	go func() {
		log.Printf("%s listening on %s", serviceName, srv.Addr)
		errc <- srv.ListenAndServe()
	}()

	select {
	case err := <-errc:
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	case <-ctx.Done():
		log.Println("shutting down")
		sctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
		defer cancel()
		if err := srv.Shutdown(sctx); err != nil {
			log.Fatalf("shutdown: %v", err)
		}
	}
}
