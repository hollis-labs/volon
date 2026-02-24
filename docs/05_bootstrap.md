---
intent: system_doc
audience: humans
---

# Iteration Bootstrap — v0.1

## Purpose
Provide a **single, small, action-oriented** file that enables:
- restarting in a fresh prompt/session without relying on chat context
- clear "what to do next"
- a stable boundary between iterations

Bootstrap is **not** PCC and **not** user docs. It is an execution handoff.

## Files
- `.forge/bootstrap.md` (current; overwritten each finalize)
- `.forge/bootstrap/history/bootstrap-iteration-<N>.md` (optional history copy)

## Bootstrap format
```markdown
---
version: 1
type: bootstrap
iteration: <N>
updated_at: <ISO-8601>
config:
  path: forge.yaml
  storage_backend: files|nanite
  pcc_path: .forge/pcc
  tasks_path: .forge/tasks
  logs_path: .forge/logs
execution:
  driver: tasks
  max_tasks_per_run: 6
  completion_token: DONE
---

# Forge Bootstrap

## What to do next
- <one-line next action>

## Current state summary
- forge.yaml: present
- PCC: present
- Tasks: X todo, Y blocked, Z done
- Last run log: <path>

## Active work
### Highest priority todos
- TASK-... — <title> (A)
- ...

### Blockers
- TASK-... — <one-line blocker>

## Guardrails
- Treat repository files as ground truth. Do not rely on chat memory.
- Minimal diffs; small verifiable steps.
- Update tasks + write a run log per loop invocation.

## Evidence pointers
- PCC index: .forge/pcc/00_project.md
- Workflow contract: docs/03_workflow-contracts.md
- Task model: docs/04_task-model.md
```

## When to generate bootstrap
At the end of each iteration ("Finalize iteration"):
1. Ensure tasks are updated (todo/doing/blocked/done).
2. Ensure a run log exists.
3. Refresh PCC if needed.
4. Generate `.forge/bootstrap.md`.
5. Copy to history.

## Why this matters
Bootstrap makes **manual runs** behave like **automated runs** by forcing clean re-grounding in repo state.
