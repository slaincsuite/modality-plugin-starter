# Coding Constraints

## Go

- Use Go 1.27.3 and the module path `github.com/slaincsuite/modality-plugin-starter`.
- Every Go file begins with one short Chinese `//` comment.
- Public identifiers have Chinese doc comments beginning with the identifier.
- Keep the plugin independent: it may import the SDK, gRPC, protobuf, and its
  own packages; it must not import modality host packages.
- Tests use stdlib `testing` (or testify when needed) and explicit waits rather
  than sleep-and-hope.

## UI

- TypeScript strict settings come from the existing Vue tsconfig files.
- Routes declared by `ui/src/index.ts` must match `plugin/plugin.json`
  and `plugin/ui/routes.json`.
- UI text is Chinese.

## Hygiene

- No generated `dist/` or `build/` artifacts.
- No secrets or machine-specific absolute paths.
