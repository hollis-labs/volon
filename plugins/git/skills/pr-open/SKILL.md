---
name: pr-open
description: Create a PR for current branch if PR mode allows (uses gh if available).
argument-hint: ""title" [body=auto]"
disable-model-invocation: true
---

# PR Open

Open a pull request from the current branch using the `gh` CLI if available,
or print manual steps if not. Reads all policy from `volon.yaml`.

---

## Arguments

- `$0` (optional): PR title. If absent, derived from current branch name.
- `body=auto|<text>` (optional, default: `auto`): PR body. `auto` generates
  from branch name and recent commits. A string value is used verbatim.

---

## Step 1 — Read config

Run: !`cat volon.yaml 2>/dev/null || cat .volon/volon.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract and note:
- `git.pr_mode` → `off`, `optional`, or `required`
- `git.pr.base_branch` → default: `auto`
- `git.pr.title_prefix` → default: `volon:`
- `git.pr.body_template` → default: `.volon/templates/pr-body.md`

If `git.pr_mode` is `off`:
output `PR creation is disabled (git.pr_mode: off in volon.yaml).` and stop.

---

## Step 2 — Guard: validate environment

Run: !`git rev-parse --is-inside-work-tree 2>/dev/null`

If output is not `true`:
output `ERROR: not a git repository. Run from within a git repo.` and stop.

Run: !`git rev-parse --abbrev-ref HEAD 2>/dev/null`

Capture as `current_branch`. If result is `HEAD` (detached state):
output `ERROR: HEAD is detached. Check out a branch before opening a PR.` and stop.

---

## Step 3 — Confirm branch follows configured prefix

Check if `current_branch` starts with `git.branch_prefix` (e.g., `volon/`).

If it does not:
output `WARN: current branch "<current_branch>" does not match prefix "<branch_prefix>". Proceeding anyway.`
(Do not stop — the user may be opening a PR from a non-Volon branch intentionally.)

---

## Step 4 — Resolve PR base branch

If `git.pr.base_branch` is `auto`:
- Run: !`git rev-parse --verify main 2>/dev/null`
  If exit 0: use `main`.
- Else run: !`git rev-parse --verify master 2>/dev/null`
  If exit 0: use `master`.
- Else: output `ERROR: cannot determine PR base branch — set git.pr.base_branch in volon.yaml.` and stop.

Otherwise: use the configured value directly as `base_branch`.

---

## Step 5 — Resolve PR title and body

**Title:**
If `$0` is provided: `pr_title` = `<title_prefix> <$0>`
Otherwise: derive from branch name by stripping the prefix and replacing hyphens with spaces.

Run: !`echo "<current_branch>" | sed "s|<branch_prefix>||" | sed 's/-/ /g'`

`pr_title` = `<title_prefix> <derived title>`

**Body:**
If `body=auto` (default):
- Run: !`git log --oneline <base_branch>..HEAD 2>/dev/null`
  Collect recent commit messages. Format as a bullet list.
- If body template file exists (check `git.pr.body_template`):
  Run: !`cat <body_template> 2>/dev/null`
  Prepend template content before commit list.
- Compose final body.

If `body=<text>`: use the provided string verbatim.

---

## Step 6 — Check `gh` CLI availability

Run: !`which gh 2>/dev/null`

### 6a — gh available: create PR

Run: !`gh pr create --title "<pr_title>" --base "<base_branch>" --body "<pr_body>"`

If the command succeeds: capture and output the PR URL.

If it fails (e.g., already exists, auth issue):
output `ERROR: gh pr create failed — see output above.`
Then print manual steps (Step 6b) as a fallback reference.

### 6b — gh not available: print manual steps

```
gh CLI not found. To open this PR manually:

1. Push the branch:
   git push -u origin <current_branch>

2. Open a PR on GitHub/GitLab pointing:
   head: <current_branch>
   base: <base_branch>
   title: <pr_title>
```

---

## Step 7 — Output

If PR was created via `gh`:
```
PR opened: <pr_url>
Base:  <base_branch>
Head:  <current_branch>
```

If manual steps were printed:
```
Manual PR steps printed above.
Base:  <base_branch>
Head:  <current_branch>
Title: <pr_title>
```

---

## Invariants

- Never create a PR if `git.pr_mode: off`.
- Never push the branch — that is the user's responsibility (or `git push` in manual steps).
- Never modify any files in the repository.
- All config values read from `volon.yaml` in Step 1 — no hardcoded defaults in commands.
- If `gh` is unavailable, always provide manual steps rather than failing silently.
