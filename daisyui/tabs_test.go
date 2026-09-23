package daisyui

import "testing"

func TestTabs(t *testing.T) {
	got := render(t, Tabs(TabsConfig{
		Items: []TabItem{
			{Label: "Tab 1"},
			{Label: "Tab 2", Active: true},
			{Label: "Tab 3", Disabled: true},
		},
	}))
	mustContain(t, got, `<div role="tablist" class="tabs">`)
	mustContain(t, got, `<button role="tab" class="tab">Tab 1</button>`)
	mustContain(t, got, `<button role="tab" class="tab tab-active">Tab 2</button>`)
	mustContain(t, got, `<button role="tab" class="tab tab-disabled" disabled>Tab 3</button>`)
}

func TestTabsStyleSizePlacement(t *testing.T) {
	got := render(t, Tabs(TabsConfig{
		Style:     "lift",
		Size:      "sm",
		Placement: "bottom",
		Items:     []TabItem{{Label: "A"}},
	}))
	mustContain(t, got, `class="tabs tabs-lift tabs-bottom tabs-sm"`)
}

func TestTabsHref(t *testing.T) {
	got := render(t, Tabs(TabsConfig{
		Items: []TabItem{{Label: "Home", Href: "/home"}},
	}))
	mustContain(t, got, `<a role="tab" class="tab" href="/home">Home</a>`)
}
