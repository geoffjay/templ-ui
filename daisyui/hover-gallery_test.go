package daisyui

import (
	"strings"
	"testing"
)

func TestHoverGallery(t *testing.T) {
	got := render(t, HoverGallery(HoverGalleryConfig{
		Class: "max-w-60",
		Images: []HoverGalleryImage{
			{Src: "https://img.daisyui.com/images/stock/daisyui-hat-1.webp", Alt: "daisyUI hat 1"},
			{Src: "https://img.daisyui.com/images/stock/daisyui-hat-2.webp", Alt: "daisyUI hat 2"},
			{Src: "https://img.daisyui.com/images/stock/daisyui-hat-3.webp", Alt: "daisyUI hat 3"},
			{Src: "https://img.daisyui.com/images/stock/daisyui-hat-4.webp", Alt: "daisyUI hat 4"},
		},
	}))
	mustContain(t, got, `<figure class="hover-gallery max-w-60">`)
	mustContain(t, got, `<img src="https://img.daisyui.com/images/stock/daisyui-hat-1.webp" alt="daisyUI hat 1">`)
	if c := strings.Count(got, "<img"); c != 4 {
		t.Errorf("expected 4 <img> elements; got %d in:\n%s", c, got)
	}
}

func TestHoverGalleryDivElement(t *testing.T) {
	got := render(t, HoverGallery(HoverGalleryConfig{
		Element: "div",
		Images:  []HoverGalleryImage{{Src: "/a.png", Alt: "A"}},
	}))
	mustContain(t, got, `<div class="hover-gallery">`)
	if strings.Contains(got, "<figure") {
		t.Errorf("did not expect <figure>; got:\n%s", got)
	}
}

func TestHoverGalleryImageClass(t *testing.T) {
	got := render(t, HoverGallery(HoverGalleryConfig{
		Images: []HoverGalleryImage{{Src: "/a.png", Alt: "A", Class: "w-40 rounded"}},
	}))
	mustContain(t, got, `<img src="/a.png" alt="A" class="w-40 rounded">`)
}
