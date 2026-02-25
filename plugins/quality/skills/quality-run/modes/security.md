# Security Mode — Steps 2S–5S

## Step 2S — Scan for security signals

Run the following four signal checks. Each finding is: `{ signal, path, detail }`.

### Signal A — Hardcoded secret patterns

Run: !`grep -rn --include="*.md" --include="*.yaml" --include="*.yml" --include="*.json" --include="*.env" -E "(password|secret|api_key|token|private_key)\s*[:=]\s*['\"]?[^$\{\(<\s]{6,}" . 2>/dev/null | grep -v "Binary\|\.volon/logs\|#" | grep -v "your_\|_here\|_value\|placeholder\|example" | grep -v ":\s*\`[A-Za-z_][A-Za-z0-9_]*\`" || true`

Filter explanations:
- `grep -v "#"` — exclude shell comment lines
- `grep -v "your_\|_here\|_value\|placeholder\|example"` — exclude common placeholder prose patterns
- `grep -v ":\s*\`[A-Za-z_][A-Za-z0-9_]*\`"` — exclude lines where the value is a backtick-wrapped identifier (markdown inline code; documentation examples)

For each match remaining after all filters:
→ Add finding: `{ signal: "hardcoded_secret", path: "<file>:<line>", detail: "Possible secret assignment: <matched-key>." }`

### Signal B — Volon config credential exposure

Run: !`cat volon.yaml 2>/dev/null || cat .volon/volon.yaml 2>/dev/null || true`

Inspect `storage.nanite.vault` value: if it looks like a real credential (length > 8, not `"NaniteVaultName"` or `auto`) flag it.
→ Add finding: `{ signal: "config_credential_exposure", path: "volon.yaml", detail: "storage.nanite.vault may contain a real credential." }`

### Signal C — Sensitive files present

Run: !`find . -maxdepth 3 \( -name ".env" -o -name "*.pem" -o -name "*.key" -o -name "credentials.*" \) 2>/dev/null | grep -v ".git" || true`

For each match found:
- Check `.gitignore` coverage: !`cat .gitignore 2>/dev/null | grep -F "<filename>" || true`
- If NOT covered by `.gitignore`:
  → Add finding: `{ signal: "sensitive_file_unignored", path: "<path>", detail: "Sensitive file not covered by .gitignore." }`

### Signal D — redact_secrets disabled

Run: !`grep "redact_secrets" volon.yaml 2>/dev/null || grep "redact_secrets" .volon/volon.yaml 2>/dev/null || true`

If the result contains `redact_secrets: false`:
→ Add finding: `{ signal: "redact_secrets_disabled", path: "volon.yaml", detail: "observability.redact_secrets is false — secrets may appear in run logs." }`

If `redact_secrets` key is absent entirely:
→ Emit `INFO: redact_secrets not set — defaults to false in v0.1. Consider setting to true.`

---

## Step 3S — Collate security findings

Build a findings table (same format as dead_code Step 3).
If findings list is empty: → Output `No security issues found.` and go to Step 5S.

---

## Step 4S — Apply on_issue action

Follow the same action logic as dead_code Step 4, with these overrides:

- **`action = log`**: always valid — print findings only.
- **`action = create_task`**: create tasks with `tags: [quality, security]`, `priority: A`
  (security findings are elevated to A regardless of config default).
- **`action = auto_fix_pr`**: security findings must **never** be auto-fixed.
  → Override to `log` only; emit: `WARN: auto_fix_pr suppressed for security findings — manual review required.`

---

## Step 5S — Output summary

Print:

```
quality-run complete (mode: security)
Findings: <N>
Action taken: <effective_action>
<findings table or "No security issues found.">
```

If `observability.write_run_log` is true: append entry to `<log_dir>/run-<YYYYMMDD>-<HHMM>-quality.md`
(same format as dead_code; append to same file if both modes ran in one session).
