# Forge Agent Convention

## Separation of prompts vs profiles
- **Boot prompts** live in `.forge/boot/*.md` and remain the canonical instruction bodies that describe how a role operates once selected.
- **Agent profiles** live in `.forge/agents/*.yaml` and encode the harness-facing defaults: which prompt to load, PCC slices to ship, write authority, tool allowlists, and envelope/workflow suggestions.

This split keeps prompts human-readable while allowing the harness to reason about permissions and context programmatically.

## File map
| Role | Boot prompt | Profile |
|---|---|---|
| Architect | `.forge/boot/architect.md` | `.forge/agents/architect.yaml` |
| Orchestrator | `.forge/boot/orchestrator.md` | `.forge/agents/orchestrator.yaml` |
| Worker | `.forge/boot/worker.md` | `.forge/agents/worker.yaml` |
| Reviewer | `.forge/boot/reviewer.md` | `.forge/agents/reviewer.yaml` |

## Profile structure
Each profile is a YAML document with the following fields:

```yaml
version: 1                   # schema version
name: orchestrator           # unique identifier
description: Primary writer…
boot_prompt: .forge/boot/orchestrator.md
defaults:
  pcc:
    includes:                # PCC files shipped by default
      - .forge/pcc/global/00_project.md
  write_scope:               # filesystem authority
    mode: allow|deny
    paths:
      - docs/**
  tool_allowlist:            # harness tool names that may execute
    - shell
    - apply_patch
  workflows:
    suggested:               # CLI workflow names to surface first
      - workflow-new-feature
  envelopes:
    suggested:               # Preferred envelope slugs (Task 2)
      - orchestrator/default
```

Profiles are intentionally declarative; enforcement stays in the harness. New fields can be added in minor versions as envelope support matures.

## Profile catalogue
- **Architect** — doc/decision focused agent with a new boot prompt (`.forge/boot/architect.md`). Write scope is limited to `docs/**`, `artifacts/plan/**`, `artifacts/knowledge/**`, and `.forge/pcc/global/05_decisions.md`, making it safe for long-lived planning sessions.
- **Orchestrator** — single-writer implementation driver. Has full repo write access and the broadest tool list.
- **Worker** — read-only bounded helper. Generates analyses, reports, or diffs but never writes to disk.
- **Reviewer** — read-only investigator optimized for reviews and scans, with PCC slices biased toward decisions and conventions.

## Selecting an agent profile
1. Determine the session objective.
2. Run `/agent use <name>` (e.g., `/agent use architect`). The command normalizes the slug, validates `.forge/agents/<name>.yaml`, displays the YAML, and prompts you to echo a structured summary (boot prompt, PCC includes, write scope, etc.). It also lists available profiles if a typo is detected (e.g., `architecht`).
3. Follow the displayed constraints for the remainder of the session. Valid slugs today: `architect`, `orchestrator`, `worker`, `reviewer`.
4. Run the command again whenever you need to switch personas mid-session; no repository state is written.
5. To start a brand-new session directly in a profile, launch the CLI with `FORGE_AGENT_PROFILE=<name> scripts/forge-cli.sh --repo <path>` or pass `--agent <name>` when invoking `scripts/forge-cli.sh`. Each CLI invocation is isolated, so concurrent sessions can choose different profiles safely.

## When to switch agents
- Use **Architect** for proofs-of-concept, ADRs, roadmap shaping, PCC updates, and any doc-only planning.
- Use **Orchestrator** for implementation, integration, and any changes that modify repo state.
- Use **Worker** when dispatching bounded read-only analyses (dead-code scans, dependency audits, etc.).
- Use **Reviewer** for structured reviews, verification gates, and targeted investigations.

Keep `docs/AGENTS.md` updated whenever new profiles are introduced so contributors know where prompts live, how permissions are set, and how to opt into the right persona.
