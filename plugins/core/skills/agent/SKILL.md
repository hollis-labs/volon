---
name: agent
description: Inspect an agent profile and instruct the current session to operate under it.
argument-hint: "use <profile-name>"
disable-model-invocation: true
---

# agent

`/agent use architect` (or worker/orchestrator/reviewer) is a **session-scoped switch**.
It surfaces the selected profile’s guardrails so the current agent can immediately
operate with the right boot prompt, PCC slices, write scope, and tooling limits.
No repository files are modified, which keeps concurrent sessions independent.

---

## Arguments

- Subcommand `use` (required)
- `profile-name` (required): slug referencing `.volon/agents/<slug>.yaml`

Example: `/agent use architect`

Valid profiles today: `architect`, `orchestrator`, `worker`, `reviewer`.

---

## Step 1 — Guard + normalize slug

1. Ensure the first argument is `use`. Otherwise output `Usage: /agent use <profile>` and stop.
2. Ensure a profile name is provided. If missing: same usage message, stop.
3. Capture the raw argument and normalize it to a slug:

```
PROFILE_RAW="$2"
PROFILE_SLUG="$(printf '%s' "$PROFILE_RAW" | tr '[:upper:]' '[:lower:]' | sed 's/[^a-z0-9]/-/g' | sed 's/-\+/-/g' | sed 's/^-//; s/-$//')"
PROFILE_FILE=".volon/agents/${PROFILE_SLUG}.yaml"
```

4. Verify the profile file exists:

```
if [[ ! -f "$PROFILE_FILE" ]]; then
  echo "ERROR: profile ${PROFILE_SLUG} not found at ${PROFILE_FILE}"
  echo "Available profiles:"
  ls .volon/agents/*.yaml 2>/dev/null | sed 's#.*/##; s/\.yaml$//'
  exit 1
fi
```

This catches typos such as `architecht` and lists valid options.

---

## Step 2 — Read the profile and print a summary

1. Show the YAML so you can reference it quickly:

```
cat "$PROFILE_FILE"
```

2. From the file, capture:
   - `name`
   - `description`
   - `boot_prompt`
   - `.defaults.pcc.includes`
   - `.defaults.write_scope.mode` and `.defaults.write_scope.paths`
   - `.defaults.tool_allowlist`
   - `.defaults.workflows.suggested`
   - `.defaults.envelopes.suggested`

3. Print a summary block:

```
Active agent: <name> (slug: $PROFILE_SLUG)
Description : <description>
Boot prompt : <boot_prompt>

## Defaults
- PCC includes     : <comma-separated list or "none">
- Write scope      : <mode> → <paths or "no direct writes">
- Tool allowlist   : <list or "none">
- Workflows        : <list or "none">
- Envelopes        : <list or "none">
```

4. Remind yourself (and the user) of the next actions:
   - Read `<boot_prompt>` immediately and follow it for this session.
   - To launch a CLI session directly in this profile, run `FORGE_AGENT_PROFILE=$PROFILE_SLUG scripts/volon-cli.sh --repo <path>` or pass `--agent $PROFILE_SLUG`.
   - This command does not persist anything; rerun `/agent use <profile>` whenever you switch personas.

DONE
