# Plugin Packaging

- `plugin.json` lives in the repository root.
- `routes.json` lives at `ui/src/routes.json` (referenced from `plugin.json`).
- `name` must be kebab-case.
- `binary` must equal the built executable name.
- Every `services` entry must equal the `service.Service.Name` registered in Go.
- `grants` is an application, not an entitlement: the host re-authorizes every
  call and denies by default.
- UI packaging requires `ui/index.html`; keep route declarations valid JSON.
- Package output is `build/plugin/$(name)-v$(version).mgpn` and is never committed.
