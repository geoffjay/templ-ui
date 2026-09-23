package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestValidatorHint(t *testing.T) {
	got := render(t, ValidatorHint(ValidatorHintConfig{Text: "Enter valid email address"}))
	mustContain(t, got, `<p class="validator-hint">`)
	mustContain(t, got, "Enter valid email address")
}

func TestValidatorHintHidden(t *testing.T) {
	got := render(t, ValidatorHint(ValidatorHintConfig{Text: "Required", Hidden: true}))
	mustContain(t, got, `class="validator-hint hidden"`)
}

func TestValidatorHintChildren(t *testing.T) {
	got := render(t, ValidatorHint(ValidatorHintConfig{
		Children: templ.Raw(`Must be 3 to 30 characters<br/>containing only letters`),
	}))
	mustContain(t, got, `<p class="validator-hint">`)
	mustContain(t, got, "Must be 3 to 30 characters<br/>")
}

func TestValidator(t *testing.T) {
	got := render(t, Validator(ValidatorConfig{
		Hint: "Must be valid URL",
	}, templ.Raw(`<input type="url" class="input validator" required/>`)))
	mustContain(t, got, `<input type="url" class="input validator" required/>`)
	mustContain(t, got, `<p class="validator-hint">`)
	mustContain(t, got, "Must be valid URL")
}

func TestValidatorHiddenHint(t *testing.T) {
	got := render(t, Validator(ValidatorConfig{
		Hint:       "Required",
		HintHidden: true,
	}, templ.Raw(`<input class="input validator"/>`)))
	mustContain(t, got, `class="validator-hint hidden"`)
}
