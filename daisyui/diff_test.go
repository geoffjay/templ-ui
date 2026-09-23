package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestDiff(t *testing.T) {
	got := render(t, Diff(DiffConfig{
		Item1: DiffItem{Children: templ.Raw(`<img alt="daisy" src="a.webp"/>`), Tabindex: true},
		Item2: DiffItem{Children: templ.Raw(`<img alt="daisy" src="a-blur.webp"/>`)},
	}))
	mustContain(t, got, `<figure class="diff rounded-field aspect-16/9" tabindex="0">`)
	mustContain(t, got, `<div class="diff-item-1" role="img" tabindex>`)
	mustContain(t, got, `<div class="diff-item-2" role="img">`)
	mustContain(t, got, `src="a.webp"`)
	mustContain(t, got, `src="a-blur.webp"`)
	mustContain(t, got, `<div class="diff-resizer"></div>`)
}

func TestDiffText(t *testing.T) {
	got := render(t, Diff(DiffConfig{
		Item1: DiffItem{Children: templ.Raw(`<div class="bg-primary text-primary-content grid place-content-center text-9xl font-black">DAISY</div>`)},
		Item2: DiffItem{Children: templ.Raw(`<div class="bg-base-200 grid place-content-center text-9xl font-black">DAISY</div>`)},
	}))
	mustContain(t, got, `bg-primary text-primary-content`)
	mustContain(t, got, `bg-base-200`)
	mustContain(t, got, `>DAISY<`)
}

func TestDiffCustomClass(t *testing.T) {
	got := render(t, Diff(DiffConfig{Class: "aspect-4/3"}))
	mustContain(t, got, `<figure class="diff aspect-4/3" tabindex="0">`)
}
