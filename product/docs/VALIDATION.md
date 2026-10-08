# MemoGuard validation record

Date: 8 October 2026.

- All four Go repositories pass local `go test ./...` and `go vet ./...` checks.
- The rules and engine fuzz targets each completed a short arbitrary-input smoke run without a panic.
- `memoguard-rules` v0.1.1 and `memoguard-engine` v0.2.1 are public tagged modules. The engine has checked-in synthetic clean and unsafe XDR envelopes, memo-type tests, simulation tests, and fail-closed extraction-limit tests.
- `memoguard-cli` v0.2.2 is a public release with Linux, macOS, and Windows binaries plus checksums. Its [release workflow](https://github.com/Memoguard8876/memoguard-cli/actions/runs/37818025204) tests and vets the exact versions in `go.mod` before building.
- The Action downloads a public CLI binary and compares it with a SHA-256 digest pinned in the Action repository. Its integration workflow checks clean, warning, and blocked fixtures on Linux, macOS, and Windows.
- Reports, SARIF results, and annotations identify rules and field paths without copying matched values. The CLI supports `--fail-on block|warning|none`.

The scanner covers supported decoded fields only. Opaque binary content and unsupported XDR shapes still need explicit coverage work. A consenting pilot must review false positives and add its own sanitized transaction shapes before MemoGuard becomes a mandatory production CI gate. See [Wave issues](WAVE.md) for scoped follow-up work.
