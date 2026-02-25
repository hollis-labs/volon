# Correctness Mode — Steps 2C–5C

## Step 2C — Scan for correctness signals

Run the following four signal checks and collect all findings into a list.
Each finding is: `{ signal, path, detail }`.

### Signal A — Broken internal links

Run: !`find . -name "*.md" -not -path "./.git/*" -not -path "./.volon/logs/*" 2>/dev/null || true`

For each `.md` file found, extract all inline links matching `[text](path)`:
Run: !`grep -oP '\[([^\]]+)\]\(([^)]+)\)' <file> 2>/dev/null || true`

For each extracted link path:
- Skip paths starting with `http://` or `https://` (external links).
- Resolve relative to the file's directory.
- Check if the resolved path exists: !`test -e <resolved-path> && echo "exists" || echo "missing"`

If path does not exist:
→ Add finding: `{ signal: "broken_internal_link", path: "<file>:<line>", detail: "Link target '<path>' does not exist." }`

### Signal B — Frontmatter schema violations

Run: !`find .volon/tasks/ -name "TASK-*.md" 2>/dev/null || true`

For each task file, check that these required fields are present in frontmatter:
`id`, `title`, `status`, `priority`, `project`, `created_at`

Run: !`grep -E "^(id|title|status|priority|project|created_at):" <file> 2>/dev/null || true`

If any required field is missing:
→ Add finding: `{ signal: "frontmatter_schema_violation", path: "<file>", detail: "Missing required field(s): <missing-fields>." }`

Also check artifact files: !`find artifacts/ -name "*.md" 2>/dev/null || true`

For each artifact file, required fields: `id`, `type`, `status`

If any required artifact field is missing:
→ Add finding: `{ signal: "frontmatter_schema_violation", path: "<file>", detail: "Missing required artifact field(s): <missing-fields>." }`

### Signal C — SKILL.md missing required sections

Run: !`find plugins/ -name "SKILL.md" 2>/dev/null || true`

For each SKILL.md found, verify:
1. Frontmatter contains `name:` field: !`grep "^name:" <file> | head -1 || true`
2. Frontmatter contains `description:` field: !`grep "^description:" <file> | head -1 || true`
3. Frontmatter contains `disable-model-invocation:` field: !`grep "^disable-model-invocation:" <file> | head -1 || true`
4. File contains at least one `## Step N` heading: !`grep -c "^## Step [0-9]" <file> 2>/dev/null || echo "0"`

If any of the above checks fail:
→ Add finding: `{ signal: "skill_missing_required_section", path: "<file>", detail: "Missing: <what is missing>." }`

### Signal D — Stale PCC evidence dates

Run: !`find .volon/pcc/ -name "*.md" 2>/dev/null || true`

For each PCC file:
Run: !`grep -n "Last refreshed:" <file> 2>/dev/null | head -1 || true`

If no `Last refreshed:` line is found:
→ Add finding: `{ signal: "stale_pcc_evidence", path: "<file>", detail: "No Evidence section with 'Last refreshed:' date found." }`

If a date is found, parse it (format: `YYYY-MM-DD`).
If the date is more than 90 days before today:
→ Add finding: `{ signal: "stale_pcc_evidence", path: "<file>", detail: "Last refreshed: <date> — older than 90 days." }`

---

## Step 3C — Collate correctness findings

Build a findings table:

| # | Signal | Path | Detail |
|---|---|---|---|
| 1 | <signal> | <path> | <detail> |
| ... | | | |

If the findings list is empty:
→ Output `No correctness issues found.` and go to Step 5C.

---

## Step 4C — Apply on_issue action

For each finding in the findings table, apply the effective action:

### action = log

Print the finding to output. Do not create tasks. No file changes.

### action = create_task

For each finding, create a task using the task-create protocol:
- `title`: `"Quality: <signal> — <path>"`
- `priority`: `B`
- `tags`: `[quality, correctness]`
- `context`: current context from config (default: `dev`)
- `description`: the finding detail

Run: !`ls <storage.files.root>/TASK-*.md 2>/dev/null | wc -l | tr -d ' ' || echo "0"`

Use this count to derive the next sequence number (zero-padded 3 digits).
Write each new task file per the standard task format.

### action = auto_fix_pr

v0.1: `auto_fix_pr` is not implemented for correctness mode.
→ Log a note: `WARN: auto_fix_pr not implemented in v0.1 — falling back to create_task.`
→ Apply `create_task` action instead.

---

## Step 5C — Output summary

Print:

```
quality-run complete (mode: correctness)
Findings: <N>
Action taken: <effective_action>
<findings table or "No correctness issues found.">
```

If `observability.write_run_log` is true: append a brief entry to
`<log_dir>/run-<YYYYMMDD>-<HHMM>-quality.md`:
- datetime
- mode: correctness
- findings count
- action taken
- list of affected paths
