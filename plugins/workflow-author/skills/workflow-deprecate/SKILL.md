---
name: workflow-deprecate
description: Mark a workflow as deprecated; records the reason and optional successor.
argument-hint: "<name> [domain=user] [successor=<name>] [reason=\"...\"]"
disable-model-invocation: true
---

# Workflow Deprecate

Mark a workflow `status: deprecated`. Non-destructive — the file is retained
as a historical record. If a successor workflow exists, record it.

---

## Arguments

- `name`: `$0` (required) — workflow to deprecate
- `domain`: `$1` — `user` (default) or `volon`
- `successor`: `$2` — name of the workflow that replaces this one (optional)
- `reason`: `$3` — one-sentence reason for deprecation (optional)

---

## Step 1 — Read config and locate file

Run: !`cat volon.yaml 2>/dev/null || cat .volon/volon.yaml 2>/dev/null || echo "NO_CONFIG"`

Resolve `name` from `$0`. Resolve `domain`; default `user`.
Target path: `workflows/<domain>/<name>.md`

Run: !`cat workflows/<domain>/<name>.md 2>/dev/null || echo "NOT_FOUND"`

If `NOT_FOUND`: output `ERROR: workflow '<name>' not found in workflows/<domain>/.` and stop.

---

## Step 2 — Pre-flight checks

Read current `status`.

If `status: deprecated`: output `WARN: '<name>' is already deprecated. No changes made.` and stop.

If `status: draft`: output `WARN: Deprecating a draft workflow that was never activated. Confirm this is intentional.` and continue.

If `successor` arg provided: verify the successor file exists.
Run: !`ls workflows/<domain>/<successor>.md 2>/dev/null || echo "NOT_FOUND"`
If successor `NOT_FOUND`: output `WARN: successor workflow '<successor>' not found in workflows/<domain>/. Proceeding without linking.` and set successor to null.

---

## Step 3 — Apply deprecation

Update the workflow file frontmatter:
- `status`: `deprecated`
- `updated_at`: today (YYYY-MM-DD)

Append a `## Deprecation` section to the file body (after Evidence, if it does not already exist):

```markdown
## Deprecation

- **Date**: <today>
- **Reason**: <reason if provided, else "Not specified.">
- **Successor**: `workflows/<domain>/<successor>.md` (if provided, else "None specified.")
```

Write the updated file.

---

## Step 4 — Update successor (if provided)

If `successor` was found: read `workflows/<domain>/<successor>.md` and set `replaces: <name>` in its frontmatter (if not already set). Update `updated_at`. Write it.

---

## Step 5 — Output

```
Deprecated: workflows/<domain>/<name>.md
Status:     deprecated (was: <previous-status>)
Reason:     <reason or "not specified">
Successor:  <successor or "none">
```

If no successor was provided:
```
ACTION REQUIRED: If a successor workflow exists, set its frontmatter:
  replaces: <name>
Or create one with: /workflow-create <new-name> [domain=<domain>]
```

---

## Invariants

- Never delete the workflow file — deprecation is always non-destructive.
- Always append the `## Deprecation` section; do not overwrite existing one.
- If successor is provided and found, always set `replaces` on the successor.
- Do not modify any body sections other than appending `## Deprecation`.
- `updated_at` always reflects today's date after deprecation.
