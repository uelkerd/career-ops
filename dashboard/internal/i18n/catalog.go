package i18n

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// Catalog holds all localized UI strings, labels, table headers, and formats
// for the Go TUI Dashboard. It provides a static, zero-dependency translation architecture.
type Catalog struct {
	// Screen banners & general
	AppTitle       string
	OffersSummary  string
	NoOffersMatch  string
	LoadingPreview string

	// Tabs & filters
	TabAll       string
	TabEvaluated string
	TabApplied   string
	TabInterview string
	TabTop       string
	TabSkip      string
	TabRejected  string
	TabDiscarded string

	// Table column headers
	ColFit      string
	ColApplied  string
	ColCompany  string
	ColRole     string
	ColStatus   string
	ColLocation string
	ColPay      string
	ColLast     string

	// Preview labels
	LabelLoc     string
	LabelPay     string
	LabelLast    string
	LabelRemote  string
	LabelOutcome string

	// Work modes
	ModeRemote     string
	ModeRemoteFlex string
	ModeHybrid     string
	ModeFull       string

	// Progress screen
	ProgressTitle   string
	ProgressSummary string
	FunnelTitle     string
	ScoresTitle     string
	RatesTitle      string
	WeeklyTitle     string
	ActiveInfo      string

	// Relative dates
	TimeToday     string
	TimeYesterday string
	TimeDaysAgo   string

	// Status display names
	StatusEvaluated string
	StatusApplied   string
	StatusResponded string
	StatusInterview string
	StatusOffer     string
	StatusRejected  string
	StatusDiscarded string
	StatusSkip      string
}

// StatusLabel returns the localized display label for a canonical status ID
// (interview, offer, responded, applied, evaluated, skip, rejected, discarded).
func (c *Catalog) StatusLabel(norm string) string {
	switch strings.ToLower(strings.TrimSpace(norm)) {
	case "interview":
		return c.StatusInterview
	case "offer":
		return c.StatusOffer
	case "responded":
		return c.StatusResponded
	case "applied":
		return c.StatusApplied
	case "evaluated":
		return c.StatusEvaluated
	case "skip":
		return c.StatusSkip
	case "rejected":
		return c.StatusRejected
	case "discarded":
		return c.StatusDiscarded
	default:
		return norm
	}
}

// FormatTimeAgo renders an ISO date as a relative duration in calendar days using localized strings:
// "today", "yesterday", or "Nd ago" (or Turkish equivalents).
func (c *Catalog) FormatTimeAgo(dateStr string) string {
	t, err := time.ParseInLocation("2006-01-02", dateStr, time.Local)
	if err != nil {
		return dateStr
	}
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	contactDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	days := int(math.Round(today.Sub(contactDay).Hours() / 24))
	switch {
	case days <= 0:
		return c.TimeToday
	case days == 1:
		return c.TimeYesterday
	default:
		return fmt.Sprintf(c.TimeDaysAgo, days)
	}
}

// En is the static English translation catalog.
var En = Catalog{
	// Screen banners & general
	AppTitle:       "CAREER PIPELINE",
	OffersSummary:  "%d offers | Avg %s/5",
	NoOffersMatch:  "No offers match this filter",
	LoadingPreview: "Loading preview...",

	// Tabs & filters
	TabAll:       "ALL",
	TabEvaluated: "EVALUATED",
	TabApplied:   "APPLIED",
	TabInterview: "INTERVIEW",
	TabTop:       "TOP ≥4",
	TabSkip:      "SKIP",
	TabRejected:  "REJECTED",
	TabDiscarded: "DISCARDED",

	// Table column headers
	ColFit:      "FIT",
	ColApplied:  "APPLIED",
	ColCompany:  "COMPANY",
	ColRole:     "ROLE",
	ColStatus:   "STATUS",
	ColLocation: "LOCATION",
	ColPay:      "PAY",
	ColLast:     "LAST",

	// Preview labels
	LabelLoc:     "Loc: ",
	LabelPay:     "Pay: ",
	LabelLast:    "Last contact: ",
	LabelRemote:  "Remote: ",
	LabelOutcome: "Outcome: ",

	// Work modes
	ModeRemote:     "Remote",
	ModeRemoteFlex: "RemoteFlex",
	ModeHybrid:     "Hybrid",
	ModeFull:       "Full",

	// Progress screen
	ProgressTitle:   "SEARCH PROGRESS",
	ProgressSummary: "%d evaluated | %.1f avg score",
	FunnelTitle:     "Pipeline Funnel",
	ScoresTitle:     "Score Distribution",
	RatesTitle:      "Conversion Rates",
	WeeklyTitle:     "Weekly Activity",
	ActiveInfo:      "%d active applications | %d total offers",

	// Relative dates
	TimeToday:     "today",
	TimeYesterday: "yesterday",
	TimeDaysAgo:   "%dd ago",

	// Status display names
	StatusEvaluated: "Evaluated",
	StatusApplied:   "Applied",
	StatusResponded: "Responded",
	StatusInterview: "Interview",
	StatusOffer:     "Offer",
	StatusRejected:  "Rejected",
	StatusDiscarded: "Discarded",
	StatusSkip:      "SKIP",
}

// Tr is the static Turkish translation catalog.
var Tr = Catalog{
	// Screen banners & general
	AppTitle:       "KARİYER PIPELINE",
	OffersSummary:  "%d ilan | Ort %s/5",
	NoOffersMatch:  "Bu filtreye uyan ilan yok",
	LoadingPreview: "Önizleme yükleniyor...",

	// Tabs & filters
	TabAll:       "TÜMÜ",
	TabEvaluated: "DEĞERLENDİRİLDİ",
	TabApplied:   "BAŞVURULDU",
	TabInterview: "MÜLAKAT",
	TabTop:       "EN İYİ ≥4",
	TabSkip:      "UYGUN DEĞİL",
	TabRejected:  "REDDEDİLDİ",
	TabDiscarded: "İPTAL",

	// Table column headers
	ColFit:      "UYUM",
	ColApplied:  "TARİH",
	ColCompany:  "ŞİRKET",
	ColRole:     "POZİSYON",
	ColStatus:   "DURUM",
	ColLocation: "KONUM",
	ColPay:      "ÜCRET",
	ColLast:     "SON",

	// Preview labels
	LabelLoc:     "Konum: ",
	LabelPay:     "Ücret: ",
	LabelLast:    "Son iletişim: ",
	LabelRemote:  "Çalışma Şekli: ",
	LabelOutcome: "Sonuç: ",

	// Work modes
	ModeRemote:     "Uzaktan",
	ModeRemoteFlex: "Uzaktan (Esnek)",
	ModeHybrid:     "Hibrit",
	ModeFull:       "Ofiste",

	// Progress screen
	ProgressTitle:   "TAKİP İLERLEMESİ",
	ProgressSummary: "%d değerlendirildi | %.1f ort. puan",
	FunnelTitle:     "Pipeline Hunisi",
	ScoresTitle:     "Puan Dağılımı",
	RatesTitle:      "Dönüşüm Oranları",
	WeeklyTitle:     "Haftalık Aktivite",
	ActiveInfo:      "%d aktif başvuru | %d toplam teklif",

	// Relative dates
	TimeToday:     "bugün",
	TimeYesterday: "dün",
	TimeDaysAgo:   "%d gün önce",

	// Status display names
	StatusEvaluated: "Değerlendirildi",
	StatusApplied:   "Başvuruldu",
	StatusResponded: "Yanıt Verildi",
	StatusInterview: "Mülakat",
	StatusOffer:     "Teklif",
	StatusRejected:  "Reddedildi",
	StatusDiscarded: "İptal Edildi",
	StatusSkip:      "Uygun Değil",
}

// Current points to the active language catalog. Defaults to English (&En).
var Current = &En

// SetLang sets the active catalog based on language code prefix (e.g., "tr", "tr_TR" -> &Tr).
func SetLang(lang string) {
	if strings.HasPrefix(strings.ToLower(strings.TrimSpace(lang)), "tr") {
		Current = &Tr
	} else {
		Current = &En
	}
}

// ToggleLang switches Current between &En and &Tr.
func ToggleLang() {
	if Current == &En {
		Current = &Tr
	} else {
		Current = &En
	}
}

// GetLang returns the active language code ("tr" if Current == &Tr, else "en").
func GetLang() string {
	if Current == &Tr {
		return "tr"
	}
	return "en"
}
