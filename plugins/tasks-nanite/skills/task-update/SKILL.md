---
name: task-update
description: Update a task's status and/or fields; optionally append a message.
argument-hint: "<id> [status=todo|doing|blocked|done] [priority=A|B|C] [message="text"]"
disable-model-invocation: true
standalone: true
---

# Task Update

Update a task in the backend defined by `storage.backend` in `forge.yaml`.
For the files backend: edits frontmatter and appends message to the task file.
For the nanite backend: pushes a status-update note referencing the task ID
(nanite items are append-only; the update note records the change for audit).

---

## Arguments

- `$0` (required): task ID — `TASK-YYYYMMDD-NNN` (files) or `NANITE-<id>` (nanite)
- `status=<value>` (optional): new status — `todo`, `doing`, `blocked`, or `done`
- `priority=<value>` (optional): new priority — `A`, `B`, or `C`
- `message=<text>` (optional): freeform text describing the update

At least one of `status`, `priority`, or `message` must be provided.
If none: output `ERROR: provide at least one of status=, priority=, or message=.` and stop.

---

## Step 1 — Read config

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract:
- `storage.backend` → default: `files`
- `storage.files.root` → default: `.forge/tasks`
- `storage.nanite.vault` → default: `null`
- `storage.nanite.tag_prefix` → default: `forge/`

---

## Step 2 — Dispatch on backend

### If `storage.backend` is `nanite`

Validate: `$0` must match `NANITE-<digits>`. If not: output
`ERROR: nanite backend requires NANITE-<id> format. Got: <$0>.` and stop.

Extract numeric ID: strip `NANITE-` prefix → `<nanite_id>`.

**Resolve vault ID** (if `storage.nanite.vault` is non-null):

Run: !`nanite vaults 2>/dev/null`

Parse JSON. Find vault by name. Capture `id` as `VAULT_ID`, or omit `--vault` if null.

Verify the item exists:
Run: !`nanite get <nanite_id> [--vault <VAULT_ID>] 2>/dev/null`

- If `ok: false` or item not found: output `ERROR: task NANITE-<nanite_id> not found.` and stop.
- Capture `data.title` for reference in the update note.

Build update summary string from provided arguments:
- If `status=` provided: include `status: <old> → <new>`
- If `priority=` provided: include `priority: → <new>`
- If `message=` provided: include message text

Push a status-update note:
Run: !`nanite push "Update: <data.title>" --source forge-agent --type note --tags <tag_prefix>update,<tag_prefix>task --notes "<p><strong>Task:</strong> NANITE-<nanite_id></p><p><strong>Changes:</strong> <update-summary></p><p><strong>Note:</strong> <message-or-empty></p>" [--vault <VAULT_ID>] 2>/dev/null`

- If `ok: false`: output `ERROR: nanite push failed — <error>.` and stop.

Output:
```
Updated: NANITE-<nanite_id>
Note:    NANITE-<update_note_id> (update record)
```

Stop — do not proceed to file-backend steps.

---

### If `storage.backend` is `files` (default)

Proceed to Steps 3–7.

---

## Step 3 — Locate task file

Construct path: `<storage.files.root>/<id>.md`

Run: !`ls <storage.files.root>/<id>.md 2>/dev/null`

If the file is not found:
- Output `ERROR: task <id> not found in <storage.files.root>.`
- Output `Hint: run /task-list to see available task IDs.`
- Stop.

---

## Step 4 — Read current task file

Read the full file contents.

Extract and hold separately:
- **Frontmatter block**: YAML between the first `---` and second `---`
- **Body block**: all content after the second `---`
- **Current status**: value of `status` field (for output)

---

## Step 5 — Apply frontmatter updates

Update only the fields explicitly provided as arguments:

| Argument | Frontmatter field | Valid values |
|---|---|---|
| `status=<v>` | `status` | `todo`, `doing`, `blocked`, `done` |
| `priority=<v>` | `priority` | `A`, `B`, `C` |

If an invalid value is provided: output `ERROR: invalid value '<v>' for <field>.` and stop.

Always update regardless of other arguments:
- `updated_at` → today's date

Run: !`date +%Y-%m-%d`

Do not modify `id`, `title`, `created_at`, `tags`, `context`, or `project`
unless those fields were explicitly passed as arguments (not supported in v0.1).

---

## Step 6 — Append message to body (if provided)

If a `message` argument was provided:

1. Locate the `## Updates` section in the body.
   - If it does not exist: add `## Updates` as a new section at the very end
     of the body.
2. Append this line inside the `## Updates` section:
   ```
   - [YYYY-MM-DD] <message text>
   ```
   Use today's date. New entries go at the bottom of the section.

Do not modify any other body section.
`## Updates` is append-only — never edit or delete prior entries.

---

## Step 7 — Write updated file

Reconstruct the full file:
1. Opening `---`
2. Updated YAML frontmatter fields (only changed lines replaced; all other
   lines preserved verbatim)
3. Closing `---`
4. Body block with `## Updates` section modified if applicable

Write to `<storage.files.root>/<id>.md`, overwriting the previous version.

---

## Step 8 — Output

```
Updated: <id>
Status: <new_status (or current status if unchanged)>
Path: <storage.files.root>/<id>.md
```

---

## Invariants

- Never change `id` or `created_at` fields (files backend).
- Never delete the body or any body section (files backend).
- `## Updates` is append-only (files backend).
- Nanite backend is append-only: push an update note; never attempt to modify the original item.
- `--vault` requires the vault ID (not the vault name). Always resolve name→ID via `nanite vaults`.
- If `storage.nanite.vault` is `null`, omit `--vault` flag entirely.
- Preserve all frontmatter fields not explicitly updated, verbatim (files backend).
