---
name: workflow-clone
description: Copy an existing workflow to a new name, resetting to draft and recording the source in replaces.
argument-hint: "<source-name> <new-name> [domain=user] [target-domain=user]"
disable-model-invocation: true
---

# Workflow Clone

Copy an existing workflow definition to a new name. The clone starts as `draft`,
records the source in `replaces`, and preserves all body sections so the author
can modify incrementally rather than starting from scratch.

Useful for: adapting a volon workflow for user context, creating a revised version
of an existing user workflow, or experimenting with variations.

---

## Arguments

- `source`: `$0` (required) — name of the workflow to clone
- `new-name`: `$1` (required) — name for the cloned workflow (kebab-case)
- `domain`: `$2` — domain of the source workflow (default: `user`)
- `target-domain`: `$3` — domain of the clone (default: same as `domain`)

---

## Step 1 — Read config and locate source

Run: !`cat volon.yaml 2>/dev/null || cat .volon/volon.yaml 2>/dev/null || echo "NO_CONFIG"`

Resolve `source` from `$0`; `new-name` from `$1`.
If either is missing: output `ERROR: both source and new-name are required.` and stop.

Normalize `new-name` to kebab-case.
Resolve `domain` (default: `user`) and `target-domain` (default: same as `domain`).

Source path: `workflows/<domain>/<source>.md`
Run: !`cat workflows/<domain>/<source>.md 2>/dev/null || echo "NOT_FOUND"`

If `NOT_FOUND`: output `ERROR: source workflow '<source>' not found in workflows/<domain>/.` and stop.

---

## Step 2 — Check target for conflicts

Target path: `workflows/<target-domain>/<new-name>.md`

Run: !`ls workflows/<target-domain>/<new-name>.md 2>/dev/null || echo "NOT_FOUND"`

If file exists and `status` is `draft` or `active`:
output `ERROR: workflow '<new-name>' already exists in workflows/<target-domain>/ (status: <status>). Choose a different name.` and stop.

Ensure target domain directory exists:
Run: !`ls workflows/<target-domain>/ 2>/dev/null || echo "MISSING"`
If `MISSING`: output `ERROR: workflows/<target-domain>/ directory not found.` and stop.

---

## Step 3 — Build clone

Read the full source file content. Produce the clone by:

1. Setting frontmatter fields:
   - `name`: `<new-name>`
   - `domain`: `<target-domain>`
   - `status`: `draft`
   - `created_at`: today (YYYY-MM-DD)
   - `updated_at`: today (YYYY-MM-DD)
   - `replaces`: `<source>` (records provenance)
   - `plugin`: `null` (clone starts unpromoted)
   - All other frontmatter fields (`version`, `intent`, `tags`, `description`, `invocation`): inherited from source

2. Preserving the full body verbatim (all sections: Arguments, Steps, Artifacts, Invariants, Evidence).

3. Appending to the Evidence section:
   ```
   - Cloned from: `workflows/<domain>/<source>.md` on <today> via /workflow-clone
   ```

---

## Step 4 — Write clone

Write `workflows/<target-domain>/<new-name>.md` with the clone content.

---

## Step 5 — Output

```
Cloned: workflows/<domain>/<source>.md
    →   workflows/<target-domain>/<new-name>.md

Status:   draft
Replaces: <source>

Next steps:
  1. Review and modify Steps/Invariants as needed
  2. Update description and invocation in frontmatter
  3. Dry-run the steps before activating
  4. Run: /workflow-edit <new-name> status=active  (when ready)
```

---

## Invariants

- Clone always starts `status: draft` — never inherits `active` or `deprecated`.
- `replaces` always set to the source workflow name.
- Source workflow is not modified.
- Body sections are inherited verbatim — do not summarize or truncate.
- If source is `deprecated`, include a note: `WARN: Cloning a deprecated workflow. Review whether the deprecation reason applies to the clone.`
