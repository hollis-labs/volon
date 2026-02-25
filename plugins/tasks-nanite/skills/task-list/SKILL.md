---
name: task-list
description: List tasks with optional filters (status/tag/priority/project/context).
argument-hint: "[status=todo|doing|blocked|done] [tag=<tag>] [priority=A|B|C] [project=<name>] [context=<ctx>]"
disable-model-invocation: true
standalone: true
---

# Task List

List tasks from the backend defined by `storage.backend` in `volon.yaml`.
If no arguments: list all tasks.

---

## Arguments (all optional, all key=value)

- `status=<value>` — filter to `todo`, `doing`, `blocked`, or `done`
- `tag=<value>` — filter to tasks whose `tags` list contains this value
- `priority=<value>` — filter to `A`, `B`, or `C`
- `project=<value>` — filter to tasks matching this project name
- `context=<value>` — filter to tasks matching this context

---

## Step 1 — Read config

Run: !`cat volon.yaml 2>/dev/null || cat .volon/volon.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract:
- `storage.backend` → default: `files`
- `storage.files.root` → default: `.volon/tasks`
- `storage.nanite.vault` → default: `null`
- `storage.nanite.tag_prefix` → default: `volon/`

---

## Step 2 — Dispatch on backend

### If `storage.backend` is `nanite`

**Resolve vault ID** (if `storage.nanite.vault` is non-null):

Run: !`nanite vaults 2>/dev/null`

Parse JSON. Find the vault entry where `name` matches `storage.nanite.vault`.
- If found: capture `id` as `VAULT_ID`.
- If not found: output `ERROR: nanite vault '<name>' not found.` and stop.
- If `storage.nanite.vault` is `null`: omit `--vault` flag entirely.

Build search query from filters. Volon tasks are tagged with `<tag_prefix>task`.

Run: !`nanite search "<project_or_tag_prefix>task" --projects <project_filter_or_empty> [--vault <VAULT_ID>] 2>/dev/null`

If `status=` filter is provided and equals `done`: use `--type note` and adjust tag filter.
If no `project=` filter: omit `--projects` flag.
If `vault` is `null`: omit `--vault` flag.

Parse JSON response:
- If `ok: false`: output `ERROR: nanite search failed — <error>.` and stop.
- If `ok: true` and `data.items` is empty: output `No tasks found.` and stop.

Apply any remaining filters (`priority=`, `tag=`, `context=`) to the returned items
by checking each item's tags list (tags are prefixed: `volon/priority:A`, `volon/context:dev`, etc.).

Format and output a markdown table with columns: `ID | Title | Status | Priority | Tags`
using `NANITE-<item.id>` as the ID value. Match column format of the files backend.

Stop — do not proceed to file-backend steps.

---

### If `storage.backend` is `files` (default)

Proceed to Steps 3–6.

---

## Step 3 — Collect task files

Run: !`ls <storage.files.root>/*.md 2>/dev/null`

If no files found: output `No tasks found.` and stop.

---

## Step 4 — Parse each file

For each file path from Step 2:
1. Read the full file contents.
2. Extract YAML frontmatter: the block between the first `---` line and the
   second `---` line.
3. Parse these fields: `id`, `title`, `status`, `priority`, `tags`, `project`,
   `context`, `created_at`.
4. If frontmatter is absent or cannot be parsed: skip the file and record:
   `WARN: skipped <filename> — malformed or missing frontmatter.`

---

## Step 5 — Apply filters

Include a task in the result only if ALL provided filters match:

| Filter argument | Match condition |
|---|---|
| `status=<v>` | `task.status == v` |
| `tag=<v>` | `v` is a member of `task.tags` list |
| `priority=<v>` | `task.priority == v` |
| `project=<v>` | `task.project == v` |
| `context=<v>` | `task.context == v` |

If no filter arguments: include all parsed tasks.

---

## Step 6 — Sort

Sort the filtered list by:
1. Priority ascending: `A` before `B` before `C`
2. Then by `created_at` ascending (oldest first)

---

## Step 7 — Output

Print a markdown table:

```
| ID | Title | Status | Priority | Tags |
|---|---|---|---|---|
| TASK-... | ... | todo | B | volon, tasks |
```

If the filtered list is empty: output `No tasks match the given filters.`

Print any WARN lines collected in Step 3 below the table (or below the
empty-result line).

---

## Invariants

- Never modify any task file.
- Read all `*.md` files in the tasks root regardless of filename format;
  skip non-task files with a WARN rather than failing.
- Do not call Nanite or any external service.
