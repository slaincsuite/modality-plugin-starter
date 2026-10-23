# Upstream Synchronization

The upstream repository set includes the modality host and SDK repositories. Their
AGENTS knowledge and plugin contracts can change; this starter repository
must not silently drift from the released SDK contract.

## Checklist

Run this checklist when modality changes any of:

1. `AGENTS.md` or `.agents/rules/` guidance useful to plugin authors;
2. `library/gopn` handshake, manifest, lifecycle, extension point, or access
   contracts;
3. `modules/plugin` plugin management behavior;
4. plugin UI mounting or route declarations;
5. plugin permission codes or host authorization semantics.

Then update this starter:

- [ ] Read the main repository change and filter out purely host-internal work.
- [ ] Update this `AGENTS.md` and `.agents/` when downstream guidance changes.
- [ ] Update SDK usage and `plugin.json` when the public contract changes.
- [ ] Update README debugging or packaging instructions.
- [ ] Run `make sync-check` and `make verify`; run `make package` if resources changed.

## Verification

`make sync-check` verifies that repository-local agent files and the upstream
synchronization rule exist.
