package daisyui

import "testing"

func TestAvatar(t *testing.T) {
	got := render(t, Avatar(AvatarConfig{
		Src:   "https://example.com/x.webp",
		Alt:   "avatar",
		Size:  "w-24",
		Class: "test",
	}))
	mustContain(t, got, `<div class="avatar test">`)
	mustContain(t, got, `<div class="rounded w-24">`)
	mustContain(t, got, `src="https://example.com/x.webp"`)
	mustContain(t, got, `alt="avatar"`)
}

func TestAvatarShapes(t *testing.T) {
	cases := []struct {
		shape, expect string
	}{
		{"rounded-xl", `class="rounded-xl w-24"`},
		{"rounded-full", `class="rounded-full w-24"`},
		{"squircle", `class="mask mask-squircle w-24"`},
		{"heart", `class="mask mask-heart w-24"`},
		{"hexagon-2", `class="mask mask-hexagon-2 w-24"`},
	}
	for _, c := range cases {
		got := render(t, Avatar(AvatarConfig{Shape: c.shape, Size: "w-24"}))
		mustContain(t, got, c.expect)
	}
}

func TestAvatarStatus(t *testing.T) {
	for _, s := range []string{"online", "offline"} {
		got := render(t, Avatar(AvatarConfig{Status: s, Size: "w-24"}))
		mustContain(t, got, `<div class="avatar avatar-`+s+`">`)
	}
}

func TestAvatarPlaceholder(t *testing.T) {
	got := render(t, Avatar(AvatarConfig{
		Placeholder:      true,
		PlaceholderText:  "D",
		PlaceholderClass: "text-3xl",
		Size:             "w-24",
		InnerClass:       "bg-neutral text-neutral-content rounded-full",
	}))
	mustContain(t, got, `<div class="avatar avatar-placeholder">`)
	mustContain(t, got, `<div class="rounded w-24 bg-neutral text-neutral-content rounded-full">`)
	mustContain(t, got, `<span class="text-3xl">D</span>`)
}

func TestAvatarRing(t *testing.T) {
	got := render(t, Avatar(AvatarConfig{
		Ring:  "primary",
		Size:  "w-24",
		Shape: "rounded-full",
	}))
	mustContain(t, got, `class="rounded-full w-24 ring-2 ring-primary ring-offset-2 ring-offset-base-100"`)
}

func TestAvatarGroup(t *testing.T) {
	got := render(t, AvatarGroup(AvatarGroupConfig{
		Class: "-space-x-6",
		Items: []AvatarConfig{
			{Src: "a.webp", Size: "w-12"},
			{Placeholder: true, PlaceholderText: "+99", Size: "w-12", InnerClass: "bg-neutral text-neutral-content"},
		},
	}))
	mustContain(t, got, `<div class="avatar-group -space-x-6">`)
	mustContain(t, got, `<div class="avatar avatar-placeholder">`)
	mustContain(t, got, `src="a.webp"`)
}
