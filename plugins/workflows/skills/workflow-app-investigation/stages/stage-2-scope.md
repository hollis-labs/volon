# Stage 2 — Scope

**Idempotency check:**
Run: !`ls artifacts/knowledge/<SLUG>/scope.md 2>/dev/null`

If file exists: output `[skip] artifacts/knowledge/<SLUG>/scope.md already exists.` and proceed to Stage 3.

Otherwise:

Run: !`mkdir -p artifacts/knowledge/<SLUG>`

Read `volon.yaml` (or `.volon/volon.yaml`), `.volon/pcc/00_project.md`, and `.volon/pcc/01_architecture.md` for context.

If `scope` argument is provided: note the target path. If not provided, default to repo root.

Create `artifacts/knowledge/<SLUG>/scope.md` with this frontmatter and sections:

```
---
id: "ka-<TODAY>-<SLUG>"
type: "knowledge_artifact"
intent: "knowledge_artifact"
status: draft
project: "<project.name>"
tags: ["investigation", "<SLUG>"]
depth: <surface|deep>
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <App Name> — Investigation Scope

## Summary
<One paragraph: what we're investigating and why. Include target path(s) and investigation goal.>

## Target paths / Files
<Bulleted list of primary paths to investigate. If scope argument was provided, list it. Otherwise list major directories/services in repo.>

## Out-of-scope
<What will NOT be investigated, why, and potential follow-up investigations.>

## Investigation questions
<Numbered list of 5–10 key questions the investigation aims to answer.>

## Evidence
- Workflow: workflow-app-investigation "<app-name>" — <TODAY>
- Scope argument: <scope value or "none (repo root)">
- Config inspected: volon.yaml, pcc/00_project.md, pcc/01_architecture.md
```

Proceed to Stage 3.
