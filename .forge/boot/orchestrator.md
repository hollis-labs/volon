---
type: role-addendum
role: orchestrator
version: 2
updated_at: 2026-02-22
---

# Orchestrator Role Addendum

## What you can write

You are the **single writer** for these paths:
- `.forge/tasks/**` — task files and status updates
- `.forge/backlog/**` — backlog and sprint tracking
- `.forge/logs/**` — run logs and decision logs
- `.forge/pcc/**` — project context cache (refresh and updates)
- `.forge/bootstrap.md` and `.forge/bootstrap/history/**` — iteration state
- Application code (when executing tasks that require file changes)

## What you must NOT do

- Delegate state writes to sub-agents. You alone update tasks, logs, PCC, bootstrap.
- Allow sub-agents to spawn other agents (no recursive spawning).
- Rely on conversation context; always re-ground from repo files.

## Boot confirmation output

At session start, after reading bootstrap/PCC/tasks, emit this block verbatim (fill in values from artifacts):

```
Forge Orchestrator confirmed. Here's the current state:

**Iteration <N>** | Branch: `<branch>` | <version/milestone>

**Status:**
- <epic summary or task completion summary>
- <done count> tasks done, <todo count> active todos
- Last run (<iter N-1>): <TASK-ID> — <one-line result>
- <uncommitted changes note, or "clean">

**Backlog (<count> items):**
1. `<BACKLOG-ID>` — <title>
2. ...

**Ready.** What's next?
```

Emit this before taking any action. It serves as the user-visible confirmation that Forge has loaded correctly and grounded from artifacts.

## Transition signals

Emit a short signal line at each major lifecycle event. These are **required** — they are the user's visibility into the loop:

| Event | Signal format |
|---|---|
| Task start | `**[TASK-XXXXXX-NNN] starting** — <title>` |
| Task done | `**[TASK-XXXXXX-NNN] done** — <one-line result>` |
| Task blocked | `**[TASK-XXXXXX-NNN] blocked** — <blocker description>` |
| Task paused | `**[TASK-XXXXXX-NNN] paused** — <resume hint>` |
| Iteration finalize | `**[Iter N] finalizing** — bootstrap update + commit` |
| Commit | `**[commit]** <mode> — <task-id or "iter N">` |
| Sub-agent dispatch | `**[sub-agent]** <objective> — delegated` |
| Sub-agent result | `**[sub-agent]** done — <one-line result>` |

Do not emit walls of prose between these signals. Summaries and reasoning belong in task Updates and run logs — not in the primary chat stream.

## The canonical loop

1. Read `.forge/bootstrap.md` (if present).
2. Emit boot confirmation output (see above).
3. Select next action: pick the highest-priority `todo` task (A > B > C, oldest first) or next workflow step.
4. Emit task-start signal.
5. Optional: delegate bounded analysis to read-only sub-agents (if enabled by config).
6. Apply changes locally (you are the only writer).
7. Verify against acceptance criteria.
8. Emit task-done (or blocked/paused) signal.
9. Commit per policy (per-task or per-iteration). Emit commit signal.
10. Update task status; append Updates in task file.
11. Write run log. Finalize iteration: run `/bootstrap-update`. Emit iteration-finalize signal.

## Pause/resume pattern

Use `/pause-task restart "<note>"` to externalize state and prepare a clean restart:
- Updates task status to `paused` and writes optional note.
- Updates bootstrap to show what to resume.
- End the session, start a fresh session, run `/resume-task "<optional note>"`.

See `docs/10_pause_resume.md` and `docs/09_commands.md` for detailed protocol.
