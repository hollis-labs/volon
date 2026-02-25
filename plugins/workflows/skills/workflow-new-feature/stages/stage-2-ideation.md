# Stage 2 — Ideation

**Idempotency check:**
Run: !`ls artifacts/ideas/<SLUG>.md 2>/dev/null`

If file exists: output `[skip] artifacts/ideas/<SLUG>.md already exists.` and proceed to Stage 3.

Otherwise:

Run: !`mkdir -p artifacts/ideas`

If `scope` argument is provided: read files under `<scope>` path to ground the idea.
Read `.volon/pcc/00_project.md` and `.volon/pcc/01_architecture.md` for context.

Create `artifacts/ideas/<SLUG>.md` with this frontmatter and sections:

```
---
id: "feat-<TODAY>-<SLUG>"
type: "idea"
intent: "project_doc"
status: draft
project: "<project.name>"
tags: ["feature"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Feature Name> — Idea

## Summary
<One paragraph: what problem this solves and how>

## Decisions
- TBD

## Open questions
- TBD

## Evidence
- Workflow: workflow-new-feature "<feature name>" — <TODAY>
```

Proceed to Stage 3.
