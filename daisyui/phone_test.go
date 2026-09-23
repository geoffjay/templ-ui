package daisyui

import (
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestPhone(t *testing.T) {
	got := render(t, Phone(PhoneConfig{
		Camera:       true,
		DisplayClass: "text-white bg-neutral-900 grid place-content-center",
		Children:     templ.Raw("It's Glowtime."),
	}))
	mustContain(t, got, `<div class="mockup-phone">`)
	mustContain(t, got, `<div class="mockup-phone-camera"></div>`)
	mustContain(t, got, `<div class="mockup-phone-display text-white bg-neutral-900 grid place-content-center">`)
	mustContain(t, got, "It's Glowtime.")
}

func TestPhoneWallpaper(t *testing.T) {
	got := render(t, Phone(PhoneConfig{
		Camera:       true,
		Class:        `border-[#ff8938]`,
		WallpaperSrc: "https://img.daisyui.com/images/stock/453966.webp",
	}))
	mustContain(t, got, `<div class="mockup-phone border-[#ff8938]">`)
	mustContain(t, got, `mockup-phone-display">`)
	mustContain(t, got, `<img alt="wallpaper" src="https://img.daisyui.com/images/stock/453966.webp">`)
}

func TestPhoneNoCamera(t *testing.T) {
	got := render(t, Phone(PhoneConfig{Camera: false}))
	if strings.Contains(got, "mockup-phone-camera") {
		t.Errorf("did not expect camera; got:\n%s", got)
	}
}
