# Volon — Agent Boot Instructions

This repo runs under **Volon**, an agentic development system for Claude Code.

## Boot into Volon Orchestrator mode

1. Read `.volon/agent-boot.md` — core rules, ground truth sources, role definitions
2. Read `.volon/bootstrap.md` — current iteration state and next actions
3. Read `.volon/boot/orchestrator.md` — Orchestrator-specific rules, boot confirmation format, transition signals
4. Emit the boot confirmation block (defined in `.volon/boot/orchestrator.md`)
5. Select the next work unit and begin the canonical loop

**You are the single writer.** Sub-agents (if used) are read-only.

---

## Inception Run Prompt (copy/paste to start a session)

```
You are operating in a Volon-managed repository in Orchestrator Mode.

Do not rely on prior chat context. Repo artifacts are the only truth:
- volon.yaml, .volon/bootstrap.md, .volon/pcc/, .volon/tasks/, .volon/backlog/, .volon/logs/

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
| `volon.yaml` | System configuration |
| `.volon/bootstrap.md` | Current iteration state — start here |
| `.volon/agent-boot.md` | Full boot reference (rules, reference map) |
| `.volon/boot/orchestrator.md` | Orchestrator role addendum |
| `.volon/pcc/` | Project context cache |
| `docs/13_inception-workflow.md` | Inception workflow spec |
| `docs/08_orchestrator.md` | Orchestrator mode doc |

---

## Installing Volon in a new repo

1. Copy `volon.yaml` to the target repo root and customize for your project.
2. Copy the `plugins/` directory or load them via `--plugin-dir`.
3. Copy `.volon/agent-boot.md` and `.volon/boot/` to the target repo.
4. Copy this `CLAUDE.md` to the target repo root.
5. Run `/bootstrap-update` to generate `.volon/bootstrap.md` for the new repo.
