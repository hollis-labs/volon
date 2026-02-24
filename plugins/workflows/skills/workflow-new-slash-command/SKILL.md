---
name: workflow-new-slash-command
description: Guided workflow for defining a new user slash command — ideation, requirements, spec, plan, tasks.
argument-hint: ""command-name""
disable-model-invocation: true
---

# Workflow: New Slash Command

Drive an agent through the full new-slash-command definition lifecycle. Produces a
spec artifact and implementation tasks. Updates `docs/09_commands.md` with the new command.
Idempotent: skip any stage whose output artifact already exists.
Stage protocols live in `stages/stage-N-*.md`.

---

## Arguments

- `$0` (required): command name without leading slash — e.g. `"task-search"` (will be invoked as `/task-search`)

---

## Step 1 — Preflight

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract and note:
- `project.name` → used in all artifact frontmatter
- `workflows.new_slash_command.enabled` → if `false`: output `workflow-new-slash-command is disabled in forge.yaml.` and stop.

If `$0` is absent: output `ERROR: command name is required.` and stop.

Invoke `pcc-refresh` (scope=all).
Read `docs/09_commands.md` to check for name collision.
If command name already exists in `docs/09_commands.md`: output `ERROR: /<command> is already documented.` and stop.

Run: !`date +%Y-%m-%d`
Capture today's date as `TODAY`.

**Derive slug:**
Run: !`echo "$0" | tr '[:upper:]' '[:lower:]' | sed 's/[^a-z0-9]/-/g' | sed 's/-\+/-/g' | sed 's/^-\|-$//g'`
Capture result as `SLUG`. Artifact ID = `cmd-<TODAY>-<SLUG>`.

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

- Never create an artifact that already exists — check first, skip if present.
- Never add a command to `docs/09_commands.md` that already exists there.
- All artifact frontmatter must use the schema from `docs/03_workflow-contracts.md`.
- All shell commands use `!` prefix.
- Output `DONE` only after Step 7 completes.
