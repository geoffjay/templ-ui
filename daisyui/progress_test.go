package daisyui

import (
	"strings"
	"testing"
)

func TestProgress(t *testing.T) {
	got := render(t, Progress(ProgressConfig{Value: "40", Max: "100", Class: "w-56"}))
	mustContain(t, got, `<progress`)
	mustContain(t, got, `class="progress w-56"`)
	mustContain(t, got, `value="40"`)
	mustContain(t, got, `max="100"`)
}

func TestProgressColor(t *testing.T) {
	got := render(t, Progress(ProgressConfig{Value: "70", Color: "primary"}))
	mustContain(t, got, `class="progress progress-primary"`)
}

func TestProgressDefaultMax(t *testing.T) {
	got := render(t, Progress(ProgressConfig{Value: "10"}))
	mustContain(t, got, `max="100"`)
}

func TestProgressIndeterminate(t *testing.T) {
	got := render(t, Progress(ProgressConfig{Class: "w-56"}))
	mustContain(t, got, `<progress`)
	mustContain(t, got, `class="progress w-56"`)
	if strings.Contains(got, `value="`) {
		t.Errorf("expected no value attribute for indeterminate; got:\n%s", got)
	}
}
