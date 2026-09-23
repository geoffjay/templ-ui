package daisyui

import "testing"

func TestFileInput(t *testing.T) {
	got := render(t, FileInput(FileInputConfig{
		Name:     "upload",
		Accept:   "image/*",
		Multiple: true,
		Color:    "primary",
		Size:     "lg",
	}))
	mustContain(t, got, `<input type="file"`)
	mustContain(t, got, `name="upload"`)
	mustContain(t, got, `accept="image/*"`)
	mustContain(t, got, `class="file-input file-input-primary file-input-lg"`)
	mustContain(t, got, `multiple`)
}

func TestFileInputDefaults(t *testing.T) {
	got := render(t, FileInput(FileInputConfig{}))
	mustContain(t, got, `<input type="file"`)
	mustContain(t, got, `class="file-input"`)
}

func TestFileInputGhost(t *testing.T) {
	got := render(t, FileInput(FileInputConfig{Style: "ghost"}))
	mustContain(t, got, `class="file-input file-input-ghost"`)
}

func TestFileInputDisabled(t *testing.T) {
	got := render(t, FileInput(FileInputConfig{Disabled: true}))
	mustContain(t, got, `disabled`)
}
