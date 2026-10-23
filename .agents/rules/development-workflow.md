# Development Workflow

1. Inventory `AGENTS.md`, this directory, `src/apps/`, and `ui/`
   before creating an artifact.
2. Extend the existing service, test, route, or package target instead of
   adding a parallel structure.
3. Update `plugin.json` whenever service names, grants, extension points, or UI
   routes change.
4. Run `make verify` for backend changes and `make package` for packaging
   changes.
5. Record durable decisions in `.agents/notes/YYYY-MM-DD-<topic>.md` and update
   the AGENTS notes index.
6. Never commit unless explicitly requested.
