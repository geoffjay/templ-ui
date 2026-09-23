package daisyui

import (
	"strings"
	"testing"
)

func TestTextRotate(t *testing.T) {
	got := render(t, TextRotate(TextRotateConfig{
		Lines: []TextRotateLine{
			{Text: "ONE"},
			{Text: "TWO"},
			{Text: "THREE"},
		},
	}))
	mustContain(t, got, `<span class="text-rotate">`)
	mustContain(t, got, `<span class="">ONE</span>`)
	mustContain(t, got, `<span class="">TWO</span>`)
	mustContain(t, got, `<span class="">THREE</span>`)
	if strings.Contains(got, "duration-") {
		t.Errorf("expected no duration class by default; got:\n%s", got)
	}
}

func TestTextRotateDuration(t *testing.T) {
	got := render(t, TextRotate(TextRotateConfig{
		Class:    "text-7xl font-title",
		Duration: 6000,
		Lines: []TextRotateLine{
			{Text: "BLAZING"},
			{Text: "FAST", Class: "font-bold italic px-2"},
		},
		InnerClass: "justify-items-center",
	}))
	mustContain(t, got, `<span class="text-rotate duration-6000 text-7xl font-title">`)
	mustContain(t, got, `<span class="justify-items-center">`)
	mustContain(t, got, `<span class="">BLAZING</span>`)
	mustContain(t, got, `<span class="font-bold italic px-2">FAST</span>`)
}

func TestTextRotateColoredWords(t *testing.T) {
	got := render(t, TextRotate(TextRotateConfig{
		Lines: []TextRotateLine{
			{Text: "Designers", Class: "bg-teal-400 text-teal-800 px-2"},
			{Text: "Developers", Class: "bg-red-400 text-red-800 px-2"},
		},
	}))
	mustContain(t, got, `<span class="bg-teal-400 text-teal-800 px-2">Designers</span>`)
	mustContain(t, got, `<span class="bg-red-400 text-red-800 px-2">Developers</span>`)
}
