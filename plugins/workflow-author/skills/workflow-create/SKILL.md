---
name: workflow-create
description: Scaffold a new user workflow definition in workflows/user/ from a guided template.
argument-hint: "\"workflow-name\" [domain=user] [description=\"...\"] [tags=tag1,tag2]"
disable-model-invocation: true
---

# Workflow Create

Generate a new workflow definition file following the format in
`docs/08_workflow-authoring.md`. The file is created with `status: draft`
and a filled-in template ready for the author to complete.

---

## Arguments

- `name`: `$0` (required) — kebab-case workflow name, e.g. `"app-investigation"`
- `domain`: `$1` — `user` (default) or `volon`
- `description`: `$2` — one-sentence description (default: TBD)
- `tags`: `$3` — comma-separated tags (default: none)

---

## Step 1 — Read config

Run: !`cat volon.yaml 2>/dev/null || cat .volon/volon.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract and note:
- `project.name` → for `project` field in frontmatter
- Any workflow-authoring config (none yet; reserved for future)

Resolve and validate `name` from `$0`:
- Must be non-empty. If empty: output `ERROR: workflow name is required.` and stop.
- Normalize to kebab-case (lowercase, spaces → hyphens).
- Resolve `domain` from `$1`; default `user`.
- Valid domains: `user`, `volon`. If invalid: output `ERROR: domain must be user or volon.` and stop.

---

## Step 2 — Check for conflicts

Determine target path: `workflows/<domain>/<name>.md`

Run: !`ls workflows/<domain>/<name>.md 2>/dev/null || echo "NOT_FOUND"`

If file exists (not `NOT_FOUND`):
→ Read its frontmatter `status` field.
→ If `status: active` or `status: draft`:
  output `ERROR: workflow '<name>' already exists in workflows/<domain>/ (status: <status>). Use /workflow-edit to modify it or choose a different name.` and stop.
→ If `status: deprecated`:
  output `WARN: A deprecated workflow named '<name>' exists. Proceeding will create a new draft alongside it.`
  Continue (do not stop).

Ensure `workflows/<domain>/` directory exists:
Run: !`ls workflows/<domain>/ 2>/dev/null || echo "MISSING"`
If `MISSING`: output `ERROR: workflows/<domain>/ directory not found. Ensure volon layout is initialized.` and stop.

---

## Step 3 — Collect workflow intent

Resolve the following from args or set to defaults:
- `description` from `$2`; if not provided, set to `"TBD — fill in before activating."`
- `tags` from `$3`; parse comma-separated into YAML list; if not provided, set to `[]`
- `invocation`: default to `"/<name> [args]"` — author fills in actual args
- Set `intent`: `user_workflow` if `domain=user`, `forge_workflow` if `domain=volon`
- Set today's date for `created_at` and `updated_at` (format: YYYY-MM-DD)

---

## Step 4 — Write workflow definition

Write `workflows/<domain>/<name>.md` with the following content:

```
---
name: <name>
version: "0.1"
domain: <domain>
intent: <intent>
status: draft
tags: <tags-as-yaml-list>
created_at: <today>
updated_at: <today>
description: "<description>"
invocation: "<invocation>"
plugin: null
replaces: null
---

# Workflow: <name>

<description>

---

## Arguments

| Arg | Type | Default | Description |
|---|---|---|---|
| (fill in) | string | — | ... |

---

## Steps

1. **Step 1** — (describe what happens; what artifact or output it produces)
2. **Step 2** — ...

> Tip: Each step should produce a verifiable output or side effect.
> Use `!` prefix for shell commands the model should run.

---

## Artifacts Produced

| Artifact | Class | Location |
|---|---|---|
| (fill in) | system_doc \| project_doc \| knowledge_artifact | ... |

---

## Invariants

- (list non-negotiable behaviors — e.g. "never overwrite existing artifacts")
- Never invent content — use **TBD** for unknowns
- Idempotent: re-running a completed step must not duplicate artifacts

---

## Evidence

- Created: <today> via /workflow-create
- Source: (describe what prompted this workflow)
- Status: draft — fill in Steps and Invariants, then set status: active
```

---

## Step 5 — Output

```
Created: workflows/<domain>/<name>.md
Status:  draft
Domain:  <domain>

Next steps:
  1. Open workflows/<domain>/<name>.md and fill in Steps + Invariants
  2. Dry-run the steps manually in a scratch session
  3. Set status: active when satisfied
  4. To promote to a plugin skill: see docs/08_workflow-authoring.md#promotion-to-plugin
```

---

## Invariants

- Always create with `status: draft` — never write `status: active` directly.
- Never overwrite an existing `draft` or `active` workflow without user confirmation.
- Always use kebab-case for the workflow name.
- The template body sections (Arguments, Steps, Artifacts, Invariants, Evidence) must always be present, even if empty.
- `domain: volon` is allowed but unusual — prefer `domain: user` for new workflows.
