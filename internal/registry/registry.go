// Package registry is the index of gallery entries: one per component or
// container documented by the templ-ui gallery app. Each entry pairs a slug
// (the URL path segment) with the page body builder. The gallery server and
// the static exporter both walk this registry, so a component appears in the
// demo app, the index page, and the GitHub Pages export as soon as its entry
// exists — one registration, three surfaces.
package registry

import (
	"sort"

	"github.com/a-h/templ"
)

// Entry describes one gallery page. Slug is the URL path segment (also the
// static export filename). Name is the human title. Description is the
// one-line summary shown on the index cards and page headers (sourced from
// the daisyUI components listing). Body builds the page content: a list of
// preview sections rendered inside the shared page shell. Category groups
// entries in the index and the sidebar navigation.
type Entry struct {
	Slug        string
	Name        string
	Description string
	Category    string
	Body        func() templ.Component
}

// Entries is the ordered registry. Component demos append to it from their
// demo_*.go files via Register; ordering is finalized by Sorted/ByCategory.
var Entries []Entry

// Register adds an entry. Called from per-component init functions in the
// demos package, so importing demos is enough to populate the gallery.
func Register(e Entry) {
	Entries = append(Entries, e)
}

// Sorted returns the registry entries ordered by category then name — the
// order used by the index page and the export loop.
func Sorted() []Entry {
	out := make([]Entry, len(Entries))
	copy(out, Entries)
	sort.SliceStable(out, func(i, j int) bool {
		if out[i].Category != out[j].Category {
			return out[i].Category < out[j].Category
		}
		return out[i].Name < out[j].Name
	})
	return out
}

// Find returns the entry with the given slug, or nil.
func Find(slug string) *Entry {
	for i := range Entries {
		if Entries[i].Slug == slug {
			return &Entries[i]
		}
	}
	return nil
}
