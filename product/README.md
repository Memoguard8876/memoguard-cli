# MemoGuard

![MemoGuard vector mark](brand/memoguard.svg)

MemoGuard is a Go-based privacy check for Stellar applications. It finds sensitive information in transaction data **before submission**, so teams can fix a leak before that information becomes public and permanent.

MemoGuard is a developer tool, not a web app. Its first delivery is a Go scanning engine, a CLI, and a GitHub Action. It checks decoded Stellar transaction envelopes, memos, operation fields, Soroban invocation arguments, and simulation output when available. It reports the location and rule that matched without copying the sensitive value into logs.

## Start here

1. [Product requirements](docs/PRD.md)
2. [Repository boundaries and build order](docs/REPOSITORIES.md)
3. [Technical architecture](docs/ARCHITECTURE.md)
4. [Security and privacy rules](docs/SECURITY.md)
5. [Delivery plan](docs/BUILD_PLAN.md)
6. [Decision log](docs/DECISIONS.md)
7. [Interfaces](docs/INTERFACES.md) and [test plan](docs/TEST_PLAN.md)
8. [Logo assets and usage](brand/README.md)
9. [Why Stellar and MVP demo](docs/WHY_STELLAR.md)
10. [Validation record](docs/VALIDATION.md)

The four Git repositories in this folder are `memoguard-rules`, `memoguard-engine`, `memoguard-cli`, and `memoguard-action`. The documents in `docs/` apply to the whole product; each repository also has its own focused README.

## Product status

The four Go repositories now contain a working rules library, Stellar scanning engine, CLI, and GitHub Action. Tagged private modules and the CLI release are published. Local tests, remote CI, and the Action's clean/blocked fixture workflow pass. Pilot teams should tune the rules with synthetic and consented transactions before enabling blocking in their own CI.
