---
name: commit-task
description: Commit changes for a completed task using the configured commit mode (iteration or isolated).
argument-hint: "[task-id=TASK-YYYYMMDD-NNN] [mode=iteration|isolated]"
disable-model-invocation: true
---

# commit-task

## Inputs
- `task-id` (optional): explicit task ID to commit for. If omitted, uses the most recently completed `doing` task.
- `mode` (optional): `iteration` or `isolated`. If omitted, reads `git.commit_mode` from `forge.yaml` (default: `iteration`).

## Required behavior

### Step 1 — Read config
Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract:
- `git.commit_mode` → `iteration` or `isolated` (default: `iteration`)
- `git.auto_commit` → if `false`, skip commit and print "git.auto_commit is false — skipping commit." then DONE.

Override with `mode` argument if provided.

### Step 2 — Identify task
If `task-id` provided: read `.forge/tasks/<task-id>.md`.
Otherwise: find the most recently updated task with `status: done` in `.forge/tasks/`.

If no task found: output `ERROR: no completed task found to commit.` and stop.

Note task `id` and `title`.

### Step 3 — Check git status
Run: !`git status --porcelain`

If output is empty: output `Nothing to commit for <task-id>.` and stop.

### Step 4 — Stage files
Run: !`git add -A`

Note: Orchestrator is responsible for ensuring only task-relevant files are staged.
If unrelated files appear in git status, warn: `WARN: unrelated changes staged — review before committing.`

### Step 5 — Commit

**Iteration mode:**
Run: !`git commit -m "forge: iteration <N> — <task-count> tasks (<task-id>)"`

Where:
- `<N>` = current iteration from `.forge/bootstrap.md`
- `<task-count>` = number of tasks completed this iteration
- `<task-id>` = this task's ID (or comma-separated list if batching)

**Isolated mode:**
Run: !`git commit -m "forge: <task-id> — <task-title>"`

### Step 6 — Output
- Print commit hash: Run: !`git log -1 --format="%H %s"`
- List changed files: Run: !`git diff HEAD~1 --name-only`
- Output `DONE`

## Notes
- This skill only commits. Push is a separate action (`pr-open` or manual).
- If commit fails (pre-commit hook, conflict), output the error verbatim and stop. Do not retry.
- See `docs/11_git-hooks.md` for commit strategy documentation.
