package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func get(t *testing.T, path string) *httptest.ResponseRecorder {
	t.Helper()
	rr := httptest.NewRecorder()
	mux().ServeHTTP(rr, httptest.NewRequest(http.MethodGet, path, nil))
	return rr
}

func decode(t *testing.T, rr *httptest.ResponseRecorder) map[string]any {
	t.Helper()
	if ct := rr.Header().Get("Content-Type"); !strings.HasPrefix(ct, "application/json") {
		t.Fatalf("Content-Type = %q, want application/json", ct)
	}
	var body map[string]any
	if err := json.Unmarshal(rr.Body.Bytes(), &body); err != nil {
		t.Fatalf("body is not a JSON object: %v (%q)", err, rr.Body.String())
	}
	return body
}

func TestHealthz(t *testing.T) {
	rr := get(t, "/healthz")
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	if body := decode(t, rr); body["status"] != "ok" {
		t.Fatalf(`status field = %v, want "ok"`, body["status"])
	}
}

func TestReadyz(t *testing.T) {
	rr := get(t, "/readyz")
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	body := decode(t, rr)
	if body["status"] != "ok" {
		t.Fatalf(`status field = %v, want "ok"`, body["status"])
	}
	if _, ok := body["checks"].(map[string]any); !ok {
		t.Fatalf("checks = %#v, want a JSON object", body["checks"])
	}
}

func TestVersion(t *testing.T) {
	rr := get(t, "/version")
	if rr.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rr.Code)
	}
	body := decode(t, rr)
	for _, k := range []string{"name", "version"} {
		if s, _ := body[k].(string); s == "" {
			t.Errorf("%s = %#v, want a non-empty string", k, body[k])
		}
	}
	for _, k := range []string{"commit", "built_at"} {
		if _, ok := body[k]; !ok {
			t.Errorf("key %q missing", k)
		}
	}
}

func TestAddr(t *testing.T) {
	t.Setenv("ZID_ADDR", "")
	if got := addr(); got != ":8080" {
		t.Fatalf("addr() with ZID_ADDR unset = %q, want :8080", got)
	}
	t.Setenv("ZID_ADDR", ":9090")
	if got := addr(); got != ":9090" {
		t.Fatalf("addr() = %q, want :9090", got)
	}
}

func TestNewServerTimeouts(t *testing.T) {
	srv := newServer(":0", mux())
	if srv.Addr != ":0" {
		t.Fatalf("Addr = %q, want :0", srv.Addr)
	}
	for name, c := range map[string]struct{ got, want time.Duration }{
		"ReadHeaderTimeout": {srv.ReadHeaderTimeout, 5 * time.Second},
		"ReadTimeout":       {srv.ReadTimeout, 10 * time.Second},
		"WriteTimeout":      {srv.WriteTimeout, 15 * time.Second},
		"IdleTimeout":       {srv.IdleTimeout, 60 * time.Second},
	} {
		if c.got != c.want {
			t.Errorf("%s = %v, want %v", name, c.got, c.want)
		}
	}
}
