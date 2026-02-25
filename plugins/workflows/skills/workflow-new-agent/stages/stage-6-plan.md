# Stage 6 — Plan

**Idempotency check:**
Run: !`ls artifacts/agents/<SLUG>/plan.md 2>/dev/null`
If file exists: output `[skip] artifacts/agents/<SLUG>/plan.md already exists.` and proceed to Stage 7.

Read `artifacts/agents/<SLUG>/interface.md`. Identify files to create/update.

Create `artifacts/agents/<SLUG>/plan.md`:

```
---
id: "agent-<TODAY>-<SLUG>"
type: "plan"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["agent", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Agent Name> — Plan

## Files to create
| Path | Purpose |
|---|---|
| .volon/boot/<SLUG>.md | Role addendum (if needed) |
| ... | ... |

## Files to update
| Path | Change |
|---|---|
| .volon/agent-boot.md | Add role entry (if new role type) |
| docs/07_subagents.md | Add agent cap or policy (if applicable) |
| ... | ... |

## Task breakdown
| # | Task | Priority |
|---|---|---|
| 1 | Create role addendum file | A/B/C |
| 2 | Update agent-boot.md reference | A/B/C |
| ... | ... | ... |

## Verification plan
- [ ] Role addendum exists at correct path
- [ ] agent-boot.md references the role
- [ ] Agent can initialize from boot pack without reading all docs

## Decisions
- TBD

## Evidence
- Input: artifacts/agents/<SLUG>/interface.md
- Workflow: workflow-new-agent "<agent name>" — <TODAY>
```

Proceed to Stage 7.
