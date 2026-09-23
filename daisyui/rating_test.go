package daisyui

import "testing"

func TestRating(t *testing.T) {
	got := render(t, Rating(RatingConfig{
		Name: "rating-1",
		Size: "lg",
		Items: []RatingItem{
			{Mask: "star-2", AriaLabel: "1 star", ColorClass: "bg-orange-400"},
			{Mask: "star-2", AriaLabel: "2 star", ColorClass: "bg-orange-400", Checked: true},
			{Mask: "star-2", AriaLabel: "3 star", ColorClass: "bg-orange-400"},
		},
	}))
	mustContain(t, got, `<div class="rating rating-lg">`)
	mustContain(t, got, `<input type="radio" name="rating-1"`)
	mustContain(t, got, `aria-label="2 star"`)
	mustContain(t, got, `class="mask mask-star-2 bg-orange-400"`)
	mustContain(t, got, `checked`)
}

func TestRatingHalf(t *testing.T) {
	got := render(t, Rating(RatingConfig{
		Name: "rating-2",
		Half: true,
		Size: "lg",
		Items: []RatingItem{
			{Hidden: true, AriaLabel: "clear"},
			{Mask: "star-2", Modifier: "half-1", AriaLabel: "0.5 star", ColorClass: "bg-green-500"},
			{Mask: "star-2", Modifier: "half-2", AriaLabel: "1 star", ColorClass: "bg-green-500"},
		},
	}))
	mustContain(t, got, `class="rating rating-lg rating-half"`)
	mustContain(t, got, `rating-hidden`)
	mustContain(t, got, `class="mask mask-star-2 mask-half-1 bg-green-500"`)
}

func TestRatingDefaults(t *testing.T) {
	got := render(t, Rating(RatingConfig{
		Name: "rating-3",
		Items: []RatingItem{
			{Mask: "star", AriaLabel: "1 star"},
		},
	}))
	mustContain(t, got, `<div class="rating">`)
	mustContain(t, got, `class="mask mask-star"`)
}
