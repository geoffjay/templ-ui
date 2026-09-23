package daisyui

import (
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestWindow(t *testing.T) {
	got := render(t, Window(WindowConfig{
		Bordered:   true,
		Background: true,
		Children:   templ.Raw(`<span>Hello!</span>`),
	}))
	mustContain(t, got, `<div class="mockup-window w-full border border-base-300 bg-base-100">`)
	mustContain(t, got, `class="grid place-content-center border-t border-base-300 h-80"`)
	mustContain(t, got, `<span>Hello!</span>`)
}

func TestWindowPlain(t *testing.T) {
	got := render(t, Window(WindowConfig{
		Bordered:   false,
		Background: false,
		BodyClass:  "h-40",
	}))
	mustContain(t, got, `<div class="mockup-window w-full">`)
	mustContain(t, got, `class="grid place-content-center h-40"`)
	if strings.Contains(got, "border") {
		t.Errorf("expected no border classes; got:\n%s", got)
	}
}
