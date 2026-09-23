package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestBadge(t *testing.T) {
	got := render(t, Badge(BadgeConfig{Text: "Badge"}))
	mustContain(t, got, `<span class="badge">`)
	mustContain(t, got, ">Badge<")
}

func TestBadgeColorSizeStyle(t *testing.T) {
	got := render(t, Badge(BadgeConfig{
		Text:  "Primary",
		Color: "primary",
		Style: "outline",
		Size:  "lg",
	}))
	mustContain(t, got, `class="badge badge-primary badge-outline badge-lg"`)
	mustContain(t, got, ">Primary<")
}

func TestBadgeGhost(t *testing.T) {
	got := render(t, Badge(BadgeConfig{Text: "ghost", Style: "ghost"}))
	mustContain(t, got, `class="badge badge-ghost"`)
}

func TestBadgeChildren(t *testing.T) {
	got := render(t, Badge(BadgeConfig{
		Color:    "info",
		Children: templ.Raw(`<svg class="size-[1em]"></svg>Info`),
	}))
	mustContain(t, got, `<span class="badge badge-info">`)
	mustContain(t, got, `<svg class="size-[1em]"></svg>Info`)
}
