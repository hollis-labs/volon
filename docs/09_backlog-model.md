---
intent: system_doc
audience: humans
---

# Backlog Model — v0.4

## Purpose

The backlog is a **low-friction capture layer** between a raw idea and an active task.
Ideas should be cheap to record and never lost. Commitment to execution (promotion)
is a deliberate, separate step.

---

## Hierarchy

```
Idea / observation
  → Backlog item  (.forge/backlog/BACKLOG-YYYYMMDD-NNN.md)
      → Active task  (.forge/tasks/TASK-YYYYMMDD-NNN.md)
          → Sprint / iteration execution
```

A **backlog item** is informal and exploratory.
An **active task** is committed and tracked in the iteration loop.

---

## Backlog Item Format

```yaml
---
id: BACKLOG-YYYYMMDD-NNN
title: "Short description of the idea"
status: captured | promoting | promoted | dropped
priority: A | B | C
project: forge
tags: []
context: dev
created_at: YYYY-MM-DD
updated_at: YYYY-MM-DD
promoted_to: null          # set to TASK-ID when promoted
---

## Description

One paragraph: what is this and why does it matter?

## Notes

Details, constraints, open questions — to be refined before promoting.
```

---

## Status Lifecycle

| Transition | Trigger | Who |
|---|---|---|
| → `captured` | `/backlog-task "Title"` | Agent or human |
| `captured` → `promoting` | Manual edit or `/backlog-task promote` | Human |
| `promoting` → `promoted` | `/backlog-task promote` succeeds | Agent |
| `captured/promoting` → `dropped` | Manual edit | Human |

**Rule: promotion is always explicit.** No auto-promotion from backlog → task.

---

## Promotion Rules

Before promoting, verify:
1. Title is clear and actionable
2. Description explains the "why" (not just "what")
3. Priority is set appropriately
4. Tags and context are correct

On promotion, `/backlog-task promote <id>` will:
1. Read the backlog item
2. Create an active task via `/task-create` with matching fields
3. Update the backlog item: `status: promoted`, `promoted_to: <TASK-ID>`

---

## Sprint & Iteration Identifiers (v0.4)

Active tasks support two optional scheduling fields:

```yaml
sprint_id: "sprint-2026-01"     # optional — sprint this task belongs to
iteration_id: 24                 # optional — forge iteration that created this task
```

**Sprint rules (from execution brief):**
- A sprint *can* invoke an iteration loop
- An iteration loop *cannot* invoke a sprint
- Sprint scope is larger than an iteration

These fields are set at task creation time and are informational — they do not
affect task selection or loop behavior in v0.4.

---

## Using the Backlog Skill

```bash
# Capture
/backlog-task "Idea title"
/backlog-task "Investigate caching options" priority=B tags=perf,research

# Review
/backlog-task list
/backlog-task list status=captured

# Promote
/backlog-task promote BACKLOG-20260221-001
```

See `plugins/backlog/skills/backlog-task/SKILL.md` for full protocol.

---

## Evidence

- Created: 2026-02-21 (EPIC-004, TASK-084, iteration 24)
- Source: `forge_v0.4_execution_brief.md` — EPIC-004 scope
- Reviewed against: `docs/04_task-model.md`, `.forge/pcc/02_conventions.md`
