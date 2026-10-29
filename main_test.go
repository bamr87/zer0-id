package main

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDiscovery(t *testing.T) {
	rr := httptest.NewRecorder()
	mux().ServeHTTP(rr, httptest.NewRequest("GET", "/.well-known/openid-configuration", nil))
	if rr.Code != 200 || !strings.Contains(rr.Body.String(), "issuer") {
		t.Fatalf("bad discovery: %d %s", rr.Code, rr.Body.String())
	}
}
