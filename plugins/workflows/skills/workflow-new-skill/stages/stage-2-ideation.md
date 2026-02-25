# Stage 2 — Ideation

**Idempotency check:**
Run: !`ls artifacts/skills/<SLUG>/ideation.md 2>/dev/null`
If file exists: output `[skip] artifacts/skills/<SLUG>/ideation.md already exists.` and proceed to Stage 3.

Run: !`mkdir -p artifacts/skills/<SLUG>`

Read `.volon/pcc/01_architecture.md` for context on existing skills.
List `plugins/<plugin>/skills/` to check for name collision.

Create `artifacts/skills/<SLUG>/ideation.md`:

```
---
id: "skill-<TODAY>-<SLUG>"
type: "idea"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["skill", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Skill Name> — Ideation

## What does this skill do?
<One sentence — complete the prompt: "This skill allows a user to...">

## What triggers it?
<User slash command, orchestrator dispatch, loop step>

## What does it produce?
<Output: task file, artifact, modified file, printed report, etc.>

## Why does it need to exist?
<Gap in current skill set that this fills>

## Plugin
Target plugin directory: plugins/<plugin>/

## Decisions
- TBD

## Open questions
- TBD

## Evidence
- Workflow: workflow-new-skill "<skill name>" — <TODAY>
```

Proceed to Stage 3.
