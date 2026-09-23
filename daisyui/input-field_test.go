package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestInput(t *testing.T) {
	got := render(t, Input(InputConfig{
		Type:        "email",
		Name:        "email",
		Placeholder: "mail@site.com",
		Color:       "primary",
		Size:        "lg",
		Required:    true,
	}))
	mustContain(t, got, `<input type="email" name="email"`)
	mustContain(t, got, `placeholder="mail@site.com"`)
	mustContain(t, got, `class="input input-primary input-lg"`)
	mustContain(t, got, `required`)
}

func TestInputDefaults(t *testing.T) {
	got := render(t, Input(InputConfig{}))
	mustContain(t, got, `<input type="text"`)
	mustContain(t, got, `class="input"`)
}

func TestInputGhost(t *testing.T) {
	got := render(t, Input(InputConfig{Style: "ghost"}))
	mustContain(t, got, `class="input input-ghost"`)
}

func TestInputDisabled(t *testing.T) {
	got := render(t, Input(InputConfig{Disabled: true}))
	mustContain(t, got, `disabled`)
}

func TestInputLabel(t *testing.T) {
	got := render(t, InputLabel(InputLabelConfig{
		Color:     "primary",
		Validator: true,
	}, templ.Raw(`<input type="email"/>`)))
	mustContain(t, got, `<label class="input input-primary validator">`)
	mustContain(t, got, `<input type="email"/>`)
}
