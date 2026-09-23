package daisyui

import "testing"

func TestToggle(t *testing.T) {
	got := render(t, Toggle(ToggleConfig{
		Name:    "remember",
		Value:   "on",
		Checked: true,
		Color:   "primary",
		Size:    "lg",
	}))
	mustContain(t, got, `<input type="checkbox"`)
	mustContain(t, got, `name="remember"`)
	mustContain(t, got, `value="on"`)
	mustContain(t, got, `class="toggle toggle-primary toggle-lg"`)
	mustContain(t, got, `checked`)
}

func TestToggleDefaults(t *testing.T) {
	got := render(t, Toggle(ToggleConfig{}))
	mustContain(t, got, `<input type="checkbox"`)
	mustContain(t, got, `class="toggle"`)
}

func TestToggleDisabled(t *testing.T) {
	got := render(t, Toggle(ToggleConfig{Disabled: true}))
	mustContain(t, got, `disabled`)
}
