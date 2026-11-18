<!-- kit: governance v0.1.0 -->
# 0002 — Missions

- **Status:** proposed
- **Date:** 2026-11-16
- **Proposed by:** Grok Bot (agent), at the maintainer's request
- **Ratified by:**
- **Affects:** constitution §1; PRD §1.1; feature ZID-008
- **Supersedes:** none
- **Related:** [PRD.md](../../PRD.md) §1.1; feature ZID-008

## Context

Constitution §1 Mission was a product description (or, for a new constitution, had not been written yet), not a three-mission direction statement. The fleet now requires every repo's constitution §1 to be one direction statement followed by exactly three numbered missions, adopted through a numbered amendment the maintainer ratifies by merging. The maintainer asked for those missions to be comprehensive and epic: zer0-id as the ultimate open home for all things open identity and single sign-on. Mission 1 follows the maintainer's wording, adapted to this domain: provide data, tools, resources and transparency for all stakeholders, domain practitioners and the general public. The ambition is the direction, not a claim that the work is already finished. The missions are aligned with this repo's existing constitution and PRD.

## Amendment

Replace constitution §1 with the following, retitled **Missions**:

zer0-id's direction is to become the ultimate open home for all things open identity and single sign-on — data, tools, resources and transparency, gathered in one place that anyone can inspect and reuse.

1. **Provide data, tools, resources and transparency for all stakeholders, identity practitioners and the general public.** Operators, application authors and anyone signing in are first-class audiences. The aim is a shared, inspectable commons: the same protocols, the same keys, and the same login flows for everyone.
2. **Build single sign-on from the specs up: OIDC discovery, authorization code plus PKCE, signed JWTs and passwordless passkeys, so every login is standards-based and auditable.** A small identity provider should be something a practitioner can read against the specifications, not a proprietary black box.
3. **Serve as SSO for the fleet's apps — the console, ragio-ai, fredgar-ai, zer0-CMS and fleet-mobile — with identities that stay under the operator's control.** Fleet single sign-on should be reusable, reviewable and local-first where it can be, without handing account control to a third-party IdP by default.

These missions set the course; they do not claim the work is finished. They serve the PRD's learning-first path: an OpenID Connect provider with passkey login, built from the specifications up for the fleet's apps.

_Source: [PRD.md](../PRD.md) §1.1._


## Consequences

Agents and reviewers judge fit against these three missions before inventing scope, treating them as the direction of the repo rather than a description of what already ships. The same text is added as PRD §1.1 so the constitution's `_Source:` line stays accurate. The original PRD purpose / overview paragraph is left in place. This PR also aligns the README and AGENTS.md taglines with the missions and links them to constitution §1, and adds feature ZID-008 to the feature catalog, traced to requirement MR1. Merging this PR ratifies amendment 0002. Until then the record stays `Status: proposed`.
