---
name: workflow-app-investigation
description: Guided investigation workflow — scoping → discovery → analysis → findings → report. Produces Knowledge Artifacts without development intent.
argument-hint: '"app-name" [scope=path] [depth=surface|deep]'
disable-model-invocation: true
model-tier: generate
---

# Workflow: App Investigation

Systematically investigate an application or codebase component, producing Knowledge
Artifacts without assuming development intent. Idempotent: skip any stage whose output
artifact already exists. Stage protocols live in `stages/stage-N-*.md`.

---

## Arguments

- `$0` (required): app/component name — quoted string, e.g. `"auth-service"`, `"frontend-build"`
- `scope=<path>` (optional): restrict investigation to this path; if omitted, investigates repo root
- `depth=surface|deep` (optional, default: `deep`): depth of analysis
  - `surface`: read entry points and config only
  - `deep`: read all discovered files comprehensively

---

## Step 1 — Preflight

Run: !`cat volon.yaml 2>/dev/null || cat .volon/volon.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract and note:
- `project.name` → used in all artifact frontmatter
- `workflows.app_investigation.enabled` → if `false`: output `workflow-app-investigation is disabled in volon.yaml.` and stop.
- `workflows.app_investigation.defaults.depth` → default depth if `depth` argument not provided (fallback: `deep`)

If `$0` is absent: output `ERROR: app/component name is required.` and stop.

Invoke `pcc-refresh` (scope=all) to ensure PCC reflects current repo state.
Read the updated `.volon/pcc/` after refresh.

Run: !`date +%Y-%m-%d`

Capture today's date as `TODAY`.

**Derive slug:**
Run: !`echo "$0" | tr '[:upper:]' '[:lower:]' | sed 's/[^a-z0-9]/-/g' | sed 's/-\+/-/g' | sed 's/^-\|-$//g'`

Capture result as `SLUG`. Artifact ID = `ka-<TODAY>-<SLUG>`.

Execute stages in sequence — read each file and follow its instructions:
- Stage 2: `stages/stage-2-scope.md`
- Stage 3: `stages/stage-3-discovery.md`
- Stage 4: `stages/stage-4-analysis.md`
- Stage 5: `stages/stage-5-findings.md`
- Stage 6: `stages/stage-6-report.md`

---

## Step 7 — Finalize

Invoke `/bootstrap-update` to update `.volon/bootstrap.md` from repo ground truth.

Ensure `.volon/bootstrap.md` and a history copy exist before completing.

Output `DONE`.

---

## Invariants

- Never create an artifact that already exists — check first, skip if present.
- All artifact frontmatter must use the Knowledge Artifact schema from `docs/03_workflow-contracts.md`.
- All shell commands use `!` prefix.
- `scope` argument affects which source files are scanned — never affects output paths.
- `depth` argument affects analysis breadth — never affects artifact structure.
- This workflow does NOT create tasks; it produces Knowledge Artifacts only.
- Output `DONE` only after Step 7 (Finalize) completes.
