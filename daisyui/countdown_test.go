package daisyui

import "testing"

func TestCountdown(t *testing.T) {
	got := render(t, Countdown(CountdownConfig{
		Class: "font-mono text-6xl",
		Items: []CountdownItem{
			{Value: "59", Digits: "2"},
		},
	}))
	mustContain(t, got, `<span class="countdown font-mono text-6xl">`)
	mustContain(t, got, `style="--value:59;--digits:2;"`)
	mustContain(t, got, `aria-live="polite"`)
	mustContain(t, got, `aria-label="59"`)
	mustContain(t, got, ">59<")
}

func TestCountdownClock(t *testing.T) {
	got := render(t, Countdown(CountdownConfig{
		Class: "font-mono text-2xl",
		Items: []CountdownItem{
			{Value: "10", Suffix: "h "},
			{Value: "24", Suffix: "m "},
			{Value: "59", Suffix: "s"},
		},
	}))
	mustContain(t, got, `>10</span> h `)
	mustContain(t, got, `>24</span> m `)
	mustContain(t, got, `>59</span> s`)
}
