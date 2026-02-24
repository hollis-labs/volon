---
name: bootstrap-update
description: Generate/update .forge/bootstrap.md from repo ground truth (tasks/logs/PCC/config).
argument-hint: "[iteration=auto] [max_tasks_per_run=auto]"
disable-model-invocation: true
---

# Bootstrap Update

Create or update:
- `.forge/bootstrap.md`

Also write (optional history):
- `.forge/bootstrap/history/bootstrap-iteration-<N>.md`

---

## Step 1 — Read config

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract and note:
- `bootstrap.path` → default: `.forge/bootstrap.md`
- `bootstrap.history_path` → default: `.forge/bootstrap/history`
- `bootstrap.max_tasks_per_run` → default: `6`
- `storage.files.root` → default: `.forge/tasks`
- `observability.log_dir` → default: `.forge/logs`
- `pcc.location` → default: `.forge/pcc`

---

## Step 2 — Read ground truth

Read the following sources to establish current repo state:

- **Tasks:** scan all `<storage.files.root>/TASK-*.md` files
  - Count by status (`todo`, `done`) for the Task Counts table
  - Collect top-3 `todo` tasks ordered by priority (A > B > C), then oldest first for ties
- **Logs:** list `<log_dir>/` to identify the latest run log filename
- **PCC:** read `<pcc.location>/04_backlog.md` for deferred items
- **Config:** `forge.yaml` for workflows, quality modes, and storage settings

---

## Step 3 — Resolve iteration number

- If an existing `<bootstrap.path>` exists: read its `**Iteration**:` line and increment by 1.
- Otherwise: infer from the latest run log name (e.g., `run-20260221-1715-iteration-10.md` → 11).
- Otherwise: default to `1`.
- If caller passes an explicit `iteration=N` argument: use that value instead.

---

## Step 4 — Compose bootstrap.md content

Build the file with these sections:

1. **Header**: `# Forge Bootstrap`, Iteration N, Generated date, source of truth note.
2. **Task Counts**: table of `todo` / `done` / total.
3. **Top 3 Next Tasks**: if `todo` tasks exist, list priority-ordered top 3 with one-line summaries.
   If no `todo` tasks remain: note deferred roadmap items from `04_backlog.md`.
4. **Blockers**: one-line each; `None.` if empty.
5. **Latest Run Log**: filename with brief summary of what that run completed.
6. **Quick Start**: code block with next-step instructions.
7. **Config Snapshot**: key config values (storage, PCC, enabled workflows, quality modes).

Content rules:
- Action-oriented; short.
- Do not copy large blocks from PCC — summarize only.

---

## Step 5 — Write files

1. Write composed content to `<bootstrap.path>`.
2. If `<bootstrap.history_path>` directory exists (or can be created):
   write a copy to `<bootstrap.history_path>/bootstrap-iteration-<N>.md`.

---

## Step 6 — Validate boot pack

Check that `.forge/agent-boot.md` exists.
- If it exists: note it in the output ("boot pack present").
- If it does NOT exist: warn "WARN: .forge/agent-boot.md missing — agents may not boot correctly. Run task to create it."

Also check `.forge/boot/` contains orchestrator.md, worker.md, reviewer.md.
- If any are missing: warn with list of missing files.

Do not create or modify the boot pack here; this step is validation only.

---

## Step 7 — Output

- List changed files.
- Print full contents of `.forge/bootstrap.md`.
- End with `DONE` when used in a loop.
