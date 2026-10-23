# Plugin Packaging

- `name` must be kebab-case and the package directory name must match it.
- `binary` must equal the built executable name.
- Every `services` entry must equal the `service.Service.Name` registered in Go.
- `grants` is an application, not an entitlement: the host re-authorizes every
  call and denies by default.
- UI packaging requires `ui/index.html`; keep route declarations valid JSON.
- Package output is `build/plugin/modality-plugin-starter/` and is never committed.
