# Stage 3 — Scope & Constraints

**Idempotency check:**
Run: !`ls artifacts/agents/<SLUG>/scope.md 2>/dev/null`
If file exists: output `[skip] artifacts/agents/<SLUG>/scope.md already exists.` and proceed to Stage 4.

Read `artifacts/agents/<SLUG>/purpose.md`.
Read `docs/07_subagents.md` and `docs/08_orchestrator.md` for write/read-only constraints.

Create `artifacts/agents/<SLUG>/scope.md`:

```
---
id: "agent-<TODAY>-<SLUG>"
type: "requirements"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["agent", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Agent Name> — Scope & Constraints

## What this agent MAY read
<List of paths/patterns it reads during execution>

## What this agent MAY write
<List of paths/patterns it writes — or "READ ONLY" if worker/reviewer>

## What this agent must NOT do
- [ ] Edit files outside its authorized paths
- [ ] Update tasks/logs/PCC/bootstrap (if worker/reviewer)
- [ ] Spawn sub-agents (if worker/reviewer)
- <add role-specific constraints>

## Execution boundary
<Max scope: single task, single file scan, bounded analysis, etc.>

## Failure behaviour
<What to return if input is missing/malformed; do not silently fail>

## Decisions
- TBD

## Evidence
- Input: artifacts/agents/<SLUG>/purpose.md
- Workflow: workflow-new-agent "<agent name>" — <TODAY>
```

Proceed to Stage 4.
