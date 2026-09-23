# templ-ui

Reusable [templ](https://github.com/a-h/templ) UI components: a full
[daisyUI 5](https://daisyui.com) component set, [shiki](https://shiki.style)
code highlighting, and layout containers — extracted from
[jughead](https://github.com/geoffjay/jughead) into a standalone public
module.

## Packages

| Package | Contents |
|---|---|
| `daisyui` | 66 daisyUI 5 components (button, card, navbar, modal, table, …) with render tests |
| `shiki` | `shiki.Code` code blocks + `shiki.Head` page-head asset helper |
| `containers` | `AppShell`: fixed navbar + collapsible sidebar layout with theme persistence |
| `internal/registry` | gallery registry (drives the demo app and the Pages export) |
| `examples/gallery` | component gallery: live server, static export, and per-component demo pages |

## Install

```sh
go get github.com/geoffjay/templ-ui
```

Consuming a component is a normal templ call:

```templ
package main

import "github.com/geoffjay/templ-ui/daisyui"

templ hello() {
	@daisyui.Button(daisyui.ButtonConfig{Label: "Click me", Color: "primary"})
}
```

daisyUI components emit daisyUI class names; they need a compiled Tailwind
CSS v4 + daisyUI stylesheet in your application. See
[docs/USAGE.md](docs/USAGE.md) for the required CSS build, the shiki setup,
and the head component.

## Gallery / demo

The component gallery shows every component live next to the templ source
that rendered it, with a light/dark theme switcher:

```sh
make assets    # build shiki.js + styles.css into examples/gallery/static/
make run-demo  # serve http://localhost:8280
```

For live development use [pitchfork](https://pitchfork.jdx.dev) (replaces
overmind-style Procfiles): it runs the templ generator, the gallery server,
and both asset watchers with restart-on-change:

```sh
pitchfork start          # all daemons (templ, gallery, css, shiki)
pitchfork logs gallery   # tail one daemon
```

The same pages export to static HTML for GitHub Pages (deployed by
`.github/workflows/pages.yml`):

```sh
go run ./examples/gallery/cmd/export examples/gallery/dist
```

## Development

```sh
make generate  # templ generate (pinned CLI version)
make test      # go test ./...
make lint      # go vet + gofmt check
make ci        # lint + test
```

- Go 1.25+, templ v0.3.1020 (pinned by `make templ`).
- Generated `*_templ.go` files are committed (matches templ-charts/jughead
  convention).
- Frontend toolchain is [bun](https://bun.sh): Tailwind v4 CLI + esbuild for
  the shiki bundle (`bun install && bun run build`).

## License

MIT
