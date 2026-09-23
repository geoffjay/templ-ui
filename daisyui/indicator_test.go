package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestIndicator(t *testing.T) {
	got := render(t, Indicator(IndicatorConfig{
		Items: []IndicatorItemConfig{
			{Class: "badge badge-primary", Children: templ.Raw("New")},
		},
		Children: templ.Raw(`<div class="grid w-32 h-32 bg-base-300 place-items-center">content</div>`),
	}))
	mustContain(t, got, `class="indicator"`)
	mustContain(t, got, `<span class="indicator-item badge badge-primary">`)
	mustContain(t, got, `>New<`)
	mustContain(t, got, `place-items-center">content`)
}

func TestIndicatorPositions(t *testing.T) {
	cases := []struct {
		h, v, expect string
	}{
		{"start", "top", "indicator-start"},
		{"center", "top", "indicator-center"},
		{"end", "top", ""},
		{"start", "middle", "indicator-start indicator-middle"},
		{"start", "bottom", "indicator-start indicator-bottom"},
	}
	for _, c := range cases {
		got := render(t, Indicator(IndicatorConfig{
			Items: []IndicatorItemConfig{
				{Horizontal: c.h, Vertical: c.v, Class: "badge"},
			},
			Children: templ.Raw(`<div></div>`),
		}))
		if c.expect == "" {
			mustContain(t, got, `<span class="indicator-item badge">`)
		} else {
			mustContain(t, got, `<span class="indicator-item `+c.expect+` badge">`)
		}
	}
}
