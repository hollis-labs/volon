# Stage 6 — Create Tasks

Read `artifacts/slash-commands/<SLUG>/plan.md`. For each task in the Task Breakdown:

Follow the `task-create` protocol:
1. Run: !`ls <storage.files.root>/TASK-<YYYYMMDD>-*.md 2>/dev/null | wc -l | tr -d ' '`
2. Next ID = count + 1, zero-padded → `TASK-<TODAY_COMPACT>-<NNN>`
3. Create `.forge/tasks/TASK-<TODAY_COMPACT>-<NNN>.md` with frontmatter:
   - `title`: from plan task breakdown
   - `status: todo`
   - `priority`: from plan (A/B/C)
   - `project`: `<project.name>`
   - `tags`: `[<SLUG>, slash-command]`
   - `context: dev`
4. Body sections: Description, Acceptance, Verification, Paths, Updates

If spec.md indicates a dependency (new skill required first), create a blocked task
with `status: blocked` and note the dependency in the task body.

Output all created task IDs.

Proceed to Step 7 (Finalize).
