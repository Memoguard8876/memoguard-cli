# MemoGuard working rules

- Use Go for product logic and CLI code.
- Keep repository boundaries in `docs/REPOSITORIES.md`; do not copy detection code into the CLI or Action.
- Treat input data as sensitive. Never log matched values or full transaction payloads by default.
- Decode Stellar XDR with a maintained Stellar Go library; do not scan opaque base64 as if it were plain text.
- Distinguish pre-submit prevention from after-the-fact audit.
- Add meaningful fixture tests for supported transaction shapes and malformed input.
- Keep the four repositories independently buildable from tagged dependencies before release.
- Do not put credentials, private keys, tokens, or real customer data in this workspace.
