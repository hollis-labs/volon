---
name: workflow-new-feature
description: Guided workflow: Ideation → Requirements → PRD → Spec → Plan → Tasks.
argument-hint: ""feature name" [scope=path] [audience=user|dev|both]"
disable-model-invocation: true
model-tier: generate
---

# Workflow: New Feature

Drive an agent through the full new-feature lifecycle, producing one artifact
per stage and ≥1 task at the end. Idempotent: skip any stage whose output
artifact already exists. Stage protocols live in `stages/stage-N-*.md`.

---

## Arguments

- `$0` (required): feature name — quoted string, e.g. `"worktree-start"`
- `scope=<path>` (optional): restrict ideation/spec file scanning to this path
- `audience=user|dev|both` (optional, default: `both`): frames requirements and PRD

---

## Step 1 — Preflight

Run: !`cat volon.yaml 2>/dev/null || cat .volon/volon.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract and note:
- `project.name` → used in all artifact frontmatter
- `storage.files.root` → task backend path
- `workflows.new_feature.enabled` → if `false`: output `workflow-new-feature is disabled in volon.yaml.` and stop.
- `workflows.new_feature.defaults.require_spec` → if `true`, Stage 5 is mandatory
- `workflows.new_feature.defaults.track_tasks` → if `true`, Stage 7 is mandatory

If `$0` is absent: output `ERROR: feature name is required.` and stop.

Invoke `pcc-refresh` (scope=all) to ensure PCC reflects current repo state.
Read the updated `.volon/pcc/` after refresh.

Run: !`date +%Y-%m-%d`

Capture today's date as `TODAY`.

**Derive slug:**
Run: !`echo "$0" | tr '[:upper:]' '[:lower:]' | sed 's/[^a-z0-9]/-/g' | sed 's/-\+/-/g' | sed 's/^-\|-$//g'`

Capture result as `SLUG`. Artifact ID = `feat-<TODAY>-<SLUG>`.

Execute stages in sequence — read each file and follow its instructions:
- Stage 2: `stages/stage-2-ideation.md`
- Stage 3: `stages/stage-3-requirements.md`
- Stage 4: `stages/stage-4-prd.md`
- Stage 5: `stages/stage-5-spec.md`
- Stage 6: `stages/stage-6-plan.md`
- Stage 7: `stages/stage-7-tasks.md`

---

## Step 8 — Finalize

Invoke `/bootstrap-update` to update `.volon/bootstrap.md` from repo ground truth.

Ensure `.volon/bootstrap.md` and a history copy exist before completing.

Output `DONE`.

---

## Invariants

- Never create an artifact that already exists — check first, skip if present.
- Never create tasks if `track_tasks: false`.
- All artifact frontmatter must use the schema from `docs/03_workflow-contracts.md`.
- All shell commands use `!` prefix.
- `audience` argument affects content framing only — never affects file paths or IDs.
- `scope` argument affects which source files are read — never affects output paths.
- Output `DONE` only after Step 8 (Finalize) completes.
