package daisyui

import "testing"

func TestCalendarCally(t *testing.T) {
	got := render(t, Calendar(CalendarConfig{
		Class: "bg-base-100 border border-base-300 shadow-lg rounded-box",
	}))
	mustContain(t, got, `<calendar-date class="cally bg-base-100 border border-base-300 shadow-lg rounded-box">`)
	mustContain(t, got, `aria-label="Previous"`)
	mustContain(t, got, `slot="previous"`)
	mustContain(t, got, `aria-label="Next"`)
	mustContain(t, got, `slot="next"`)
	mustContain(t, got, `<calendar-month></calendar-month>`)
}

func TestCalendarCallyValue(t *testing.T) {
	got := render(t, Calendar(CalendarConfig{
		Value: "2024-01-15",
	}))
	mustContain(t, got, `value="2024-01-15"`)
}

func TestCalendarPikaday(t *testing.T) {
	got := render(t, Calendar(CalendarConfig{
		Variant: "pikaday",
		Class:   "w-full",
	}))
	mustContain(t, got, `<input type="text" class="input pika-single w-full"`)
}

func TestCalendarPikadayValue(t *testing.T) {
	got := render(t, Calendar(CalendarConfig{
		Variant: "pikaday",
		Value:   "Pick a day",
	}))
	mustContain(t, got, `value="Pick a day"`)
}

func TestCalendarVariantDefault(t *testing.T) {
	if calendarVariant("") != "cally" {
		t.Errorf("expected default variant to be cally")
	}
	if calendarVariant("unknown") != "cally" {
		t.Errorf("expected unknown variant to fall back to cally")
	}
}
