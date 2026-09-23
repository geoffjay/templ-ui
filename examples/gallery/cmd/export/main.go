// Command export prerenders the templ-ui gallery to static HTML for GitHub
// Pages: one index.html plus components/<slug>.html per registry entry, with
// the built assets copied into the output tree. templ components render to
// strings server-side, so the export is a plain render loop — no headless
// browser needed.
package main

import (
	"bytes"
	"context"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/a-h/templ"
	"github.com/geoffjay/templ-ui/examples/gallery"
	_ "github.com/geoffjay/templ-ui/examples/gallery/demos"
	"github.com/geoffjay/templ-ui/internal/registry"
)

func main() {
	outDir := "dist"
	if len(os.Args) > 1 {
		outDir = os.Args[1]
	}

	ctx := context.Background()

	// Render the index.
	if err := render(ctx, filepath.Join(outDir, "index.html"), gallery.IndexPage()); err != nil {
		log.Fatal(err)
	}

	// Render one page per component.
	for _, e := range registry.Sorted() {
		path := filepath.Join(outDir, "components", e.Slug+".html")
		if err := render(ctx, path, gallery.ComponentPage(e)); err != nil {
			log.Fatal(err)
		}
	}

	// Copy the built assets (styles.css, shiki.js) into the export tree.
	staticSrc := "examples/gallery/static"
	if err := filepath.WalkDir(staticSrc, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, err := filepath.Rel(staticSrc, path)
		if err != nil {
			return err
		}
		dst := filepath.Join(outDir, "static", rel)
		if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
			return err
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(dst, data, 0o644)
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("export: wrote %d pages + assets to %s\n", len(registry.Sorted())+1, outDir)
}

// render writes one templ component as an HTML file, creating directories as
// needed.
func render(ctx context.Context, path string, c templ.Component) error {
	var buf bytes.Buffer
	if err := c.Render(ctx, &buf); err != nil {
		return fmt.Errorf("render %s: %w", path, err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), 0o644)
}
