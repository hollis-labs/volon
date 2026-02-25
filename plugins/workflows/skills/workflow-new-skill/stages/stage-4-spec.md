# Stage 4 — Spec (SKILL.md Schema)

**Idempotency check:**
Run: !`ls artifacts/skills/<SLUG>/spec.md 2>/dev/null`
If file exists: output `[skip] artifacts/skills/<SLUG>/spec.md already exists.` and proceed to Stage 5.

Read `artifacts/skills/<SLUG>/requirements.md`.
Read an existing SKILL.md from the target plugin to understand frontmatter conventions.

Create `artifacts/skills/<SLUG>/spec.md` — this is the draft SKILL.md content:

```
---
id: "skill-<TODAY>-<SLUG>"
type: "spec"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["skill", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Skill Name> — Spec

## SKILL.md frontmatter (draft)
```yaml
---
name: <SLUG>
description: <one sentence from ideation>
argument-hint: "<arg hint string>"
disable-model-invocation: true|false
standalone: true|false   # omit if false
model-tier: read_scan|summarize|generate|plan|orchestrate|complex_reasoning
---
```

## Step-by-step behaviour
<Number each step. For each: what it reads, what it decides, what it writes/outputs.>

1. Preflight: read config, validate inputs.
2. ...
N. Output: list changed files and print DONE.

## Shell commands used
| Purpose | Command |
|---|---|
| ... | !`...` |

## Invariants
- <from requirements>

## Evidence
- Input: artifacts/skills/<SLUG>/requirements.md
- Workflow: workflow-new-skill "<skill name>" — <TODAY>
```

Proceed to Stage 5.
