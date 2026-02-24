---
name: workflow-inception
description: Run the Forge inception loop — read bootstrap, select next work unit, execute, verify, log, finalize. The canonical self-building loop.
argument-hint: "[max_tasks=3] [commit_mode=iteration|task]"
disable-model-invocation: true
---

# Workflow: Inception

The "system builds itself" loop. Forge uses its own artifacts (tasks, PCC, bootstrap, skills, workflows) to evolve a repo with high throughput and low drift.

**Reference:** `docs/13_inception-workflow.md`

---

## Step 0 — Preflight

Read ground truth from repo artifacts:

1. Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`
2. Read `.forge/bootstrap.md` (if present) — note iteration number, next actions, blockers.
3. Read `.forge/pcc/00_project.md` (high-level only) — note version, stage, goals.
4. Read `.forge/tasks/TASK-*.md` files — collect status counts and top-3 `todo` tasks by priority.
5. Run: !`git status --short` and !`git branch --show-current`

Emit boot confirmation block:
```
Forge Orchestrator confirmed. Here's the current state:
**Iteration <N>** | Branch: `<branch>` | <version/milestone>
**Status:** <summary>
**Backlog (<n> items):** <top items>
**Ready.** <next action in one sentence>
```

Exit criteria: can state "what's next" in one sentence from artifacts alone.

---

## Step 1 — Select next work unit

Selection order (strict):
1. If bootstrap names a specific next task/step → use that.
2. Else pick highest-priority `todo` task (A > B > C), oldest first for ties.
3. If no `todo` tasks → promote a backlog item: read `.forge/backlog/`, select highest priority, create a TASK file, set status `todo`.
4. If no backlog → run `/workflow-new-feature` or ask the user for direction.

Emit: `**[TASK-XXXXXX-NNN] starting** — <title>`

---

## Step 2 — Execute

For the selected task:

1. Set task `status: doing` in the task file.
2. Read the task's Description, Acceptance, and Verification sections.
3. Implement the smallest coherent change that satisfies all acceptance criteria.
   - Prefer editing existing files over creating new ones.
   - No large refactors without an explicit planning artifact.
4. Verify against the task's Verification steps.
5. Append a timestamped `Updates` entry to the task file with factual results.

**Sub-agent delegation (optional):** If analysis benefits from parallel reads, dispatch a read-only sub-agent with a single objective, explicit inputs, and a strict output format. Emit `**[sub-agent]** <objective> — delegated` before dispatch and `**[sub-agent]** done — <result>` on return. Integrate output in the Orchestrator; do not let sub-agents write files.

Constraint: complete ≤ `max_tasks` tasks per run (default: 3). Stop and finalize if the limit is reached.

---

## Step 3 — Close work unit

Mark the task:
- `done` if all acceptance criteria are verified
- `blocked` if an external dependency prevents completion (note blocker in task Updates)
- `paused` if intentionally suspended (use `/pause-task`)

Emit the appropriate signal:
- `**[TASK-XXXXXX-NNN] done** — <one-line result>`
- `**[TASK-XXXXXX-NNN] blocked** — <blocker>`
- `**[TASK-XXXXXX-NNN] paused** — <resume hint>`

Write a run log entry at `.forge/logs/run-<date>-iter<N>.md`:
```
# Run Log — Iter <N> — <date>

## Tasks
- TASK-XXXXXX-NNN: <title> → <done|blocked|paused>

## Deliverables
- <files changed/added, one per line>

## Notable decisions
- <key decisions made, if any>

## Constraints / risks
- <anything that limited scope or introduced risk>

## Next actions
- <what should happen next>
```

If `max_tasks` reached: note "Stopped after <N> tasks (limit reached). Next: <task-id>."

---

## Step 4 — Finalize iteration

1. Run `/bootstrap-update` to regenerate `.forge/bootstrap.md` from repo ground truth.
2. Emit: `**[Iter N] finalizing** — bootstrap update + commit`
3. Commit per policy:
   - `commit_mode=iteration` (default): one commit covering all changes this run. Message: `iter <N>: <one-line summary>`
   - `commit_mode=task`: one commit per task already done during this run. Message: `<task-id>: <title>`
4. Emit: `**[commit]** <mode> — iter <N>` (or task-id if task mode)
5. End with `DONE`.

---

## Arguments

| Argument | Default | Description |
|---|---|---|
| `max_tasks` | `3` | Stop after completing this many tasks in one run |
| `commit_mode` | `iteration` | `iteration` = one commit per run; `task` = one commit per task |
