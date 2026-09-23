package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestMenu(t *testing.T) {
	got := render(t, Menu(MenuConfig{
		Items: []MenuItem{
			{Label: "Item 1"},
			{Label: "Item 2", Active: true},
			{Label: "Item 3", Disabled: true},
		},
		Class: "bg-base-200 rounded-box w-56",
	}))
	mustContain(t, got, `<ul class="menu bg-base-200 rounded-box w-56">`)
	mustContain(t, got, `<button class="">Item 1 </button>`)
	mustContain(t, got, `<button class="menu-active">Item 2 </button>`)
	mustContain(t, got, `<li class="menu-disabled">`)
	mustContain(t, got, `disabled>Item 3 </button>`)
}

func TestMenuTitle(t *testing.T) {
	got := render(t, Menu(MenuConfig{
		Title: "Title",
		Items: []MenuItem{{Label: "Item 1", Href: "/x"}},
	}))
	mustContain(t, got, `<li class="menu-title">Title</li>`)
	mustContain(t, got, `<a href="/x"`)
	mustContain(t, got, `>Item 1 </a>`)
}

func TestMenuHorizontalAndSize(t *testing.T) {
	got := render(t, Menu(MenuConfig{
		Items:     []MenuItem{{Label: "A"}},
		Direction: "horizontal",
		Size:      "lg",
	}))
	mustContain(t, got, `class="menu menu-horizontal menu-lg"`)
}

func TestMenuSubmenu(t *testing.T) {
	got := render(t, Menu(MenuConfig{
		Items: []MenuItem{
			{Label: "Parent", Submenu: []MenuItem{{Label: "Sub 1"}}},
		},
	}))
	mustContain(t, got, `>Parent </button>`)
	mustContain(t, got, `<ul>`)
	mustContain(t, got, `>Sub 1 </button>`)
}

func TestMenuCollapsible(t *testing.T) {
	got := render(t, Menu(MenuConfig{
		Items: []MenuItem{
			{Label: "Parent", Collapsible: true, Open: true, Submenu: []MenuItem{{Label: "Sub 1"}}},
		},
	}))
	mustContain(t, got, `<details open>`)
	mustContain(t, got, "<summary")
	mustContain(t, got, `>Parent<`)
	mustContain(t, got, `>Sub 1 </button>`)
}

func TestMenuBadgesAndIcons(t *testing.T) {
	got := render(t, Menu(MenuConfig{
		Items: []MenuItem{{
			Label: "Inbox",
			Badge: "99+",
			Icon:  templ.Raw(`<svg class="h-5 w-5"></svg>`),
		}},
	}))
	mustContain(t, got, `<svg class="h-5 w-5"></svg>`)
	mustContain(t, got, `>Inbox `)
	mustContain(t, got, `<span class="badge badge-xs">99+</span>`)
}

func TestMenuTooltips(t *testing.T) {
	got := render(t, Menu(MenuConfig{
		Items:       []MenuItem{{Label: "A", Tooltip: "Home"}},
		UseTooltips: true,
	}))
	mustContain(t, got, `class="tooltip"`)
	mustContain(t, got, `data-tip="Home"`)
}
