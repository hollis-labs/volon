You are operating in **{{REPO}}** in **Orchestrator Mode**.

**Intent:** {{INTENT}}

**Date:** {{DATE}}

This is an **ideation session**. Your goal is to produce a structured ideation artifact.
Do NOT write application code or modify non-artifact source files during this session.

Do not rely on prior chat context. Repo artifacts are the only truth:
- `forge.yaml`, `.forge/pcc/`, `.forge/bootstrap.md`, `artifacts/`

---

## Rules

- You are the **single writer** for the output artifact.
- {{SUBAGENTS_NOTE}}
- Do not modify application source code during this session.
- Do not create TASK files unless explicitly instructed.
- Keep ideas grounded in the repo's actual context (read PCC before generating).

---

## Run

### 1) Preflight

1. Run: `cat forge.yaml 2>/dev/null || echo "NO_CONFIG"`
2. Read `.forge/pcc/` (scan `global/` for project goals and constraints).
3. Scan `artifacts/ideas/` for existing related ideation (skip duplicates).
4. Run: `git branch --show-current`

Note any relevant constraints from PCC before generating ideas.

### 2) Derive slug

From the intent, derive a short slug (lowercase, hyphens, no special chars).
Example: "brainstorm names with vibe constraints" → `names-vibe-constraints`

Artifact path: `artifacts/ideas/<slug>-{{DATE}}.md`

**Idempotency check:** If `artifacts/ideas/<slug>-{{DATE}}.md` already exists, output
`[skip] artifact already exists.` and proceed to Step 4 (Finalize).

### 3) Generate ideas

Generate ideas aligned with the intent: **{{INTENT}}**

For each idea, include:
- **Name / title** — concise label
- **One-sentence description** — what it is
- **Rationale** — why it fits the project/intent
- **Risks or trade-offs** — any downsides (be honest)

Aim for 5–15 concrete ideas before filtering. Group by theme if natural groupings emerge.
Rank ideas within each group (strongest first).

Write the artifact to `artifacts/ideas/<slug>-{{DATE}}.md`:

```
---
id: idea-{{DATE}}-<slug>
type: idea
intent: project_doc
status: draft
project: {{REPO}}
tags: [ideation, <slug>]
created_at: {{DATE}}
updated_at: {{DATE}}
---

# Ideation: <short title>

## Summary
<1-2 sentence overview of the ideation goal and approach>

## Ideas

### <Theme A> (if applicable)

#### 1. <Name>
**Description:** ...
**Rationale:** ...
**Risks:** ...

#### 2. <Name>
...

## Open Questions
- <unresolved questions that affect the ideas>

## Next Steps
- <what to do with these ideas — e.g., "review with team", "run /forge-prompt planning for top candidate">
```

### 4) Finalize

1. Run `/bootstrap-update`.
2. Commit: `iter <N>: ideation — <slug>` (policy: {{COMMIT_POLICY}})

---

## Constraints

{{CONSTRAINTS}}

---

## Expected Deliverables

{{DELIVERABLES}}

If not specified above:
- `artifacts/ideas/<slug>-{{DATE}}.md` with 5–15 ideas
- `.forge/bootstrap.md` updated
- Commit

---

## Guardrails

- No application code changes during this session.
- No TASK file creation unless explicitly asked.
- Artifact output scope: `artifacts/ideas/` only.
- **Single writer**: only this session writes artifacts/bootstrap.
- **{{SUBAGENTS_NOTE}}**
- Ideas must be grounded in repo context (PCC) — not invented without basis.

---

End with `{{DONE_TOKEN}}`.
