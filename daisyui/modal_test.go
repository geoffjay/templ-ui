package daisyui

import (
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestModalDialog(t *testing.T) {
	got := render(t, Modal(ModalConfig{
		ID:           "my_modal_1",
		Title:        "Hello!",
		Message:      "Press ESC key or click the button below to close",
		TriggerLabel: "open modal",
		Actions:      templ.Raw(`<form method="dialog"><button class="btn">Close</button></form>`),
	}))
	mustContain(t, got, `<button class="btn"`)
	mustContain(t, got, `showModal`)
	mustContain(t, got, `<dialog id="my_modal_1" class="modal">`)
	mustContain(t, got, `<div class="modal-box">`)
	mustContain(t, got, `<h3 class="text-lg font-bold">Hello!</h3>`)
	mustContain(t, got, `Press ESC key or click the button below to close`)
	mustContain(t, got, `<div class="modal-action">`)
}

func TestModalDialogBackdrop(t *testing.T) {
	got := render(t, Modal(ModalConfig{
		ID:           "my_modal_2",
		Title:        "Hello!",
		Message:      "Click outside to close",
		TriggerLabel: "open modal",
		Backdrop:     true,
	}))
	mustContain(t, got, `<form method="dialog" class="modal-backdrop">`)
	mustContain(t, got, `<button>close</button>`)
}

func TestModalDialogCloseButton(t *testing.T) {
	got := render(t, Modal(ModalConfig{
		ID:          "my_modal_3",
		Title:       "Hello!",
		Message:     "Click on ✕ button to close",
		CloseButton: true,
	}))
	mustContain(t, got, `class="btn btn-sm btn-circle btn-ghost absolute right-2 top-2"`)
	mustContain(t, got, `>✕<`)
}

func TestModalDialogCustomWidth(t *testing.T) {
	got := render(t, Modal(ModalConfig{
		ID:       "my_modal_4",
		Title:    "Hello!",
		BoxClass: "w-11/12 max-w-5xl",
	}))
	mustContain(t, got, `class="modal-box w-11/12 max-w-5xl"`)
}

func TestModalDialogPlacement(t *testing.T) {
	got := render(t, Modal(ModalConfig{
		ID:        "my_modal_5",
		Placement: "bottom sm:modal-middle",
		Class:     "modal-bottom sm:modal-middle",
	}))
	mustContain(t, got, `modal-bottom`)
	mustContain(t, got, `sm:modal-middle`)
}

func TestModalCheckbox(t *testing.T) {
	got := render(t, Modal(ModalConfig{
		ID:           "my_modal_6",
		Method:       "checkbox",
		Title:        "Hello!",
		Message:      "This modal works with a hidden checkbox!",
		TriggerLabel: "open modal",
		Actions:      templ.Raw(`<label for="my_modal_6" class="btn">Close!</label>`),
	}))
	mustContain(t, got, `<label for="my_modal_6" class="btn">open modal</label>`)
	mustContain(t, got, `<input type="checkbox" id="my_modal_6" class="modal-toggle">`)
	mustContain(t, got, `<div class="modal" role="dialog">`)
}

func TestModalCheckboxBackdrop(t *testing.T) {
	got := render(t, Modal(ModalConfig{
		ID:           "my_modal_7",
		Method:       "checkbox",
		Title:        "Hello!",
		TriggerLabel: "open modal",
		Backdrop:     true,
	}))
	mustContain(t, got, `<label class="modal-backdrop" for="my_modal_7">Close</label>`)
}

func TestModalAnchor(t *testing.T) {
	got := render(t, Modal(ModalConfig{
		ID:           "my_modal_8",
		Method:       "anchor",
		Title:        "Hello!",
		Message:      "This modal works with anchor links",
		TriggerLabel: "open modal",
		Actions:      templ.Raw(`<a href="#" class="btn">Yay!</a>`),
	}))
	mustContain(t, got, `<a href="#my_modal_8" class="btn" rel="external">open modal</a>`)
	mustContain(t, got, `<div class="modal" role="dialog" id="my_modal_8">`)
}

func TestModalOpenStatic(t *testing.T) {
	got := render(t, Modal(ModalConfig{
		ID:     "my_modal_open",
		Method: "checkbox",
		Open:   true,
		Title:  "Open by default",
	}))
	mustContain(t, got, `class="modal modal-open"`)
	mustContain(t, got, `checked`)
}

func TestModalDialogOpen(t *testing.T) {
	got := render(t, Modal(ModalConfig{
		ID:    "my_modal_dialog_open",
		Open:  true,
		Title: "Open dialog",
	}))
	if !strings.Contains(got, "open") {
		t.Errorf("expected dialog to be open; got:\n%s", got)
	}
}

func TestModalBody(t *testing.T) {
	got := render(t, Modal(ModalConfig{
		ID:   "my_modal_body",
		Body: templ.Raw(`<span class="badge">Custom</span>`),
	}))
	mustContain(t, got, `<span class="badge">Custom</span>`)
}
