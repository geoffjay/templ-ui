package daisyui

import "testing"

func TestRadialProgress(t *testing.T) {
	got := render(t, RadialProgress(RadialProgressConfig{
		Value: "70",
		Label: "70%",
	}))
	mustContain(t, got, `class="radial-progress"`)
	mustContain(t, got, `style="--value:70;"`)
	mustContain(t, got, `aria-valuenow="70"`)
	mustContain(t, got, `role="progressbar"`)
	mustContain(t, got, ">70%<")
}

func TestRadialProgressColor(t *testing.T) {
	got := render(t, RadialProgress(RadialProgressConfig{
		Value:      "70",
		ColorClass: "text-primary",
	}))
	mustContain(t, got, `class="radial-progress text-primary"`)
}

func TestRadialProgressSizeThickness(t *testing.T) {
	got := render(t, RadialProgress(RadialProgressConfig{
		Value:     "70",
		Size:      "12rem",
		Thickness: "2px",
	}))
	mustContain(t, got, `style="--value:70;--size:12rem;--thickness:2px;"`)
}
