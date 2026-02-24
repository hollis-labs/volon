# Stage 5 — Plan

**Idempotency check:**
Run: !`ls artifacts/plugins/<SLUG>/plan.md 2>/dev/null`
If file exists: output `[skip] artifacts/plugins/<SLUG>/plan.md already exists.` and proceed to Stage 6.

Read `artifacts/plugins/<SLUG>/spec.md`.

Create `artifacts/plugins/<SLUG>/plan.md`:

```
---
id: "plugin-<TODAY>-<SLUG>"
type: "plan"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["plugin", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Plugin Name> — Plan

## Files to create
| Path | Purpose |
|---|---|
| plugins/<SLUG>/plugin.json | Plugin manifest |
| plugins/<SLUG>/skills/<skill-1>/SKILL.md | Skill stub |
| plugins/<SLUG>/skills/<skill-2>/SKILL.md | Skill stub |

## Files to update
| Path | Change |
|---|---|
| forge.yaml | Add plugin enable flags (if applicable) |
| .forge/pcc/01_architecture.md | Add plugin to architecture summary |

## Task breakdown
| # | Task | Priority |
|---|---|---|
| 1 | Create plugin scaffold (plugin.json + dirs) | B |
| 2 | Write skill stubs | B |
| 3 | Update forge.yaml | C |
| 4 | Update PCC architecture | C |

## Post-scaffold
After scaffold tasks complete, use workflow-new-skill for each skill that needs full implementation.

## Verification plan
- [ ] plugins/<SLUG>/plugin.json exists and is valid JSON
- [ ] Each skill has a SKILL.md stub
- [ ] Plugin loadable via --plugin-dir
- [ ] forge.yaml references plugin (if applicable)

## Decisions
- TBD

## Evidence
- Input: artifacts/plugins/<SLUG>/spec.md
- Workflow: workflow-new-plugin "<plugin name>" — <TODAY>
```

Proceed to Stage 6.
