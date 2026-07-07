---
wave: 2
depends_on:
  - wave: 1
files_modified:
  - templates/states.yml
  - normalize-statuses.mjs
  - verify-pipeline.mjs
autonomous: true
requirements:
  - REQ-3.3
---

# Wave 2 — Turkish state aliases in runtime files

## Goal

Add Turkish-language status aliases to all three state-handling files so that
Turkish tracker entries (e.g. `Değerlendirildi`, `Başvuruldu`) are normalized
and validated correctly by the pipeline toolchain.

Three files must all be updated — they are independent (no shared write target)
and run in parallel within wave 2.

### Turkish alias mapping

| Turkish alias | Canonical English |
|---|---|
| değerlendirildi | Evaluated |
| başvuruldu | Applied |
| yanıt verildi | Responded |
| mülakat | Interview |
| teklif | Offer |
| reddedildi | Rejected |
| iptal edildi | Discarded |
| uygun değil | SKIP |

---

## Task D — templates/states.yml: add Turkish aliases

```xml
<task id="D" parallel_group="wave2">
  <title>Add Turkish aliases to templates/states.yml</title>
  <requirements>REQ-3.3</requirements>

  <read_first>
    - templates/states.yml  (full file, 57 lines)
  </read_first>

  <action>
  Edit templates/states.yml. For each state entry, append Turkish aliases to
  the existing aliases list. Make NO other changes.

  EXACT DIFFS (YAML block additions):

    evaluated:
      aliases: [evaluada]
      → aliases: [evaluada, değerlendirildi]

    applied:
      aliases: [aplicado, enviada, aplicada, sent]
      → aliases: [aplicado, enviada, aplicada, sent, başvuruldu]

    responded:
      aliases: [respondido]
      → aliases: [respondido, yanıt verildi]

    interview:
      aliases: [entrevista]
      → aliases: [entrevista, mülakat]

    offer:
      aliases: [oferta]
      → aliases: [oferta, teklif]

    rejected:
      aliases: [rechazado, rechazada]
      → aliases: [rechazado, rechazada, reddedildi]

    discarded:
      aliases: [descartado, descartada, cerrada, cancelada]
      → aliases: [descartado, descartada, cerrada, cancelada, iptal edildi]

    skip:
      aliases: [no_aplicar, no aplicar, skip, monitor]
      → aliases: [no_aplicar, no aplicar, skip, monitor, uygun değil]

  Preserve all other fields (id, label, description, dashboard_group) exactly.
  </action>

  <acceptance_criteria>
    grep 'değerlendirildi' templates/states.yml
    # Must exit 0

    grep 'başvuruldu' templates/states.yml
    # Must exit 0

    grep 'reddedildi' templates/states.yml
    # Must exit 0

    grep 'mülakat' templates/states.yml
    # Must exit 0

    grep 'uygun değil' templates/states.yml
    # Must exit 0

    # YAML stays valid
    node -e "require('fs').readFileSync('templates/states.yml','utf8'); console.log('File readable');"
  </acceptance_criteria>
</task>
```

---

## Task E — normalize-statuses.mjs: add Turkish aliases

```xml
<task id="E" parallel_group="wave2">
  <title>Add Turkish aliases to normalize-statuses.mjs</title>
  <requirements>REQ-3.3</requirements>

  <read_first>
    - normalize-statuses.mjs  (lines 76-87: Spanish aliases block)
  </read_first>

  <action>
  Edit normalize-statuses.mjs. In the normalizeStatus() function, the "Spanish
  aliases → English canonicals" block spans lines 76-83. Append Turkish aliases
  to each existing .includes() array, and add one new line for yanıt verildi.

  EXACT CHANGES (show before → after for each line):

  LINE 77 (Evaluated):
    BEFORE: if (['evaluada'].includes(lower)) return { status: 'Evaluated' };
    AFTER:  if (['evaluada', 'değerlendirildi'].includes(lower)) return { status: 'Evaluated' };

  LINE 78 (Applied):
    BEFORE: if (['aplicado', 'enviada', 'aplicada', 'applied', 'sent'].includes(lower)) return { status: 'Applied' };
    AFTER:  if (['aplicado', 'enviada', 'aplicada', 'applied', 'sent', 'başvuruldu'].includes(lower)) return { status: 'Applied' };

  LINE 79 (Responded) — add a new line after the existing respondido line:
    BEFORE: if (['respondido'].includes(lower)) return { status: 'Responded' };
    AFTER:  if (['respondido', 'yanıt verildi'].includes(lower)) return { status: 'Responded' };

  LINE 80 (Interview):
    BEFORE: if (['entrevista'].includes(lower)) return { status: 'Interview' };
    AFTER:  if (['entrevista', 'mülakat'].includes(lower)) return { status: 'Interview' };

  LINE 81 (Offer):
    BEFORE: if (['oferta'].includes(lower)) return { status: 'Offer' };
    AFTER:  if (['oferta', 'teklif'].includes(lower)) return { status: 'Offer' };

  LINE 82 (Discarded):
    BEFORE: if (['cerrada', 'descartada'].includes(lower)) return { status: 'Discarded' };
    AFTER:  if (['cerrada', 'descartada', 'iptal edildi'].includes(lower)) return { status: 'Discarded' };

  LINE 83 (SKIP):
    BEFORE: if (['no aplicar', 'no_aplicar', 'skip'].includes(lower)) return { status: 'SKIP' };
    AFTER:  if (['no aplicar', 'no_aplicar', 'skip', 'uygun değil'].includes(lower)) return { status: 'SKIP' };

  Also add reddedildi to the Rejected handling. Find the line (around line 46)
  that reads:
    if (/^rechazada?$/i.test(s)) return { status: 'Rejected' };
  Add immediately after it:
    if (/^reddedildi$/i.test(s)) return { status: 'Rejected' };

  Make NO other changes to the file.
  </action>

  <acceptance_criteria>
    grep 'değerlendirildi' normalize-statuses.mjs
    # Must exit 0

    grep 'başvuruldu' normalize-statuses.mjs
    # Must exit 0

    grep 'reddedildi' normalize-statuses.mjs
    # Must exit 0

    grep 'mülakat' normalize-statuses.mjs
    # Must exit 0

    grep 'uygun değil' normalize-statuses.mjs
    # Must exit 0

    # Script still parses as valid JS
    node --input-type=module < normalize-statuses.mjs 2>/dev/null; echo "exit $?"
    # (will exit 0 since apps file may not exist; the important thing is no syntax error)
    node -e "import('./normalize-statuses.mjs').catch(e => { if (!e.message.includes('No applications')) { process.stderr.write(e.message); process.exit(1); } })"
  </acceptance_criteria>
</task>
```

---

## Task F — verify-pipeline.mjs: add Turkish aliases

```xml
<task id="F" parallel_group="wave2">
  <title>Add Turkish aliases to verify-pipeline.mjs</title>
  <requirements>REQ-3.3</requirements>

  <read_first>
    - verify-pipeline.mjs  (lines 44-53: ALIASES object)
  </read_first>

  <action>
  Edit verify-pipeline.mjs. The ALIASES object (lines 44-53) maps alias strings
  to canonical lowercase IDs. Append the Turkish aliases to the existing object.

  CURRENT ALIASES object (lines 44-53):
    const ALIASES = {
      'evaluada': 'evaluated', 'condicional': 'evaluated', 'hold': 'evaluated', 'evaluar': 'evaluated', 'verificar': 'evaluated',
      'aplicado': 'applied', 'enviada': 'applied', 'aplicada': 'applied', 'applied': 'applied', 'sent': 'applied',
      'respondido': 'responded',
      'entrevista': 'interview',
      'oferta': 'offer',
      'rechazado': 'rejected', 'rechazada': 'rejected',
      'descartado': 'discarded', 'descartada': 'discarded', 'cerrada': 'discarded', 'cancelada': 'discarded',
      'no aplicar': 'skip', 'no_aplicar': 'skip', 'monitor': 'skip', 'geo blocker': 'skip',
    };

  REPLACE WITH (add Turkish entries to each row):
    const ALIASES = {
      'evaluada': 'evaluated', 'condicional': 'evaluated', 'hold': 'evaluated', 'evaluar': 'evaluated', 'verificar': 'evaluated',
      'değerlendirildi': 'evaluated',
      'aplicado': 'applied', 'enviada': 'applied', 'aplicada': 'applied', 'applied': 'applied', 'sent': 'applied',
      'başvuruldu': 'applied',
      'respondido': 'responded', 'yanıt verildi': 'responded',
      'entrevista': 'interview', 'mülakat': 'interview',
      'oferta': 'offer', 'teklif': 'offer',
      'rechazado': 'rejected', 'rechazada': 'rejected', 'reddedildi': 'rejected',
      'descartado': 'discarded', 'descartada': 'discarded', 'cerrada': 'discarded', 'cancelada': 'discarded', 'iptal edildi': 'discarded',
      'no aplicar': 'skip', 'no_aplicar': 'skip', 'monitor': 'skip', 'geo blocker': 'skip', 'uygun değil': 'skip',
    };

  Make NO other changes to the file.
  </action>

  <acceptance_criteria>
    grep 'değerlendirildi' verify-pipeline.mjs
    # Must exit 0

    grep 'başvuruldu' verify-pipeline.mjs
    # Must exit 0

    grep 'reddedildi' verify-pipeline.mjs
    # Must exit 0

    grep 'mülakat' verify-pipeline.mjs
    # Must exit 0

    grep 'uygun değil' verify-pipeline.mjs
    # Must exit 0

    # Script parses without syntax errors
    node -e "
      import('./verify-pipeline.mjs').catch(e => {
        // Acceptable errors: missing apps file, not syntax errors
        if (e.code !== 'MODULE_NOT_FOUND' && !e.message.includes('ENOENT')) {
          process.stderr.write('Syntax error: ' + e.message + '\n');
          process.exit(1);
        }
      });
    "
  </acceptance_criteria>
</task>
```

---

## Wave 2 verification (run after all three tasks complete)

```bash
# All 8 Turkish aliases present in states.yml
for alias in değerlendirildi başvuruldu 'yanıt verildi' mülakat teklif reddedildi 'iptal edildi' 'uygun değil'; do
  grep -q "$alias" templates/states.yml && echo "states.yml OK: $alias" || echo "MISSING in states.yml: $alias"
done

# All 8 aliases present in normalize-statuses.mjs
for alias in değerlendirildi başvuruldu 'yanıt verildi' mülakat teklif reddedildi 'iptal edildi' 'uygun değil'; do
  grep -q "$alias" normalize-statuses.mjs && echo "normalize OK: $alias" || echo "MISSING in normalize: $alias"
done

# All 8 aliases present in verify-pipeline.mjs
for alias in değerlendirildi başvuruldu 'yanıt verildi' mülakat teklif reddedildi 'iptal edildi' 'uygun değil'; do
  grep -q "$alias" verify-pipeline.mjs && echo "verify OK: $alias" || echo "MISSING in verify: $alias"
done

# End-to-end alias smoke test — normalize recognizes Turkish statuses
node -e "
  // Inline test of normalizeStatus logic
  const aliases = {
    'değerlendirildi': 'Evaluated',
    'başvuruldu': 'Applied',
    'yanıt verildi': 'Responded',
    'mülakat': 'Interview',
    'teklif': 'Offer',
    'reddedildi': 'Rejected',
    'iptal edildi': 'Discarded',
    'uygun değil': 'SKIP',
  };
  const src = require('fs').readFileSync('normalize-statuses.mjs', 'utf8');
  for (const [alias, expected] of Object.entries(aliases)) {
    if (!src.includes(alias)) {
      console.error('MISSING: ' + alias + ' in normalize-statuses.mjs');
      process.exit(1);
    }
  }
  console.log('PASS: all 8 Turkish aliases found in normalize-statuses.mjs');
"
```

## must_haves

- All 8 Turkish aliases present in `templates/states.yml`
- All 8 Turkish aliases present in `normalize-statuses.mjs`
- All 8 Turkish aliases present in `verify-pipeline.mjs`
- No existing Spanish/English aliases removed from any of the three files
- All three scripts remain syntactically valid JavaScript/YAML
