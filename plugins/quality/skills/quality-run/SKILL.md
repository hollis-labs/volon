---
name: quality-run
description: Run configured quality modes and take configured on_issue actions (log, create_task, auto_fix_pr).
argument-hint: "[mode=auto|dead_code|security|correctness|perf_smells] [action=auto|log|create_task|auto_fix_pr]"
disable-model-invocation: true
standalone: true
model-tier: summarize
---

# Quality Run

Scan the repo for quality issues using the modes configured in `forge.yaml`.
For each finding, apply the configured `on_issue.default_action`.
v0.1 implements **`dead_code`**, **`security`**, **`correctness`**, and **`perf_smells`** modes.
Mode protocols live in `modes/<mode>.md` (read and execute the relevant file for each mode).

---

## Inputs

- `mode`: `$0` (default: `auto` — runs all enabled modes in `forge.yaml`)
- `action`: `$1` (default: `auto` — uses `quality.on_issue.default_action` from config)

---

## Step 1 — Read config and guard

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

If `NO_CONFIG`: use defaults and continue.

Extract and note:
- `quality.modes` → list of enabled modes (e.g. `[dead_code, security, correctness, perf_smells]`)
- `quality.on_issue.default_action` → `log` | `create_task` | `auto_fix_pr` (default: `log`)
- `storage.files.root` → default: `.forge/tasks`
- `observability.log_dir` → default: `.forge/logs`

If `mode` arg is `auto`: run all modes listed in `quality.modes`.
If `mode` arg is a specific mode name: run only that mode.

Mode dispatch — read the file listed and execute the steps within it:

| Mode | File to read |
|---|---|
| `dead_code` | `modes/dead_code.md` — execute Steps 2–5 |
| `security` | `modes/security.md` — execute Steps 2S–5S |
| `correctness` | `modes/correctness.md` — execute Steps 2C–5C |
| `perf_smells` | `modes/perf_smells.md` — execute Steps 2P–5P |

If the resolved mode list contains no implemented modes: output `No implemented quality modes to run.` and stop.

Resolve effective action:
- If `action` arg is not `auto`: use it (overrides config).
- Otherwise: use `quality.on_issue.default_action`.

Note: `auto_fix_pr` is not implemented in v0.1 — falls back to `create_task`.

---

## Invariants

- Never modify any source file — this is a read-only scan. Only task files may be created.
- Do not create duplicate tasks for the same `path` if an open task for it already exists.
  Run: !`grep -r "<path>" <storage.files.root>/ 2>/dev/null || true`
  If a match is found with `status: todo`, skip task creation for that path.
- `auto_fix_pr` always degrades to `create_task` in v0.1, **except for security mode** where it degrades to `log`.
- Security findings must never be auto-fixed — always require human review.
- Emit no output for signals that produce zero findings.
- Each mode's findings scope to that mode's signals only.
- Correctness findings at priority B; perf_smells findings at priority C (advisory).
- `pcc_over_word_limit` findings escalate to priority B when creating tasks.
