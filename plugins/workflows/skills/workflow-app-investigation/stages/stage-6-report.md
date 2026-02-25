# Stage 6 — Report

**Idempotency check:**
Run: !`ls artifacts/knowledge/<SLUG>/report.md 2>/dev/null`

If file exists: output `[skip] artifacts/knowledge/<SLUG>/report.md already exists.` and stop.

Otherwise:

Read `artifacts/knowledge/<SLUG>/scope.md` and `artifacts/knowledge/<SLUG>/findings.md`.

Create `artifacts/knowledge/<SLUG>/report.md` — the final, shareable artifact.
Note: status is `complete` (unlike earlier stage artifacts which are `draft`).

```
---
id: "ka-<TODAY>-<SLUG>"
type: "knowledge_artifact"
intent: "knowledge_artifact"
status: complete
project: "<project.name>"
tags: ["investigation", "<SLUG>"]
depth: <surface|deep>
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <App Name> — Investigation Report

## Executive Summary
<2–3 sentences: what was investigated, why, and the top 1–2 findings.>

## Scope Recap
- Target(s): <path(s) from scope.md>
- Depth: <surface|deep>
- Focus areas: <investigation questions from scope.md>

## Key Findings
<Copy HIGH-confidence findings from findings.md with confidence level and evidence.
Include MEDIUM findings if relevant to decisions.>

## Gaps & Unknowns
<Copy unresolved items from findings.md; note what needs follow-up.>

## Next Steps (only if recommendations were requested)
<Actionable recommendations:
- Investigation to pursue
- Architecture review/changes to consider
- Documentation needed>

## Evidence Trail
- Stage 2 (Scope): artifacts/knowledge/<SLUG>/scope.md
- Stage 3 (Discovery): artifacts/knowledge/<SLUG>/discovery.md
- Stage 4 (Analysis): artifacts/knowledge/<SLUG>/analysis.md
- Stage 5 (Findings): artifacts/knowledge/<SLUG>/findings.md
- Investigation depth: <surface|deep>
- Workflow: workflow-app-investigation "<app-name>" — <TODAY>
```

Proceed to Step 7 (Finalize).
