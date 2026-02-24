# workflows/forge/

This directory is a **reference index** for core Forge workflows.

The authoritative implementations live in `plugins/*/skills/*/SKILL.md`.
Files here are reference definitions following the workflow authoring format
(`docs/08_workflow-authoring.md`) for discoverability and documentation.

## Core Forge Workflows

| Name | Invocation | Plugin Location | Status |
|---|---|---|---|
| workflow-new-feature | `/workflow-new-feature "Name"` | `plugins/workflows/skills/workflow-new-feature/` | active |
| workflow-docs-review | `/workflow-docs-review [focus]` | `plugins/workflows/skills/workflow-docs-review/` | active |
| docs-sync | `/docs-sync [focus]` | `plugins/docsmith/skills/docs-sync/` | active |
| pcc-refresh | `/pcc-refresh [scope]` | `plugins/core/skills/pcc-refresh/` | active |
| worktree-start | `/worktree-start [slug]` | `plugins/git/skills/worktree-start/` | active |
| pr-open | `/pr-open [title] [body]` | `plugins/git/skills/pr-open/` | active |
| quality-run | `/quality-run [mode]` | `plugins/quality/skills/quality-run/` | active |
| backlog-task | `/backlog-task "Title"` | `plugins/backlog/skills/backlog-task/` | active |
| workflow-create | `/workflow-create "name" [domain] [description]` | `plugins/workflow-author/skills/workflow-create/` | active |
| workflow-edit | `/workflow-edit <name> [field=value]` | `plugins/workflow-author/skills/workflow-edit/` | active |
| workflow-clone | `/workflow-clone <source> <new-name>` | `plugins/workflow-author/skills/workflow-clone/` | active |
| workflow-deprecate | `/workflow-deprecate <name> [successor]` | `plugins/workflow-author/skills/workflow-deprecate/` | active |

See `docs/08_workflow-authoring.md` for the full workflow definition format.
