# Stage 5 — Interface Spec

**Idempotency check:**
Run: !`ls artifacts/agents/<SLUG>/interface.md 2>/dev/null`
If file exists: output `[skip] artifacts/agents/<SLUG>/interface.md already exists.` and proceed to Stage 6.

Read `artifacts/agents/<SLUG>/context.md`.

Create `artifacts/agents/<SLUG>/interface.md`:

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

# <Agent Name> — Interface Spec

## Invocation
How is this agent launched?
[ ] Via Task tool (subagent_type or prompt)
[ ] Via slash command: /<command-name>
[ ] Via orchestrator dispatch in loop
[ ] Scheduled/automated

## Inputs
| Input | Type | Required | Description |
|---|---|---|---|
| ... | ... | ... | ... |

## Outputs
| Output | Format | Description |
|---|---|---|
| ... | ... | ... |

## Return format
<Exact structure the orchestrator expects: bullets/table/JSON/markdown>

## Error conditions
| Condition | Behaviour |
|---|---|
| Missing required input | ... |
| Target not found | ... |

## Decisions
- TBD

## Evidence
- Input: artifacts/agents/<SLUG>/context.md
- Workflow: workflow-new-agent "<agent name>" — <TODAY>
```

Proceed to Stage 6.
