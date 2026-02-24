---
name: task-create
description: Create a task in the configured backend (files or nanite).
argument-hint: ""title" [priority=A|B|C] [tags=comma,list] [context=dev] [project=auto]"
disable-model-invocation: true
---

# Task Create

Create a task in the backend defined by `storage.backend` in `forge.yaml`.
- `files` (default): writes a markdown file at `<storage.files.root>/<id>.md`
- `nanite`: pushes an item to the configured nanite vault via the `nanite` CLI

---

## Arguments

- `$0` (required): task title — quoted string, e.g. `"Implement worktree skill"`
- `priority=A|B|C` (optional, default: `B`)
- `tags=comma,separated` (optional, default: `[]`)
- `context=<string>` (optional, default: `dev`)
- `project=<string>` (optional, default: `auto`)

---

## Step 1 — Read config

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

Extract and note:
- `storage.backend` → default: `files`
- `storage.files.root` → default: `.forge/tasks`
- `storage.nanite.vault` → default: `null`
- `storage.nanite.tag_prefix` → default: `forge/`
- `storage.nanite.default_priority` → default: `B`
- `project.name` → used when `project=auto`

---

## Step 2 — Resolve arguments

- **title**: `$0` — strip surrounding quotes if present. If absent:
  output `ERROR: title is required.` and stop.
- **priority**: from `priority=` argument. Default: `storage.nanite.default_priority`
  or `B`. Must be `A`, `B`, or `C`. If invalid: output `ERROR: priority must be A, B, or C.` and stop.
- **tags**: from `tags=` argument, split on comma. Default `[]`.
- **context**: from `context=` argument. Default `dev`.
- **project**: from `project=` argument. If `auto` or absent: use `project.name`
  from config. If still unresolved: use current directory name.

---

## Step 3 — Dispatch on backend

### If `storage.backend` is `nanite`

**Resolve vault ID** (if `storage.nanite.vault` is non-null):

Run: !`nanite vaults 2>/dev/null`

Parse JSON. Find the vault entry where `name` matches `storage.nanite.vault`.
- If found: capture `id` as `VAULT_ID`. The `--vault <VAULT_ID>` flag will be appended.
- If not found: output `ERROR: nanite vault '<name>' not found. Run 'nanite vaults' to list available vaults.` and stop.
- If `storage.nanite.vault` is `null`: omit `--vault` flag entirely (nanite uses its active vault).

Build the `nanite push` command:

```
nanite push "<title>" \
  --source forge-agent \
  --type todo \
  --priority <priority> \
  --tags <tag_prefix>task,<tags-comma-separated-if-any> \
  --projects <project> \
  --contexts <context> \
  [--vault <VAULT_ID>]   ← omit if vault is null
```

Run: !`nanite push "<title>" --source forge-agent --type todo --priority <priority> --tags <tag_prefix>task<,tags> --projects <project> --contexts <context> [--vault <VAULT_ID>] 2>/dev/null`

Parse JSON response:
- If `ok: false`: output `ERROR: nanite push failed — <error>.` and stop.
- If `ok: true`: capture `data.id` as `NANITE_ID`.

Output:
```
Created: NANITE-<NANITE_ID>
Vault:   <vault or "(active)">
```

Stop — do not proceed to file-backend steps.

---

### If `storage.backend` is `files` (default)

Proceed to Steps 4–6.

---

## Step 4 — Generate file task ID

Run: !`date +%Y%m%d`

Capture result as `YYYYMMDD`.

Run: !`ls <storage.files.root>/TASK-<YYYYMMDD>-*.md 2>/dev/null | wc -l | tr -d ' '`

Capture result as `N`. Next sequence = `N + 1`. Zero-pad to 3 digits → `NNN`.

Task ID = `TASK-<YYYYMMDD>-<NNN>`

If the resulting file path already exists: increment NNN by 1 and check again.
Retry once only. If still collides: output `ERROR: ID collision — try again.` and stop.

---

## Step 5 — Create tasks directory and write task file

Run: !`mkdir -p <storage.files.root>`

Create file at `<storage.files.root>/<id>.md`:

```
---
id: <id>
title: "<title>"
status: todo
priority: <priority>
project: <project>
tags: [<tags as comma-separated YAML inline list, or empty []>]
context: <context>
created_at: <YYYY-MM-DD>
updated_at: <YYYY-MM-DD>
---

## Description

## Acceptance

## Verification

## Paths

## Updates
```

`created_at` and `updated_at` both equal today's date.
`status` is always `todo` on creation.

---

## Step 6 — Output

```
Created: <id>
Path: <storage.files.root>/<id>.md
```

---

## Invariants

- Never overwrite an existing task file (files backend).
- `status` is always `todo` on creation.
- `--vault` requires the vault ID (integer/hex), not the vault name. Always resolve name→ID via `nanite vaults` before passing `--vault`.
- If `storage.nanite.vault` is `null`, omit `--vault` flag entirely — nanite uses its active vault.
- All frontmatter fields from the task model must be present, even if empty (files backend).
