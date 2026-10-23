# AGENTS.md

Entry point for AI agents working in `modality-plugin-starter`. Read this file
first, then load only the rules and skills relevant to the current task.

| Section | Path |
| --- | --- |
| Rules (must follow) | [.agents/rules/](.agents/rules/) |
| Skills (recurring procedures) | [.agents/skills/](.agents/skills/) |
| Notes (session records) | [.agents/notes/](.agents/notes/) |

## Project Overview

`modality-plugin-starter` is a Go template for third-party modality plugin
authors. It separates the backend entry, implementation tests, plugin resources,
and optional UI while using a Go-native build and packaging workflow.

| Aspect | Value |
| --- | --- |
| Language | Go 1.27.3 |
| Module | `github.com/slaincsuite/modality-plugin-starter` |
| SDK | `github.com/slaincsuite/modality-gopn-sdk-go` |
| Plugin process | standalone binary launched by the modality host through gopn |
| UI | optional Vue 3 + Vite project under `ui/` |
| License | GPL-3.0-only |
| State | bootstrap template; not published |

## Technology Stack

| Concern | Choice |
| --- | --- |
| Backend | Go, gRPC health service |
| Plugin serving | `modality-gopn-sdk-go/service` |
| UI | Vue 3, Vite, TypeScript, Vitest |
| Package | Makefile (Go-native) |

## Directory Map

```text
modality-plugin-starter/
├── AGENTS.md
├── Makefile
├── README.md
├── go.mod / go.sum
├── plugin.json                    # plugin manifest (root)
├── src/                           # Go source
│   ├── starter.go                # implementation and tests
│   └── apps/main.go               # process entry
├── ui/                           # optional Vue plugin UI (routes at ui/src/routes.json)
└── .agents/                      # repository-local agent knowledge
```

## Rules Index

| File | Scope |
| --- | --- |
| [coding-constraints.md](.agents/rules/coding-constraints.md) | Go and TypeScript style |
| [development-workflow.md](.agents/rules/development-workflow.md) | inventory, implementation, verification |
| [plugin-packaging.md](.agents/rules/plugin-packaging.md) | manifest, binary, UI package rules |
| [synchronization.md](.agents/rules/synchronization.md) | upstream sync checklist |

## Skills Index

| Skill | Use when |
| --- | --- |
| [add-plugin-service](.agents/skills/add-plugin-service/SKILL.md) | adding a gRPC service to the starter |
| [update-plugin-ui](.agents/skills/update-plugin-ui/SKILL.md) | changing routes or console pages |

## Notes Index

| File | Topic |
| --- | --- |
| [2026-10-23-starter-bootstrap.md](.agents/notes/2026-10-23-starter-bootstrap.md) | Go-native starter decisions and layout |

## Commands

| Command | Purpose |
| --- | --- |
| `make verify` | Go tests, vet, sync checks |
| `make package` | build UI + binary + plugin directory |
| `make sync-check` | verify local agent knowledge and sync checklist |
| `make tidy` | tidy Go dependencies |

## Conventions

- Agent docs are English; READMEs, code comments, and UI text are Chinese.
- Keep this repository independent: never import `modality` host packages.
- The SDK is imported from the published module version.
- Do not commit build output.
