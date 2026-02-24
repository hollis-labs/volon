---
name: backlog-task
description: Capture an idea or work item to the backlog; optionally promote a captured item to an active task.
argument-hint: "\"Title\" [priority=B] [tags=tag1,tag2] [context=dev] | promote <backlog-id> | list"
disable-model-invocation: true
---

# Backlog Task

Capture ideas and work items to `.forge/backlog/` without committing to execution.
Backlog items are distinct from active tasks — they require a deliberate **promote**
step before they enter the task loop.

---

## Modes

- **capture** (default): `/backlog-task "Title" [priority] [tags] [context]`
- **list**: `/backlog-task list [status=captured|promoting|promoted|dropped]`
- **promote**: `/backlog-task promote <BACKLOG-ID>`

Resolve mode from the first argument: if it starts with `list`, run list mode;
if it starts with `promote`, run promote mode; otherwise run capture mode.

---

## Step 1 — Read config

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract:
- `storage.files.root` → default `.forge/tasks`
- `observability.write_run_log` → default `false`
- `observability.log_dir` → default `.forge/logs`

Set backlog dir: `.forge/backlog/`

---

## Step 2 — Execute mode

### Capture mode

1. Determine next BACKLOG ID:
   - Run: !`ls .forge/backlog/BACKLOG-*.md 2>/dev/null | sort | tail -1 || true`
   - Extract sequence number; increment by 1 (zero-pad to 3 digits).
   - Format: `BACKLOG-YYYYMMDD-NNN` (today's date + sequence).
   - If no existing files: start at `BACKLOG-YYYYMMDD-001`.

2. Write `.forge/backlog/BACKLOG-YYYYMMDD-NNN.md`:

```yaml
---
id: BACKLOG-YYYYMMDD-NNN
title: "<title from arg>"
status: captured
priority: <priority, default B>
project: <project from forge.yaml>
tags: [<tags>]
context: <context, default dev>
created_at: YYYY-MM-DD
updated_at: YYYY-MM-DD
promoted_to: null
---

## Description

<title> — captured via /backlog-task. Add details here before promoting.

## Notes

(empty — to be filled before promotion)
```

3. Output:
```
Captured: BACKLOG-YYYYMMDD-NNN
Title: <title>
File: .forge/backlog/BACKLOG-YYYYMMDD-NNN.md
Next: edit the file to add details, then run `/backlog-task promote BACKLOG-YYYYMMDD-NNN`
```

---

### List mode

Run: !`ls .forge/backlog/BACKLOG-*.md 2>/dev/null || echo "NO_BACKLOG"`

If `NO_BACKLOG`: output `No backlog items.` and stop.

For each file, read frontmatter (`id`, `title`, `status`, `priority`).
If `status` filter provided: show only matching items.

Output table:
```
ID                     | Status    | P | Title
BACKLOG-YYYYMMDD-001   | captured  | B | My idea
...
```

---

### Promote mode

1. Read `.forge/backlog/<backlog-id>.md`. Fail fast if not found.
2. Verify `status` is `captured` or `promoting` (not already `promoted` or `dropped`).
3. Display the backlog item (id, title, description, notes, priority, tags, context).
4. Apply `/task-create` protocol using the backlog item's fields:
   - title → task title
   - priority, tags, context → task fields
   - description + notes → task body
5. Note the created task ID.
6. Update the backlog item:
   - `status: promoted`
   - `promoted_to: <TASK-ID>`
   - `updated_at: today`
7. Output:
```
Promoted: BACKLOG-YYYYMMDD-NNN → TASK-YYYYMMDD-NNN
Backlog item status: promoted
Task created: TASK-YYYYMMDD-NNN
```

---

## Invariants

- Backlog items live in `.forge/backlog/` — never in `.forge/tasks/`.
- Capture is always fast — do not ask for details at capture time.
- Promote requires reading the backlog item first — fail fast if not found.
- Never auto-promote. Promotion is always an explicit user action.
- If `.forge/backlog/` does not exist: create it before writing the first item.
- `sprint_id` and `iteration_id` may be set on the promoted task (not the backlog item).
