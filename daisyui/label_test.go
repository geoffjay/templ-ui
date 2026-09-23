package daisyui

import "testing"

func TestLabelSpan(t *testing.T) {
	got := render(t, Label(LabelConfig{Text: "Optional"}))
	mustContain(t, got, `<span class="label">`)
	mustContain(t, got, ">Optional<")
}

func TestLabelFor(t *testing.T) {
	got := render(t, Label(LabelConfig{Text: "Email", ForID: "email"}))
	mustContain(t, got, `<label for="email" class="label">`)
	mustContain(t, got, ">Email<")
}

func TestLabelClass(t *testing.T) {
	got := render(t, Label(LabelConfig{Text: "X", Class: "text-sm"}))
	mustContain(t, got, `class="label text-sm"`)
}

func TestFloatingLabel(t *testing.T) {
	got := render(t, FloatingLabel(FloatingLabelConfig{
		Label:       "Your Email",
		InputType:   "email",
		Placeholder: "mail@site.com",
		InputClass:  "input input-md",
	}))
	mustContain(t, got, `<label class="floating-label">`)
	mustContain(t, got, `<span>Your Email</span>`)
	mustContain(t, got, `<input type="email"`)
	mustContain(t, got, `placeholder="mail@site.com"`)
	mustContain(t, got, `class="input input-md"`)
}
