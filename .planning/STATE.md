# Project State

## Current Position
- **Current Phase:** 04-localize-templates-and-data-states
- **Current Plan:** 01 (Wave 1 — HTML templates localization)
- **Status:** Plan 01 completed

## Key Decisions
- Added empty `html[lang="tr"] {}` blocks to serve as correctness markers without font/direction overrides (Turkish is standard Latin LTR).
- Replaced hardcoded `lang="en"` and `"Cover Letter:"` in `cover-letter-template.html` with dynamic `{{LANG}}` and `{{COVER_LETTER_LABEL}}` placeholders.

## Metrics
- **Last Updated:** 2026-07-07
