# Stage 3 — Requirements

**Idempotency check:**
Run: !`ls artifacts/plugins/<SLUG>/requirements.md 2>/dev/null`
If file exists: output `[skip] artifacts/plugins/<SLUG>/requirements.md already exists.` and proceed to Stage 4.

Read `artifacts/plugins/<SLUG>/ideation.md`.

Create `artifacts/plugins/<SLUG>/requirements.md`:

```
---
id: "plugin-<TODAY>-<SLUG>"
type: "requirements"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["plugin", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Plugin Name> — Requirements

## Skills to include (initial)
| Skill name | Purpose | User-facing? |
|---|---|---|
| ... | ... | yes/no |

## Config this plugin needs (forge.yaml additions)
```yaml
# proposed additions under workflows: or a top-level key
```

## Dependencies on other plugins
| Plugin | What is needed from it |
|---|---|
| core | pcc-refresh, bootstrap-update |
| ... | ... |

## Constraints
- Plugin must be loadable independently (no hard deps on other non-core plugins)
- Skills must follow SKILL.md frontmatter schema

## Acceptance criteria
- [ ] plugin.json valid
- [ ] All listed skills have SKILL.md stubs
- [ ] Plugin loads without errors when added to --plugin-dir
- [ ] forge.yaml updated with enable flag

## Decisions
- TBD

## Evidence
- Input: artifacts/plugins/<SLUG>/ideation.md
- Workflow: workflow-new-plugin "<plugin name>" — <TODAY>
```

Proceed to Stage 4.
