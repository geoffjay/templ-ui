package daisyui

import (
	"strings"
	"testing"
)

func TestRange(t *testing.T) {
	got := render(t, Range(RangeConfig{
		Name:  "vol",
		Min:   "0",
		Max:   "100",
		Value: "40",
		Step:  "10",
		Color: "primary",
		Size:  "lg",
	}))
	mustContain(t, got, `<input type="range"`)
	mustContain(t, got, `name="vol"`)
	mustContain(t, got, `min="0"`)
	mustContain(t, got, `max="100"`)
	mustContain(t, got, `value="40"`)
	mustContain(t, got, `step="10"`)
	mustContain(t, got, `class="range range-primary range-lg"`)
}

func TestRangeDefaults(t *testing.T) {
	got := render(t, Range(RangeConfig{}))
	mustContain(t, got, `<input type="range"`)
	mustContain(t, got, `class="range"`)
	if strings.Contains(got, `step=`) {
		t.Errorf("expected no step attr; got:\n%s", got)
	}
}

func TestRangeDisabled(t *testing.T) {
	got := render(t, Range(RangeConfig{Disabled: true}))
	mustContain(t, got, `disabled`)
}
