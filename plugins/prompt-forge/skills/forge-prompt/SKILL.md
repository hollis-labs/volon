---
name: forge-prompt
aliases: ["prompt-forge", "make-prompt"]
description: Generate a copy/paste-ready Forge prompt from a natural language intent. Infers the correct template (inception/task/ideation/planning/patch) and fills all guardrails automatically.
argument-hint: '"<intent>" [--mode inception|task|ideation|planning|patch] [--repo <path>] [--constraints "<text>"] [--deliverables "<text>"]'
disable-model-invocation: true
---

# forge-prompt

Generate a single, copy/paste-ready Forge prompt from a natural language intent description.
Applies the correct template, reads repo context, fills all placeholders, and enforces
non-negotiable guardrails (single writer, no destructive ops, bounded sub-agents, DONE token).

**Aliases:** `/prompt-forge`, `/make-prompt` — all invoke this skill identically.

---

## Arguments

| Argument | Required | Description |
|---|---|---|
| `$0` | Yes | Natural language intent — what the prompt should accomplish |
| `--mode` | No | Override template: `inception` \| `task` \| `ideation` \| `planning` \| `patch` |
| `--repo <path>` | No | Target repo path or name. Default: current directory |
| `--constraints "<text>"` | No | Additional scope/guardrail text to inject into the prompt |
| `--deliverables "<text>"` | No | Explicit expected outputs to inject into the prompt |

---

## Step 1 — Preflight

Run: !`cat forge.yaml 2>/dev/null || cat .forge/forge.yaml 2>/dev/null || echo "NO_CONFIG"`

If config found, extract and note:
- `project.name` → REPO_NAME (fallback: "this repo")
- `bootstrap.completion_token` → DONE_TOKEN (fallback: `DONE`)
- `prompt_generator.commit_policy` → COMMIT_POLICY (fallback: `iteration`)
- `prompt_generator.pr_mode` → PR_MODE (fallback: `optional`)
- `prompt_generator.subagents_enabled` → SUBAGENTS_ENABLED (fallback: `false`)
- `prompt_generator.verbosity` → VERBOSITY (fallback: `standard`)

Run: !`date +%Y-%m-%d`
Capture as DATE.

If `$0` is absent: output `ERROR: intent is required. Usage: /forge-prompt "<intent>"` and stop.

---

## Step 2 — Parse arguments

From the invocation, capture:
- INTENT = `$0` (the full quoted intent string)
- MODE = value after `--mode` if present, else `infer`
- REPO = value after `--repo` if present, else REPO_NAME from config, else "this repo"
- CONSTRAINTS = value after `--constraints` if present, else `""`
- DELIVERABLES = value after `--deliverables` if present, else `""`

---

## Step 3 — Infer mode (only if MODE == "infer")

Apply these rules in order — first match wins:

| If INTENT contains any of these words/phrases... | → MODE |
|---|---|
| "inception", "install forge", "boot loop", "run loop", "iter", "run inception" | `inception` |
| "ideation", "brainstorm", "brainstorming", "names", "naming", "vibe", "explore ideas", "ideas for" | `ideation` |
| "plan", "prd", "spec", "design", "roadmap", "requirements", "backlog planning" | `planning` |
| "patch", "apply zip", "apply patch", "hotfix", "apply diff", "apply fix" | `patch` |
| (no match) | `task` |

Output one line: `[forge-prompt] mode inferred: <MODE>`

---

## Step 4 — Load template

Read the template file corresponding to the selected mode from the `templates/` directory
(co-located with this SKILL.md):

| MODE | Template file |
|---|---|
| `inception` | `templates/inception.md` |
| `task` | `templates/task.md` |
| `ideation` | `templates/ideation.md` |
| `planning` | `templates/planning.md` |
| `patch` | `templates/patch.md` |

Read the full contents of the selected template file.

---

## Step 5 — Fill placeholders

The template contains `{{PLACEHOLDER}}` tokens. Replace each with its resolved value:

| Token | Resolved value |
|---|---|
| `{{INTENT}}` | INTENT |
| `{{REPO}}` | REPO |
| `{{DATE}}` | DATE |
| `{{CONSTRAINTS}}` | CONSTRAINTS if non-empty, else `"(none — use Forge defaults)"` |
| `{{DELIVERABLES}}` | DELIVERABLES if non-empty, else the template's own default deliverables hint |
| `{{COMMIT_POLICY}}` | COMMIT_POLICY |
| `{{PR_MODE}}` | PR_MODE |
| `{{SUBAGENTS_NOTE}}` | If SUBAGENTS_ENABLED=false: `"Sub-agents: disabled. This session is the single writer."` else `"Sub-agents: read-only, max 2 per run. No recursive spawning."` |
| `{{DONE_TOKEN}}` | DONE_TOKEN |

After replacement, no `{{...}}` tokens should remain in the output. If any remain (unresolved), replace them with their token name in angle brackets, e.g., `<INTENT>`.

---

## Step 6 — Apply guardrail check

Scan the filled prompt text. If any of the following patterns appear, remove or rewrite them:

- `rm -rf` without an explicit `confirm:` qualifier → add qualifier or remove
- `git push --force` or `git push -f` without `confirm:` → add qualifier or remove
- `git reset --hard` without `confirm:` → add qualifier or remove
- `spawn unlimited sub-agents` or `unlimited delegation` → rewrite to bounded form
- `write outside repo root` or `write to /` → remove entirely

If CONSTRAINTS is non-empty, append its content verbatim to the generated prompt's
`## Constraints` section (on a new bullet line).

---

## Step 7 — Output

Print the following, exactly in this format:

```
---
Generated Forge Prompt
Mode: <MODE>
Intent: <INTENT>
Date: <DATE>
---
Copy/paste the prompt below into Claude Code or your agent session:

---PROMPT-START---
<filled template content>
---PROMPT-END---
```

Then print one final line:
`[forge-prompt] done — <MODE> prompt generated for: <INTENT>`

Do NOT output `DONE`. This skill is a generator, not a workflow runner.

---

## Invariants

- Never produce a prompt that omits the single-writer rule.
- Never produce a prompt that omits explicit stop conditions (DONE token).
- Never produce a prompt that allows `git push --force` without a `confirm:` note.
- Always include a `## Guardrails` section in the output prompt.
- The output prompt must be self-contained: no unresolved file references.
- When REPO context cannot be determined, use safe defaults (repo="this repo", commit_policy=iteration).
- The filled prompt must be valid Markdown (no broken fencing, no unclosed backtick blocks).
