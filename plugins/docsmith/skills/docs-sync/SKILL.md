---
name: docs-sync
description: Sync repo docs with minimal diffs driven by git change signals and focus scope; sync PCC; optional PR.
argument-hint: "[focus=user|dev|feature|release|all] [slug=<feature-slug>] [pr=auto]"
disable-model-invocation: true
---

# Docs Sync

Apply **minimal, accurate** doc updates based on recent repo changes and the
`focus` argument. Never invent content. Mark unknowns as **TBD**.
Calls `pcc-refresh` after updates to keep PCC in sync.

---

## Inputs

- `focus`: `$0` (default: `all`)
- `slug`: `$1` — feature slug, required when `focus=feature` (default: none)
- `pr`: `$2` — `auto` triggers PR creation if `pr_mode` is not `off` (default: none)

### Focus → Target mapping

| Focus | Target files |
|---|---|
| `user` | `docs/` public-facing files, `README.md` |
| `dev` | `docs/` technical files (config, PCC, workflow, task model), plugin SKILL.md description lines |
| `feature` | `artifacts/<type>/<slug>.md` files matching `slug`, `README.md` feature section if present |
| `release` | `CHANGELOG.md`, `README.md` version/status block |
| `all` | Union of all rows above |

---

## Step 1 — Read config

Run: !`cat volon.yaml 2>/dev/null || cat .volon/volon.yaml 2>/dev/null || echo "NO_CONFIG"`

If `NO_CONFIG`: use defaults and continue.

Extract and note:
- `pcc.location` → default: `.volon/pcc`
- `observability.write_run_log` → default: `false`
- `observability.log_dir` → default: `.volon/logs`
- `git.pr_mode` → default: `optional`
- `git.pr.base_branch` → default: `auto`
- `git.pr.title_prefix` → default: `volon:`
- `git.pr.body_template` → default: `.volon/templates/pr-body.md`

Resolve focus from `$0` arg; default `all` if not provided.
If `focus=feature` and no `slug` arg: output `ERROR: focus=feature requires a slug argument.` and stop.

---

## Step 2 — Collect git signals

Run all four commands; capture full output:

- !`git status --porcelain=v1 2>/dev/null || true`
- !`git diff --name-only 2>/dev/null || true`
- !`git diff --stat HEAD 2>/dev/null || true`
- !`git log --oneline -10 2>/dev/null || true`

Collect all unique changed file paths.

If no git repo or all commands return empty:
→ Emit: `WARN: no git repository detected — git signals unavailable. Proceeding with focus-based scan only.`
→ Continue (do not stop).

---

## Step 3 — Determine target files

Apply the focus→target mapping table to build the **target set**.

For each target category implied by the focus value:

### user
- !`ls docs/ 2>/dev/null || true` → include all `.md` files in `docs/` flagged as public-facing (heuristic: filename does not start with a digit or `_`)
- Include `README.md` if it exists

### dev
- !`ls docs/ 2>/dev/null || true` → include all `docs/*.md`
- !`find plugins/ -name "SKILL.md" 2>/dev/null || true` → include each SKILL.md (description-line updates only — do not rewrite full protocol)

### feature
- !`find artifacts/ -name "*<slug>*" 2>/dev/null || true` → include all artifact files matching the slug
- If `README.md` contains a section referencing the slug: include `README.md`

### release
- Include `CHANGELOG.md` if it exists: !`ls CHANGELOG.md 2>/dev/null || true`
- Include `README.md`

### all
Union of all above.

Intersect with changed-file paths from Step 2 when git signals are available.
If git signals unavailable: use the full focus-based target set.

If target set is empty after intersection:
→ Output `No doc updates required — no target files changed.` and proceed to Step 5.

---

## Step 4 — Read and update each target file

For each file in the target set:

1. Read current contents.
2. Identify which sections are stale relative to the changes signaled in Step 2.
   - Cross-reference changed file paths against section content.
   - Only update sections that reference changed components, files, or versions.
3. Apply **minimal diff** updates:
   - Do not reorder sections or restructure headings.
   - Do not rewrite text that is still accurate.
   - Do not add new sections unless a clear gap exists in the document.
   - Mark any unknown or unverifiable facts as **TBD**.
4. For SKILL.md files (focus=dev): update only the `description:` frontmatter line
   if it is stale. Do not modify the protocol steps.
5. Write the updated file.

---

## Step 5 — Sync PCC and output

Run pcc-refresh to update `.volon/pcc/` after doc changes:
- !`cat plugins/core/skills/pcc-refresh/SKILL.md 2>/dev/null | head -5 || true`
  (confirms pcc-refresh is available)

Apply pcc-refresh protocol with `scope=backlog` to capture doc-update activity.

If `pr` arg is `auto` and `git.pr_mode` is not `off`:
→ Apply pr-open protocol (see `plugins/git/skills/pr-open/SKILL.md`).

If `observability.write_run_log` is true: write entry to `<log_dir>/run-<YYYYMMDD>-<HHMM>-docs-sync.md`:
- datetime, focus, target files updated, lines changed (estimate), pcc-refresh result

Output:
```
docs-sync complete
Focus: <focus>
Files updated: <N> (<list>)
PCC sync: done (scope=backlog)
PR: <created|skipped>
```

---

## Invariants

- **Never invent content.** If a fact cannot be verified from the repo, write **TBD**.
- **Minimal diff only.** Do not restructure, reformat, or rewrite accurate text.
- **Focus boundary is strict.** Do not update files outside the resolved target set.
- **SKILL.md protocols are read-only** when focus=dev. Only `description:` frontmatter may change.
- **Append-only for decision logs.** Do not modify `05_decisions.md` — that is pcc-refresh's domain.
- pcc-refresh runs after doc updates, not before (this skill is called from workflow-docs-review which runs pcc-refresh as its preflight).
- If git signals are unavailable: proceed with focus-based scan; always emit the WARN.
- `slug` is required for `focus=feature`. Fail fast with a clear error message if absent.
