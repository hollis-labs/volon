# Stage 7 — Create Tasks

If `track_tasks: false`: output `[skip] track_tasks disabled — no tasks created.` and proceed to Step 8.

Read `artifacts/plan/<SLUG>.md`. For each task in the Task Breakdown:

Follow the `task-create` protocol:
1. Run: !`ls <storage.files.root>/TASK-<YYYYMMDD>-*.md 2>/dev/null | wc -l | tr -d ' '`
2. Next ID = count + 1, zero-padded → `TASK-<TODAY_COMPACT>-<NNN>`
3. Create `.volon/tasks/TASK-<TODAY_COMPACT>-<NNN>.md` with frontmatter:
   - `title`: from plan task breakdown
   - `status: todo`
   - `priority`: from plan (A/B/C)
   - `project`: `<project.name>`
   - `tags`: `[<SLUG>, feature]`
   - `context: dev`
4. Body sections: Description, Acceptance, Verification, Paths, Updates

Output all created task IDs.

Proceed to Step 8.
