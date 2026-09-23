package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestDock(t *testing.T) {
	got := render(t, Dock(DockConfig{
		Items: []DockItem{
			{Label: "Home", Icon: templ.Raw(`<svg class="size-[1.2em]"></svg>`)},
			{Label: "Inbox", Active: true, Icon: templ.Raw(`<svg></svg>`)},
			{Label: "Settings", Href: "/settings"},
		},
	}))
	mustContain(t, got, `<div class="dock">`)
	mustContain(t, got, `<button class="">`)
	mustContain(t, got, `<span class="dock-label">Home</span>`)
	mustContain(t, got, `<button class="dock-active">`)
	mustContain(t, got, `<a class="" href="/settings">`)
	mustContain(t, got, `<span class="dock-label">Settings</span>`)
}

func TestDockSize(t *testing.T) {
	got := render(t, Dock(DockConfig{
		Size:  "sm",
		Items: []DockItem{{Label: "A"}},
	}))
	mustContain(t, got, `class="dock dock-sm"`)
}
