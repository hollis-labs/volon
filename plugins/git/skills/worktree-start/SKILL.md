---
name: worktree-start
description: Start a worktree for a task/feature with standard branch naming.
argument-hint: "<task-id-or-slug> [base=auto]"
disable-model-invocation: true
---

# Worktree Start

Create an isolated git worktree under the configured root with a consistently-named
branch. Transitions a task from `todo` to `doing` if a task ID is provided.

---

## Arguments

- `$0` (required): task ID (`TASK-YYYYMMDD-NNN`) or plain slug string
- `base=<branch>` (optional, default: `auto`): base branch for the new worktree

---

## Step 1 — Read config

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract and note:
- `git.use_worktrees` → if `false`: go to **E2** and stop
- `git.worktree_root` → default: `.worktrees`
- `git.branch_prefix` → default: `forge/`
- `storage.files.root` → default: `.forge/tasks`

---

## Step 2 — Guard: validate environment

If `$0` is absent or empty: output `ERROR: provide a task ID or slug.` and stop. **(E1)**

Run: !`git rev-parse --is-inside-work-tree 2>/dev/null`

If output is not `true`: output `ERROR: not a git repository. Run from within a git repo.` and stop. **(E3)**

If `git.use_worktrees` is `false`:
output `ERROR: worktrees disabled in forge.yaml. Create branch manually: git checkout -b <branch_prefix><slug>` and stop. **(E2)**

---

## Step 3 — Derive slug

### 3a — Detect input type
If `$0` matches the pattern `TASK-[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]-[0-9][0-9][0-9]`:
- This is a task ID. Read `<storage.files.root>/$0.md`.
- If file not found: output `WARN: task $0 not found in <storage.files.root> — continuing without task update.` **(W1)**
  Set `task_found=false`. Derive slug from `$0` directly (use the ID as slug).
- If file found: extract `title` from frontmatter. Set `task_found=true`.
  Slugify the title:
  Run: !`echo "<title>" | tr '[:upper:]' '[:lower:]' | sed 's/[^a-z0-9]/-/g' | sed 's/-\+/-/g' | sed 's/^-\|-$//g'`
  Capture result as `slug`.

### 3b — Plain slug
Otherwise: use `$0` directly as `slug`. Set `task_found=false`.

---

## Step 4 — Resolve base branch

If `base` argument is provided and not `auto`: use it directly as `base_branch`.

Otherwise:
Run: !`git rev-parse --abbrev-ref HEAD 2>/dev/null`

If result is `HEAD` (detached state):
- Run: !`git rev-parse --verify main 2>/dev/null`
  If exit 0: use `main` as `base_branch`.
- Else run: !`git rev-parse --verify master 2>/dev/null`
  If exit 0: use `master` as `base_branch`.
- Else: output `ERROR: cannot determine base branch — provide base=<branch> explicitly.` and stop.

Otherwise: use the result as `base_branch`.

---

## Step 5 — Check for conflicts

Construct:
- `branch_name` = `<branch_prefix><slug>` (e.g., `forge/worktree-start`)
- `worktree_path` = `<worktree_root>/<slug>` (e.g., `.worktrees/worktree-start`)

Run: !`git rev-parse --verify <branch_name> 2>/dev/null`

If exit 0 (branch exists):
output `ERROR: branch <branch_name> already exists. To resume, use: git worktree add <worktree_path> <branch_name>` and stop. **(E4)**

Run: !`ls <worktree_path> 2>/dev/null`

If exit 0 (path exists):
output `ERROR: worktree path <worktree_path> already exists. Remove it first: git worktree remove <worktree_path>` and stop. **(E5)**

---

## Step 6 — Create worktree

Run: !`mkdir -p <worktree_root>`

Run: !`git worktree add <worktree_path> -b <branch_name> <base_branch>`

If the command fails: output `ERROR: git worktree add failed — see output above.` and stop.

---

## Step 7 — Update task status (conditional)

Only if `task_found=true` (task file was found in Step 3a):

Run: !`date +%Y-%m-%d`

In `<storage.files.root>/$0.md`:
- Change `status: todo` → `status: doing`
- Change `updated_at: <old>` → `updated_at: <today>`
- Append to `## Updates` section:
  `- [<today>] worktree-start: branch <branch_name>`

Preserve all other frontmatter fields verbatim. Do not modify any body section
other than `## Updates`.

---

## Step 8 — Output

```
Worktree: <worktree_path>
Branch:   <branch_name>
```

If task was updated (Step 7):
```
Task <id>: todo → doing
```

---

## Invariants

- Never hardcode `.worktrees/` or `forge/` — always read from config.
- Never push to remote — that is `pr-open`'s responsibility.
- Never `cd` into the worktree — output the path only.
- If task update fails (file write error), output a WARN but do not roll back the worktree.
- All config values read from `forge.yaml` in Step 1 — no defaults embedded in commands.
