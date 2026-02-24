# Forge — Agent Boot Instructions

This repo runs under **Forge**, an agentic development system for Claude Code.

## Boot into Forge Orchestrator mode

1. Read `.forge/agent-boot.md` — core rules, ground truth sources, role definitions
2. Read `.forge/bootstrap.md` — current iteration state and next actions
3. Read `.forge/boot/orchestrator.md` — Orchestrator-specific rules, boot confirmation format, transition signals
4. Emit the boot confirmation block (defined in `.forge/boot/orchestrator.md`)
5. Select the next work unit and begin the canonical loop

**You are the single writer.** Sub-agents (if used) are read-only.

---

## Inception Run Prompt (copy/paste to start a session)

```
You are operating in a Forge-managed repository in Orchestrator Mode.

Do not rely on prior chat context. Repo artifacts are the only truth:
- forge.yaml, .forge/bootstrap.md, .forge/pcc/, .forge/tasks/, .forge/backlog/, .forge/logs/

Rules:
- You are the single writer. Sub-agents (if enabled) are read-only.
- Work in small, verifiable steps. Prefer completing 1–3 tasks this run.
- Emit the boot confirmation block before taking any action.
- Emit transition signals at task-start, task-done, and iteration-finalize.

Run:
1) Preflight: read bootstrap/PCC/tasks; emit boot confirmation.
2) Execute: select next task (A>B>C, oldest first). doing → verify → done/blocked/paused.
3) Log: write run log entry.
4) Finalize: run /bootstrap-update.
5) Commit per policy (iteration if in iteration mode; task if isolated).

Stop after 3 tasks or when no todo remains. End with DONE.
```

---

## Key files

| File | Purpose |
|---|---|
| `forge.yaml` | System configuration |
| `.forge/bootstrap.md` | Current iteration state — start here |
| `.forge/agent-boot.md` | Full boot reference (rules, reference map) |
| `.forge/boot/orchestrator.md` | Orchestrator role addendum |
| `.forge/pcc/` | Project context cache |
| `docs/13_inception-workflow.md` | Inception workflow spec |
| `docs/08_orchestrator.md` | Orchestrator mode doc |

---

## Installing Forge in a new repo

1. Copy `forge.yaml` to the target repo root and customize for your project.
2. Copy the `plugins/` directory or load them via `--plugin-dir`.
3. Copy `.forge/agent-boot.md` and `.forge/boot/` to the target repo.
4. Copy this `CLAUDE.md` to the target repo root.
5. Run `/bootstrap-update` to generate `.forge/bootstrap.md` for the new repo.
