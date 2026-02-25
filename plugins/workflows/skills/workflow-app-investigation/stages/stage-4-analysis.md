# Stage 4 — Analysis

**Idempotency check:**
Run: !`ls artifacts/knowledge/<SLUG>/analysis.md 2>/dev/null`

If file exists: output `[skip] artifacts/knowledge/<SLUG>/analysis.md already exists.` and proceed to Stage 5.

Otherwise:

Read `artifacts/knowledge/<SLUG>/scope.md` and `artifacts/knowledge/<SLUG>/discovery.md`.

**Depth-based analysis scope:**
- If `depth=surface`: read only entry points, main config files, and README/primary docs
- If `depth=deep` (default): read all files listed in discovery.md File Inventory; compile comprehensive understanding

Read each relevant source file according to depth setting. Extract:
- Code structure: classes, functions, APIs, key patterns
- Data flows: where data enters/exits, transformations, storage
- Dependencies: how components interact, call chains
- Anomalies: unusual patterns, technical debt markers, error handling gaps
- Risk indicators: security boundaries, input validation, resource management

**Run analysis hook point** (advisory; not required):
Hook `on_analysis_complete` may trigger:
- AST scan for complexity metrics
- Dependency graph generation
- Config validation against known schemas

Create `artifacts/knowledge/<SLUG>/analysis.md` with this frontmatter and sections:

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

# <App Name> — Analysis

## Summary
<One paragraph: what the system does, its primary responsibilities, key technologies.>

## Component Breakdown
<Each major module/service:
- Name, Responsibility, Key files, Public interface, Internal patterns>

## Data Flows (if applicable)
<Inputs (user, API, files, events) → Processing/transformation → Outputs (responses, files, events, storage)>

## Key Patterns Observed
<Recurring design patterns: DI, factory, observer, middleware, etc. Note consistency/inconsistency.>

## Anomalies / Risks
<Issues found: missing error handling, unvalidated inputs, resource leaks, circular deps, dead code, unclear security boundaries>

## Open Questions
<Unresolved: Why was this designed this way? What is the deployment model? Who owns this? What is the SLA?>

## Evidence
- Input: artifacts/knowledge/<SLUG>/scope.md, artifacts/knowledge/<SLUG>/discovery.md
- Depth setting: <surface|deep>
- Files analyzed: [count]
- Analysis hook point: on_analysis_complete (if available)
- Workflow: workflow-app-investigation "<app-name>" — <TODAY>
```

Proceed to Stage 5.
