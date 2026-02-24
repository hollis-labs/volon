# Stage 2 — Ideation

**Idempotency check:**
Run: !`ls artifacts/plugins/<SLUG>/ideation.md 2>/dev/null`
If file exists: output `[skip] artifacts/plugins/<SLUG>/ideation.md already exists.` and proceed to Stage 3.

Run: !`mkdir -p artifacts/plugins/<SLUG>`

Read `.forge/pcc/01_architecture.md` and list `plugins/` to understand existing plugin domains.

Create `artifacts/plugins/<SLUG>/ideation.md`:

```
---
id: "plugin-<TODAY>-<SLUG>"
type: "idea"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["plugin", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Plugin Name> — Ideation

## What domain does this plugin cover?
<One paragraph: the capability domain (e.g. "integrations with external issue trackers")>

## Why a new plugin (vs. adding to existing)?
<Justify the separation: different concern, different deployment unit, optional extension>

## Likely skills this plugin will contain
<Bulleted list of skill names with one-line descriptions>

## Expected consumers
[ ] Users via slash commands  [ ] Orchestrator dispatch  [ ] Both  [ ] Automated only

## Decisions
- TBD

## Open questions
- TBD

## Evidence
- Workflow: workflow-new-plugin "<plugin name>" — <TODAY>
```

Proceed to Stage 3.
