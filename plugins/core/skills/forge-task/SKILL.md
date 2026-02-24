---
name: forge-task
description: Use the Go-based forge task CLI to create/list/update tasks while preserving markdown as the source of truth.
argument-hint: "<subcommand> [...]"
disable-model-invocation: true
---

# forge-task

## Purpose
Give agents a deterministic checklist for when and how to run the compiled Tasks CLI (`forge task ...`). The CLI automates common lifecycle operations but **never** replaces `/pause-task`, `/resume-task`, or PCC workflows.

## When to use
- Creating a new task file with canonical frontmatter (`forge task create "Title"`).
- Moving a task from `todo → doing` (`forge task start <id>`) or `doing → done` (`forge task done <id>`).
- Listing or inspecting tasks without manually grepping `.forge/tasks/` (`forge task list`, `forge task show`).
- Rebuilding the SQLite cache after manual edits (`forge task reindex`).

## Guardrails
- Markdown files remain canonical: never edit SQLite directly, and never rely on the DB as the primary record.
- Pause/resume transitions still go through `/pause-task` and `/resume-task`.
- `todo.db`, `todo.db-shm`, `todo.db-wal` at repo root belong to the Nanite app. Do **not** read, delete, or migrate them.
- Only call `forge task start/done` if you own that transition for the identified task.
- Run commands from the repo root (or pass `--repo <path>`) so the CLI picks up the correct `forge.yaml`.

## Procedure

### Step 1 — Confirm intent
- Re-read `.forge/tasks/<id>.md` or `forge task show <id>` to ensure you have the latest state.
- For create/start/done, verify no other agent is mid-transition (single-writer rule).

### Step 2 — Run the appropriate command
- `forge task create "…" [--type ... --priority ... --tags ... --parent ...]`
- `forge task start <id>` (only from `status: todo`).
- `forge task done <id>` (only from `status: doing`).
- `forge task list [filters]` to find candidates.
- `forge task reindex` after manual edits, schema bumps, or deleting `.forge/state/forge.db`.

### Step 3 — Validate outputs
- Re-open the task file to confirm frontmatter + `## Updates` were edited as expected.
- For start/done, ensure the appended update line captures the action and timestamp.
- If the CLI prints “SQLite unavailable,” note the warning and continue (file writes already succeeded). Consider running `forge task reindex` once the underlying issue is fixed.

### Step 4 — Continue the workflow
- Update PCC / bootstrap / tasks artifacts as required by the active task (the CLI does not touch those files).
- If the transition unblocks downstream work, emit the appropriate orchestration signal (`**[TASK-...] done**`, etc.).

## Reindex triggers
- Manual edits to `.forge/tasks/`.
- Corrupt/missing `.forge/state/forge.db`.
- CLI prompts indicating a schema mismatch.

## Invariants
- Markdown tasks are always authoritative; SQLite is rebuildable cache data.
- The CLI must never modify `todo.db*`.
- Commands are deterministic and idempotent where possible (e.g., reindex can run multiple times safely).
