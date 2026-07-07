---
phase: 03
slug: translate-minor-utility-modes-update-docs
status: draft
nyquist_compliant: false
wave_0_complete: false
created: 2026-07-07
---

# Phase 03 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | None — Validation is via file existence and regex grep |
| **Config file** | none |
| **Quick run command** | `ls modes/tr/*.md` |
| **Full suite command** | `grep -E "Applied|Rejected|Evaluated" modes/tr/tracker.md` |
| **Estimated runtime** | ~1 seconds |

---

## Sampling Rate

- **After every task commit:** Run `ls modes/tr/*.md`
- **After every plan wave:** Run `grep` assertions
- **Before `/gsd-verify-work`:** Full suite must be green
- **Max feedback latency:** 1 seconds

---

## Per-Task Verification Map

| Task ID | Plan | Wave | Requirement | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|-----------|-------------------|-------------|--------|
| 03-01-01 | 01 | 1 | null | file_check | `ls modes/tr/project.md` | ✅ / ❌ W0 | ⬜ pending |
| 03-01-02 | 01 | 1 | null | file_check | `ls modes/tr/training.md` | ✅ / ❌ W0 | ⬜ pending |
| 03-01-03 | 01 | 1 | null | grep | `grep "Applied" modes/tr/tracker.md` | ✅ / ❌ W0 | ⬜ pending |
| 03-01-04 | 01 | 1 | null | file_check | `ls modes/tr/auto-pipeline.md` | ✅ / ❌ W0 | ⬜ pending |
| 03-01-05 | 01 | 2 | null | grep | `grep "%100" modes/tr/README.md` | ✅ / ❌ W0 | ⬜ pending |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

- [ ] None needed.

*If none: "Existing infrastructure covers all phase requirements."*

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| Tone Check | null | LLM Tone | Read modes/tr/README.md to verify proud and formal tone. |

---

## Validation Sign-Off

- [ ] All tasks have `<automated>` verify or Wave 0 dependencies
- [ ] Sampling continuity: no 3 consecutive tasks without automated verify
- [ ] Wave 0 covers all MISSING references
- [ ] No watch-mode flags
- [ ] Feedback latency < 5s
- [ ] `nyquist_compliant: true` set in frontmatter

**Approval:** pending
