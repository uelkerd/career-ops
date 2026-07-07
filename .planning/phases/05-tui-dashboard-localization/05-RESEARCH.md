# Phase 05 Research: TUI Dashboard Localization

## Standard Stack

The Go TUI Dashboard (`dashboard/`) is built on a lightweight, statically compiled command-line stack:
- **`github.com/charmbracelet/bubbletea v1.3.10`** — The Elm-architecture TUI framework.
- **`github.com/charmbracelet/lipgloss v1.1.0`** — Terminal layout and styling.
- **`github.com/charmbracelet/x/ansi v0.11.7`** — ANSI escape sequences and string width measurement.
- **`github.com/muesli/termenv v0.16.0`** — Terminal color profile detection.

### Prescriptive Library Strategy for i18n/l10n
**Use a lightweight, compile-time struct-based translation catalog (`internal/i18n`).**
- **Do NOT add heavy external i18n libraries** such as `go-i18n`, GNU `gettext`, or YAML/JSON runtime translation loading engines. 
- **Reason:** `career-dashboard` is compiled into a single static binary. External translation files or runtime reflection/file-system lookups add unnecessary overhead, failure points, and breaking dependencies for ~50 UI strings.
- **Implementation:** Use native Go structs with static instances for English (`En`) and Turkish (`Tr`), leveraging standard Go packages (`strings`, `unicode`, `fmt`).

---

## Architecture Patterns

### 1. Package-Level Translation Catalog (`internal/i18n`)
Create an `internal/i18n` package defining a `Catalog` struct that holds all UI labels, prompts, table headers, filter tabs, empty states, and status display names.
- Provide canonical static instances: `i18n.En` (default) and `i18n.Tr`.
- Maintain a package-level pointer or accessor (`i18n.Current` or `i18n.Get(lang string)`) initialized once at application startup in `main.go`.

### 2. Locale Resolution Strategy (Hierarchical Fallback)
In `dashboard/main.go`, determine the active locale using a deterministic 3-tier priority sequence:
1. **CLI Flag:** Check for a `-lang` flag (e.g., `career-dashboard -lang=tr` or `-lang=en`).
2. **Environment Variables:** If flag is empty, inspect `CAREER_OPS_LANG`, `LANG`, and `LC_ALL` (matching `tr`, `tr_TR`, etc.).
3. **User Profile Config:** If still unresolved, check if `config/profile.yml` exists in `careerOpsPath` and whether `language.modes_dir` is set to `modes/tr` or `tr`.
4. **Default:** Fallback to canonical English (`i18n.En`).

### 3. SOTA Turkish Engineering Terminology (Domain Rule)
Per `modes/tr/README.md`, Turkish engineering teams in Istanbul and Ankara speak fluent Turkish mixed with established English technical terminology.
- **NEVER Translate (Technical Jargon):** `Pipeline` (never use "boru hattı"), `Archetype`, `TL;DR`, `CV`, `PDF`, `Playwright`, `Tracker`, `Report`, `Score`, `Post`, `Bash`.
- **ALWAYS Translate (Domain & UI Text):**
  - `Job posting` → `İş ilanı` / `İlanlar`
  - `Application` → `Başvuru` / `Başvurular`
  - `Skills` → `Beceriler`
  - `Interview` → `Mülakat`
  - `Remote / Hybrid` → `Uzaktan / Hibrit`
  - `Conversion Rates` → `Dönüşüm Oranları`
  - `Weekly Activity` → `Haftalık Aktivite`

### 4. Separation of Data Storage vs. Display Layer
- **Storage Layer (`applications.md` & `career.go`):** All underlying application data, analytics grouping, and file writing MUST remain canonical English (`Evaluated`, `Applied`, `Responded`, `Interview`, `Offer`, `Rejected`, `Discarded`, `SKIP`).
- **Input Normalization Layer (`NormalizeStatus`):** When reading from tracker files or parsing user input, `NormalizeStatus(status string)` in `data/career.go` MUST recognize Turkish status aliases (from Phase 4 Wave 2: `değerlendirildi`, `başvuruldu`, `mülakat`, `teklif`, `reddedildi`, `iptal edildi`, `uygun değil`) and map them to lowercase canonical English IDs.
- **Display Layer (`ui/screens/...`):** When rendering tables, status pickers, or group headers, UI components call `i18n.Current.StatusLabel(norm)` to display localized Turkish strings.

---

## Don't Hand-Roll

### 1. Do NOT Hand-Roll Turkish Unicode Case Folding
- Standard ASCII `strings.ToLower` and `strings.ToUpper` fail on Turkish dotted/dotless I (`i` vs `İ` and `ı` vs `I`).
- When checking status strings in `NormalizeStatus`, compare against exact lowercase UTF-8 literals (e.g., `"değerlendirildi"`, `"uygun değil"`). Do not attempt byte-by-byte ASCII lowercasing on UTF-8 strings.

### 2. Do NOT Hand-Roll String Width Calculation
- Turkish UTF-8 characters (`ç`, `ğ`, `ı`, `İ`, `ö`, `ş`, `ü`) occupy 2 bytes each. Using `len(str)` for table column padding or truncation breaks terminal alignment.
- **Always use:** `[]rune(str)` for character truncation (as implemented in `truncateRunes`) and `ansi.StringWidth()` or `lipgloss.Width()` to calculate visual column widths.

### 3. Do NOT Hand-Roll Runtime File Loaders
- Do not build custom YAML or JSON file loaders to read external translation files at runtime. Keep all strings compiled in static Go structs.

---

## Common Pitfalls

### Pitfall 1: Truncating Turkish Status Labels in Table Columns
- **Issue:** In `pipeline.go` (`columnWidths()`), the default status column width (`cw.status`) is hardcoded to `12` runes. The Turkish translation for `Evaluated` is `DEĞERLENDİRİLDİ` (15 runes in uppercase group headers or table cells), causing display truncation (`DEĞERLENDİR...`) or table misalignment.
- **Verification / Fix:** When designing the catalog, use concise display labels for narrow table columns (e.g., `Değerlend.` or `Değerl.` for 12-rune columns), OR adjust `cw.status` budget from `12` to `14` runes in `columnWidths()`.

### Pitfall 2: Status Picker Writing Localized Strings to Disk
- **Issue:** If the interactive status picker (`overlayStatusPicker`) displays Turkish names (`Mülakat`, `Teklif`) and emits them directly in `PipelineUpdateStatusMsg.NewStatus`, `data.UpdateApplicationStatus` will write Turkish text to `applications.md`, breaking data grouping.
- **Verification / Fix:** Ensure `overlayStatusPicker` displays localized strings (`i18n.Current.StatusLabel(opt)`) to the user, but emits canonical English strings (`"Interview"`, `"Offer"`, etc.) in `PipelineUpdateStatusMsg`.

### Pitfall 3: Relative Date Grammatical Mismatches (`formatTimeAgo`)
- **Issue:** Direct string replacement of `"today"`, `"yesterday"`, and `"%dd ago"` without checking grammatical suffixes can look jarring.
- **Verification / Fix:** Implement `formatTimeAgo` localized formatting in the catalog: `"bugün"`, `"dün"`, and `"%d gün önce"`.

### Pitfall 4: Over-translating UI Brand and Command Names
- **Issue:** Translating commands or brand text like `"CAREER PIPELINE"` to `"KARİYER BORU HATTI"` or `"TL;DR:"` to `"ÖZETLE:"` violates user design constraints.
- **Verification / Fix:** Retain exact SOTA hybrid terminology: `"KARİYER PIPELINE"`, `"TL;DR:"`, `"Archetype:"` (or `"Arketip:"`), `"CV"`, `"PDF"`.

---

## Code Examples

### 1. The Translation Catalog (`internal/i18n/catalog.go`)
```go
package i18n

type Catalog struct {
	// Screen Banners & General
	AppTitle       string // e.g. "CAREER PIPELINE" vs "KARİYER PIPELINE"
	OffersSummary  string // e.g. "%d offers | Avg %s/5" vs "%d ilan | Ort %s/5"
	NoOffersMatch  string // e.g. "No offers match this filter" vs "Bu filtreye uyan ilan yok"
	LoadingPreview string // e.g. "Loading preview..." vs "Önizleme yükleniyor..."

	// Tabs & Filters
	TabAll       string // "ALL" vs "TÜMÜ"
	TabEvaluated string // "EVALUATED" vs "DEĞERLENDİRİLDİ"
	TabApplied   string // "APPLIED" vs "BAŞVURULDU"
	TabInterview string // "INTERVIEW" vs "MÜLAKAT"
	TabTop       string // "TOP ≥4" vs "EN İYİ ≥4"
	TabSkip      string // "SKIP" vs "UYGUN DEĞİL"
	TabRejected  string // "REJECTED" vs "REDDEDİLDİ"
	TabDiscarded string // "DISCARDED" vs "İPTAL"

	// Table Headers
	ColFit      string // "FIT" vs "UYUM"
	ColApplied  string // "APPLIED" vs "TARİH"
	ColCompany  string // "COMPANY" vs "ŞİRKET"
	ColRole     string // "ROLE" vs "POZİSYON"
	ColStatus   string // "STATUS" vs "DURUM"
	ColLocation string // "LOCATION" vs "KONUM"
	ColPay      string // "PAY" vs "ÜCRET"
	ColLast     string // "LAST" vs "SON"

	// Preview Labels
	LabelLoc      string // "Loc: " vs "Konum: "
	LabelPay      string // "Pay: " vs "Ücret: "
	LabelLast     string // "Last contact: " vs "Son iletişim: "
	LabelRemote   string // "Remote: " vs "Çalışma Şekli: "
	LabelOutcome  string // "Outcome: " vs "Sonuç: "

	// Work Modes
	ModeRemote     string // "Remote" vs "Uzaktan"
	ModeRemoteFlex string // "RemoteFlex" vs "Uzaktan (Esnek)"
	ModeHybrid     string // "Hybrid" vs "Hibrit"
	ModeFull       string // "Full" vs "Ofiste"

	// Progress Screen
	ProgressTitle   string // "SEARCH PROGRESS" vs "TAKİP İLERLEMESİ"
	ProgressSummary string // "%d evaluated | %.1f avg score" vs "%d değerlendirildi | %.1f ort. puan"
	FunnelTitle     string // "Pipeline Funnel" vs "Pipeline Hunisi"
	ScoresTitle     string // "Score Distribution" vs "Puan Dağılımı"
	RatesTitle      string // "Conversion Rates" vs "Dönüşüm Oranları"
	WeeklyTitle     string // "Weekly Activity" vs "Haftalık Aktivite"
	ActiveInfo      string // "%d active applications | %d total offers" vs "%d aktif başvuru | %d toplam teklif"

	// Relative Dates
	TimeToday     string // "today" vs "bugün"
	TimeYesterday string // "yesterday" vs "dün"
	TimeDaysAgo   string // "%dd ago" vs "%d gün önce"
}

var En = Catalog{
	AppTitle:       "CAREER PIPELINE",
	OffersSummary:  "%d offers | Avg %s/5",
	NoOffersMatch:  "No offers match this filter",
	LoadingPreview: "Loading preview...",
	TabAll:         "ALL",
	TabEvaluated:   "EVALUATED",
	TabApplied:     "APPLIED",
	TabInterview:   "INTERVIEW",
	TabTop:         "TOP ≥4",
	TabSkip:        "SKIP",
	TabRejected:    "REJECTED",
	TabDiscarded:   "DISCARDED",
	ColFit:         "FIT",
	ColApplied:     "APPLIED",
	ColCompany:     "COMPANY",
	ColRole:        "ROLE",
	ColStatus:      "STATUS",
	ColLocation:    "LOCATION",
	ColPay:         "PAY",
	ColLast:        "LAST",
	LabelLoc:       "Loc: ",
	LabelPay:       "Pay: ",
	LabelLast:      "Last contact: ",
	LabelRemote:    "Remote: ",
	LabelOutcome:   "Outcome: ",
	ModeRemote:     "Remote",
	ModeRemoteFlex: "RemoteFlex",
	ModeHybrid:     "Hybrid",
	ModeFull:       "Full",
	ProgressTitle:  "SEARCH PROGRESS",
	ProgressSummary: "%d evaluated | %.1f avg score",
	FunnelTitle:    "Pipeline Funnel",
	ScoresTitle:    "Score Distribution",
	RatesTitle:     "Conversion Rates",
	WeeklyTitle:    "Weekly Activity",
	ActiveInfo:     "%d active applications | %d total offers",
	TimeToday:      "today",
	TimeYesterday:  "yesterday",
	TimeDaysAgo:    "%dd ago",
}

var Tr = Catalog{
	AppTitle:       "KARİYER PIPELINE",
	OffersSummary:  "%d ilan | Ort %s/5",
	NoOffersMatch:  "Bu filtreye uyan ilan yok",
	LoadingPreview: "Önizleme yükleniyor...",
	TabAll:         "TÜMÜ",
	TabEvaluated:   "DEĞERLENDİRİLDİ",
	TabApplied:     "BAŞVURULDU",
	TabInterview:   "MÜLAKAT",
	TabTop:         "EN İYİ ≥4",
	TabSkip:        "UYGUN DEĞİL",
	TabRejected:    "REDDEDİLDİ",
	TabDiscarded:   "İPTAL",
	ColFit:         "UYUM",
	ColApplied:     "TARİH",
	ColCompany:     "ŞİRKET",
	ColRole:        "POZİSYON",
	ColStatus:      "DURUM",
	ColLocation:    "KONUM",
	ColPay:         "ÜCRET",
	ColLast:        "SON",
	LabelLoc:       "Konum: ",
	LabelPay:       "Ücret: ",
	LabelLast:      "Son iletişim: ",
	LabelRemote:    "Çalışma Şekli: ",
	LabelOutcome:   "Sonuç: ",
	ModeRemote:     "Uzaktan",
	ModeRemoteFlex: "Uzaktan (Esnek)",
	ModeHybrid:     "Hibrit",
	ModeFull:       "Ofiste",
	ProgressTitle:  "TAKİP İLERLEMESİ",
	ProgressSummary: "%d değerlendirildi | %.1f ort. puan",
	FunnelTitle:    "Pipeline Hunisi",
	ScoresTitle:    "Puan Dağılımı",
	RatesTitle:     "Dönüşüm Oranları",
	WeeklyTitle:    "Haftalık Aktivite",
	ActiveInfo:     "%d aktif başvuru | %d toplam teklif",
	TimeToday:      "bugün",
	TimeYesterday:  "dün",
	TimeDaysAgo:    "%d gün önce",
}

var Current = &En
```

### 2. Status Normalization with Turkish Aliases (`data/career.go`)
```go
func NormalizeStatus(status string) string {
	s := strings.ToLower(strings.TrimSpace(status))
	switch {
	case strings.Contains(s, "no aplicar") || strings.Contains(s, "no_aplicar") || s == "skip" ||
		strings.Contains(s, "geo blocker") || strings.Contains(s, "uygun değil") || strings.Contains(s, "uygun_değil"):
		return "skip"
	case strings.Contains(s, "interview") || strings.Contains(s, "entrevista") || strings.Contains(s, "mülakat"):
		return "interview"
	case s == "offer" || strings.Contains(s, "oferta") || strings.Contains(s, "teklif"):
		return "offer"
	case strings.Contains(s, "responded") || strings.Contains(s, "respondido") || strings.Contains(s, "yanıt verildi") || strings.Contains(s, "yanıt_verildi"):
		return "responded"
	case strings.Contains(s, "applied") || strings.Contains(s, "aplicado") || s == "enviada" || s == "aplicada" || s == "sent" || strings.Contains(s, "başvuruldu"):
		return "applied"
	case strings.Contains(s, "rejected") || strings.Contains(s, "rechazado") || s == "rechazada" || strings.Contains(s, "reddedildi"):
		return "rejected"
	case strings.Contains(s, "discarded") || strings.Contains(s, "descartado") || s == "descartada" || s == "cerrada" || s == "cancelada" ||
		strings.HasPrefix(s, "duplicado") || strings.HasPrefix(s, "dup") || strings.Contains(s, "iptal edildi") || strings.Contains(s, "iptal_edildi"):
		return "discarded"
	case strings.Contains(s, "evaluated") || strings.Contains(s, "evaluada") || s == "condicional" || s == "hold" || s == "monitor" || s == "evaluar" || s == "verificar" || strings.Contains(s, "değerlendirildi"):
		return "evaluated"
	default:
		return s
	}
}
```

### 3. Localized Status Display (`ui/screens/pipeline.go`)
```go
func statusLabel(norm string) string {
	if i18n.Current == &i18n.Tr {
		switch norm {
		case "interview":
			return "Mülakat"
		case "offer":
			return "Teklif"
		case "responded":
			return "Yanıt Verildi"
		case "applied":
			return "Başvuruldu"
		case "evaluated":
			return "Değerlendirildi"
		case "skip":
			return "Uygun Değil"
		case "rejected":
			return "Reddedildi"
		case "discarded":
			return "İptal Edildi"
		}
	}
	// Default English fallback
	switch norm {
	case "interview":
		return "Interview"
	case "offer":
		return "Offer"
	case "responded":
		return "Responded"
	case "applied":
		return "Applied"
	case "evaluated":
		return "Evaluated"
	case "skip":
		return "Skip"
	case "rejected":
		return "Rejected"
	case "discarded":
		return "Discarded"
	default:
		return norm
	}
}
```

---

## Validation Architecture

### 1. Automated Unit Tests (`dashboard/internal/...`)
- **Status Normalization Test (`career_test.go`):**
  - Verify that `NormalizeStatus` correctly maps all 8 Turkish status strings (`değerlendirildi`, `başvuruldu`, `yanıt verildi`, `mülakat`, `teklif`, `reddedildi`, `iptal edildi`, `uygun değil`) to their canonical English lowercase IDs.
- **Rune Width & Truncation Test (`viewer_test.go` / `pipeline_test.go`):**
  - Ensure `truncateRunes` and `lipgloss.Width` correctly handle Turkish multi-byte characters (`ç, ğ, ı, İ, ö, ş, ü`) without breaking string length budgets or table borders.
- **Relative Date Formatting Test:**
  - Verify `formatTimeAgo` outputs localized strings (`bugün`, `dün`, `%d gün önce`) when `i18n.Current` is set to `i18n.Tr`.

### 2. UI / Component Integration Tests
- **Header & Tab Rendering:**
  - Assert that setting `i18n.Current = &i18n.Tr` renders table column headers (`UYUM`, `TARİH`, `ŞİRKET`, `POZİSYON`, `DURUM`) and tab banners (`KARİYER PIPELINE`, `TÜMÜ`, `DEĞERLENDİRİLDİ`).
- **Status Picker Integrity:**
  - Verify that selecting a Turkish status in `overlayStatusPicker` (e.g., `Mülakat`) emits a `PipelineUpdateStatusMsg` containing the canonical English status (`"Interview"`), preserving tracker file integrity.

### 3. Execution Commands for Verification
```bash
# Run unit and regression tests across the dashboard package
cd dashboard && go test -v ./...
```

## RESEARCH COMPLETE
