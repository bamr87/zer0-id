package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func issuer() string {
	if v := os.Getenv("ZID_ISSUER"); v != "" {
		return v
	}
	return "http://localhost:8080"
}

func discovery(w http.ResponseWriter, _ *http.Request) {
	iss := issuer()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"issuer":                 iss,
		"authorization_endpoint": iss + "/authorize",
		"token_endpoint":         iss + "/token",
		"jwks_uri":               iss + "/jwks.json",
		"response_types_supported": []string{"code"},
	})
}

func mux() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte("ok")) })
	m.HandleFunc("/.well-known/openid-configuration", discovery)
	return m
}

func main() {
	log.Println("zer0-id listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux()))
}
