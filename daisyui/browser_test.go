package daisyui

import (
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestBrowser(t *testing.T) {
	got := render(t, Browser(BrowserConfig{
		URL:        "https://daisyui.com",
		Bordered:   true,
		Background: true,
		Children:   templ.Raw(`<span>Hello!</span>`),
	}))
	mustContain(t, got, `<div class="mockup-browser w-full border border-base-300 bg-base-100">`)
	mustContain(t, got, `<div class="mockup-browser-toolbar">`)
	mustContain(t, got, `<div class="input">https://daisyui.com</div>`)
	mustContain(t, got, `class="grid place-content-center border-t border-base-300 h-80"`)
	mustContain(t, got, `<span>Hello!</span>`)
}

func TestBrowserPlain(t *testing.T) {
	got := render(t, Browser(BrowserConfig{
		URL:        "",
		Bordered:   false,
		Background: false,
		BodyClass:  "h-40",
	}))
	mustContain(t, got, `<div class="mockup-browser w-full">`)
	mustContain(t, got, `class="grid place-content-center h-40"`)
	if strings.Contains(got, "border") {
		t.Errorf("expected no border classes; got:\n%s", got)
	}
}

func TestBrowserCustomToolbar(t *testing.T) {
	got := render(t, Browser(BrowserConfig{
		Toolbar: templ.Raw(`<div class="input">custom-url</div><div class="ml-2">star</div>`),
	}))
	mustContain(t, got, `<div class="mockup-browser-toolbar">`)
	mustContain(t, got, `custom-url`)
	mustContain(t, got, `star`)
}
