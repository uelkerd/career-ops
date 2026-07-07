---
phase: 05-tui-dashboard-localization
plan: 01
subsystem: ui
tags: [tui, go, i18n, turkish, dashboard, bubbletea]

# Dependency graph
requires:
  - phase: 04-localize-templates-and-data-states
    provides: [canonical Turkish status strings in states.yml]
provides:
  - Zero-dependency static i18n translation catalog for TUI dashboard
  - Turkish translation strings for all dashboard headers, tabs, and status labels
  - Runtime locale switching via SetLang and ToggleLang
  - Enhanced status normalization supporting UTF-8 and ASCII Turkish tracking aliases
affects: [05-tui-dashboard-localization, dashboard, tui]

# Tech tracking
tech-stack:
  added: []
  patterns: [Static struct-based i18n catalog without external libraries or runtime I/O, UTF-8 safe status string normalization with ASCII fallbacks]

key-files:
  created:
    - dashboard/internal/i18n/catalog.go
    - dashboard/internal/i18n/i18n_test.go
  modified:
    - dashboard/internal/data/career.go
    - dashboard/internal/data/career_test.go

key-decisions:
  - "Implemented a zero-dependency static struct Catalog in internal/i18n instead of using heavyweight external libraries or runtime filesystem loading."
  - "Added both exact UTF-8 Turkish status strings (e.g., değerlendirildi, yanıt verildi) and ASCII-normalized fallbacks (degerlendirildi, yanit verildi) to NormalizeStatus to prevent Go strings.ToLower casing issues with dotted/undotted I."

patterns-established:
  - "TUI i18n: Use static pointer i18n.Current pointing to predefined static struct catalogs (&i18n.En, &i18n.Tr)."

requirements-completed: [REQ-3.4]

# Metrics
duration: 15min
completed: 2026-07-07
---

# Phase 05 Plan 01 Summary: TUI Dashboard Translation Catalog & Status Normalization

## Overview
We established the core localization foundation for the Go TUI Dashboard by creating a zero-dependency translation catalog (`internal/i18n`) and extending the data layer (`internal/data/career.go`) to recognize Turkish status strings.

## Key Accomplishments
1. **Translation Catalog (`internal/i18n`)**:
   - Implemented `Catalog` struct containing required UI strings (Headers, Tabs, Statuses, TimeAgo).
   - Defined static instances `i18n.En` and `i18n.Tr`.
   - Added runtime locale management methods: `SetLang`, `ToggleLang`, and `GetLang`.
   - Created full unit test suite in `i18n_test.go` verifying label lookups and runtime switching.
2. **Status Normalization (`internal/data/career.go`)**:
   - Extended `NormalizeStatus` with Turkish tracking aliases matching `states.yml`.
   - Added ASCII-normalized fallbacks (e.g., `yanit verildi`, `mulakat`, `degerlendirildi`) to handle Go's standard `strings.ToLower` ASCII casing conversions for uppercase Turkish characters (e.g., dotted/undotted I).
   - Verified normalization across 8 Turkish status aliases and standard English IDs in `career_test.go`.

## Test Verification
- Ran `go test -v ./internal/i18n/...` (All tests PASSED).
- Ran `go test -v ./internal/data/...` (All tests PASSED).
- Ran `go test -v ./...` across the entire `dashboard/` module (All test suites PASSED).
