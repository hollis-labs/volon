# Stage 3 — Discovery

**Idempotency check:**
Run: !`ls artifacts/knowledge/<SLUG>/discovery.md 2>/dev/null`

If file exists: output `[skip] artifacts/knowledge/<SLUG>/discovery.md already exists.` and proceed to Stage 4.

Otherwise:

Read `artifacts/knowledge/<SLUG>/scope.md` to understand investigation scope.

**Run discovery hook point** (advisory; not required):
Hook `on_discovery_start` may trigger external tools. If available, orchestrator may run:
- !`git log --oneline -20` — last 20 commits for context
- !`find . -name "*.config.*" -not -path "*/node_modules/*"` — discover config files
- !`ls -la` — directory listing

Enumerate all files under target paths from scope.md:
Run: !`find <target-paths> -type f | head -100` (or similar, respecting scope)

For each file, categorize:
- Type: source code, config, doc, data, build artifact, test, other
- Purpose: 1-line description based on path and common patterns
- Dependencies: any imports/references to external packages (scan headers/frontmatter)

Create `artifacts/knowledge/<SLUG>/discovery.md` with this frontmatter and sections:

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

# <App Name> — Discovery

## Summary
<One paragraph: what we found. How many files, main categories, key entry points.>

## File Inventory
<Table: Path | Type | Purpose — sorted by criticality (entry points first, then config, then supporting)>

## Dependencies Found
<External packages, libraries, services, data sources. List with version/URL if discoverable.>

## Config Files
<All configuration files discovered (.json, .yaml, .env, package.json, etc.) with location.>

## Entry Points
<Primary files that serve as starting points: main.ts, index.js, __init__.py, Dockerfile, etc.>

## Evidence
- Input: artifacts/knowledge/<SLUG>/scope.md
- Discovery hook point: on_discovery_start (if available)
- Workflow: workflow-app-investigation "<app-name>" — <TODAY>
```

Proceed to Stage 4.
