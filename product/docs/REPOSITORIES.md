# MemoGuard repository map

```text
memoguard-rules ──► memoguard-engine ──► memoguard-cli ──► memoguard-action
                           │
                           └──► Go applications may import the engine directly
```

The arrows show use of a released artifact or Go package. The repositories have separate responsibilities and release independently.

| Repository | Owns | Does not own | Depends on |
| --- | --- | --- | --- |
| `memoguard-rules` | Policy schema, validation, built-in rule pack, policy versioning | XDR decoding, CLI, GitHub behavior | None |
| `memoguard-engine` | XDR decoding, public-field extraction, scanning, safe findings | Policy authoring, terminal UX, CI annotations | `memoguard-rules`, Stellar Go SDK |
| `memoguard-cli` | Flags, local file/stdin input, output formats, exit codes | Detection algorithms, rule definitions | `memoguard-engine` |
| `memoguard-action` | Action manifest, CLI download/pinning, GitHub annotations | Scanner, policy rules, XDR parser | Released `memoguard-cli` binary |

## Release order

1. Tag `memoguard-rules` v0.1.0.
2. Pin that tag in `memoguard-engine`, test, then tag v0.1.0.
3. Pin the engine tag in `memoguard-cli`, publish binaries, then tag v0.1.0.
4. Pin a CLI release and checksum in `memoguard-action`, then tag v0.1.0.

Use a parent `go.work` file for local development only. Published `go.mod` files must resolve released versions without the parent workspace.

## Repository names and remotes

- `https://github.com/Memoguard8876/memoguard-rules.git`
- `https://github.com/Memoguard8876/memoguard-engine.git`
- `https://github.com/Memoguard8876/memoguard-cli.git`
- `https://github.com/Memoguard8876/memoguard-action.git`
