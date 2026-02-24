# Stage 5 — Spec

**Idempotency check:**
Run: !`ls artifacts/spec/<SLUG>.md 2>/dev/null`

If file exists: output `[skip] artifacts/spec/<SLUG>.md already exists.` and proceed to Stage 6.

If `require_spec: true` and this stage is being skipped, warn:
`WARN: require_spec is true but spec artifact already exists — verify it is current.`

Otherwise:

Run: !`mkdir -p artifacts/spec`

Read `artifacts/prd/<SLUG>.md`. If `scope` provided, read files under `<scope>`.
Derive interfaces, shell commands, argument tables, failure codes.

Create `artifacts/spec/<SLUG>.md`:

```
---
id: "feat-<TODAY>-<SLUG>"
type: "spec"
intent: "project_doc"
status: draft
project: "<project.name>"
tags: ["feature"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Feature Name> — Spec

## Summary
<One sentence>

## Interface
<Arguments table, invocation format, output format>

## Shell Commands Used
<Table: purpose → command>

## Failure Modes
<Table: code → trigger → message>

## Decisions
<Implementation decisions>

## Open questions
- TBD

## Evidence
- Input: artifacts/prd/<SLUG>.md
- Workflow: workflow-new-feature "<feature name>" — <TODAY>
```

Proceed to Stage 6.
