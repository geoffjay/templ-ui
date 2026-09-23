package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestStat(t *testing.T) {
	got := render(t, Stat(StatConfig{
		Direction: "horizontal",
		Class:     "shadow",
		Items: []StatItem{
			{Title: "Total Page Views", Value: "89,400", Desc: "21% more than last month"},
		},
	}))
	mustContain(t, got, `class="stats stats-horizontal shadow"`)
	mustContain(t, got, `<div class="stat">`)
	mustContain(t, got, `<div class="stat-title">Total Page Views</div>`)
	mustContain(t, got, `<div class="stat-value">89,400</div>`)
	mustContain(t, got, `<div class="stat-desc">21% more than last month</div>`)
}

func TestStatFigureActionsAndColors(t *testing.T) {
	got := render(t, Stat(StatConfig{
		Items: []StatItem{
			{
				Title:       "Total Likes",
				Value:       "25.6K",
				ValueClass:  "text-primary",
				Desc:        "21% more than last month",
				DescClass:   "text-secondary",
				Figure:      templ.Raw(`<svg class="w-8 h-8"></svg>`),
				FigureClass: "text-secondary",
				Actions:     templ.Raw(`<button class="btn btn-xs">Add</button>`),
			},
		},
	}))
	mustContain(t, got, `<div class="stat-figure text-secondary">`)
	mustContain(t, got, `<svg class="w-8 h-8"></svg>`)
	mustContain(t, got, `<div class="stat-value text-primary">25.6K</div>`)
	mustContain(t, got, `<div class="stat-desc text-secondary">21% more than last month</div>`)
	mustContain(t, got, `<div class="stat-actions">`)
	mustContain(t, got, `<button class="btn btn-xs">Add</button>`)
}

func TestStatVertical(t *testing.T) {
	got := render(t, Stat(StatConfig{Direction: "vertical"}))
	mustContain(t, got, `class="stats stats-vertical"`)
}
