# zer0-id

> A small OpenID Connect identity provider in Go with passkey (WebAuthn) login.

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

## Quick start

```bash
go vet ./... && go test ./...
```

## How it grows

The roadmap lives in [PRD.md](PRD.md) and [features/features.yml](features/features.yml) (ids `ZID-NNN`, most `status: planned`). Work arrives as pull requests from the hub's evolution and fan-out loops, the `@claude` handler, or the local `evolve` lane (off until `EVOLVE_ENABLED=true`). A human merges every PR. Rules: [docs/constitution.md](docs/constitution.md) · agent guide: [AGENTS.md](AGENTS.md).

## License

[MIT](LICENSE)
