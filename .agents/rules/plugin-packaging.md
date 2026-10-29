# Plugin Packaging

- `plugin.json` lives in the repository root.
- `routes.json` lives at `ui/src/routes.json` (referenced from `plugin.json`).
- `name` must be kebab-case.
- `binary` must equal the built executable name.
- Every `services` entry must equal the `service.Service.Name` declared in Go.
  `grpc.health.v1.Health` is served by go-plugin's gRPC server itself: declare it,
  but never register it again from the plugin (duplicate registration is FATAL).
- `grants` is an application, not an entitlement: the host re-authorizes every
  call and denies by default.
- `apps` declares the applications the plugin exposes to the **host policy system**:
  policy 「服务」 = an app code, policy 「操作」 = one app operation. App and
  operation codes are kebab-case and become sam function names
  (`<app code>.<operation name>`), so both segments must stay lowercase and stable.
- Every `operations[]` entry needs `name/title/kind`, with `kind` one of
  `read|write|list|other` (读/写/列表/其他); `resource` must reference a `resources[].code`.
- Removing an operation is a **breaking policy change**: the host removes the
  function from its index; existing policies referencing it are flagged by the
  policy analyzer until edited.
- UI packaging requires `ui/index.html`; keep route declarations valid JSON.
- Package output is `build/plugin/$(name)-v$(version).mgpn` and is never committed.
- A `.mgpn` is a **gzip-compressed** tar containing `plugin.json` at its root
  (optionally under one top-level directory); the host rejects a plain tar.
  Package with `tar -czf` — `tar -caf` would infer no compression from the
  `.mgpn` suffix and produce an unreadable archive.
