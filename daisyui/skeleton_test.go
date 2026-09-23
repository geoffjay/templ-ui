package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestSkeleton(t *testing.T) {
	got := render(t, Skeleton(SkeletonConfig{Class: "w-32 h-32"}))
	mustContain(t, got, `<div class="skeleton w-32 h-32"></div>`)
}

func TestSkeletonText(t *testing.T) {
	got := render(t, Skeleton(SkeletonConfig{
		Class: "text-2xl",
		Text:  "AI is thinking harder...",
	}))
	mustContain(t, got, `<span class="skeleton skeleton-text text-2xl">`)
	mustContain(t, got, ">AI is thinking harder...<")
}

func TestSkeletonChildren(t *testing.T) {
	got := render(t, Skeleton(SkeletonConfig{
		Class:    "h-4 w-28",
		Children: templ.Raw(`<span>inner</span>`),
	}))
	mustContain(t, got, `<div class="skeleton h-4 w-28">`)
	mustContain(t, got, `<span>inner</span>`)
}
