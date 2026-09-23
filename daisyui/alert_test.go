package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestAlert(t *testing.T) {
	got := render(t, Alert(AlertConfig{
		Message: "12 unread messages. Tap to see.",
		Color:   "info",
		Class:   "w-full",
	}))
	mustContain(t, got, `role="alert"`)
	mustContain(t, got, `class="alert alert-info w-full"`)
	mustContain(t, got, ">12 unread messages. Tap to see.<")
}

func TestAlertTitleDesc(t *testing.T) {
	got := render(t, Alert(AlertConfig{
		Title:     "New message!",
		Desc:      "You have 1 unread message",
		Direction: "vertical",
		Class:     "w-full sm:alert-horizontal",
	}))
	mustContain(t, got, `class="alert alert-vertical w-full sm:alert-horizontal"`)
	mustContain(t, got, `<h3 class="font-bold">New message!</h3>`)
	mustContain(t, got, `<div class="text-xs">You have 1 unread message</div>`)
}

func TestAlertStyles(t *testing.T) {
	for _, s := range []string{"outline", "dash", "soft"} {
		got := render(t, Alert(AlertConfig{Style: s}))
		mustContain(t, got, "alert-"+s)
	}
}

func TestAlertActionsAndIcon(t *testing.T) {
	got := render(t, Alert(AlertConfig{
		Message: "we use cookies for no reason.",
		Icon:    templ.Raw(`<svg class="stroke-info shrink-0 w-6 h-6"></svg>`),
		Actions: templ.Raw(`<div><button class="btn btn-sm">Deny</button></div>`),
	}))
	mustContain(t, got, `<svg class="stroke-info shrink-0 w-6 h-6"></svg>`)
	mustContain(t, got, `<button class="btn btn-sm">Deny</button>`)
}
