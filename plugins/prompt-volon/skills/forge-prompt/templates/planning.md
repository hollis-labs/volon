You are operating in **{{REPO}}** in **Orchestrator Mode**.

**Intent:** {{INTENT}}

**Date:** {{DATE}}

This is a **planning session**. Your goal is to produce structured planning artifacts
(requirements, PRD, spec, plan, and/or backlog). Do NOT implement features or modify
application source code during this session.

Do not rely on prior chat context. Repo artifacts are the only truth:
- `volon.yaml`, `.volon/pcc/`, `.volon/bootstrap.md`, `artifacts/`

---

## Rules

- You are the **single writer** for all artifacts.
- {{SUBAGENTS_NOTE}}
- Do not modify application source code during this session.
- Prefer editing existing artifacts over creating new ones (check first).
- All artifact frontmatter must use Volon schema (see `docs/03_workflow-contracts.md`).

---

## Run

### 1) Preflight

1. Run: `cat volon.yaml 2>/dev/null || echo "NO_CONFIG"`
2. Read `.volon/pcc/` (global context — goals, architecture, constraints).
3. Scan `artifacts/` for existing related artifacts — skip any phase whose artifact exists.
4. Run: `git branch --show-current`

Derive slug from intent (lowercase, hyphens). Example: "sprint-based workflow" → `sprint-workflow`

### 2) Produce artifacts (phased)

Execute each phase in order. **Skip a phase if its artifact already exists.**

---

**Phase A — Requirements** (`artifacts/requirements/<slug>-{{DATE}}.md`)
```
---
id: req-{{DATE}}-<slug>
type: requirements
intent: project_doc
status: draft
project: {{REPO}}
tags: [planning, <slug>]
created_at: {{DATE}}
updated_at: {{DATE}}
---

# Requirements: <title>

## Problem statement
<what problem are we solving and for whom>

## Goals
- <goal 1>
- <goal 2>

## Non-goals
- <what is explicitly out of scope>

## User stories / Acceptance criteria
- As a <role>, I want <feature> so that <benefit>.
  - AC: <specific acceptance criterion>

## Open questions
- <unresolved questions>
```

---

**Phase B — PRD** (`artifacts/prd/<slug>-{{DATE}}.md`) — requires Phase A
```
---
id: prd-{{DATE}}-<slug>
type: prd
intent: project_doc
status: draft
project: {{REPO}}
tags: [planning, <slug>]
created_at: {{DATE}}
updated_at: {{DATE}}
---

# PRD: <title>

## Summary
<1-2 sentence product description>

## Feature list (prioritized)
| Priority | Feature | Notes |
|---|---|---|
| P0 | <must-have> | |
| P1 | <should-have> | |
| P2 | <nice-to-have> | |

## Success metrics
- <how we measure success>

## Dependencies
- <systems, people, or capabilities required>

## Risks
- <technical or product risks>
```

---

**Phase C — Spec** (`artifacts/spec/<slug>-{{DATE}}.md`) — requires Phase B
```
---
id: spec-{{DATE}}-<slug>
type: spec
intent: project_doc
status: draft
project: {{REPO}}
tags: [planning, <slug>]
created_at: {{DATE}}
updated_at: {{DATE}}
---

# Spec: <title>

## Interfaces
<describe key interfaces, APIs, or contracts>

## Data model
<schema or data structure descriptions>

## Sequence / flow
<step-by-step flow descriptions; use plain text or ASCII diagrams>

## Open questions
- <technical unresolved items>

## Decisions
- <key technical decisions made and rationale>
```

---

**Phase D — Plan** (`artifacts/plan/<slug>-{{DATE}}.md`) — requires Phase C
```
---
id: plan-{{DATE}}-<slug>
type: plan
intent: project_doc
status: draft
project: {{REPO}}
tags: [planning, <slug>]
created_at: {{DATE}}
updated_at: {{DATE}}
---

# Plan: <title>

## Task breakdown
| Task | Phase | Dependencies | Priority |
|---|---|---|---|
| <task 1> | 1 | none | A |
| <task 2> | 1 | task 1 | A |

## Phases
- Phase 1: <description, deliverables>
- Phase 2: <description, deliverables>

## Risk register
| Risk | Likelihood | Impact | Mitigation |
|---|---|---|---|
```

---

**Phase E — Backlog tasks** — requires Phase D (only if `track_tasks: true` in volon.yaml)

Create TASK files in `.volon/tasks/` for each implementation task from the plan.
Use the schema from `docs/04_task-model.md`.

### 3) Finalize

1. Run `/bootstrap-update`.
2. Commit: `iter <N>: planning — <slug>` (policy: {{COMMIT_POLICY}})

---

## Constraints

{{CONSTRAINTS}}

---

## Expected Deliverables

{{DELIVERABLES}}

If not specified above:
- Artifacts for all phases completed (requirements → PRD → spec → plan)
- Backlog tasks if plan produced (per volon.yaml `track_tasks` setting)
- `.volon/bootstrap.md` updated
- Commit

---

## Guardrails

- No application code changes during this session.
- All artifacts must use Volon frontmatter schema.
- Do not create tasks for phases not yet completed.
- **Single writer**: only this session writes artifacts/tasks/bootstrap.
- **{{SUBAGENTS_NOTE}}**
- Phase skipping is idempotent — never overwrite an existing artifact.

---

End with `{{DONE_TOKEN}}`.
