// Command templ-ui provides build-time helpers for applications consuming
// templ-ui. Run it through `go run` so the output always matches the
// templ-ui version pinned in the application's go.mod:
//
//	go run github.com/geoffjay/templ-ui/cmd/templ-ui safelist -o assets/vendor/templ-ui/safelist.css
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/geoffjay/templ-ui/daisyui"
)

const usage = `usage: templ-ui <command> [flags]

commands:
  safelist [-o file]   write the Tailwind v4 class safelist for templ-ui
                       components (stdout when -o is omitted)
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
	switch os.Args[1] {
	case "safelist":
		if err := safelist(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, "templ-ui safelist:", err)
			os.Exit(1)
		}
	case "-h", "-help", "--help", "help":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "templ-ui: unknown command %q\n\n%s", os.Args[1], usage)
		os.Exit(2)
	}
}

func safelist(args []string) error {
	fs := flag.NewFlagSet("safelist", flag.ContinueOnError)
	out := fs.String("o", "", "output file (parent directories are created); stdout when empty")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if fs.NArg() > 0 {
		return fmt.Errorf("unexpected arguments: %v", fs.Args())
	}
	if *out == "" {
		_, err := io.WriteString(os.Stdout, daisyui.Safelist)
		return err
	}
	if err := os.MkdirAll(filepath.Dir(*out), 0o755); err != nil {
		return err
	}
	// Write to a temp file and rename so a failed run never leaves a
	// truncated safelist for Tailwind to pick up.
	tmp, err := os.CreateTemp(filepath.Dir(*out), ".safelist-*.css")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	if _, err := io.WriteString(tmp, daisyui.Safelist); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Chmod(0o644); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), *out)
}
