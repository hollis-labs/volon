# Stage 4 — Context Requirements

**Idempotency check:**
Run: !`ls artifacts/agents/<SLUG>/context.md 2>/dev/null`
If file exists: output `[skip] artifacts/agents/<SLUG>/context.md already exists.` and proceed to Stage 5.

Read `artifacts/agents/<SLUG>/scope.md`.
Read `.volon/agent-boot.md` to understand what the core boot pack already provides.

Create `artifacts/agents/<SLUG>/context.md`:

```
---
id: "agent-<TODAY>-<SLUG>"
type: "spec"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["agent", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Agent Name> — Context Requirements

## What this agent needs at init
<List what it must read before executing any task>

## Provided by core boot (.volon/agent-boot.md)
<Items already covered — do not duplicate>
- Ground truth paths (volon.yaml, bootstrap, PCC, tasks, logs)
- Core rules (single-writer, minimal diffs, bootstrap boundaries)
- Role addendum location

## Additional context this agent requires
<Items NOT in core boot that this agent specifically needs>

## Role addendum
Does this agent need a new entry in `.volon/boot/`?
[ ] Yes — describe what to add
[ ] No — covered by existing orchestrator.md / worker.md / reviewer.md

## Decisions
- TBD

## Evidence
- Input: artifacts/agents/<SLUG>/scope.md
- Workflow: workflow-new-agent "<agent name>" — <TODAY>
```

Proceed to Stage 5.
