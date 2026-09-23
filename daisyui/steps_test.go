package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestSteps(t *testing.T) {
	got := render(t, Steps(StepsConfig{
		Items: []StepItem{
			{Label: "Register", Color: "primary"},
			{Label: "Choose plan", Color: "primary"},
			{Label: "Purchase"},
			{Label: "Receive Product"},
		},
	}))
	mustContain(t, got, `<ul class="steps">`)
	mustContain(t, got, `<li class="step step-primary"`)
	mustContain(t, got, `>Register</li>`)
	mustContain(t, got, `>Choose plan</li>`)
	mustContain(t, got, `<li class="step"`)
	mustContain(t, got, `>Purchase</li>`)
}

func TestStepsVertical(t *testing.T) {
	got := render(t, Steps(StepsConfig{
		Direction: "vertical",
		Items:     []StepItem{{Label: "A"}},
	}))
	mustContain(t, got, `<ul class="steps steps-vertical">`)
}

func TestStepsDataContent(t *testing.T) {
	got := render(t, Steps(StepsConfig{
		Items: []StepItem{{Label: "Step 1", Color: "neutral", DataContent: "?"}},
	}))
	mustContain(t, got, `<li class="step step-neutral" data-content="?">Step 1</li>`)
}

func TestStepsIcon(t *testing.T) {
	got := render(t, Steps(StepsConfig{
		Items: []StepItem{{Label: "Step 1", Icon: templ.Raw(`<span>😕</span>`)}},
	}))
	mustContain(t, got, `<span class="step-icon"><span>😕</span></span>`)
}
