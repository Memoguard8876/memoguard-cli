# Changelog

## v0.2.2 — 2026-10-08

- Make CI and release builds use the exact public module versions pinned in `go.mod`.
- Run tests and vet within the release job before publishing binaries.

## v0.2.1 — 2026-10-08

- Upgrade the engine so supported input beyond scan limits fails explicitly rather than passing as clean.
- Scan numeric values in decoded JSON with configured policy rules.

## v0.2.0 — 2026-10-08

- Add redacted SARIF output and configurable `--fail-on` threshold.
- Build public release binaries without cross-repository read tokens.
- Pick up expanded XDR memo coverage from `memoguard-engine` v0.2.0.

## v0.1.2 — 2026-10-07

- Corrected private module resolution in CI and release builds.

## v0.1.0 — 2026-10-07

- Added local and CI scanning, safe reports, policy input, and stable exit codes.
