---
phase: 05
slug: tui-dashboard-localization
status: draft
nyquist_compliant: true
wave_0_complete: true
created: 2026-07-07
---

# Phase 05 — Validation Strategy

> Per-phase validation contract for feedback sampling during execution.

---

## Test Infrastructure

| Property | Value |
|----------|-------|
| **Framework** | go test |
| **Config file** | dashboard/go.mod |
| **Quick run command** | `cd dashboard && go test -short ./...` |
| **Full suite command** | `cd dashboard && go test -v ./...` |
| **Estimated runtime** | ~5 seconds |

---

## Sampling Rate

- **After every task commit:** Run `cd dashboard && go test -short ./...`
- **After every plan wave:** Run `cd dashboard && go test -v ./...`
- **Before `/gsd-verify-work`:** Full suite must be green
- **Max feedback latency:** 5 seconds

---

## Per-Task Verification Map

| Task ID | Plan | Wave | Requirement | Test Type | Automated Command | File Exists | Status |
|---------|------|------|-------------|-----------|-------------------|-------------|--------|
| 05-01-01 | 01 | 1 | N/A (i18n catalog) | unit | `cd dashboard && go test -v ./internal/i18n/...` | ❌ W1 | ⬜ pending |
| 05-01-02 | 01 | 1 | N/A (data layer) | unit | `cd dashboard && go test -v ./internal/data/...` | ✅ | ⬜ pending |
| 05-02-01 | 02 | 2 | N/A (UI screens) | unit/integration | `cd dashboard && go test -v ./internal/ui/...` | ✅ | ⬜ pending |
| 05-02-02 | 02 | 2 | N/A (main wiring) | regression | `cd dashboard && go test -v ./...` | ✅ | ⬜ pending |

*Status: ⬜ pending · ✅ green · ❌ red · ⚠️ flaky*

---

## Wave 0 Requirements

- [x] Existing infrastructure covers all phase requirements.

---

## Manual-Only Verifications

| Behavior | Requirement | Why Manual | Test Instructions |
|----------|-------------|------------|-------------------|
| TUI visual alignment | N/A | Terminal UI visual check | Run `cd dashboard && go run . -lang=tr` and check table column widths and headers |
| Runtime keyboard toggle | N/A | Interactive keypress | Press 'L' inside dashboard and observe immediate language switch |

---

## Validation Sign-Off

- [x] All tasks have `<automated>` verify or Wave 0 dependencies
- [x] Sampling continuity: no 3 consecutive tasks without automated verify
- [x] Wave 0 covers all MISSING references
- [x] No watch-mode flags
- [x] Feedback latency < 5s
- [x] `nyquist_compliant: true` set in frontmatter

**Approval:** approved 2026-07-07
