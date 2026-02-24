# Stage 4 — Spec

**Idempotency check:**
Run: !`ls artifacts/slash-commands/<SLUG>/spec.md 2>/dev/null`
If file exists: output `[skip] artifacts/slash-commands/<SLUG>/spec.md already exists.` and proceed to Stage 5.

Read `artifacts/slash-commands/<SLUG>/requirements.md`.

Determine implementation approach:
- **Backed by existing skill**: maps to `plugins/<plugin>/skills/<skill>/SKILL.md`
- **New skill required**: needs workflow-new-skill first (note as dependency)
- **Standalone command**: implemented directly in SKILL.md with same name

Create `artifacts/slash-commands/<SLUG>/spec.md`:

```
---
id: "cmd-<TODAY>-<SLUG>"
type: "spec"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["slash-command", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# /<Command Name> — Spec

## Implementation
[ ] Backed by existing skill: `plugins/<plugin>/skills/<skill>/SKILL.md`
[ ] New skill required: run workflow-new-skill "<skill-name>" first
[ ] Standalone (SKILL.md is the command implementation)

## Backing skill path (if applicable)
`plugins/<plugin>/skills/<SLUG>/SKILL.md`

## docs/09_commands.md entry (draft)
```markdown
## /<SLUG> [args...]

### Purpose
<one sentence>

### Inputs
- `arg`: description

### Outputs
- <what is produced>
```

## Decisions
- TBD

## Evidence
- Input: artifacts/slash-commands/<SLUG>/requirements.md
- Workflow: workflow-new-slash-command "<command name>" — <TODAY>
```

Proceed to Stage 5.
