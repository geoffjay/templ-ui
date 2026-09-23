package daisyui

import "testing"

func TestSelect(t *testing.T) {
	got := render(t, Select(SelectConfig{
		Name:        "color",
		Placeholder: "Pick a color",
		Options: []SelectOption{
			{Value: "Crimson", Label: "Crimson"},
			{Value: "Amber", Label: "Amber"},
			{Value: "Velvet", Label: "Velvet"},
		},
		Selected: "Amber",
		Color:    "primary",
		Size:     "lg",
	}))
	mustContain(t, got, `<select name="color"`)
	mustContain(t, got, `class="select select-primary select-lg"`)
	mustContain(t, got, `<option disabled selected value="">Pick a color</option>`)
	mustContain(t, got, `<option value="Amber"`)
	mustContain(t, got, `selected`)
}

func TestSelectDefaults(t *testing.T) {
	got := render(t, Select(SelectConfig{
		Options: []SelectOption{{Value: "Zig"}},
	}))
	mustContain(t, got, `<select name="" class="select"`)
	mustContain(t, got, `<option value="Zig"`)
	mustContain(t, got, `>Zig<`)
}

func TestSelectGhost(t *testing.T) {
	got := render(t, Select(SelectConfig{Style: "ghost"}))
	mustContain(t, got, `class="select select-ghost"`)
}

func TestSelectDisabled(t *testing.T) {
	got := render(t, Select(SelectConfig{Disabled: true}))
	mustContain(t, got, `disabled`)
}
