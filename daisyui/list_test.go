package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestList(t *testing.T) {
	got := render(t, List(ListConfig{
		Class:  "bg-base-100 rounded-box shadow-md",
		Header: templ.Raw("Most played songs this week"),
		Items: []ListItem{
			{
				Items: []templ.Component{
					templ.Raw(`<div><img class="size-10 rounded-box" src="x.webp"/></div>`),
					ListColGrow(templ.Raw(`<div>Dio Lupa</div>`)),
					templ.Raw(`<button class="btn btn-square btn-ghost">play</button>`),
				},
			},
		},
	}))
	mustContain(t, got, `<ul class="list bg-base-100 rounded-box shadow-md">`)
	mustContain(t, got, `<li class="p-4 pb-2 text-xs opacity-60 tracking-wide">`)
	mustContain(t, got, "Most played songs this week")
	mustContain(t, got, `<li class="list-row">`)
	mustContain(t, got, `<div class="list-col-grow"><div>Dio Lupa</div></div>`)
	mustContain(t, got, `<button class="btn btn-square btn-ghost">play</button>`)
}

func TestListOrdered(t *testing.T) {
	got := render(t, List(ListConfig{
		Ordered: true,
		Items: []ListItem{
			{Row: templ.Raw(`<div>A</div>`)},
		},
	}))
	mustContain(t, got, `<ol class="list">`)
	mustContain(t, got, `<div>A</div>`)
}

func TestListColWrap(t *testing.T) {
	got := render(t, ListColWrap(templ.Raw(`description`)))
	mustContain(t, got, `<p class="list-col-wrap">`)
	mustContain(t, got, "description")
}
