package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestFieldset(t *testing.T) {
	got := render(t, Fieldset(FieldsetConfig{
		Legend: "Page title",
		Class:  "w-xs",
	}, templ.Raw(`<input type="text" class="input"/>`)))
	mustContain(t, got, `<fieldset class="fieldset w-xs">`)
	mustContain(t, got, `<legend class="fieldset-legend">Page title</legend>`)
	mustContain(t, got, `<input type="text" class="input"/>`)
}

func TestFieldsetDefaults(t *testing.T) {
	got := render(t, Fieldset(FieldsetConfig{}, templ.Raw("")))
	mustContain(t, got, `<fieldset class="fieldset">`)
}
