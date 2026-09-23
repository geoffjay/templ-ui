package daisyui

import (
	"strings"
	"testing"
)

func TestTextarea(t *testing.T) {
	got := render(t, Textarea(TextareaConfig{
		Name:        "bio",
		Placeholder: "Bio",
		Color:       "primary",
		Size:        "lg",
		Rows:        4,
		Class:       "h-24",
	}))
	mustContain(t, got, `<textarea name="bio"`)
	mustContain(t, got, `placeholder="Bio"`)
	mustContain(t, got, `class="textarea textarea-primary textarea-lg h-24"`)
	mustContain(t, got, `rows="4"`)
}

func TestTextareaDefaults(t *testing.T) {
	got := render(t, Textarea(TextareaConfig{}))
	mustContain(t, got, `<textarea name="" placeholder="" class="textarea"`)
	if strings.Contains(got, `rows=`) {
		t.Errorf("expected no rows attr; got:\n%s", got)
	}
}

func TestTextareaGhost(t *testing.T) {
	got := render(t, Textarea(TextareaConfig{Style: "ghost"}))
	mustContain(t, got, `class="textarea textarea-ghost"`)
}

func TestTextareaDisabled(t *testing.T) {
	got := render(t, Textarea(TextareaConfig{Disabled: true}))
	mustContain(t, got, `disabled`)
}
