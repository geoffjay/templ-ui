package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestDrawer(t *testing.T) {
	got := render(t, Drawer(DrawerConfig{
		ButtonLabel: "Open drawer",
		SideItems:   []DrawerSidebarItem{{Label: "Sidebar Item 1", Href: "/1"}},
	}))
	mustContain(t, got, `class="drawer"`)
	mustContain(t, got, `id="drawer-toggle"`)
	mustContain(t, got, `class="drawer-toggle"`)
	mustContain(t, got, `<label for="drawer-toggle" class="btn drawer-button">Open drawer</label>`)
	mustContain(t, got, `<label for="drawer-toggle" aria-label="close sidebar" class="drawer-overlay"></label>`)
	mustContain(t, got, `<a href="/1">Sidebar Item 1</a>`)
}

func TestDrawerEndOpen(t *testing.T) {
	got := render(t, Drawer(DrawerConfig{
		End:       true,
		Open:      true,
		ToggleID:  "my-drawer",
		SideItems: []DrawerSidebarItem{{Label: "X"}},
	}))
	mustContain(t, got, `class="drawer drawer-end drawer-open"`)
	mustContain(t, got, `id="my-drawer"`)
}

func TestDrawerSide(t *testing.T) {
	got := render(t, Drawer(DrawerConfig{
		Side: templ.Raw(`<aside class="my-side"></aside>`),
	}))
	mustContain(t, got, `<aside class="my-side"></aside>`)
}
