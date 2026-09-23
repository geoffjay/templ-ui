package daisyui

import (
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestCard(t *testing.T) {
	got := render(t, Card(CardConfig{
		ImageSrc: "https://example.com/x.webp",
		ImageAlt: "Shoes",
		Title:    "Card Title",
		Desc:     "A card component has a figure, a body part.",
		Class:    "w-96 bg-base-100 shadow-sm",
		Actions:  templ.Raw(`<button class="btn btn-primary">Buy Now</button>`),
	}))
	mustContain(t, got, `<div class="card w-96 bg-base-100 shadow-sm">`)
	mustContain(t, got, `<figure>`)
	mustContain(t, got, `src="https://example.com/x.webp"`)
	mustContain(t, got, `alt="Shoes"`)
	mustContain(t, got, `<h2 class="card-title">Card Title</h2>`)
	mustContain(t, got, `class="justify-end card-actions"`)
	mustContain(t, got, `<button class="btn btn-primary">Buy Now</button>`)
}

func TestCardBottomImage(t *testing.T) {
	got := render(t, Card(CardConfig{
		ImageSrc:    "https://example.com/x.webp",
		ImageBottom: true,
		Title:       "Title",
	}))
	idxFig := strings.Index(got, "<figure>")
	idxBody := strings.Index(got, `class="card-body"`)
	if idxFig < 0 || idxBody < 0 {
		t.Fatalf("missing figure or body:\n%s", got)
	}
	if idxFig < idxBody {
		t.Errorf("expected figure after body; figure=%d body=%d", idxFig, idxBody)
	}
}

func TestCardStyleAndSize(t *testing.T) {
	got := render(t, Card(CardConfig{
		Style: "border",
		Size:  "lg",
		Class: "bg-base-100",
	}))
	mustContain(t, got, `class="card card-border card-lg bg-base-100"`)
}

func TestCardDashAndImageFull(t *testing.T) {
	got := render(t, Card(CardConfig{Style: "dash", Modifier: "image-full"}))
	mustContain(t, got, `class="card card-dash image-full"`)
}

func TestCardSide(t *testing.T) {
	got := render(t, Card(CardConfig{Modifier: "side"}))
	mustContain(t, got, `class="card card-side"`)
}

func TestCardBody(t *testing.T) {
	got := render(t, Card(CardConfig{
		Body: templ.Raw(`<span class="badge badge-warning">Popular</span>`),
	}))
	mustContain(t, got, `<span class="badge badge-warning">Popular</span>`)
}

func TestCardFigureClass(t *testing.T) {
	got := render(t, Card(CardConfig{
		ImageSrc:   "x",
		ImageClass: "px-10 pt-10",
	}))
	mustContain(t, got, `<figure class="px-10 pt-10">`)
}
