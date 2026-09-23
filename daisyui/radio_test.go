package daisyui

import "testing"

func TestRadio(t *testing.T) {
	got := render(t, Radio(RadioConfig{
		Name:  "radio-1",
		Color: "primary",
		Size:  "lg",
		Options: []RadioOption{
			{Value: "a", Label: "Option A", Checked: true},
			{Value: "b", Label: "Option B"},
		},
	}))
	mustContain(t, got, `<input type="radio"`)
	mustContain(t, got, `name="radio-1"`)
	mustContain(t, got, `value="a"`)
	mustContain(t, got, `class="radio radio-primary radio-lg"`)
	mustContain(t, got, `aria-label="Option A"`)
	mustContain(t, got, `checked`)
}

func TestRadioDefaults(t *testing.T) {
	got := render(t, Radio(RadioConfig{
		Name: "radio-1",
		Options: []RadioOption{
			{Value: "a"},
		},
	}))
	mustContain(t, got, `<input type="radio"`)
	mustContain(t, got, `class="radio"`)
}

func TestRadioSingle(t *testing.T) {
	got := render(t, RadioSingle(RadioSingleConfig{
		Name:     "radio-2",
		Value:    "x",
		Checked:  true,
		Color:    "secondary",
		Disabled: true,
	}))
	mustContain(t, got, `<input type="radio"`)
	mustContain(t, got, `name="radio-2"`)
	mustContain(t, got, `value="x"`)
	mustContain(t, got, `class="radio radio-secondary"`)
	mustContain(t, got, `checked`)
	mustContain(t, got, `disabled`)
}
