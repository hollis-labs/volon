---
name: workflow-edit
description: Apply structured edits to an existing workflow definition; updates updated_at and validates the result.
argument-hint: "<name> [domain=user] [field=value ...]"
disable-model-invocation: true
---

# Workflow Edit

Open an existing workflow definition for structured editing. Reads the current
file, applies the requested field changes, validates the result against the
format in `docs/08_workflow-authoring.md`, and writes the updated file.

Supported field args: `description=`, `status=`, `tags=`, `invocation=`, `version=`

---

## Arguments

- `name`: `$0` (required) — workflow name to edit
- `domain`: `$1` — `user` (default) or `forge`
- `field=value`: remaining args — explicit field updates to apply

---

## Step 1 — Read config and locate file

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

Resolve `name` from `$0`. Resolve `domain` from `$1`; default `user`.
Target path: `workflows/<domain>/<name>.md`

Run: !`cat workflows/<domain>/<name>.md 2>/dev/null || echo "NOT_FOUND"`

If `NOT_FOUND`: output `ERROR: workflow '<name>' not found in workflows/<domain>/. Use /workflow-create to scaffold a new one.` and stop.

---

## Step 2 — Read and display current state

Parse the file's frontmatter fields: name, version, domain, status, description, tags, invocation, plugin, replaces.

Output a summary of current state:
```
Workflow: <name>  (workflows/<domain>/<name>.md)
Status:      <status>
Description: <description>
Invocation:  <invocation>
Tags:        <tags>
Version:     <version>
```

If `status: deprecated`: output `WARN: This workflow is deprecated. Editing a deprecated workflow is allowed but unusual. Consider /workflow-clone to create an active successor.`

---

## Step 3 — Resolve and validate changes

Parse remaining args for `field=value` pairs. Recognized fields:
- `description` — string; update frontmatter `description`
- `status` — must be `draft`, `active`, or `deprecated`; if `deprecated`, recommend `/workflow-deprecate` instead
- `tags` — comma-separated; parse to YAML list
- `invocation` — string; update frontmatter `invocation`
- `version` — string; bump semver or set explicitly

If no field args provided: output the current file content in full, then output:
```
No changes specified. Provide field=value args to update, or edit the file directly.
Example: /workflow-edit <name> status=active description="New description"
```
and stop.

Validate requested `status` transition:
- `draft` → `active`: allowed
- `active` → `deprecated`: redirect to `/workflow-deprecate` with a message
- Any → `draft`: allowed (rollback)

---

## Step 4 — Apply changes

For each resolved field change:
1. Update the corresponding frontmatter field in the file.
2. Set `updated_at` to today (YYYY-MM-DD).
3. Do not modify the body sections (Steps, Invariants, etc.) — those are author-managed.

Write the updated file.

---

## Step 5 — Output

```
Updated: workflows/<domain>/<name>.md
Changed fields:
  <field>: <old> → <new>
  updated_at: <old> → <today>
```

If `status` was set to `active`:
```
NOTE: Workflow is now active. Ensure you have completed:
  [ ] Steps section is fully written
  [ ] Invariants section is complete
  [ ] Dry-run tested (see docs/08_workflow-authoring.md#dry-run-testing)
```

---

## Invariants

- Never modify body sections (Steps, Artifacts, Invariants, Evidence) — only frontmatter.
- Always update `updated_at` on any change.
- Redirect `status: deprecated` transitions to `/workflow-deprecate`.
- If no valid changes are resolved from args, do not write the file.
