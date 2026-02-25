# Stage 4 — Spec

**Idempotency check:**
Run: !`ls artifacts/plugins/<SLUG>/spec.md 2>/dev/null`
If file exists: output `[skip] artifacts/plugins/<SLUG>/spec.md already exists.` and proceed to Stage 5.

Read `artifacts/plugins/<SLUG>/requirements.md`.
Read an existing `plugin.json` for schema reference.

Create `artifacts/plugins/<SLUG>/spec.md`:

```
---
id: "plugin-<TODAY>-<SLUG>"
type: "spec"
intent: "system_doc"
status: draft
project: "<project.name>"
tags: ["plugin", "<SLUG>"]
priority: B
created_at: "<TODAY>"
updated_at: "<TODAY>"
---

# <Plugin Name> — Spec

## plugin.json (draft)
```json
{
  "name": "<SLUG>",
  "version": "0.1.0",
  "description": "<one sentence>",
  "skills": ["<skill-1>", "<skill-2>"]
}
```

## Directory structure
```
plugins/<SLUG>/
  plugin.json
  skills/
    <skill-1>/
      SKILL.md
    <skill-2>/
      SKILL.md
```

## volon.yaml additions (proposed)
```yaml
workflows:
  <SLUG>_<skill>:
    enabled: true
```

## Skill stubs
For each skill: a minimal SKILL.md with name, description, and placeholder body.
Full skill implementation via workflow-new-skill in subsequent tasks.

## Decisions
- TBD

## Evidence
- Input: artifacts/plugins/<SLUG>/requirements.md
- Workflow: workflow-new-plugin "<plugin name>" — <TODAY>
```

Proceed to Stage 5.
