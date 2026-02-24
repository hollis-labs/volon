# Stage 5 — Plan

**Idempotency check:**
Run: !`ls artifacts/slash-commands/<SLUG>/plan.md 2>/dev/null`
If file exists: output `[skip] artifacts/slash-commands/<SLUG>/plan.md already exists.` and proceed to Stage 6.

Read `artifacts/slash-commands/<SLUG>/spec.md`.

Create `artifacts/slash-commands/<SLUG>/plan.md`:

```
---
id: "cmd-<TODAY>-<SLUG>"
type: "plan"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["slash-command", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# /<Command Name> — Plan

## Files to create
| Path | Purpose |
|---|---|
| plugins/<plugin>/skills/<SLUG>/SKILL.md | Skill implementation (if new) |

## Files to update
| Path | Change |
|---|---|
| plugins/<plugin>/plugin.json | Add "<SLUG>" to skills array |
| docs/09_commands.md | Add /<SLUG> command entry |

## Task breakdown
| # | Task | Priority |
|---|---|---|
| 1 | Create/update SKILL.md | B |
| 2 | Update plugin.json | B |
| 3 | Add entry to docs/09_commands.md | B |

## Verification plan
- [ ] /<SLUG> appears in docs/09_commands.md
- [ ] Backing skill file exists
- [ ] plugin.json lists the skill

## Dependencies
<List any prerequisite tasks (e.g. workflow-new-skill first if new skill needed)>

## Decisions
- TBD

## Evidence
- Input: artifacts/slash-commands/<SLUG>/spec.md
- Workflow: workflow-new-slash-command "<command name>" — <TODAY>
```

Proceed to Stage 6.
