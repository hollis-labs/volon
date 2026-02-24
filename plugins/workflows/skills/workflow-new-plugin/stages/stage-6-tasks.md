# Stage 6 — Create Tasks

Read `artifacts/plugins/<SLUG>/plan.md`. For each task in the Task Breakdown:

Follow the `task-create` protocol:
1. Run: !`ls <storage.files.root>/TASK-<YYYYMMDD>-*.md 2>/dev/null | wc -l | tr -d ' '`
2. Next ID = count + 1, zero-padded → `TASK-<TODAY_COMPACT>-<NNN>`
3. Create `.forge/tasks/TASK-<TODAY_COMPACT>-<NNN>.md` with frontmatter:
   - `title`: from plan task breakdown
   - `status: todo`
   - `priority`: from plan (A/B/C)
   - `project`: `<project.name>`
   - `tags`: `[<SLUG>, plugin]`
   - `context: dev`
4. Body sections: Description, Acceptance, Verification, Paths, Updates

Note: scaffold tasks (plugin.json + dirs + stubs) are the immediate output.
Full skill implementation tasks should reference workflow-new-skill for follow-up.

Output all created task IDs.

Proceed to Step 7 (Finalize).
