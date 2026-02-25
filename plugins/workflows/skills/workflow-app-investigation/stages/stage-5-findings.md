# Stage 5 — Findings

**Idempotency check:**
Run: !`ls artifacts/knowledge/<SLUG>/findings.md 2>/dev/null`

If file exists: output `[skip] artifacts/knowledge/<SLUG>/findings.md already exists.` and proceed to Stage 6.

Otherwise:

Read `artifacts/knowledge/<SLUG>/analysis.md`.

Synthesize key insights from analysis:
- Patterns that enable/constrain future work
- Risks that affect reliability/security
- Gaps in understanding or implementation
- Dependencies that affect scope

Assign confidence level to each finding:
- `HIGH`: validated by code review, multiple sources, or clear evidence
- `MEDIUM`: inferred from patterns or single source; needs verification
- `LOW`: hypothesis based on limited data; needs further investigation

Create `artifacts/knowledge/<SLUG>/findings.md` with this frontmatter and sections:

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

# <App Name> — Findings

## Summary
<One paragraph: the most important takeaway(s) for someone deciding what to do next.>

## Key Findings
<Numbered list with confidence level and supporting evidence:
1. [Finding title] (HIGH|MEDIUM|LOW)
   - Evidence: ...
   - Implication: ...>

## Gaps & Unknowns
<Unresolved questions:
- What remains unclear?
- What needs follow-up investigation?
- What is missing from the codebase?
- What cannot be determined without external input?>

## Recommendations (optional)
<Include only if investigation was specifically requested with recommendations.
- Recommend actions for addressing findings
- Suggest follow-up investigations
- Propose architectural changes if warranted>

## Evidence
- Input: artifacts/knowledge/<SLUG>/analysis.md
- Analysis depth: <surface|deep>
- Confidence levels assigned: HIGH / MEDIUM / LOW
- Workflow: workflow-app-investigation "<app-name>" — <TODAY>
```

Proceed to Stage 6.
