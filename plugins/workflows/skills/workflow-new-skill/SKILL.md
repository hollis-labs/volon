---
name: workflow-new-skill
description: Guided workflow for creating a new Forge skill — ideation, requirements, spec, plan, tasks.
argument-hint: ""skill-name" [plugin=<plugin-dir>] [standalone=true|false]"
disable-model-invocation: true
---

# Workflow: New Skill

Drive an agent through the full new-skill definition lifecycle, producing one artifact
per stage and creating the implementation scaffold at the end. Idempotent: skip any
stage whose output artifact already exists.
Stage protocols live in `stages/stage-N-*.md`.

---

## Arguments

- `$0` (required): skill name — quoted string, e.g. `"task-search"`
- `plugin=<dir>` (optional): target plugin directory (e.g. `plugin=core`). If omitted, prompt at preflight.
- `standalone=true|false` (optional, default `false`): marks skill as standalone (exempt from dead_code scan)

---

## Step 1 — Preflight

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract and note:
- `project.name` → used in all artifact frontmatter
- `workflows.new_skill.enabled` → if `false`: output `workflow-new-skill is disabled in forge.yaml.` and stop.

If `$0` is absent: output `ERROR: skill name is required.` and stop.

If `plugin` argument is absent: list `plugins/` subdirectories and ask which plugin this skill belongs to.

Invoke `pcc-refresh` (scope=all).
Read `.forge/pcc/01_architecture.md` to understand existing plugin structure.

Run: !`date +%Y-%m-%d`
Capture today's date as `TODAY`.

**Derive slug:**
Run: !`echo "$0" | tr '[:upper:]' '[:lower:]' | sed 's/[^a-z0-9]/-/g' | sed 's/-\+/-/g' | sed 's/^-\|-$//g'`
Capture result as `SLUG`. Artifact ID = `skill-<TODAY>-<SLUG>`.

Execute stages in sequence:
- Stage 2: `stages/stage-2-ideation.md`
- Stage 3: `stages/stage-3-requirements.md`
- Stage 4: `stages/stage-4-spec.md`
- Stage 5: `stages/stage-5-plan.md`
- Stage 6: `stages/stage-6-tasks.md`

---

## Step 7 — Finalize

Invoke `/bootstrap-update`.
Ensure `.forge/bootstrap.md` and a history copy exist before completing.
Output `DONE`.

---

## Invariants

- Never create an artifact or skill file that already exists — check first, skip if present.
- All artifact frontmatter must use the schema from `docs/03_workflow-contracts.md`.
- All shell commands use `!` prefix.
- Output `DONE` only after Step 7 completes.
