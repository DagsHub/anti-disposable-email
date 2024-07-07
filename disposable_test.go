package disposable

import "testing"

func Test_ParseEmail_Normalize_gmail(t *testing.T) {
	email := "r2d2@gmail.com"
	p, err := ParseEmail(email)
	if err != nil {
		t.Fatal(err)
	}

	if p.Normalized != "r2d2" {
		t.Errorf("expected normalized to be 'r2d2' but got %s", p.Normalized)
	}
	if p.Preferred != "r2d2" {
		t.Errorf("expected preferred to be 'r2d2' but got %s", p.Preferred)
	}
	if p.Extra != "" {
		t.Errorf("expected extra to be '' but got %s", p.Extra)
	}
	if p.Disposable {
		t.Errorf("expected disposable to be false but got true")
	}
	if p.Domain != "gmail.com" {
		t.Errorf("expected domain to be 'gmail.com' but got %s", p.Domain)
	}
	if p.LocalPart != "r2d2" {
		t.Errorf("expected local part to be 'r2d2' but got %s", p.LocalPart)
	}
	if p.Email != email {
		t.Errorf("expected email to be %s but got %s", email, p.Email)
	}
}

func Test_ParseEmail_Normalize_gmail_plus(t *testing.T) {
	email := "R2.D2+junk@gmail.com"
	p, err := ParseEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	if p.Normalized != "r2d2" {
		t.Errorf("expected normalized to be 'r2d2' but got %s", p.Normalized)
	}
	if p.Preferred != "R2.D2" {
		t.Errorf("expected preferred to be 'R2.D2' but got %s", p.Preferred)
	}
	if p.Extra != "junk" {
		t.Errorf("expected extra to be 'junk' but got %s", p.Extra)
	}
	if p.Disposable {
		t.Errorf("expected disposable to be false but got true")
	}
	if p.Domain != "gmail.com" {
		t.Errorf("expected domain to be 'gmail.com' but got %s", p.Domain)
	}
	if p.LocalPart != "R2.D2+junk" {
		t.Errorf("expected local part to be 'R2.D2+junk' but got %s", p.LocalPart)
	}
	if p.Email != email {
		t.Errorf("expected email to be %s but got %s", email, p.Email)
	}
}

func Test_ParseEmail_Normalize_gmail_multi_plus(t *testing.T) {
	email := "R2.D2+junk+morejunk@gmail.com"
	p, err := ParseEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	if p.Normalized != "r2d2" {
		t.Errorf("expected normalized to be 'r2d2' but got %s", p.Normalized)
	}
	if p.Preferred != "R2.D2" {
		t.Errorf("expected preferred to be 'R2.D2' but got %s", p.Preferred)
	}
	if p.Extra != "junk+morejunk" {
		t.Errorf("expected extra to be 'junk+morejunk' but got %s", p.Extra)
	}
	if p.Disposable {
		t.Errorf("expected disposable to be false but got true")
	}
	if p.Domain != "gmail.com" {
		t.Errorf("expected domain to be 'gmail.com' but got %s", p.Domain)
	}
	if p.LocalPart != "R2.D2+junk+morejunk" {
		t.Errorf("expected local part to be 'R2.D2+junk+morejunk' but got %s", p.LocalPart)
	}
	if p.Email != email {
		t.Errorf("expected email to be %s but got %s", email, p.Email)
	}
}

func Test_ParseEmail_Normalize_non_gmail_plus(t *testing.T) {
	email := "R2.D2+junk@yahoo.com"
	p, err := ParseEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	if p.Normalized != "r2.d2+junk" {
		t.Errorf("expected normalized to be 'r2.d2+junk' but got %s", p.Normalized)
	}
	if p.Preferred != "R2.D2+junk" {
		t.Errorf("expected preferred to be 'R2.D2+junk' but got %s", p.Preferred)
	}
	if p.Extra != "" {
		t.Errorf("expected extra to be '' but got %s", p.Extra)
	}
	if p.Disposable {
		t.Errorf("expected disposable to be false but got true")
	}
	if p.Domain != "yahoo.com" {
		t.Errorf("expected domain to be 'gmail.com' but got %s", p.Domain)
	}
	if p.LocalPart != "R2.D2+junk" {
		t.Errorf("expected local part to be 'R2.D2+junk' but got %s", p.LocalPart)
	}
	if p.Email != email {
		t.Errorf("expected email to be %s but got %s", email, p.Email)
	}
}

func Test_ParseEmail_Disposable_domain(t *testing.T) {
	email := "example@mailto.plus"
	p, err := ParseEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	if !p.Disposable {
		t.Errorf("expected disposable to be true but got false")
	}
}

func Test_IsDisposable_Subdomain(t *testing.T) {
	listAtomic.Store(map[string]struct{}{
		"somewhere.eu.org": {},
	})

	p, err := ParseEmail("someone@someplace.somewhere.eu.org")
	if err != nil {
		t.Fatal(err)
	}
	if !p.Disposable {
		t.Errorf("expected disposable to be true but got false")
	}

	p, err = ParseEmail("someone@somewhere.eu.org")
	if err != nil {
		t.Fatal(err)
	}
	if !p.Disposable {
		t.Errorf("expected disposable to be true but got false")
	}

	p, err = ParseEmail("someone@eu.org")
	if err != nil {
		t.Fatal(err)
	}
	if p.Disposable {
		t.Errorf("expected disposable to be false but got true")
	}

	p, err = ParseEmail("someone@eu.org")
	if err != nil {
		t.Fatal(err)
	}
	if p.Disposable {
		t.Errorf("expected disposable to be false but got true")
	}

	p, err = ParseEmail("someone@elsewhere.eu.org")
	if err != nil {
		t.Fatal(err)
	}
	if p.Disposable {
		t.Errorf("expected disposable to be false but got true")
	}
}
