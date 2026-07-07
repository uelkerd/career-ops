# Wave 2 Execution Summary

## What Was Done
- Migrated `dashboard/internal/ui/screens/pipeline.go` to use `i18n.Current` for all table column headers, filter tab labels, footer shortcut texts, relative dates (`formatTimeAgo`), and empty state messages.
- Expanded the status column width in `colWidths` from 12 to 15 to accommodate Turkish aliases (e.g., "DEĞERLENDİRİLDİ").
- Updated `viewer.go` and `progress.go` UI strings and shortcuts to rely on the static `i18n.Current` catalog.
- Rewrote `dashboard/main.go` to handle the `-lang` command-line flag and `LANG` environment variables, automatically initializing `i18n.Current`.
- Wired the `"t"` / `"T"` global keyboard event interceptor in `main.go` to invoke `i18n.ToggleLang()` at runtime, enabling seamless hot-switching between English and Turkish.

## Validation
- [x] Run `go build -o dashboard_test ./dashboard` and ensure it compiles successfully.
- [x] Run `cd dashboard && go test ./...` and ensure tests pass.
- [x] Run `./dashboard_test -lang tr` and verify the UI text is in Turkish.
- [x] Press `t` to toggle back to English.

## Notes
- The static zero-dependency localization architecture is complete. All `statusLabel` calls now pass through `i18n.Current.StatusLabel`, decoupling canonical ASCII data keys from their localized display strings.
