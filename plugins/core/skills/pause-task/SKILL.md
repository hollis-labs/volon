---
name: pause-task
description: Pause current work safely by updating task state, writing a Task PCC resume capsule, and updating bootstrap with a resume point.
argument-hint: "[mode=soft|restart|compact] [note='...']"
disable-model-invocation: true
---

# pause-task

## Inputs
- mode: soft|restart|compact (default soft)
- note: optional additional context

## Required behavior

### Step 1 — Identify the active task

- Prefer a task with `status: doing` and most recent update.
- If none, prefer the highest priority `doing/paused/blocked` task.
- If still none, do not invent: write bootstrap with "No active task; resume by selecting next todo".

### Step 2 — Update the task file

- Set `status: paused` (preferred) OR `status: blocked` with a pause marker.
- Append to `## Updates`:
  - timestamp
  - mode
  - short note (include any user-supplied note)
  - evidence pointers (git status/diff if available)

### Step 3 — Write Task PCC capsule (L2)

Write (or overwrite) `.forge/pcc/tasks/<task_id>.md` using the schema from
`.forge/templates/task-pcc-capsule.md`.

Populate each field as follows:

| Field | Source |
|---|---|
| `task_id` | Task file frontmatter `id` |
| `paused_commit` | Run: `git rev-parse HEAD` |
| `paused_at` | Current ISO-8601 timestamp |
| `confidence` | Agent's assessment: high \| medium \| low |
| **Goal** | Task `## Description` (1 sentence summary) |
| **Acceptance criteria** | Task `## Acceptance` checklist (copy verbatim, trimmed to fit cap) |
| **Current plan / hypothesis** | Agent's current working theory or approach |
| **Worktree / branch** | `git rev-parse --abbrev-ref HEAD` + worktree path if applicable |
| **Key files touched** | Files modified since task start (from `git diff --name-only`) |
| **Commands run + outcomes** | Last 3–5 significant commands and their pass/fail result |
| **Git diff stat** | Output of `git diff --stat HEAD` (or `git diff --stat <task_branch_base>..HEAD`) |
| **Next 1–3 actions** | The agent's explicit, verifiable next steps |
| **Open questions / blockers** | Any unresolved questions or external blockers |

**Cap:** Max 400 words. Trim prose, not structure. Never omit headings or the next actions section.

**Overwrite policy:** This file is always overwritten on pause. It is not append-only.
The git history preserves prior versions.

### Step 4 — Update `.forge/bootstrap.md`

- Set "What to do next" to resume the paused task.
- Add or update frontmatter keys:

```yaml
paused_task_id: "TASK-YYYYMMDD-###"
paused_at: "ISO-8601"
resume_hint: "one line describing where to pick up"
```

### Step 5 — Write run log (if enabled)

If `observability.write_run_log` is enabled:
- Append a pause entry to the current run log OR create a minimal pause log in `.forge/logs/`.

### Step 6 — Output Resume Instructions

- If mode=restart: instruct user to start a new session and run `/resume-task`.
- If mode=compact: instruct user to compact context if supported, then run `/resume-task`.
- If mode=soft: instruct that `/resume-task` can be used any time.

## Output

List files changed:
- `.forge/tasks/<task_id>.md` (status → paused, Updates appended)
- `.forge/pcc/tasks/<task_id>.md` (Task PCC capsule written)
- `.forge/bootstrap.md` (paused_task_id, paused_at, resume_hint updated)
- `.forge/logs/<run-log>.md` (if observability enabled)

Print updated bootstrap "Quick Start" block.
Print "Resume Instructions" per mode.
DONE
