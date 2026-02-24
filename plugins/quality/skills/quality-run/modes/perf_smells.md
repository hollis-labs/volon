# Perf Smells Mode — Steps 2P–5P

## Step 2P — Scan for perf_smells signals

Run the following four signal checks and collect all findings into a list.
Each finding is: `{ signal, path, detail }`.

### Signal A — Oversized SKILL.md files

Run: !`find plugins/ -name "SKILL.md" 2>/dev/null || true`

For each SKILL.md found, count lines:
Run: !`wc -l <file> 2>/dev/null || true`

If line count exceeds 300:
→ Add finding: `{ signal: "oversized_skill", path: "<file>", detail: "<N> lines — exceeds 300-line guideline; consider splitting into sub-protocols." }`

### Signal B — Unbounded find/grep in SKILL.md steps

Run: !`grep -n "!\`find \." <file> 2>/dev/null || true` for each SKILL.md.

A `find .` command (scanning from repo root) without a path-scoping argument is unbounded.
Distinct from scoped finds like `find plugins/`, `find .forge/tasks/`, etc.

For each match of `!\`find .` in a SKILL.md step:
→ Add finding: `{ signal: "unbounded_scan", path: "<file>:<line>", detail: "find . without path scope — may scan entire repo; consider scoping to a subdirectory." }`

### Signal C — PCC files approaching word limit

Run: !`wc -w .forge/pcc/*.md 2>/dev/null || true`

For each PCC file with word count ≥ 350:
→ Add finding: `{ signal: "pcc_near_word_limit", path: "<file>", detail: "<N> words — within 50 words of 400-word cap; trim before overflow." }`

For each PCC file with word count > 400:
→ Add finding: `{ signal: "pcc_over_word_limit", path: "<file>", detail: "<N> words — exceeds 400-word cap; must trim immediately." }`

### Signal D — Run log accumulation

Run: !`ls <observability.log_dir>/ 2>/dev/null | wc -l | tr -d ' ' || echo "0"`

If the count exceeds 20:
→ Add finding: `{ signal: "log_accumulation", path: "<log_dir>/", detail: "<N> log files — exceeds 20-file guideline; consider archiving older logs." }`

---

## Step 3P — Collate perf_smells findings

Build a findings table:

| # | Signal | Path | Detail |
|---|---|---|---|
| 1 | <signal> | <path> | <detail> |
| ... | | | |

If the findings list is empty:
→ Output `No perf smell issues found.` and go to Step 5P.

---

## Step 4P — Apply on_issue action

For each finding in the findings table, apply the effective action:

### action = log

Print the finding to output. Do not create tasks. No file changes.

### action = create_task

For each finding, create a task using the task-create protocol:
- `title`: `"Quality: <signal> — <path>"`
- `priority`: `C`
- `tags`: `[quality, perf_smells]`
- `context`: current context from config (default: `dev`)
- `description`: the finding detail

Run: !`ls <storage.files.root>/TASK-*.md 2>/dev/null | wc -l | tr -d ' ' || echo "0"`

Use this count to derive the next sequence number (zero-padded 3 digits).
Write each new task file per the standard task format.

### action = auto_fix_pr

v0.1: `auto_fix_pr` is not implemented for perf_smells mode.
→ Log a note: `WARN: auto_fix_pr not implemented in v0.1 — falling back to create_task.`
→ Apply `create_task` action instead.

---

## Step 5P — Output summary

Print:

```
quality-run complete (mode: perf_smells)
Findings: <N>
Action taken: <effective_action>
<findings table or "No perf smell issues found.">
```

If `observability.write_run_log` is true: append a brief entry to
`<log_dir>/run-<YYYYMMDD>-<HHMM>-quality.md`:
- datetime
- mode: perf_smells
- findings count
- action taken
- list of affected paths
