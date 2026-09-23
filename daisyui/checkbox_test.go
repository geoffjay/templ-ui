package daisyui

import "testing"

func TestCheckbox(t *testing.T) {
	got := render(t, Checkbox(CheckboxConfig{
		Name:    "remember",
		Value:   "yes",
		Checked: true,
		Color:   "primary",
		Size:    "lg",
	}))
	mustContain(t, got, `<input type="checkbox"`)
	mustContain(t, got, `name="remember"`)
	mustContain(t, got, `value="yes"`)
	mustContain(t, got, `class="checkbox checkbox-primary checkbox-lg"`)
	mustContain(t, got, `checked`)
}

func TestCheckboxDefaults(t *testing.T) {
	got := render(t, Checkbox(CheckboxConfig{}))
	mustContain(t, got, `<input type="checkbox"`)
	mustContain(t, got, `class="checkbox"`)
}

func TestCheckboxDisabled(t *testing.T) {
	got := render(t, Checkbox(CheckboxConfig{Disabled: true}))
	mustContain(t, got, `disabled`)
}
