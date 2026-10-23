---
name: add-plugin-service
description: Add or update a gRPC service exposed by the starter plugin.
---

1. Inspect `src/apps/starters.go` and the existing service registration.
2. Define or import a stable proto-generated service contract.
3. Implement it in `src/apps`, keeping host access behind `host.Client`.
4. Add the exact service name to `plugin/plugin.json`.
5. Add focused tests and run `make verify`.
