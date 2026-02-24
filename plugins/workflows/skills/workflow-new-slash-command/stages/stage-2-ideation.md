# Stage 2 — Ideation

**Idempotency check:**
Run: !`ls artifacts/slash-commands/<SLUG>/ideation.md 2>/dev/null`
If file exists: output `[skip] artifacts/slash-commands/<SLUG>/ideation.md already exists.` and proceed to Stage 3.

Run: !`mkdir -p artifacts/slash-commands/<SLUG>`

Read `docs/09_commands.md` for existing commands and naming conventions.

Create `artifacts/slash-commands/<SLUG>/ideation.md`:

```
---
id: "cmd-<TODAY>-<SLUG>"
type: "idea"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["slash-command", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# /<Command Name> — Ideation

## Command signature
`/<SLUG> [args...]`

## What does the user type and what happens?
<One paragraph: user-visible behaviour from invocation to output>

## Who uses this command?
[ ] Developer/operator  [ ] Automated loop  [ ] Both

## What problem does it solve?
<Gap in current command set>

## Is it safe by default?
[ ] Read-only (no side effects)  [ ] Writes state (describe)  [ ] Destructive (requires confirmation)

## Decisions
- TBD

## Open questions
- TBD

## Evidence
- Workflow: workflow-new-slash-command "<command name>" — <TODAY>
```

Proceed to Stage 3.
