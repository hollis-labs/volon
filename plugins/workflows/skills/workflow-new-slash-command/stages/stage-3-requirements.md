# Stage 3 — Requirements

**Idempotency check:**
Run: !`ls artifacts/slash-commands/<SLUG>/requirements.md 2>/dev/null`
If file exists: output `[skip] artifacts/slash-commands/<SLUG>/requirements.md already exists.` and proceed to Stage 4.

Read `artifacts/slash-commands/<SLUG>/ideation.md`.

Create `artifacts/slash-commands/<SLUG>/requirements.md`:

```
---
id: "cmd-<TODAY>-<SLUG>"
type: "requirements"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["slash-command", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# /<Command Name> — Requirements

## Argument spec
| Arg | Required | Description |
|---|---|---|
| ... | yes/no | ... |

## Expected behaviour (step by step)
1. ...
2. ...
N. Output: ...

## Output format
<What is printed/written when the command succeeds>

## Error cases
| Condition | Response |
|---|---|
| Missing required arg | ERROR: <message> |
| ... | ... |

## Idempotency
[ ] Idempotent — describe check:
[ ] Not idempotent — describe risk:

## Orchestrator Mode compliance
[ ] Respects single-writer rule
[ ] Does not spawn agents
[ ] Externalizes state before any context reset

## Decisions
- TBD

## Evidence
- Input: artifacts/slash-commands/<SLUG>/ideation.md
- Workflow: workflow-new-slash-command "<command name>" — <TODAY>
```

Proceed to Stage 4.
