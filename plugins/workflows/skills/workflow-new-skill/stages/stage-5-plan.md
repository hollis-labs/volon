# Stage 5 — Plan

**Idempotency check:**
Run: !`ls artifacts/skills/<SLUG>/plan.md 2>/dev/null`
If file exists: output `[skip] artifacts/skills/<SLUG>/plan.md already exists.` and proceed to Stage 6.

Read `artifacts/skills/<SLUG>/spec.md`.
Read `plugins/<plugin>/plugin.json` to understand current skill registration.

Create `artifacts/skills/<SLUG>/plan.md`:

```
---
id: "skill-<TODAY>-<SLUG>"
type: "plan"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["skill", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Skill Name> — Plan

## Files to create
| Path | Content |
|---|---|
| plugins/<plugin>/skills/<SLUG>/SKILL.md | Skill implementation |

## Files to update
| Path | Change |
|---|---|
| plugins/<plugin>/plugin.json | Add "<SLUG>" to skills array |
| docs/09_commands.md | Add /<command> entry (if user-facing) |

## Standalone flag
[ ] standalone: true (add to quality-run exemption list)
[ ] standalone: false (default — included in dead_code scan)

## Task breakdown
| # | Task | Priority |
|---|---|---|
| 1 | Write SKILL.md | B |
| 2 | Update plugin.json | B |
| 3 | Update docs/09_commands.md (if user-facing) | C |

## Verification plan
- [ ] SKILL.md exists at correct path
- [ ] plugin.json lists the skill
- [ ] Skill invocable from plugin dir
- [ ] dead_code scan: standalone flag correct

## Decisions
- TBD

## Evidence
- Input: artifacts/skills/<SLUG>/spec.md
- Workflow: workflow-new-skill "<skill name>" — <TODAY>
```

Proceed to Stage 6.
