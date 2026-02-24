---
name: pcc-refresh
description: Refresh Project Context Cache (PCC) based on repo changes (minimal updates).
argument-hint: "[scope=all|project|arch|conventions|workflows|backlog|decisions]"
disable-model-invocation: true
model-tier: read_scan
---

# PCC Refresh

Update **L0 Global PCC** (`.forge/pcc/global/*`) with **minimal diffs** based on repo reality.
Touch only affected sections. Never invent content.

> **Note:** This skill refreshes L0 (Global) PCC only.
> L2 Task PCC capsules (`.forge/pcc/tasks/`) are managed by `/pause-task` and `/resume-task`.

---

## Inputs

- `scope`: `$0` (default: `all`)

| Scope value | PCC file targeted |
|---|---|
| `all` | All 6 required files (evaluated by git change set) |
| `project` | `00_project.md` only |
| `arch` | `01_architecture.md` only |
| `conventions` | `02_conventions.md` only |
| `workflows` | `03_workflows.md` only |
| `backlog` | `04_backlog.md` only |
| `decisions` | `05_decisions.md` only |

---

## Step 1 — Read config

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

If `NO_CONFIG`: use these defaults and continue.

Extract and note:
- `pcc.location` → default: `.forge/pcc`
- `pcc.global_dir` → default: `global`
- Derive `global_path` = `<pcc.location>/<pcc.global_dir>` → default: `.forge/pcc/global`
- `pcc.limits.max_section_words` → default: `400`
- `observability.write_run_log` → default: `false`
- `observability.log_dir` → default: `.forge/logs`

---

## Step 2 — Collect git signals

Run all four commands; capture full output:

- !`git status --porcelain=v1 2>/dev/null || true`
- !`git diff --name-only 2>/dev/null || true`
- !`git diff --stat HEAD 2>/dev/null || true`
- !`git log --oneline -10 2>/dev/null || true`

Collect all unique file paths appearing in any output.
If no git repo or all commands return empty: treat change set as empty.

If all four commands returned empty output:
→ Emit: `WARN: no git repository detected — git signals unavailable. PCC refresh running without change-set data. Use explicit scope to force a section update.`
→ Continue to Step 3 (do not stop).

---

## Step 3 — Determine affected PCC sections

### 3a — If scope is specific (not `all`)
Skip git-based mapping. The affected set is the single file named by the scope value.
Proceed to Step 4.

### 3b — If scope is `all`
For each path in the change set, apply this mapping (a path may match multiple rows):

All PCC files in this table live under `<global_path>/` (default: `.forge/pcc/global/`).

| Path pattern | Affects PCC file(s) |
|---|---|
| `README.md` | `00_project.md` |
| `forge.yaml` | `00_project.md`, `01_architecture.md` |
| `.forge/forge.yaml` | `00_project.md`, `01_architecture.md` |
| `docs/01_config.md` | `00_project.md` |
| `docs/02_pcc.md` | `01_architecture.md`, `02_conventions.md` |
| `docs/pcc_layers.md` | `01_architecture.md`, `02_conventions.md` |
| `docs/03_workflow-contracts.md` | `02_conventions.md`, `03_workflows.md` |
| `docs/04_task-model.md` | `02_conventions.md` |
| `plugins/core/**` | `01_architecture.md` |
| `plugins/workflows/**` | `01_architecture.md`, `03_workflows.md` |
| `plugins/git/**` | `01_architecture.md`, `03_workflows.md` |
| `plugins/tasks-nanite/**` | `01_architecture.md`, `03_workflows.md` |
| `plugins/docsmith/**` | `01_architecture.md` |
| `plugins/quality/**` | `01_architecture.md` |
| `artifacts/**` | `04_backlog.md` |
| `.forge/tasks/**` | `04_backlog.md` |
| `CHANGELOG.md` | `05_decisions.md` |
| *(unmatched — any other path)* | `04_backlog.md` |

Build the **affected set**: union of all matched PCC files.

If the change set is empty and no scope was forced:
→ Output `No PCC changes required.` and stop.

---

## Step 4 — Read affected PCC files

For each PCC file in the affected set:
- Read its current contents from `<global_path>/<filename>`
- Note its current word count

Do **not** read PCC files outside the affected set.

---

## Step 5 — Update affected PCC files

Apply the update rules below for each file in the affected set.
**Minimal diff rule**: change only the content that is stale. Do not restructure
headings, reorder sections, or rewrite text that is still accurate.

### 00_project.md
Update when: `README.md`, `forge.yaml`, or `docs/01_config.md` changed.
- Refresh Active Config block if `forge.yaml` changed.
- Refresh Goals/Non-goals if `README.md` changed.
- Do not modify Identity block unless project name or version changed.

### 01_architecture.md
Update when: any `plugins/**`, `forge.yaml`, or `docs/02_pcc.md` changed.
- Update only the row(s) in the plugin/skill table that correspond to changed plugins.
- Update storage or data flow sections only if `forge.yaml` changed.
- Do not rewrite rows for unchanged plugins.

### 02_conventions.md
Update when: `docs/03_workflow-contracts.md`, `docs/04_task-model.md`, or `docs/02_pcc.md` changed.
- Rewrite only the specific convention block whose source changed.
- Preserve all other convention blocks verbatim.

### 03_workflows.md
Update when: any `plugins/workflows/**`, `plugins/git/**`, `plugins/tasks-nanite/**`, or `docs/03_workflow-contracts.md` changed.
- Update only the workflow or skill row that corresponds to the changed file.
- Do not rewrite steps for workflows whose source files did not change.

### 04_backlog.md
Update when: `artifacts/**`, `.forge/tasks/**`, or any unmatched changed path.
- Move items to Completed if their artifact or task file now exists and is complete.
- Add new items if new tasks or artifacts were created.
- Do not rewrite existing pending items unless their status changed.

### 05_decisions.md — APPEND ONLY
Update when: `CHANGELOG.md` changed, or scope explicitly set to `decisions`.
- Never rewrite or delete any existing ADR entry.
- Only append a new ADR entry at the bottom if the change represents a
  decision not already recorded.
- If no new decision is detected, skip this file entirely (do not touch it).
- ADR format:
  ```
  ## ADR-NNN — <title>
  - **Date:** YYYY-MM-DD
  - **Status:** Accepted
  - **Context:** <one sentence>
  - **Decision:** <one sentence>
  - **Consequence:** <one sentence>
  ```

---

## Step 6 — Enforce word limits

After updating each file, count total words.

If count exceeds `max_section_words` (default 400):
- For `05_decisions.md`: do **not** trim. Output warning:
  `WARN: 05_decisions.md exceeds word limit — manual trim required.`
- For all other files:
  - Remove the least informative bullet points from body sections only.
  - Never remove headings, table rows, or the Evidence section.
  - Never compress below structural minimum (all headings + Evidence).
  - Recount after trimming. Repeat if still over.

---

## Step 7 — Update Evidence section in each modified file

For each PCC file that was modified, update its `## Evidence` section.

Replace the "Last refreshed" line (or add if absent). Append changed paths line.
Preserve all prior evidence lines below the replaced/new lines.

Evidence block format:
```
## Evidence
- Last refreshed: YYYY-MM-DD (pcc-refresh, scope=<scope>)
- Trigger: <comma-separated list of changed source paths>
- Git signals: status, diff --name-only, diff --stat HEAD, log --oneline -10
- <all prior evidence lines preserved below here>
```

---

## Step 8 — Write observability log

If `observability.write_run_log` is `true`:
- Create `<log_dir>/` if it does not exist.
- Append to `<log_dir>/run.log`:
  ```
  [YYYY-MM-DD] pcc-refresh scope=<scope>
    trigger_paths: <changed paths, comma-separated, or "none">
    updated: <updated PCC filenames, comma-separated, or "none">
    unchanged: <skipped PCC filenames, comma-separated>
  ```

---

## Step 9 — Output

If one or more PCC files were updated:
```
PCC refresh complete.
Updated: <comma-separated PCC filenames>
Unchanged: <comma-separated PCC filenames>
```

If no PCC files were updated:
```
No PCC changes required.
```

---

## Invariants (never violate)

- Never read or write a PCC file not in the affected set for this run.
- Never invent content — only reflect verified, observable repo state.
- Never delete or overwrite `05_decisions.md` entries — only append.
- Never rewrite sections whose source files did not change.
- Mark all unknowns as **TBD** rather than guessing.
- If a PCC file does not exist yet, create it from scratch using the section
  rules above rather than failing.
