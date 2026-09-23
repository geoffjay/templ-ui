package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestFooter(t *testing.T) {
	got := render(t, Footer(FooterConfig{
		Sections: []FooterSection{
			{
				Title: "Services",
				Links: []FooterLink{{Label: "Branding", Href: "/branding"}},
			},
			{
				Title: "Company",
				Links: []FooterLink{{Label: "About us", Href: "/about"}},
			},
		},
		Class: "p-10 bg-neutral text-neutral-content",
	}))
	mustContain(t, got, `<footer class="footer p-10 bg-neutral text-neutral-content">`)
	mustContain(t, got, `<h6 class="footer-title">Services</h6>`)
	mustContain(t, got, `<a class="link link-hover" href="/branding">Branding</a>`)
}

func TestFooterDirectionAndCenter(t *testing.T) {
	got := render(t, Footer(FooterConfig{
		Direction: "horizontal",
		Center:    true,
		Sections:  []FooterSection{{Title: "X"}},
	}))
	mustContain(t, got, `class="footer footer-horizontal footer-center"`)
}

func TestFooterAsideAndCopyright(t *testing.T) {
	got := render(t, Footer(FooterConfig{
		Aside:     templ.Raw(`<aside><svg class="logo"></svg></aside>`),
		Copyright: "© 2026 Acme",
		Sections:  []FooterSection{{Title: "S"}},
	}))
	mustContain(t, got, `<svg class="logo"></svg>`)
	mustContain(t, got, `© 2026 Acme`)
}
