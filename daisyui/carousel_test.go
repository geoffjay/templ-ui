package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestCarouselSnapStart(t *testing.T) {
	got := render(t, Carousel(CarouselConfig{
		Class: "rounded-box",
		Slides: []CarouselSlide{
			{Src: "https://example.com/1.webp", Alt: "Burger"},
			{Src: "https://example.com/2.webp", Alt: "Burger"},
		},
	}))
	mustContain(t, got, `<div class="carousel rounded-box">`)
	mustContain(t, got, `<div class="carousel-item">`)
	mustContain(t, got, `src="https://example.com/1.webp"`)
	mustContain(t, got, `src="https://example.com/2.webp"`)
}

func TestCarouselSnapCenter(t *testing.T) {
	got := render(t, Carousel(CarouselConfig{
		Snap: "center",
	}))
	mustContain(t, got, `class="carousel carousel-center"`)
}

func TestCarouselSnapEnd(t *testing.T) {
	got := render(t, Carousel(CarouselConfig{
		Snap: "end",
	}))
	mustContain(t, got, `class="carousel carousel-end"`)
}

func TestCarouselVertical(t *testing.T) {
	got := render(t, Carousel(CarouselConfig{
		Direction: "vertical",
		Class:     "h-96 rounded-box",
		Slides:    []CarouselSlide{{Src: "https://example.com/1.webp", Class: "h-full"}},
	}))
	mustContain(t, got, `class="carousel carousel-vertical h-96 rounded-box"`)
	mustContain(t, got, `class="carousel-item h-full"`)
}

func TestCarouselFullWidth(t *testing.T) {
	got := render(t, Carousel(CarouselConfig{
		Class: "w-64",
		Slides: []CarouselSlide{
			{Src: "https://example.com/1.webp", Full: true, ImageClass: "w-full"},
		},
	}))
	mustContain(t, got, `<div class="carousel-item w-full">`)
	mustContain(t, got, `class="w-full"`)
}

func TestCarouselHalfWidth(t *testing.T) {
	got := render(t, Carousel(CarouselConfig{
		Class: "w-96",
		Slides: []CarouselSlide{
			{Src: "https://example.com/1.webp", Width: "w-1/2", ImageClass: "w-full"},
		},
	}))
	mustContain(t, got, `<div class="carousel-item w-1/2">`)
}

func TestCarouselContent(t *testing.T) {
	got := render(t, Carousel(CarouselConfig{
		Slides: []CarouselSlide{
			{Content: templ.Raw(`<img src="https://example.com/x.webp" class="rounded-box"/>`)},
		},
	}))
	mustContain(t, got, `<img src="https://example.com/x.webp" class="rounded-box"/>`)
}

func TestCarouselIndicators(t *testing.T) {
	got := render(t, Carousel(CarouselConfig{
		Indicators: true,
		Slides: []CarouselSlide{
			{ID: "item1", Src: "https://example.com/1.webp", Full: true, ImageClass: "w-full"},
			{ID: "item2", Src: "https://example.com/2.webp", Full: true, ImageClass: "w-full"},
		},
	}))
	mustContain(t, got, `id="item1"`)
	mustContain(t, got, `id="item2"`)
	mustContain(t, got, `<div class="flex w-full justify-center gap-2 py-2">`)
	mustContain(t, got, `<a href="#item1" class="btn btn-xs">item1</a>`)
	mustContain(t, got, `<a href="#item2" class="btn btn-xs">item2</a>`)
}

func TestCarouselControls(t *testing.T) {
	got := render(t, Carousel(CarouselConfig{
		Controls: true,
		Slides: []CarouselSlide{
			{ID: "slide1", Src: "https://example.com/1.webp", Full: true, ImageClass: "w-full"},
			{ID: "slide2", Src: "https://example.com/2.webp", Full: true, ImageClass: "w-full"},
			{ID: "slide3", Src: "https://example.com/3.webp", Full: true, ImageClass: "w-full"},
		},
	}))
	mustContain(t, got, `<div id="slide1" class="carousel-item w-full relative">`)
	mustContain(t, got, `<a href="#slide3" class="btn btn-circle">❮</a>`)
	mustContain(t, got, `<a href="#slide2" class="btn btn-circle">❯</a>`)
	mustContain(t, got, `<a href="#slide1" class="btn btn-circle">❮</a>`)
}

func TestCarouselControlHrefWrap(t *testing.T) {
	got := carouselControlHref(3, []CarouselSlide{
		{ID: "s1"}, {ID: "s2"}, {ID: "s3"},
	})
	if got != "#s1" {
		t.Errorf("expected wrap to #s1, got %q", got)
	}
	got = carouselControlHref(-1, []CarouselSlide{
		{ID: "s1"}, {ID: "s2"}, {ID: "s3"},
	})
	if got != "#s3" {
		t.Errorf("expected wrap to #s3, got %q", got)
	}
}
