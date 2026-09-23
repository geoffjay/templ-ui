package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestHeroCentered(t *testing.T) {
	got := render(t, Hero(HeroConfig{
		Title:        "Hello there",
		Subtitle:     "Provident cupiditate voluptatem et in.",
		Class:        "min-h-screen bg-base-200",
		ContentClass: "text-center",
		Actions:      templ.Raw(`<button class="btn btn-primary">Get Started</button>`),
	}))
	mustContain(t, got, `<div class="hero min-h-screen bg-base-200">`)
	mustContain(t, got, `<div class="hero-content text-center">`)
	mustContain(t, got, `<h1 class="text-5xl font-bold">Hello there</h1>`)
	mustContain(t, got, `<p class="py-6">Provident cupiditate voluptatem et in.</p>`)
	mustContain(t, got, `<button class="btn btn-primary">Get Started</button>`)
}

func TestHeroWithFigure(t *testing.T) {
	got := render(t, Hero(HeroConfig{
		Class:        "min-h-screen bg-base-200",
		ContentClass: "flex-col lg:flex-row",
		Content: templ.Raw(`
			<img src="https://example.com/photo.webp" class="max-w-sm rounded-lg shadow-2xl" alt="hero"/>
			<div>
				<h1 class="text-5xl font-bold">Box Office News!</h1>
				<button class="btn btn-primary">Get Started</button>
			</div>
		`),
	}))
	mustContain(t, got, `<div class="hero min-h-screen bg-base-200">`)
	mustContain(t, got, `<div class="hero-content flex-col lg:flex-row">`)
	mustContain(t, got, `src="https://example.com/photo.webp"`)
	mustContain(t, got, `Box Office News!`)
}

func TestHeroOverlay(t *testing.T) {
	got := render(t, Hero(HeroConfig{
		Class:           "min-h-screen",
		Overlay:         true,
		BackgroundImage: "https://example.com/bg.webp",
		ContentClass:    "text-center text-neutral-content",
		Title:           "Hello there",
	}))
	mustContain(t, got, `style="background-image: url(https://example.com/bg.webp);"`)
	mustContain(t, got, `<div class="hero-overlay"></div>`)
	mustContain(t, got, `class="hero-content text-center text-neutral-content"`)
}

func TestHeroEmpty(t *testing.T) {
	got := render(t, Hero(HeroConfig{}))
	mustContain(t, got, `<div class="hero">`)
	mustContain(t, got, `<div class="hero-content">`)
}
