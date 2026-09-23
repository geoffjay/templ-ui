package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestTooltip(t *testing.T) {
	got := render(t, Tooltip(TooltipConfig{
		Tip:      "hello",
		Children: templ.Raw(`<button class="btn">Hover me</button>`),
	}))
	mustContain(t, got, `class="tooltip"`)
	mustContain(t, got, `data-tip="hello"`)
	mustContain(t, got, `<button class="btn">Hover me</button>`)
}

func TestTooltipPosition(t *testing.T) {
	for _, p := range []string{"top", "bottom", "left", "right"} {
		got := render(t, Tooltip(TooltipConfig{Tip: "x", Position: p}))
		if p == "top" {
			mustContain(t, got, `class="tooltip"`)
		} else {
			mustContain(t, got, "tooltip-"+p)
		}
	}
}

func TestTooltipColorAndOpen(t *testing.T) {
	got := render(t, Tooltip(TooltipConfig{
		Tip:   "primary",
		Color: "primary",
		Open:  true,
	}))
	mustContain(t, got, `class="tooltip tooltip-primary tooltip-open"`)
	mustContain(t, got, `data-tip="primary"`)
}

func TestTooltipContent(t *testing.T) {
	got := render(t, Tooltip(TooltipConfig{
		Content:  templ.Raw(`<div class="animate-bounce text-orange-400">Wow!</div>`),
		Children: templ.Raw(`<button class="btn">Hover me</button>`),
	}))
	mustContain(t, got, `<div class="tooltip-content">`)
	mustContain(t, got, `<div class="animate-bounce text-orange-400">Wow!</div>`)
}
