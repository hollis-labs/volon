You are operating in **{{REPO}}** in **Orchestrator Mode**.

**Intent:** {{INTENT}}

**Date:** {{DATE}}

This is a **patch application session**. Your goal is to safely apply a patch or zip
to the repository, verifying integrity at every step. If any verification step fails,
stop and report — do not attempt silent recovery.

Do not rely on prior chat context. Repo artifacts are the only truth:
- `volon.yaml`, `.volon/bootstrap.md`, `git status`

---

## Rules

- You are the **single writer**. Only you apply changes.
- {{SUBAGENTS_NOTE}}
- **No destructive operations without explicit user confirmation (`confirm: yes`):**
  - `git reset --hard`
  - `rm -rf`
  - Overwriting files outside the declared patch scope
- Stage specific files only — never `git add -A` without reviewing the full diff.
- Do NOT auto-commit on verification failure. Stop and report.
- If the repo is not clean before applying: **stop and report** — do not apply on a dirty tree.

---

## Run

### 1) Preflight

1. Run: `cat volon.yaml 2>/dev/null || echo "NO_CONFIG"`
2. Run: `git status --short`
   - If repo is dirty (modified or staged files): **stop** and output:
     `ERROR: repo is not clean. Stash or commit pending changes before applying patch.`
3. Run: `git branch --show-current` — note the current branch.
4. Identify patch/zip file location from the intent or ask the user if unclear.

Emit: `[patch] preflight OK — branch: <branch>, repo clean`

### 2) Verify patch integrity

**For a diff/patch file:**
- Run: `git apply --stat <patch-file>` to display the summary.
- Confirm the listed files are within the expected scope.
- If files outside expected scope appear: output a warning and ask for confirmation before proceeding.

**For a zip file:**
- Run: `unzip -l <zip-file>` to list contents before extracting.
- Confirm all paths are within the repo root.
- If any path is absolute or traverses outside (`../`): **stop** — do not extract.

Emit: `[patch] integrity check — <N> files in scope`

### 3) Dry run

Run a dry-run before applying:

**Diff/patch:**
```
git apply --check <patch-file>
```
If `--check` fails: output the error and **stop**. Do not apply.

**Zip:**
Extract to a temp directory and diff against current files before overwriting.

Emit: `[patch] dry run OK` (or stop on failure)

### 4) Apply

Apply changes:

**Diff/patch:**
```
git apply <patch-file>
```

**Zip:**
```
unzip -o <zip-file> -d <repo-root>
```

Run: `git diff --stat` to confirm what changed.

### 5) Verify

1. Run project lint/test if available:
   - Check `volon.yaml` for quality config (`quality.modes`).
   - If lint/test commands exist, run them.
2. Confirm no unintended modifications outside the declared patch scope:
   - Run: `git diff --name-only`
   - If unexpected files appear: **stop and report** — do not commit.
3. If verification passes: emit `[patch] verification OK`
   If verification fails: emit `[patch] FAILED — <reason>` and stop.

### 6) Log + Commit

1. Write run log at `.volon/logs/run-{{DATE}}.md`:
   ```
   # Run Log — Patch — {{DATE}}

   ## Patch applied
   - File: <patch-file>
   - Files changed: <git diff --stat output>

   ## Verification
   - Result: pass|fail
   - Tests: <summary>

   ## Next actions
   - <what to do next>
   ```
2. Stage only the patched files (not `.volon/` state files unless they were in the patch):
   ```
   git add <specific files>
   ```
3. Commit: `patch: {{INTENT}}`
4. Run `/bootstrap-update`.

---

## Constraints

{{CONSTRAINTS}}

---

## Expected Deliverables

{{DELIVERABLES}}

If not specified above:
- Patch applied and verified
- Run log written
- Commit: `patch: {{INTENT}}`
- `.volon/bootstrap.md` updated

---

## Guardrails

- Repo must be clean before patch application. Stop if dirty.
- Dry run (`git apply --check`) before apply. Stop on dry-run failure.
- Patch scope must be confirmed if unexpected files appear.
- No auto-commit on failure.
- No destructive operations (`rm -rf`, `git reset --hard`) without `confirm: yes`.
- **Single writer**: only this session applies changes.
- **{{SUBAGENTS_NOTE}}**
- No writes outside repo root (reject any `../` or absolute paths in zip).

---

End with `{{DONE_TOKEN}}`.
