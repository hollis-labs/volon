# Stage 4 — PRD

**Idempotency check:**
Run: !`ls artifacts/prd/<SLUG>.md 2>/dev/null`

If file exists: output `[skip] artifacts/prd/<SLUG>.md already exists.` and proceed to Stage 5.

Otherwise:

Run: !`mkdir -p artifacts/prd`

Read `artifacts/requirements/<SLUG>.md`. Derive user flows, success criteria,
failure flows. Frame for `audience` (same mapping as Stage 3).

Create `artifacts/prd/<SLUG>.md`:

```
---
id: "feat-<TODAY>-<SLUG>"
type: "prd"
intent: "project_doc"
status: draft
project: "<project.name>"
tags: ["feature"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Feature Name> — PRD

## Summary
<One paragraph>

## User / Agent Flow
<Step-by-step invocation flow with example inputs/outputs>

## Success Criteria
<Measurable table or list>

## Failure Flows
<Table: condition → behaviour>

## Decisions
<Architectural or scoping decisions>

## Open questions
- TBD

## Evidence
- Input: artifacts/requirements/<SLUG>.md
- Workflow: workflow-new-feature "<feature name>" — <TODAY>
```

Proceed to Stage 5.
