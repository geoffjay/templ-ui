package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestTable(t *testing.T) {
	got := render(t, Table(TableConfig{
		Headers: []TableCell{{Text: ""}, {Text: "Name"}, {Text: "Job"}},
		Rows: []TableRow{
			{
				Cells: []TableCell{
					{Text: "1", Header: true},
					{Text: "Cy"},
					{Text: "QC Specialist"},
				},
			},
			{
				Class: "bg-base-200",
				Cells: []TableCell{
					{Text: "2", Header: true},
					{Text: "Hart"},
					{HTML: templ.Raw(`<span class="badge badge-ghost">Support</span>`)},
				},
			},
		},
	}))
	mustContain(t, got, `<table class="table">`)
	mustContain(t, got, `<thead>`)
	mustContain(t, got, `<th>Name</th>`)
	mustContain(t, got, `<th>1</th>`)
	mustContain(t, got, `<td>Cy</td>`)
	mustContain(t, got, `<tr class="bg-base-200">`)
	mustContain(t, got, `<span class="badge badge-ghost">Support</span>`)
}

func TestTableModifiers(t *testing.T) {
	cases := []struct {
		mod, expect string
	}{
		{"zebra", "table table-zebra"},
		{"pin-rows", "table table-pin-rows"},
		{"pin-cols", "table table-pin-cols"},
	}
	for _, c := range cases {
		got := render(t, Table(TableConfig{Modifier: c.mod}))
		mustContain(t, got, `class="`+c.expect+`"`)
	}
}

func TestTableSize(t *testing.T) {
	got := render(t, Table(TableConfig{Size: "xs"}))
	mustContain(t, got, `class="table table-xs"`)
}

func TestTableWrap(t *testing.T) {
	got := render(t, Table(TableConfig{
		Wrap:      true,
		WrapClass: "h-96",
		Size:      "xs",
		Modifier:  "pin-rows",
	}))
	mustContain(t, got, `<div class="overflow-x-auto h-96">`)
	mustContain(t, got, `class="table table-pin-rows table-xs"`)
}

func TestTableFooters(t *testing.T) {
	got := render(t, Table(TableConfig{
		Headers: []TableCell{{Text: "Name"}},
		Rows:    []TableRow{{Cells: []TableCell{{Text: "Cy"}}}},
		Footers: []TableCell{{Text: "Name"}},
	}))
	mustContain(t, got, `<tfoot>`)
}

func TestTableCaption(t *testing.T) {
	got := render(t, Table(TableConfig{
		Caption: "Sales figures",
	}))
	mustContain(t, got, `<caption class="sr-only">Sales figures</caption>`)
}
