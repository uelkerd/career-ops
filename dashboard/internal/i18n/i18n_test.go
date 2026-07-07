package i18n

import (
	"testing"
	"time"
)

func TestStatusLabel(t *testing.T) {
	tests := []struct {
		norm string
		en   string
		tr   string
	}{
		{"interview", "Interview", "Mülakat"},
		{"offer", "Offer", "Teklif"},
		{"responded", "Responded", "Yanıt Verildi"},
		{"applied", "Applied", "Başvuruldu"},
		{"evaluated", "Evaluated", "Değerlendirildi"},
		{"skip", "SKIP", "Uygun Değil"},
		{"rejected", "Rejected", "Reddedildi"},
		{"discarded", "Discarded", "İptal Edildi"},
		{"unknown", "unknown", "unknown"},
	}

	for _, tt := range tests {
		if got := En.StatusLabel(tt.norm); got != tt.en {
			t.Errorf("En.StatusLabel(%q) = %q; want %q", tt.norm, got, tt.en)
		}
		if got := Tr.StatusLabel(tt.norm); got != tt.tr {
			t.Errorf("Tr.StatusLabel(%q) = %q; want %q", tt.norm, got, tt.tr)
		}
	}
}

func TestFormatTimeAgo(t *testing.T) {
	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	threeDaysAgo := time.Now().AddDate(0, 0, -3).Format("2006-01-02")
	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")

	// English tests
	if got := En.FormatTimeAgo(today); got != "today" {
		t.Errorf("En.FormatTimeAgo(today) = %q; want \"today\"", got)
	}
	if got := En.FormatTimeAgo(yesterday); got != "yesterday" {
		t.Errorf("En.FormatTimeAgo(yesterday) = %q; want \"yesterday\"", got)
	}
	if got := En.FormatTimeAgo(threeDaysAgo); got != "3d ago" {
		t.Errorf("En.FormatTimeAgo(3d ago) = %q; want \"3d ago\"", got)
	}
	if got := En.FormatTimeAgo(tomorrow); got != "today" {
		t.Errorf("En.FormatTimeAgo(tomorrow) = %q; want \"today\"", got)
	}
	if got := En.FormatTimeAgo("not-a-date"); got != "not-a-date" {
		t.Errorf("En.FormatTimeAgo(invalid) = %q; want \"not-a-date\"", got)
	}

	// Turkish tests
	if got := Tr.FormatTimeAgo(today); got != "bugün" {
		t.Errorf("Tr.FormatTimeAgo(today) = %q; want \"bugün\"", got)
	}
	if got := Tr.FormatTimeAgo(yesterday); got != "dün" {
		t.Errorf("Tr.FormatTimeAgo(yesterday) = %q; want \"dün\"", got)
	}
	if got := Tr.FormatTimeAgo(threeDaysAgo); got != "3 gün önce" {
		t.Errorf("Tr.FormatTimeAgo(3d ago) = %q; want \"3 gün önce\"", got)
	}
	if got := Tr.FormatTimeAgo(tomorrow); got != "bugün" {
		t.Errorf("Tr.FormatTimeAgo(tomorrow) = %q; want \"bugün\"", got)
	}
	if got := Tr.FormatTimeAgo("not-a-date"); got != "not-a-date" {
		t.Errorf("Tr.FormatTimeAgo(invalid) = %q; want \"not-a-date\"", got)
	}
}

func TestRuntimeLanguageManagement(t *testing.T) {
	// Reset to En initially
	Current = &En

	if got := GetLang(); got != "en" {
		t.Errorf("initial GetLang() = %q; want \"en\"", got)
	}

	SetLang("tr")
	if Current != &Tr || GetLang() != "tr" {
		t.Errorf("after SetLang(\"tr\"), GetLang() = %q; want \"tr\"", GetLang())
	}

	SetLang("tr_TR")
	if Current != &Tr || GetLang() != "tr" {
		t.Errorf("after SetLang(\"tr_TR\"), GetLang() = %q; want \"tr\"", GetLang())
	}

	SetLang("en")
	if Current != &En || GetLang() != "en" {
		t.Errorf("after SetLang(\"en\"), GetLang() = %q; want \"en\"", GetLang())
	}

	SetLang("fr") // unknown language falls back to en
	if Current != &En || GetLang() != "en" {
		t.Errorf("after SetLang(\"fr\"), GetLang() = %q; want \"en\"", GetLang())
	}

	// Test ToggleLang
	ToggleLang()
	if Current != &Tr || GetLang() != "tr" {
		t.Errorf("after ToggleLang() from En, GetLang() = %q; want \"tr\"", GetLang())
	}

	ToggleLang()
	if Current != &En || GetLang() != "en" {
		t.Errorf("after ToggleLang() from Tr, GetLang() = %q; want \"en\"", GetLang())
	}
}
