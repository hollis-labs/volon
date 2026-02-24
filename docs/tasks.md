---
intent: system_doc
audience: humans
---

# Forge Tasks CLI

`forge task` is a deterministic Go-based CLI for working with `.forge/tasks/TASK-*.md`. Markdown remains the source of truth; the CLI simply automates common lifecycle transitions and maintains an optional SQLite index at `.forge/state/forge.db` for fast queries.

## Commands

| Command | Description | Example |
|---|---|---|
| `forge task create "<title>" [--type ...] [--priority A|B|C] [--tags a,b] [--parent ID]` | Generates the next `TASK-YYYYMMDD-###`, writes a new markdown file with the standard sections, and indexes it. | `forge task create "Add sprint workflow" --type feature --priority A --tags workflow,v0.5` |
| `forge task start <id>` | Validates the task is `todo`, sets `status: doing`, and appends a timestamped `## Updates` entry. | `forge task start TASK-20260224-001` |
| `forge task done <id>` | Validates the task is `doing`, sets `status: done`, and appends a completion update. | `forge task done TASK-20260224-001` |
| `forge task show <id>` | Prints the entire markdown file to stdout. | `forge task show TASK-20260221-001` |
| `forge task list [--status ...] [--type ...] [--tag ...] [--priority ...] [--limit N]` | Lists tasks from the SQLite index (fallback to file scan if the DB is unavailable). | `forge task list --status todo --priority A --limit 10` |
| `forge task reindex` | Rebuilds `.forge/state/forge.db` from every `.forge/tasks/TASK-*.md`. Run this after manual edits or if the DB schema changes. | `forge task reindex` |

Flags are CSV-friendly: `--tags` accepts `a,b,c` and `--tag` can be repeated to require multiple substrings.

## State machine

| Transition | Owner | Notes |
|---|---|---|
| `create → todo` | `forge task create` | Creates file + index row. |
| `todo → doing` | `forge task start` | Validates current status; rejects other states. |
| `doing → done` | `forge task done` | Validates current status; rejects other states. |
| `doing → paused` | `/pause-task` skill | CLI intentionally does **not** implement pause/resume; skills continue to update PCC/bootstrap. |
| `paused → doing` | `/resume-task` skill | Use `/resume-task` after `/pause-task restart` to regain orchestrator context. |

Pause/resume and PCC capsule generation are still owned by the existing skills. `forge task` never touches `.forge/bootstrap.md`, `.forge/pcc/`, or `.forge/logs/`.

## File layout

- `.forge/tasks/` — Canonical markdown files (`TASK-YYYYMMDD-###.md`). Create/start/done only modify frontmatter and append to `## Updates`.
- `.forge/state/forge.db` — SQLite index (metadata only). Safe to delete; any CLI command will recreate it, and `forge task reindex` fully rebuilds it.
- `todo.db`, `todo.db-shm`, `todo.db-wal` — **Nanite app database files** at repo root. They are unrelated to the Forge Tasks CLI and must remain untouched.

## Reindexing

Run `forge task reindex` when:
- You edited task files manually (outside the CLI).
- The SQLite schema version changes and the CLI prompts you to reindex.
- `.forge/state/forge.db` was deleted or corrupted.

The command scans every `.forge/tasks/TASK-*.md`, parses frontmatter, and repopulates the `tasks` table. Any files that fail to parse are skipped with an error message so you can fix them manually.

## Examples

```sh
# Create a new task and start it
forge task create "Implement Automation Framework" --type feature --priority A --tags automation,v0.5
forge task start TASK-20260224-004

# List active todos
forge task list --status todo --priority A --limit 5

# Finish work and inspect the final file
forge task done TASK-20260224-004
forge task show TASK-20260224-004 | less

# Rebuild the cache after a manual edit
rm -f .forge/state/forge.db
forge task reindex
```

When in doubt, run any command with `--help` to see the available flags and usage hints.
