// Command export prerenders the templ-ui gallery to static HTML for GitHub
// Pages: one index.html plus components/<slug>.html per registry entry, with
// the built assets copied into the output tree. templ components render to
// strings server-side, so the export is a plain render loop — no headless
// browser needed.
//
//	export [-base /templ-ui] [outdir]
//
// -base is the URL path prefix the site is served under (GitHub Pages
// project sites live at /<repo>/); every gallery link and asset URL is
// rendered with it.
package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
	"github.com/geoffjay/templ-ui/examples/gallery"
	_ "github.com/geoffjay/templ-ui/examples/gallery/demos"
	"github.com/geoffjay/templ-ui/internal/registry"
	"github.com/geoffjay/templ-ui/shiki"
)

func main() {
	base := flag.String("base", "", `URL path prefix the site is served under, e.g. "/templ-ui"`)
	flag.Parse()
	outDir := "dist"
	if flag.NArg() > 0 {
		outDir = flag.Arg(0)
	}

	// Normalize to "" or "/prefix" (no trailing slash) so URL("/x") joins
	// cleanly.
	b := strings.TrimSuffix(*base, "/")
	if b != "" && !strings.HasPrefix(b, "/") {
		b = "/" + b
	}
	gallery.Base = b
	shiki.ScriptSrc = gallery.URL("/static/shiki.js")

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
