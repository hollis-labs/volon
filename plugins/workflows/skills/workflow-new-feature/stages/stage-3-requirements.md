# Stage 3 — Requirements

**Idempotency check:**
Run: !`ls artifacts/requirements/<SLUG>.md 2>/dev/null`

If file exists: output `[skip] artifacts/requirements/<SLUG>.md already exists.` and proceed to Stage 4.

Otherwise:

Run: !`mkdir -p artifacts/requirements`

Read `artifacts/ideas/<SLUG>.md`. Use `audience` to frame acceptance criteria:
- `user`: focus on end-user observable behaviour
- `dev`: focus on API contracts, error handling, internal behaviour
- `both`: include both user-visible and developer-visible criteria

Create `artifacts/requirements/<SLUG>.md`:

```
---
id: "feat-<TODAY>-<SLUG>"
type: "requirements"
intent: "project_doc"
status: draft
project: "<project.name>"
tags: ["feature"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Feature Name> — Requirements

## Summary
<Restate problem + success definition in 2–3 sentences>

## Acceptance Criteria
<Numbered or bulleted AC list derived from ideation, audience-framed>

## Decisions
<Key scoping decisions made>

## Open questions
- TBD

## Evidence
- Input: artifacts/ideas/<SLUG>.md
- Workflow: workflow-new-feature "<feature name>" — <TODAY>
```

Proceed to Stage 4.
