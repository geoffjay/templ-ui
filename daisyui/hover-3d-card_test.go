package daisyui

import (
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestHover3DCard(t *testing.T) {
	got := render(t, Hover3DCard(Hover3DConfig{
		Class:    "my-12 mx-2",
		Children: templ.Raw(`<figure class="max-w-100 rounded-2xl"><img src="x"/></figure>`),
	}))
	mustContain(t, got, `<div class="hover-3d my-12 mx-2">`)
	mustContain(t, got, `<figure class="max-w-100 rounded-2xl">`)
	// Must render exactly 8 empty <div>s for the hover zones.
	if c := strings.Count(got, "<div></div>"); c != 8 {
		t.Errorf("expected 8 empty <div>s for hover zones; got %d in:\n%s", c, got)
	}
	if strings.Contains(got, "<a ") {
		t.Errorf("did not expect an anchor; got:\n%s", got)
	}
}

func TestHover3DCardAsLink(t *testing.T) {
	got := render(t, Hover3DCard(Hover3DConfig{
		Href:  "/cards",
		Class: "cursor-pointer",
	}))
	mustContain(t, got, `<a href="/cards" class="hover-3d cursor-pointer">`)
	mustContain(t, got, "<div></div>")
}
