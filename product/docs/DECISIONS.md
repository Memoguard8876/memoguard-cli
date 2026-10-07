# MemoGuard decisions

| ID | Decision | Reason | Revisit when |
| --- | --- | --- | --- |
| MG-001 | Go is the implementation language for core logic and CLI. | Shared engine for local tools and Go services. | Another language has proven demand. |
| MG-002 | The scanner runs offline by default. | The input may contain the very data being protected. | A hosted mode has explicit user demand and privacy design. |
| MG-003 | The Action invokes the released CLI. | One detection implementation and reproducible CI behavior. | Platform packaging requires a different delivery method. |
| MG-004 | Blocking applies to pre-submit input only. | Historic on-chain data cannot be withdrawn. | Never for past data; audit mode may be added. |
| MG-005 | No Soroban contract is needed for MVP. | This is a developer safety tool. | An on-chain requirement appears that Go cannot satisfy. |
