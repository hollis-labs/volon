# Dead Code Mode — Steps 2–5

## Step 2 — Scan for dead code signals

Run the following four signal checks and collect all findings into a list.
Each finding is: `{ signal, path, detail }`.

### Signal A — Plugins missing plugin.json

Run: !`find plugins/ -mindepth 1 -maxdepth 1 -type d 2>/dev/null || true`

For each plugin directory found:
Run: !`ls plugins/<plugin>/ 2>/dev/null || true`

If the directory contains no `plugin.json`:
→ Add finding: `{ signal: "missing_plugin_json", path: "plugins/<plugin>/", detail: "No plugin.json found — plugin unregistered." }`

### Signal B — SKILL.md files not referenced in any workflow

Run: !`find plugins/ -name "SKILL.md" 2>/dev/null || true`

For each SKILL.md found, extract `name:` from frontmatter:
Run: !`grep "^name:" <path>/SKILL.md 2>/dev/null | head -1 || true`

Also check for `standalone: true` in frontmatter:
Run: !`grep "^standalone:" <path>/SKILL.md 2>/dev/null | head -1 || true`

If `standalone: true` is present: skip this skill — standalone skills are intentionally
invoked directly by users and are not expected to be referenced by workflows.

Then check if the skill name appears in any workflow SKILL.md under `plugins/workflows/`:
Run: !`grep -r "<skill-name>" plugins/workflows/ 2>/dev/null || true`

If no match found:
→ Add finding: `{ signal: "unreferenced_skill", path: "<path>/SKILL.md", detail: "Skill '<skill-name>' not referenced by any workflow." }`

Note: Workflow skills themselves are excluded from this check.
Note: Skills with `standalone: true` in frontmatter are exempt from Signal B.

### Signal C — Stale draft artifacts (> 30 days)

Run: !`find artifacts/ -name "*.md" 2>/dev/null || true`

For each artifact file:
Run: !`grep "^status:" <path> 2>/dev/null | head -1 || true`
Run: !`grep "^created_at:" <path> 2>/dev/null | head -1 || true`

If `status: draft` AND `created_at` is more than 30 days before today:
→ Add finding: `{ signal: "stale_draft_artifact", path: "<path>", detail: "status: draft, created_at: <date> (>30 days stale)." }`

### Signal D — Stale todo tasks (> 60 days)

Run: !`find <storage.files.root> -name "TASK-*.md" 2>/dev/null || true`

For each task file:
Run: !`grep "^status:" <path> 2>/dev/null | head -1 || true`
Run: !`grep "^created_at:" <path> 2>/dev/null | head -1 || true`

If `status: todo` AND `created_at` is more than 60 days before today:
→ Add finding: `{ signal: "stale_todo_task", path: "<path>", detail: "status: todo, created_at: <date> (>60 days stale)." }`

---

## Step 3 — Collate findings

Build a findings table:

| # | Signal | Path | Detail |
|---|---|---|---|
| 1 | <signal> | <path> | <detail> |
| ... | | | |

If the findings list is empty:
→ Output `No dead code issues found.` and go to Step 5.

---

## Step 4 — Apply on_issue action

For each finding in the findings table, apply the effective action:

### action = log

Print the finding to output. Do not create tasks. No file changes.

### action = create_task

For each finding, create a task using the task-create protocol:
- `title`: `"Quality: <signal> — <path>"`
- `priority`: `C`
- `tags`: `[quality, dead_code]`
- `context`: current context from config (default: `dev`)
- `description`: the finding detail

Run: !`ls <storage.files.root>/TASK-*.md 2>/dev/null | wc -l | tr -d ' ' || echo "0"`

Use this count to derive the next sequence number (zero-padded 3 digits).
Write each new task file per the standard task format.

### action = auto_fix_pr

v0.1: `auto_fix_pr` is not implemented.
→ Log a note: `WARN: auto_fix_pr not implemented in v0.1 — falling back to create_task.`
→ Apply `create_task` action instead.

---

## Step 5 — Output summary

Print:

```
quality-run complete (mode: dead_code)
Findings: <N>
Action taken: <effective_action>
<findings table or "No dead code issues found.">
```

If `observability.write_run_log` is true: append a brief entry to
`<log_dir>/run-<YYYYMMDD>-<HHMM>-quality.md`:
- datetime
- mode
- findings count
- action taken
- list of affected paths
