package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestToast(t *testing.T) {
	got := render(t, Toast(ToastConfig{
		Horizontal: "start",
		Vertical:   "top",
		Class:      "absolute",
		Children:   templ.Raw(`<div class="alert alert-info"><span>New mail arrived.</span></div>`),
	}))
	mustContain(t, got, `class="toast toast-start toast-top absolute"`)
	mustContain(t, got, `<div class="alert alert-info">`)
	mustContain(t, got, `New mail arrived.`)
}

func TestToastDefaults(t *testing.T) {
	got := render(t, Toast(ToastConfig{
		Horizontal: "end",
		Vertical:   "bottom",
	}))
	mustContain(t, got, `class="toast"`)
}

func TestToastCenterMiddle(t *testing.T) {
	got := render(t, Toast(ToastConfig{
		Horizontal: "center",
		Vertical:   "middle",
	}))
	mustContain(t, got, `class="toast toast-center toast-middle"`)
}
