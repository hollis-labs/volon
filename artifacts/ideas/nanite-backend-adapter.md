---
id: "feat-2026-02-21-nanite-backend-adapter"
type: "idea"
status: draft
project: forge
tags: ["feature"]
priority: B
created_at: "2026-02-21"
updated_at: "2026-02-21"
---

# Nanite Backend Adapter — Idea

## Summary

Forge currently stores all tasks as local markdown files under `.forge/tasks/`.
This means task state is not visible outside the development environment and
cannot be accessed from other devices or tools. The Nanite backend adapter would
implement the `storage.backend: nanite` path in `forge.yaml`, routing task
create/read/update operations through the user's Nanite instance (a personal
note-taking and knowledge-base system). Tasks created by `/task-create` would
appear in the user's Nanite inbox; `/task-list` and `/task-update` would read
and write from Nanite rather than from disk. The file backend would remain the
default; Nanite becomes an opt-in alternative. This closes the gap between
forge's task model and the user's existing personal workflow system.

## Decisions

- TBD: Nanite API surface (HTTP? local IPC? MCP?)
- TBD: mapping between forge task frontmatter fields and Nanite item schema
- TBD: whether forge task IDs are stored as Nanite tags or in item body

## Open questions

- How does Nanite authenticate? Token in `forge.yaml`? Environment variable?
- What happens when a Nanite item is deleted outside of forge — does `/task-list` error?
- Should `/task-list` return items created outside of forge (non-forge Nanite items)?

## Evidence

- Workflow: workflow-new-feature "Nanite backend adapter" — 2026-02-21
- Deferred roadmap item: `.forge/pcc/04_backlog.md` (Nanite backend adapter)
