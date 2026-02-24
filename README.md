# Forge — Agentic Development System (v0.5)

A **project-agnostic** set of Claude Code plugins, skills, and workflow contracts for high-throughput, low-drift agentic development.

Forge automates the full development lifecycle:
**Ideation → Requirements → PRD → Spec → Plan → Tasks → Execute (PR/worktree) → QA/Docs → Investigate**

---

## Agent entry points

**Starting a Forge session (agents):** See [`CLAUDE.md`](CLAUDE.md) — contains the boot sequence and copy-paste Inception Run Prompt.

**Authoritative boot doc:** [`.forge/agent-boot.md`](.forge/agent-boot.md) — ground truth sources, core rules, reference map.

**Current state:** [`.forge/bootstrap.md`](.forge/bootstrap.md) — iteration, active tasks, next actions.

---

## What's in this repo

- `forge.yaml` — system configuration
- `plugins/` — Claude Code plugins (8 plugins, 22 skills)
- `docs/` — system specs (config, PCC, workflow contracts, task model, orchestrator, inception)
- `.forge/` — runtime state (bootstrap, PCC, tasks, logs, boot pack)
- `examples/` — sample configs and workflow traces
- `user-docs-mini-site/` — static user-facing documentation site

## Key capabilities (v0.5 — all epics complete)

- **Project Context Cache (PCC)** — low-token, repo-local agent context
- **Inception workflow** (`/workflow-inception`) — canonical self-building loop
- **Full workflow lifecycle** — new-feature, docs-review, app-investigation, extension authoring
- **Workflow authoring** — create/edit/clone/deprecate workflows from within Forge
- **Task + backlog management** — nanite backend, priority-ordered execution
- **Git integration** — worktrees, PR mode, commit-per-task or commit-per-iteration
- **Pause/resume** — deterministic state externalization; any session can restart cleanly
- **Quality scans** — correctness, security, dead code, performance smells
- **Bootstrap boundaries** — every iteration ends with a clean, restartable state file

## Quick start (loading plugins)

Run the helper script (keeps Forge up to date, then launches Claude with every plugin registered):

```bash
./scripts/forge-cli.sh --repo /path/to/target-repo
```

- Set `FORGE_NO_SYNC=1` to skip the git fetch/pull step.
- Override the Claude binary with `CLAUDE_BIN=/path/to/claude`.
- The `--repo` flag is handled by the script (it `cd`s into that path before launching). Pass additional Claude args after `--`, e.g. `./scripts/forge-cli.sh --repo /path -- --model claude-sonnet-4-6`.

Manual invocation (equivalent, shown for reference):

```bash
claude \
  --plugin-dir /path/to/forge/plugins/core \
  --plugin-dir /path/to/forge/plugins/workflows \
  --plugin-dir /path/to/forge/plugins/git \
  --plugin-dir /path/to/forge/plugins/tasks-nanite \
  --plugin-dir /path/to/forge/plugins/docsmith \
  --plugin-dir /path/to/forge/plugins/quality \
  --plugin-dir /path/to/forge/plugins/backlog \
  --plugin-dir /path/to/forge/plugins/workflow-author \
  --plugin-dir /path/to/forge/plugins/prompt-forge
```

Then in a target repo with `forge.yaml`:

```
/workflow-inception          # run the canonical self-building loop
/workflow-new-feature "..."  # full ideation → tasks pipeline
/workflow-docs-review        # PCC-grounded documentation scan
/quality-run                 # correctness / security / dead-code / perf scan
```

## Docs

- [`docs/13_inception-workflow.md`](docs/13_inception-workflow.md) — the canonical loop
- [`docs/08_orchestrator.md`](docs/08_orchestrator.md) — Orchestrator mode
- [`docs/01_config.md`](docs/01_config.md) — configuration reference
- [`docs/03_workflow-contracts.md`](docs/03_workflow-contracts.md) — workflow execution contracts
- [`docs/04_task-model.md`](docs/04_task-model.md) — task model
- [`docs/09_commands.md`](docs/09_commands.md) — all available commands
