# zer0-id

> zer0-id's direction is to become the ultimate open home for all things open identity and single sign-on: data, tools, resources and transparency, with standards-based OIDC, PKCE, signed JWTs and passwordless passkeys for everyone ([constitution](docs/constitution.md) §1).

Single sign-on for the fleet's apps, built from the specs up: OIDC discovery, authorization code + PKCE, signed JWTs and passwordless passkeys.

**Status:** foundation (phase 1). Part of the [bamr87 fleet](https://github.com/bamr87/bamr87).

## Goals

- Spec-compliant OIDC discovery, JWKS and token endpoints
- Authorization code flow with PKCE
- WebAuthn passkey registration and login

## Skills learned

- OAuth 2.1 / OIDC protocol details
- JWT signing and key rotation
- WebAuthn / passkeys
- Secure Go HTTP services

## Fleet tie-in

SSO for the fleet console, ragio-ai, fredgar-ai and zer0-CMS; identities for fleet-mobile.

## Configuration

Configuration comes from environment variables; nothing is read from files.

| Variable | Default | Purpose |
| --- | --- | --- |
| `ZID_ISSUER` | `http://localhost:8080` | The issuer URL published in the OIDC discovery document. |
| `ZID_ADDR` | `:8080` | The address the HTTP server listens on (`host:port` or `:port`). |

## Run

```bash
ZID_ADDR=:9090 go run .              # logs "zer0-id listening on :9090"
curl -s localhost:9090/healthz       # {"status":"ok"}
curl -s localhost:9090/readyz        # {"checks":{},"status":"ok"}
curl -s localhost:9090/version       # {"built_at":…,"commit":…,"name":"zer0-id","version":…}
curl -s localhost:9090/.well-known/openid-configuration
```

The server sets read-header, read, write and idle timeouts (5s, 10s, 15s, 60s) and shuts down gracefully on SIGINT or SIGTERM, waiting up to 10s for in-flight requests. `/version` reads the build stamp Go embeds; a release build can set it explicitly with `go build -ldflags "-X main.version=v0.1.0 -X main.commit=$(git rev-parse HEAD) -X main.builtAt=$(date -u +%FT%TZ)"`.

## Quick start

```bash
go vet ./... && go test ./...
```

## How it grows

The roadmap lives in [PRD.md](PRD.md) and [features/features.yml](features/features.yml) (ids `ZID-NNN`, most `status: planned`). Work arrives as pull requests from the hub's evolution and fan-out loops, the `@claude` handler, or the local `evolve` lane (off until `EVOLVE_ENABLED=true`). A human merges every PR. Rules: [docs/constitution.md](docs/constitution.md) · agent guide: [AGENTS.md](AGENTS.md).

## License

[MIT](LICENSE)
