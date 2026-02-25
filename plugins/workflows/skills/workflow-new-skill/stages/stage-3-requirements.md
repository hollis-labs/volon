# Stage 3 — Requirements

**Idempotency check:**
Run: !`ls artifacts/skills/<SLUG>/requirements.md 2>/dev/null`
If file exists: output `[skip] artifacts/skills/<SLUG>/requirements.md already exists.` and proceed to Stage 4.

Read `artifacts/skills/<SLUG>/ideation.md`.

Create `artifacts/skills/<SLUG>/requirements.md`:

```
---
id: "skill-<TODAY>-<SLUG>"
type: "requirements"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["skill", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Skill Name> — Requirements

## Arguments
| Arg | Required | Type | Description |
|---|---|---|---|
| $0 | yes/no | string | ... |
| key=value | yes/no | string | ... |

## Inputs (what the skill reads)
- <path or resource>

## Outputs (what the skill writes/returns)
- <path, printed output, or side effect>

## Invariants
- <things that must always be true regardless of inputs>
- Idempotent: <yes/no — if yes, describe check>
- `disable-model-invocation`: true/false

## Acceptance criteria
- [ ] <observable outcome 1>
- [ ] <observable outcome 2>

## Decisions
- TBD

## Evidence
- Input: artifacts/skills/<SLUG>/ideation.md
- Workflow: workflow-new-skill "<skill name>" — <TODAY>
```

Proceed to Stage 4.
