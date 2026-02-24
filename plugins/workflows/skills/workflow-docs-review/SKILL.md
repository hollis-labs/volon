---
name: workflow-docs-review
description: Guided workflow: refresh PCC then do a full PCC-grounded doc review with minimal diffs (optional PR).
argument-hint: "[focus=user|dev|feature|release|all] [scope=path]"
disable-model-invocation: true
---

# Workflow: Docs Review

Refresh PCC, apply minimal doc updates targeted by `focus`, sync PCC again,
then optionally open a PR. Idempotent: if docs are already current, output
`No doc changes required.` and stop.

---

## Arguments

- `focus=user|dev|feature|release|all` (optional, default: `all`)
- `scope=<path>` (optional): restrict doc scanning to this subdirectory

---

## Focus → target mapping

| focus value | Targets |
|---|---|
| `user` | `README.md`, `docs/` files with user-facing content |
| `dev` | `docs/` technical files, `CONTRIBUTING.md`, plugin `SKILL.md` files |
| `feature` | `artifacts/` for the feature named by `scope`; related `docs/` entries |
| `release` | `CHANGELOG.md`, release notes, version strings |
| `all` | All of the above |

---

## Step 1 — Preflight

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract and note:
- `git.pr_mode` → determines whether Step 4 runs
- `workflows.update_docs.enabled` → if `false`: output `workflow-docs-review is disabled in forge.yaml.` and stop.

Resolve `focus` argument. If not provided, default to `all`.
Resolve `scope` argument. If not provided, default to repo root.

Invoke `pcc-refresh` (scope=all) to ensure PCC is current before inspecting docs.

---

## Step 2 — Identify and apply doc updates

Using the Focus → target mapping above, identify the target file set.

If `scope` is provided: further restrict the target set to files under `<scope>`.

For each target file:
1. Read the file.
2. Read the relevant PCC section(s) to understand current ground truth.
3. Compare doc content against PCC and current repo state.
4. If the doc is already accurate: mark it `unchanged` and continue.
5. If the doc needs updating: apply **minimal diffs only**.
   - Do not restructure documents.
   - Do not rewrite sections that are already accurate.
   - Do not add sections not already present in the document.

If no target files needed changes:
output `No doc changes required.` and stop.

Track all files that were modified.

---

## Step 3 — Sync PCC

After applying doc changes, invoke `pcc-refresh` with the most specific scope
that covers the changed files:

| Changed file(s) | pcc-refresh scope |
|---|---|
| `README.md`, `forge.yaml` | `project` |
| `docs/02_pcc.md`, `docs/03_workflow-contracts.md` | `conventions` |
| `docs/` (any) | `arch` |
| `plugins/` (any) | `arch` |
| `artifacts/` (any) | `backlog` |
| Mixed / multiple | `all` |

If no doc changes were made (stopped in Step 2): skip this step.

---

## Step 4 — Optional PR

If `git.pr_mode` is `off`: output `PR skipped (git.pr_mode: off).` and stop.

If `git.pr_mode` is `optional`: ask whether to open a PR before proceeding.
If running non-interactively: skip PR and note it in output.

If `git.pr_mode` is `required` or user confirmed: invoke `pr-open` with:
- title derived from `focus` and changed file count
- body listing the changed doc files

---

## Step 5 — Output

```
Docs updated: <comma-separated list of changed files, or "none">
PCC synced: <scope used>
PR: <URL or "skipped">
```

---

## Step 6 — Finalize iteration

Invoke `/bootstrap-update` to update `.forge/bootstrap.md` from repo ground truth.

Ensure `.forge/bootstrap.md` and a history copy exist before completing.

Output `DONE`.

---

## Invariants

- Never rewrite doc sections that are already accurate — minimal diffs only.
- Never call `pcc-refresh` Step 3 if no doc changes were made in Step 2.
- Never open a PR if `git.pr_mode: off`.
- All config read from `forge.yaml` in Step 1 — no hardcoded paths or modes.
- If docs are already current, output `No doc changes required.` and stop before Step 3.
- Output `DONE` only after Step 6 (Finalize iteration) completes.
