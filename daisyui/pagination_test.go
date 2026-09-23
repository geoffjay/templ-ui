package daisyui

import (
	"strconv"
	"testing"
)

func TestPagination(t *testing.T) {
	got := render(t, Pagination(PaginationConfig{
		Current: 2,
		Total:   4,
	}))
	mustContain(t, got, `class="join"`)
	mustContain(t, got, `<button class="join-item btn">1</button>`)
	mustContain(t, got, `<button class="join-item btn btn-active">2</button>`)
	mustContain(t, got, `<button class="join-item btn">3</button>`)
	mustContain(t, got, `<button class="join-item btn">4</button>`)
}

func TestPaginationSize(t *testing.T) {
	got := render(t, Pagination(PaginationConfig{
		Current: 1,
		Total:   2,
		Size:    "sm",
	}))
	mustContain(t, got, `<button class="join-item btn btn-sm btn-active">1</button>`)
	mustContain(t, got, `<button class="join-item btn btn-sm">2</button>`)
}

func TestPaginationPrevNext(t *testing.T) {
	got := render(t, Pagination(PaginationConfig{
		Current:   2,
		Total:     3,
		PrevLabel: "«",
		NextLabel: "»",
	}))
	mustContain(t, got, `<button class="join-item btn">«</button>`)
	mustContain(t, got, `<button class="join-item btn">»</button>`)
}

func TestPaginationPrevNextDisabled(t *testing.T) {
	got := render(t, Pagination(PaginationConfig{
		Current:   1,
		Total:     3,
		PrevLabel: "«",
		NextLabel: "»",
	}))
	mustContain(t, got, `<button class="join-item btn btn-disabled" disabled>«</button>`)
}

func TestPaginationHrefFor(t *testing.T) {
	got := render(t, Pagination(PaginationConfig{
		Current: 2,
		Total:   3,
		HrefFor: func(p int) string { return "/page/" + strconv.Itoa(p) },
	}))
	mustContain(t, got, `<a href="/page/1" class="join-item btn">1</a>`)
	mustContain(t, got, `<a href="/page/2" class="join-item btn btn-active">2</a>`)
}
