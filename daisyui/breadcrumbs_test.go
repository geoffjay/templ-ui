package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestBreadcrumbs(t *testing.T) {
	got := render(t, Breadcrumbs(BreadcrumbsConfig{
		Items: []BreadcrumbItem{
			{Label: "Home", Href: "/"},
			{Label: "Documents", Href: "/docs"},
			{Label: "Add Document"},
		},
		Class: "text-sm",
	}))
	mustContain(t, got, `class="breadcrumbs text-sm"`)
	mustContain(t, got, `<a href="/">`)
	mustContain(t, got, `>Home<`)
	mustContain(t, got, `<a href="/docs">`)
	mustContain(t, got, `>Documents<`)
	mustContain(t, got, `<span class="inline-flex items-center gap-2">`)
	mustContain(t, got, `>Add Document<`)
}

func TestBreadcrumbsIcons(t *testing.T) {
	got := render(t, Breadcrumbs(BreadcrumbsConfig{
		Items: []BreadcrumbItem{
			{Label: "Home", Href: "/", Icon: templ.Raw(`<svg class="h-4 w-4"></svg>`)},
		},
	}))
	mustContain(t, got, `<svg class="h-4 w-4"></svg>`)
	mustContain(t, got, `>Home<`)
}
