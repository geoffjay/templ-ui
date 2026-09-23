package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestDropdown(t *testing.T) {
	got := render(t, Dropdown(DropdownConfig{
		Label: "Click to open",
		Items: []DropdownItem{
			{Label: "Item 1", Href: "/1"},
			{Label: "Item 2"},
		},
	}))
	mustContain(t, got, `<div class="dropdown">`)
	mustContain(t, got, `<div tabindex="0" role="button" class="btn m-1">Click to open</div>`)
	mustContain(t, got, `class="dropdown-content menu bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm"`)
	mustContain(t, got, `<a href="/1">Item 1</a>`)
	mustContain(t, got, `<button>Item 2</button>`)
}

func TestDropdownDetails(t *testing.T) {
	got := render(t, Dropdown(DropdownConfig{
		Label:      "open",
		UseDetails: true,
		Items:      []DropdownItem{{Label: "A"}},
	}))
	mustContain(t, got, `<details>`)
	mustContain(t, got, "<summary")
	mustContain(t, got, `>open<`)
}

func TestDropdownPlacement(t *testing.T) {
	got := render(t, Dropdown(DropdownConfig{
		Placement: "top end",
		Items:     []DropdownItem{{Label: "A"}},
	}))
	mustContain(t, got, `class="dropdown dropdown-top dropdown-end"`)
}

func TestDropdownHover(t *testing.T) {
	got := render(t, Dropdown(DropdownConfig{
		Modifier: "hover",
		Items:    []DropdownItem{{Label: "A"}},
	}))
	mustContain(t, got, `class="dropdown dropdown-hover"`)
}

func TestDropdownContent(t *testing.T) {
	got := render(t, Dropdown(DropdownConfig{
		Label:   "Click",
		Content: templ.Raw(`<div class="card">Custom</div>`),
	}))
	mustContain(t, got, `<div class="card">Custom</div>`)
}
