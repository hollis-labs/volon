---
name: workflow-new-agent
description: Guided workflow for defining a new Forge agent role — purpose, scope, context, interface, plan, tasks.
argument-hint: ""agent-name""
disable-model-invocation: true
---

# Workflow: New Agent

Drive an agent through the full new-agent-role definition lifecycle, producing one artifact
per stage. Idempotent: skip any stage whose output artifact already exists.
Stage protocols live in `stages/stage-N-*.md`.

---

## Arguments

- `$0` (required): agent name — quoted string, e.g. `"code-reviewer"`

---

## Step 1 — Preflight

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract and note:
- `project.name` → used in all artifact frontmatter
- `workflows.new_agent.enabled` → if `false`: output `workflow-new-agent is disabled in forge.yaml.` and stop.

If `$0` is absent: output `ERROR: agent name is required.` and stop.

Invoke `pcc-refresh` (scope=all) to ensure PCC reflects current repo state.
Read `.forge/agent-boot.md` and `.forge/boot/` role addenda for current role definitions.

Run: !`date +%Y-%m-%d`
Capture today's date as `TODAY`.

**Derive slug:**
Run: !`echo "$0" | tr '[:upper:]' '[:lower:]' | sed 's/[^a-z0-9]/-/g' | sed 's/-\+/-/g' | sed 's/^-\|-$//g'`
Capture result as `SLUG`. Artifact ID = `agent-<TODAY>-<SLUG>`.

Execute stages in sequence:
- Stage 2: `stages/stage-2-purpose.md`
- Stage 3: `stages/stage-3-scope.md`
- Stage 4: `stages/stage-4-context.md`
- Stage 5: `stages/stage-5-interface.md`
- Stage 6: `stages/stage-6-plan.md`
- Stage 7: `stages/stage-7-tasks.md`

---

## Step 8 — Finalize

Invoke `/bootstrap-update`.
Ensure `.forge/bootstrap.md` and a history copy exist before completing.
Output `DONE`.

---

## Invariants

- Never create an artifact that already exists — check first, skip if present.
- All artifact frontmatter must use the schema from `docs/03_workflow-contracts.md`.
- All shell commands use `!` prefix.
- Output `DONE` only after Step 8 completes.
