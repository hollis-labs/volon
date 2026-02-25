---
name: resume-task
description: Resume work from bootstrap and repo state, loading Task PCC capsule and running preflight before re-entering execution.
argument-hint: "[note='...']"
disable-model-invocation: true
---

# resume-task

## Inputs
- note: optional user instruction; may override "what to do next" but must remain consistent with repo state.

## Required behavior

### Step 1 — Identify the target task

Read `.volon/bootstrap.md` and determine the task to resume:
- Prefer `paused_task_id` in bootstrap frontmatter if present.
- Otherwise select the most recent `paused/doing/blocked` task (by `updated_at`).
- If none, select next highest priority `todo` task.

### Step 2 — Re-ground from repo artifacts

Read in this order:
1. `volon.yaml` — config (pcc.location, pcc.global_dir, pcc.tasks_dir)
2. `.volon/pcc/global/` — L0 Global PCC (project-wide context)
3. `.volon/pcc/tasks/<task_id>.md` — L2 Task PCC capsule (if exists; see Step 3)
4. `.volon/tasks/<task_id>.md` — authoritative task file (source of truth)
5. `.volon/logs/` — recent run log (last entry)

### Step 3 — Load Task PCC capsule (L2)

Attempt to read `.volon/pcc/tasks/<task_id>.md`.

**If capsule exists:**
- Note `paused_commit` and compare to current HEAD:
  - Run: `git rev-parse HEAD`
  - If HEAD != `paused_commit`: emit `WARN: N commits since pause — review delta before proceeding` and run `git log --oneline <paused_commit>..HEAD` to surface changes.
  - If HEAD == `paused_commit`: no drift, clean resume.
- Extract "Next 1–3 actions" as the primary resume plan.
- Note confidence level; if `low`, flag for human review before proceeding.
- Note open questions / blockers; if any blocker is unresolved, route to `blocked` status.

**If capsule does not exist:**
- Emit: `INFO: No Task PCC capsule found for <task_id> — falling back to task file analysis.`
- Derive next actions from the task file `## Acceptance` checklist and `## Updates` history.

### Step 4 — Preflight checklist

Run these checks before executing any work. A failed gate must be surfaced before proceeding.

| Gate | Check | Failure action |
|---|---|---|
| Git cleanliness | `git status --porcelain` — expect empty or only expected staged/unstaged changes | WARN: list unexpected changes; ask agent to stash or commit before resuming |
| Branch match | Current branch matches capsule `Worktree / branch` (if capsule exists and specifies a branch) | WARN: branch mismatch — confirm correct worktree before proceeding |
| Commit drift | HEAD == `paused_commit` (if capsule exists) | WARN: emit `git log --oneline <paused_commit>..HEAD`; agent adjusts next actions if needed |
| Open blockers | No unresolved `blocker_type: external` on task | BLOCK: cannot resume until external dependency resolved |
| Confidence | Capsule confidence != `low` (if capsule exists) | FLAG: low confidence — emit capsule's open questions; human review recommended |

**Preflight result:**
- All gates pass → proceed to Step 5.
- WARN gates → proceed with warnings noted in run log.
- BLOCK gates → set task `status: blocked`, update bootstrap, do not execute, emit explanation.

### Step 5 — Continue execution

- Transition task `status: doing`.
- If capsule exists: execute the "Next 1–3 actions" from the capsule in order.
- If no capsule: analyze the task file and derive the next concrete action.
- Apply normal loop rules: small verifiable steps, update task on each step.
- Write a run log entry describing: resume trigger, task selected, capsule loaded (yes/no), preflight result, first action taken.
- Do not write to multiple tasks at once.

### Step 6 — End condition

End with DONE after establishing the resumed execution context and completing the first concrete verifiable step (or after completing 1–3 actions per `bootstrap.max_tasks_per_run` policy).

## Output

- Selected task id
- Capsule loaded: yes/no
- Preflight result: pass / warn (details) / blocked (reason)
- First next action taken (or blocked reason)
- DONE
