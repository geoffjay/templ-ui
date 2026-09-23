package daisyui

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func render(t *testing.T, c templ.Component) string {
	t.Helper()
	var buf bytes.Buffer
	if err := c.Render(context.Background(), &buf); err != nil {
		t.Fatalf("render: %v", err)
	}
	return buf.String()
}

func mustContain(t *testing.T, got, want string) {
	t.Helper()
	if !strings.Contains(got, want) {
		t.Errorf("expected output to contain %q, got:\n%s", want, got)
	}
}

func TestJoin(t *testing.T) {
	item := templ.Raw("<button class=\"btn join-item\">A</button>")
	got := render(t, Join([]templ.Component{item}, "vertical", "w-full"))
	mustContain(t, got, `class="join join-vertical w-full"`)
	mustContain(t, got, "btn join-item")
}

func TestStack(t *testing.T) {
	a := templ.Raw("<div>A</div>")
	b := templ.Raw("<div>B</div>")
	got := render(t, Stack([]templ.Component{a, b}, "top", "h-20 w-32", false))
	mustContain(t, got, `class="stack stack-top h-20 w-32"`)
	mustContain(t, got, "<div>A</div>")
}

func TestStackGlass(t *testing.T) {
	a := templ.Raw("<div>A</div>")
	got := render(t, Stack([]templ.Component{a}, "", "h-20", true))
	mustContain(t, got, `class="stack glass h-20"`)
}

func TestAccordionGlass(t *testing.T) {
	items := []AccordionItem{{Title: "T", Content: "C"}}
	got := render(t, Accordion("a", items, "", false, true))
	mustContain(t, got, `glass`)
}

func TestDivider(t *testing.T) {
	got := render(t, Divider("OR", "horizontal", "primary", "start", ""))
	mustContain(t, got, `class="divider divider-horizontal divider-primary divider-start"`)
	mustContain(t, got, ">OR<")
}

func TestDividerDefaults(t *testing.T) {
	got := render(t, Divider("", "", "", "", ""))
	mustContain(t, got, `class="divider"`)
}

func TestMask(t *testing.T) {
	got := render(t, Mask("heart", "", "w-40", templ.Raw("inner")))
	mustContain(t, got, `class="mask mask-heart w-40"`)
	mustContain(t, got, "inner")
}

func TestMaskImage(t *testing.T) {
	got := render(t, MaskImage("/x.png", "alt", "circle", "half-1", "h-20"))
	mustContain(t, got, `src="/x.png"`)
	mustContain(t, got, `alt="alt"`)
	mustContain(t, got, `class="mask mask-circle mask-half-1 h-20"`)
}

func TestLink(t *testing.T) {
	got := render(t, Link("/about", "About", "primary", "hover", "", nil))
	mustContain(t, got, `href="/about"`)
	mustContain(t, got, `class="link link-primary link-hover"`)
	mustContain(t, got, ">About<")
}

func TestLinkEmptyHref(t *testing.T) {
	got := render(t, Link("", "Home", "", "", "", nil))
	mustContain(t, got, `href="#"`)
	mustContain(t, got, `class="link"`)
}

func TestKbd(t *testing.T) {
	got := render(t, Kbd("F", "sm", ""))
	mustContain(t, got, `<kbd class="kbd kbd-sm">F</kbd>`)
}

func TestKbdCombo(t *testing.T) {
	got := render(t, KbdCombo([]string{"ctrl", "shift", "del"}, "", ""))
	mustContain(t, got, "<kbd class=\"kbd\">ctrl</kbd>")
	mustContain(t, got, " + ")
	mustContain(t, got, "<kbd class=\"kbd\">del</kbd>")
}

func TestStatus(t *testing.T) {
	got := render(t, Status("error", "xl", "animate-ping", "server down", ""))
	mustContain(t, got, `aria-label="server down"`)
	mustContain(t, got, `class="status status-error status-xl animate-ping"`)
}

func TestLoading(t *testing.T) {
	got := render(t, Loading("spinner", "lg", "text-primary", ""))
	mustContain(t, got, `class="loading loading-spinner loading-lg text-primary"`)
}

func TestCollapseFocus(t *testing.T) {
	got := render(t, Collapse(CollapseConfig{Icon: "arrow", Class: "bg-base-100"}, "Q?", "A."))
	mustContain(t, got, `tabindex="0"`)
	mustContain(t, got, `class="collapse collapse-arrow bg-base-100"`)
	mustContain(t, got, `class="collapse-title font-semibold">Q?`)
	mustContain(t, got, `class="collapse-content text-sm">A.`)
}

func TestCollapseCheckbox(t *testing.T) {
	got := render(t, Collapse(CollapseConfig{Mode: "checkbox", Open: true}, "Q?", "A."))
	mustContain(t, got, `class="collapse collapse-open"`)
	mustContain(t, got, `<input type="checkbox"`)
}

func TestCollapseDetails(t *testing.T) {
	got := render(t, Collapse(CollapseConfig{Mode: "details", Open: true}, "Q?", "A."))
	mustContain(t, got, `<details open`)
	mustContain(t, got, "<summary")
}

func TestSwap(t *testing.T) {
	got := render(t, Swap(SwapConfig{Effect: "rotate"}, templ.Raw("ON"), templ.Raw("OFF")))
	mustContain(t, got, `<label class="swap swap-rotate">`)
	mustContain(t, got, `<input type="checkbox"`)
	mustContain(t, got, `<div class="swap-on">ON</div>`)
	mustContain(t, got, `<div class="swap-off">OFF</div>`)
}

func TestSwapActiveNoCheckbox(t *testing.T) {
	got := render(t, Swap(SwapConfig{Active: true, NoCheckbox: true}, templ.Raw("A"), templ.Raw("B")))
	mustContain(t, got, `class="swap swap-active"`)
	if strings.Contains(got, "checkbox") {
		t.Errorf("expected no checkbox; got:\n%s", got)
	}
}

func TestThemeControllerCheckbox(t *testing.T) {
	got := render(t, ThemeController(ThemeControllerConfig{
		InputClass: "toggle",
		Value:      "synthwave",
	}))
	mustContain(t, got, `<input type="checkbox"`)
	mustContain(t, got, `class="toggle theme-controller"`)
	mustContain(t, got, `value="synthwave"`)
}

func TestThemeControllerRadio(t *testing.T) {
	got := render(t, ThemeController(ThemeControllerConfig{
		AsRadio:    true,
		Name:       "theme-radios",
		Value:      "retro",
		InputClass: "radio radio-sm",
		Checked:    true,
	}))
	mustContain(t, got, `<input type="radio" name="theme-radios"`)
	mustContain(t, got, `value="retro"`)
	mustContain(t, got, `class="radio radio-sm theme-controller"`)
	mustContain(t, got, `checked`)
}

func TestButton(t *testing.T) {
	got := render(t, Button(ButtonConfig{Label: "Default"}))
	mustContain(t, got, `<button type="button" class="btn">`)
	mustContain(t, got, ">Default<")
}

func TestButtonColorSizeStyle(t *testing.T) {
	got := render(t, Button(ButtonConfig{
		Label:    "Primary",
		Color:    "primary",
		Style:    "outline",
		Size:     "lg",
		Modifier: "circle",
	}))
	mustContain(t, got, `class="btn btn-primary btn-outline btn-lg btn-circle"`)
	mustContain(t, got, ">Primary<")
}

func TestButtonLink(t *testing.T) {
	got := render(t, Button(ButtonConfig{
		Label: "Home",
		Href:  "/",
	}))
	mustContain(t, got, `<a role="button"`)
	mustContain(t, got, `href="/"`)
	mustContain(t, got, `class="btn"`)
}

func TestButtonDisabled(t *testing.T) {
	got := render(t, Button(ButtonConfig{Label: "X", Disabled: true}))
	mustContain(t, got, `disabled`)
}

func TestButtonLinkDisabled(t *testing.T) {
	got := render(t, Button(ButtonConfig{Label: "X", Href: "/x", Disabled: true}))
	mustContain(t, got, `class="btn btn-disabled"`)
	mustContain(t, got, `tabindex="-1"`)
	mustContain(t, got, `aria-disabled="true"`)
}

func TestButtonChildren(t *testing.T) {
	got := render(t, Button(ButtonConfig{
		Label:    "Like",
		Children: templ.Raw(`<svg class="size-4"></svg>`),
	}))
	mustContain(t, got, `<svg class="size-4"></svg>`)
	mustContain(t, got, ">Like<")
}

func TestFABSingle(t *testing.T) {
	got := render(t, FAB(FABConfig{
		Trigger: FABTrigger{Label: "F", Color: "primary"},
		Class:   "absolute z-1",
	}))
	mustContain(t, got, `class="fab absolute z-1"`)
	mustContain(t, got, `<button class="btn btn-primary btn-lg btn-circle">`)
	mustContain(t, got, ">F<")
}

func TestFABSpeedDial(t *testing.T) {
	got := render(t, FAB(FABConfig{
		Trigger: FABTrigger{Label: "F", Color: "primary"},
		Items: []FABItem{
			{Button: templ.Raw(`<button class="btn btn-lg btn-circle">A</button>`)},
			{Button: templ.Raw(`<button class="btn btn-lg btn-circle">B</button>`)},
		},
	}))
	mustContain(t, got, `class="fab"`)
	mustContain(t, got, `tabindex="0" role="button"`)
	mustContain(t, got, `class="btn btn-primary btn-lg btn-circle"`)
	mustContain(t, got, `<button class="btn btn-lg btn-circle">A</button>`)
}

func TestFABFlower(t *testing.T) {
	got := render(t, FAB(FABConfig{
		Trigger: FABTrigger{Label: "F", Color: "success"},
		Flower:  true,
		Items: []FABItem{
			{Button: templ.Raw(`<button class="btn btn-lg btn-circle">A</button>`)},
		},
	}))
	mustContain(t, got, `class="fab fab-flower"`)
	mustContain(t, got, `btn-success`)
}

func TestFABCloseButton(t *testing.T) {
	got := render(t, FAB(FABConfig{
		Trigger:     FABTrigger{Label: "F", Color: "info"},
		CloseButton: templ.Raw(`<span class="btn btn-circle btn-lg btn-error">✕</span>`),
	}))
	mustContain(t, got, `<div class="fab-close">`)
	mustContain(t, got, `✕`)
}
