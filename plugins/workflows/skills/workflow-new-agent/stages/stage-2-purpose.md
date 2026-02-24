# Stage 2 — Purpose

**Idempotency check:**
Run: !`ls artifacts/agents/<SLUG>/purpose.md 2>/dev/null`
If file exists: output `[skip] artifacts/agents/<SLUG>/purpose.md already exists.` and proceed to Stage 3.

Run: !`mkdir -p artifacts/agents/<SLUG>`

Read `.forge/pcc/00_project.md` and `.forge/agent-boot.md` for context on existing agent roles.
Read `.forge/boot/` to understand current role definitions.

Create `artifacts/agents/<SLUG>/purpose.md`:

```
---
id: "agent-<TODAY>-<SLUG>"
type: "idea"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["agent", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Agent Name> — Purpose

## What problem does this agent solve?
<One paragraph: the gap in current agent roles that this fills>

## Role type
[ ] Orchestrator  [ ] Worker  [ ] Reviewer/Investigator  [ ] Hybrid (describe)

## When is this agent invoked?
<Trigger conditions: user request, task type, scheduled, sub-agent dispatch>

## What does it NOT do?
<Explicit out-of-scope to prevent scope creep>

## Decisions
- TBD

## Open questions
- TBD

## Evidence
- Workflow: workflow-new-agent "<agent name>" — <TODAY>
```

Proceed to Stage 3.
