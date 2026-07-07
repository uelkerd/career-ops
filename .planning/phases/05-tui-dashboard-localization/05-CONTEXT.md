# Phase 5: TUI Dashboard Localization - Context

**Gathered:** 2026-07-07
**Status:** Ready for planning

<domain>
## Phase Boundary

Add Turkish string support and localization architecture to the Go TUI Dashboard (`dashboard/`). This covers UI labels, column headers, footer keyboard shortcuts, view tabs, empty states, and relative timestamp localization. Modifying core pipeline evaluation logic or external scraping engines is out of scope.
</domain>

<decisions>
## Implementation Decisions

### Language Configuration & Switching
- Auto-detect language via standard `LANG`/`LC_ALL` environment variables on startup.
- Support explicit CLI flag `-lang tr` (and `-lang en`) to override environment variables.
- Implement a runtime keyboard toggle (e.g., pressing `L` or `T` in the TUI) to switch languages on the fly without restarting the application.

### Localization Implementation Architecture
- Implement a zero-dependency internal `i18n` package (e.g., `dashboard/internal/i18n`).
- Use typed Go structs/dictionaries for EN and TR translations to guarantee compile-time type safety, maximum rendering performance, and zero runtime file disk I/O.

### Scope of UI Translation
- Comprehensive UI localization:
  - Column headers and table titles.
  - Footer shortcuts (`Press q to quit` -> `Çıkmak için q`, etc.).
  - View tabs and navigation bar elements.
  - Empty state notifications and error messages.
  - Relative timestamps in `timeago` formatting (`3 days ago` -> `3 gün önce`, `just now` -> `az önce`).
- Ensure Turkish characters (`ç ğ ı İ ö ş ü`) render cleanly across standard terminal viewports without alignment or width calculation bugs (utilizing `lipgloss` / `ansi` width utilities).

### Agent's Discretion
- Internal struct naming and field layout within the `i18n` package.
- Exact keyboard shortcut key assigned for runtime language toggling (e.g., `l` vs `t`), provided it does not conflict with existing navigation keys.
- Handling of unrecognized language codes (defaulting gracefully to English).
</decisions>

<canonical_refs>
## Canonical References

### Localization Standards & Pipeline State
- `.planning/ROADMAP.md` § Phase 5 — Defines core phase scope and requirements.
- `templates/states.yml` — Canonical state definitions and Turkish status aliases (`Değerlendirildi`, `Başvuruldu`, `Yanıt Verildi`, `Mülakat`, `Teklif`, `Reddedildi`, `Uygun Değil`).
- `normalize-statuses.mjs` — Runtime status normalization reference mapping Turkish aliases to canonical states.
- `dashboard/internal/ui/screens/` — Existing TUI screen implementations (viewer, pipeline, progress) that consume text labels.
- `dashboard/internal/data/` — Data layer and career model parsing logic.
</canonical_refs>

<specifics>
## Specific Ideas

- Ensure runtime language toggling updates the active view immediately without losing current selection or scroll position.
- Relative date formatting in `timeago` must read naturally in Turkish (e.g., suffix-based words like `önce` instead of prefix-based `ago`).
</specifics>

<deferred>
## Deferred Ideas

- Support for additional languages beyond English and Turkish — defer to future backlog.
- External user-configurable translation JSON/YAML overrides — defer to future enhancement.
</deferred>

---

*Phase: 05-tui-dashboard-localization*
*Context gathered: 2026-07-07*
