# Stage 6 — Plan

**Idempotency check:**
Run: !`ls artifacts/plan/<SLUG>.md 2>/dev/null`

If file exists: output `[skip] artifacts/plan/<SLUG>.md already exists.` and proceed to Stage 7.

Otherwise:

Run: !`mkdir -p artifacts/plan`

Read `artifacts/spec/<SLUG>.md`. Break implementation into ≥3 small, independently
verifiable tasks. Each task must have a title, priority, and verification criteria.

Create `artifacts/plan/<SLUG>.md`:

```
---
id: "feat-<TODAY>-<SLUG>"
type: "plan"
intent: "project_doc"
status: draft
project: "<project.name>"
tags: ["feature"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Feature Name> — Plan

## Summary
<One paragraph>

## Task Breakdown
<Table: task title | priority | depends on>

## Task Detail
<Per-task: goal, steps, verification criteria>

## Verification Plan
<End-to-end checklist>

## Decisions
<Sequencing decisions>

## Evidence
- Input: artifacts/spec/<SLUG>.md
- Workflow: workflow-new-feature "<feature name>" — <TODAY>
```

Proceed to Stage 7.
