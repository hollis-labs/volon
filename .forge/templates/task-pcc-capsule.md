---
type: task-pcc-capsule
schema_version: 1
task_id: "TASK-YYYYMMDD-NNN"
paused_commit: ""
paused_at: ""
confidence: "high|medium|low"
---

# Task PCC Capsule — {{task_id}}

> **This is a resume capsule, not a transcript.**
> It contains only what is needed to restart work deterministically.
> Max 400 words. Keep all sections present; trim prose aggressively.

---

## Goal

One-sentence goal for this task.

**Acceptance criteria (brief):**
- [ ] criterion 1
- [ ] criterion 2

---

## Current plan / hypothesis

One paragraph or bullet list describing the current working approach.
State what you believe is true and what you intend to try next.

---

## State markers

**Worktree / branch:**
- Branch: `<branch-name>` (e.g., `main` or `forge/TASK-20260222-007`)
- Worktree: `<path>` or `main worktree`

**Key files touched:**
- `path/to/file1.md` — what changed / why
- `path/to/file2.md` — what changed / why

**Commands run + outcomes:**
```
command run → outcome (pass/fail/partial)
command run → outcome
```

---

## Evidence

**Git diff stat at pause:**
```
 path/to/file1.md | 12 +++++++-----
 path/to/file2.md |  4 ++--
 2 files changed, 16 insertions(+), 7 deletions(-)
```

**paused_commit:** `<HEAD hash at pause>`
*(Compare on resume: if HEAD has diverged, review delta before proceeding.)*

---

## Next 1–3 actions

> Explicit and verifiable. Do exactly these next.

1. **Action 1** — `<specific file or command>` — `<expected outcome>`
2. **Action 2** — `<specific file or command>` — `<expected outcome>`
3. **Action 3** (optional) — `<specific file or command>` — `<expected outcome>`

---

## Open questions / blockers

- [ ] question or blocker (if none, write "None.")

---

## Non-goal reminders

- This capsule is not the task spec. See `.forge/tasks/{{task_id}}.md` for full acceptance criteria.
- This capsule is not the git diff. See `git diff <paused_commit>` for the actual delta.
- This capsule is not an approval gate. Resume and proceed unless a blocker is listed above.
