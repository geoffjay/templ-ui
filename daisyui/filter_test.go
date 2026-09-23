package daisyui

import (
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	got := render(t, Filter(FilterConfig{
		Name: "frameworks",
		Options: []FilterOption{
			{Label: "Svelte"},
			{Label: "Vue", Checked: true},
			{Label: "React"},
		},
		Reset:  true,
		AsForm: true,
	}))
	mustContain(t, got, `<form class="filter">`)
	mustContain(t, got, `<input class="btn btn-square" type="reset" value="×">`)
	mustContain(t, got, `aria-label="Svelte"`)
	mustContain(t, got, `aria-label="Vue"`)
	mustContain(t, got, `aria-label="React"`)
	mustContain(t, got, `name="frameworks"`)
}

func TestFilterWithoutForm(t *testing.T) {
	got := render(t, Filter(FilterConfig{
		Name:    "meta",
		Options: []FilterOption{{Label: "A"}},
		Reset:   true,
		AsForm:  false,
	}))
	mustContain(t, got, `<div class="filter">`)
	mustContain(t, got, `<input class="btn filter-reset" type="radio"`)
	mustContain(t, got, `aria-label="All"`)
}

func TestFilterNoReset(t *testing.T) {
	got := render(t, Filter(FilterConfig{
		Name:    "x",
		Options: []FilterOption{{Label: "A"}},
	}))
	if strings.Contains(got, "reset") {
		t.Errorf("expected no reset, got:\n%s", got)
	}
}
