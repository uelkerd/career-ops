---
wave: 1
depends_on: []
files_modified:
  - templates/cv-template.html
  - templates/resume-template.html
  - templates/cover-letter-template.html
autonomous: true
requirements:
  - REQ-3.1
  - REQ-3.2
---

# Wave 1 — Turkish lang support in HTML templates

## Goal

Add Turkish (`tr`) language support to all three CV/resume/cover-letter HTML
templates. Turkish is LTR — **no `direction: rtl`**, **no font-family
override**. The `html[lang="tr"]` block is intentionally empty — its presence
is the correctness marker signalling the template was updated.

Two additional gaps exist only in `cover-letter-template.html`:
1. Hardcoded `lang="en"` on `<html>` → replace with `{{LANG}}`
2. Hardcoded `"Cover Letter:"` prefix in the role-title div → replace with `{{COVER_LETTER_LABEL}}`

All three tasks are independent and run in parallel.

---

## Task A — cv-template.html: add `html[lang="tr"]` block

```xml
<task id="A" parallel_group="wave1">
  <title>Add html[lang="tr"] block to cv-template.html</title>
  <requirements>REQ-3.1</requirements>

  <read_first>
    - templates/cv-template.html  (lines 426-449: Japanese block comment through </style>)
  </read_first>

  <action>
  Insert the following block immediately after the closing brace of the
  html[lang="ja"] .header h1 / .section-title / … selector group
  (currently ends at line 448), before the closing </style> tag (line 449).

  EXACT TEXT TO INSERT (after the closing brace of the ja block):

    /* === Turkish (tr) — LTR, no font override === */
    /* Turkish uses standard Latin characters with extras (ç ğ ı İ ö ş ü).
       The default Liberation Sans / Helvetica stack covers these glyphs.
       No direction or font-family override needed (LTR language).
       Block kept intentionally empty — its presence signals Turkish support. */
    html[lang="tr"] {}

  Use "tr" not "tr-TR".
  Do NOT add direction:rtl or any font-family property inside this block.
  </action>

  <acceptance_criteria>
    grep -n 'html\[lang="tr"\]' templates/cv-template.html
    # Must exit 0 and print: html[lang="tr"] {}

    node -e "
      const s = require('fs').readFileSync('templates/cv-template.html','utf8');
      const m = s.match(/html\[lang=\"tr\"\]\s*\{([^}]*)\}/);
      if (!m) { process.stderr.write('FAIL: tr block not found\n'); process.exit(1); }
      if (m[1].includes('font-family')) { process.stderr.write('FAIL: font-family in tr block\n'); process.exit(1); }
      if (m[1].includes('direction')) { process.stderr.write('FAIL: direction in tr block\n'); process.exit(1); }
      console.log('PASS: tr block empty and correct');
    "
  </acceptance_criteria>
</task>
```

---

## Task B — resume-template.html: add `html[lang="tr"]` block

```xml
<task id="B" parallel_group="wave1">
  <title>Add html[lang="tr"] block to resume-template.html</title>
  <requirements>REQ-3.1</requirements>

  <read_first>
    - templates/resume-template.html  (lines 377-400: Japanese block through </style>)
  </read_first>

  <action>
  Insert the following block immediately after the closing brace of the
  html[lang="ja"] .header h1 / .section-title / … selector group
  (currently ends at line ~399), before the closing </style> tag (line 400).

  EXACT TEXT TO INSERT:

    /* === Turkish (tr) — LTR, no font override === */
    /* Turkish uses standard Latin characters with extras (ç ğ ı İ ö ş ü).
       The default Space Grotesk / DM Sans stacks cover these glyphs.
       No direction or font-family override needed (LTR language).
       Block kept intentionally empty — its presence signals Turkish support. */
    html[lang="tr"] {}

  Use "tr" not "tr-TR".
  Do NOT add direction:rtl or any font-family property inside this block.
  </action>

  <acceptance_criteria>
    grep -n 'html\[lang="tr"\]' templates/resume-template.html
    # Must exit 0

    node -e "
      const s = require('fs').readFileSync('templates/resume-template.html','utf8');
      const m = s.match(/html\[lang=\"tr\"\]\s*\{([^}]*)\}/);
      if (!m) { process.stderr.write('FAIL: tr block not found\n'); process.exit(1); }
      if (m[1].includes('font-family')) { process.stderr.write('FAIL: font-family in tr block\n'); process.exit(1); }
      if (m[1].includes('direction')) { process.stderr.write('FAIL: direction in tr block\n'); process.exit(1); }
      console.log('PASS');
    "
  </acceptance_criteria>
</task>
```

---

## Task C — cover-letter-template.html: fix lang + label + add tr block

```xml
<task id="C" parallel_group="wave1">
  <title>Fix cover-letter-template.html: dynamic lang, dynamic label, Turkish block</title>
  <requirements>REQ-3.2</requirements>

  <read_first>
    - templates/cover-letter-template.html  (full file, 139 lines)
  </read_first>

  <action>
  THREE edits to templates/cover-letter-template.html:

  EDIT 1 — line 2, replace hardcoded lang:
    BEFORE: <html lang="en">
    AFTER:  <html lang="{{LANG}}">

  EDIT 2 — line 124, replace hardcoded "Cover Letter:" prefix:
    BEFORE: <div class="role-title">Cover Letter: {{ROLE_TITLE}}</div>
    AFTER:  <div class="role-title">{{COVER_LETTER_LABEL}}{{ROLE_TITLE}}</div>

  EDIT 3 — Add Turkish lang block inside <style>, just before the closing
  </style> tag (currently line 116). Insert before it:

    /* === Turkish (tr) — LTR, no font override === */
    /* Turkish uses standard Latin characters. Helvetica/Arial covers ç ğ ı İ ö ş ü.
       No direction or font-family override needed. */
    html[lang="tr"] {}

  NOTE: The generator (generate-pdf.mjs / modes/pdf.md) substitutes {{LANG}}
  and {{COVER_LETTER_LABEL}} at generation time. The template only holds the
  placeholders.
  </action>

  <acceptance_criteria>
    grep -n '{{LANG}}' templates/cover-letter-template.html
    # Must exit 0 — confirms placeholder present on the html element

    grep -n 'lang="en"' templates/cover-letter-template.html
    # Must exit 1 — confirms hardcoded "en" is gone

    grep -n '{{COVER_LETTER_LABEL}}' templates/cover-letter-template.html
    # Must exit 0

    grep -n '"Cover Letter:' templates/cover-letter-template.html
    # Must exit 1 — confirms hardcoded prefix is gone

    grep -n 'html\[lang="tr"\]' templates/cover-letter-template.html
    # Must exit 0
  </acceptance_criteria>
</task>
```

---

## Wave 1 verification (run after all tasks complete)

```bash
# All three templates have the Turkish block
grep -l 'html\[lang="tr"\]' templates/cv-template.html templates/resume-template.html templates/cover-letter-template.html
# Expected: 3 filenames

# cover-letter placeholders
grep '{{LANG}}' templates/cover-letter-template.html
grep '{{COVER_LETTER_LABEL}}' templates/cover-letter-template.html

# No RTL pollution in tr blocks
node -e "
  const fs = require('fs');
  const files = [
    'templates/cv-template.html',
    'templates/resume-template.html',
    'templates/cover-letter-template.html',
  ];
  for (const f of files) {
    const s = fs.readFileSync(f, 'utf8');
    const m = s.match(/html\[lang=\"tr\"\]\s*\{([^}]*)\}/);
    if (!m) { console.error('FAIL: tr block missing in ' + f); process.exit(1); }
    if (m[1].trim()) { console.error('FAIL: tr block not empty in ' + f + ': ' + m[1]); process.exit(1); }
    console.log('PASS: ' + f);
  }
"
```

## must_haves

- `html[lang="tr"] {}` present in all three templates
- No `direction: rtl` or `font-family` inside any `html[lang="tr"]` block
- `cover-letter-template.html`: `lang="{{LANG}}"` on the html element (not `lang="en"`)
- `cover-letter-template.html`: `{{COVER_LETTER_LABEL}}` placeholder in the role-title div
- `cv-template.html` and `resume-template.html`: `lang="{{LANG}}"` unchanged (pre-existing — verify not broken)
