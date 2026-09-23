// Command gallery runs the templ-ui component gallery: a stdlib net/http
// server serving one page per registered component plus the index. The same
// pages are exported to static HTML by cmd/export for GitHub Pages.
package main

import (
	"log"
	"net/http"

	"github.com/geoffjay/templ-ui/examples/gallery"
	_ "github.com/geoffjay/templ-ui/examples/gallery/demos"
	"github.com/geoffjay/templ-ui/internal/registry"
)

func main() {
	mux := http.NewServeMux()

	// Static assets (compiled stylesheet + shiki bundle) built by `make assets`.
	// The path is relative to the repo root (the server is run via
	// `go run ./examples/gallery/cmd/server` or pitchfork from there).
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("examples/gallery/static"))))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		if err := gallery.IndexPage().Render(r.Context(), w); err != nil {
			log.Printf("index render: %v", err)
		}
	})

	mux.HandleFunc("GET /components/{slug}", func(w http.ResponseWriter, r *http.Request) {
		e := registry.Find(r.PathValue("slug"))
		if e == nil {
			http.NotFound(w, r)
			return
		}
		if err := gallery.ComponentPage(*e).Render(r.Context(), w); err != nil {
			log.Printf("component render: %v", err)
		}
	})

	log.Println("gallery: http://localhost:8280")
	log.Fatal(http.ListenAndServe("localhost:8280", mux))
}
