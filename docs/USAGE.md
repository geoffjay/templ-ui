# Using templ-ui components in your application

templ-ui components render HTML with daisyUI/Tailwind class names. They
carry no styling of their own beyond a few library-owned CSS rules, so your
application owns the CSS build. This guide covers the two integration points
— the Tailwind stylesheet and the shiki highlighter — plus the head helper.

## 1. Tailwind CSS + daisyUI

daisyUI 5 is a Tailwind CSS v4 plugin: component classes only exist after
your stylesheet is compiled with the plugin. Use Tailwind v4's CLI:

```sh
bun add -d tailwindcss @tailwindcss/cli daisyui @tailwindcss/typography
bunx tailwindcss -i ./assets/styles.css -o ./static/styles.css
```

Your CSS entry needs (copy `assets/styles.css` from this repo as a starting
point):

```css
@import "tailwindcss";

/* Every class templ-ui components can render (see "The safelist" below). */
@import "./vendor/templ-ui/safelist.css";   /* e.g. assets/vendor/ next to this file */

@plugin "@tailwindcss/typography";   /* optional: for prose content */

@plugin "daisyui" {
  themes:
    light --default,
    dark --prefersdark,
    nord,
    nord-dark;   /* list every theme you switch between at runtime */
}
```

Notes:

- daisyUI emits each theme's CSS custom properties at build time; **themes
  not listed are not available at runtime**. The `theme-controller` and
  `data-theme` switching used by `containers.AppShell` only works for
  themes included here.
- The `.glass` and `.component-preview` rules in `assets/styles.css` are
  owned by templ-ui (the `Glass` flag on several daisyui components relies
  on `.glass`). Keep them in your compiled stylesheet.

### The safelist (required)

Tailwind and daisyUI only emit CSS for class names they find in scanned
sources. That breaks for templ-ui in two ways: the component sources live in
the Go module cache, which Tailwind never scans, and many daisyUI modifiers
are assembled at runtime (`"btn-" + cfg.Color`), so no scanner could find
them anyway. The symptom is silent: components render with missing styles
(e.g. every `tab-content` panel visible at once, colorless buttons).

templ-ui embeds a generated list of Tailwind v4 `@source inline(...)`
directives covering every class the components emit (`daisyui.Safelist`).
Write it into your CSS build inputs as the first step of your CSS build:

```sh
go run github.com/geoffjay/templ-ui/cmd/templ-ui safelist -o assets/vendor/templ-ui/safelist.css
```

`go run` resolves templ-ui at the version pinned in your `go.mod`, so the
safelist always matches the components you render: bump templ-ui and the
next CSS build picks up any new classes. Treat the output as a build
artifact (gitignore it) or commit it; either way, regenerate it in the build
rather than copying it by hand. For example, in `package.json`:

```json
"scripts": {
  "safelist": "go run github.com/geoffjay/templ-ui/cmd/templ-ui safelist -o assets/vendor/templ-ui/safelist.css",
  "css": "bun run safelist && tailwindcss -i ./assets/styles.css -o ./static/styles.css"
}
```

Avoid a top-level `vendor/` directory in a Go module: its presence switches
the go command into vendoring mode.

Classes you write in your own templates (including a component's `Class`
field) are still picked up by Tailwind's normal scan of your sources; add
`@source` directives if your templ files live outside the source root.

## 2. shiki code highlighting

`shiki.Code` renders a plain escaped `<pre><code>` block server-side and
progressively enhances it client-side with a shiki bundle. The bundle is
ESM, built with esbuild from `assets/shiki.js` (langs: go, html, bash,
templ, javascript, css; themes: nord, github-dark, vitesse-dark,
vitesse-light; JavaScript-regex engine so it satisfies `script-src 'self'`
CSP).

Two ways to get the bundle:

**Option A — use the prebuilt recipe.** Copy `assets/shiki.js` and build:

```sh
bun add -d esbuild shiki @shikijs/langs @shikijs/themes
bunx esbuild assets/shiki.js --bundle --format=esm --target=es2022 \
  --outfile=static/shiki.js
```

Then serve `static/shiki.js` from your app and point templ-ui at it by
setting the package variable before rendering:

```go
shiki.ScriptSrc = "/static/shiki.js"   // default; change to your path
```

**Option B — vendor your own subset.** The loader script is ~80 lines; the
only contract is `[data-shiki][data-lang][data-theme]` elements containing
`<code>` text. Import additional languages/themes from `@shikijs/langs` /
`@shikijs/themes` to grow the set.

## 3. The Head helper

`shiki.Head` renders the common head assets — one stylesheet link and (when
enabled) the highlighter module — so a layout component stays tidy:

```templ
package layout

import "github.com/geoffjay/templ-ui/shiki"

templ Page(title string) {
	<!DOCTYPE html>
	<html lang="en">
		<head>
			@shiki.Head(shiki.HeadConfig{
				Stylesheet: "/static/styles.css",
				Shiki:      true,
			})
			<title>{ title }</title>
		</head>
		<body>
			{ children... }
		</body>
	</html>
}
```

`Shiki: true` loads the bundle once in the head (it begins fetching while
the body parses). `shiki.Code` also emits its own module script tag; the
browser deduplicates module scripts by `src`, so both tags coexist safely.

## 4. AppShell container

```templ
@containers.AppShell(containers.AppShellConfig{
	Title:    "My App",
	NavSections: []containers.NavSection{
		{Title: "Main", Items: []daisyui.MenuItem{
			{Label: "Dashboard", Href: "/", Icon: dashboardIcon},
		}},
	},
	Content: mainContent,
	Theme: containers.ThemeConfig{
		LightTheme: "nord",
		DarkTheme:  "nord-dark",
		StorageKey: "myapp.theme",
	},
})
```

- `MenuItem.Icon` is any `templ.Component`; `github.com/iota-uz/icons`
  (phosphor) works out of the box and is what AppShell's own toggle icons
  use.
- `Theme` renders a sun/moon toggle and restore/persist scripts; both
  themes must be in your daisyUI build (see §1).
- `Resolver daisyui.URLResolver` optionally rewrites all hrefs (jughead
  uses it for site-prefix-aware links).

The shell is responsive with three breakpoint ranges (Tailwind `md`/`lg`):

- **small** (<768px): the sidebar is hidden, the header shows a hamburger
  button instead of the sidebar toggle, and the hamburger opens a menu that
  expands from the top down below the header.
- **medium** (768–1023px): the sidebar renders as a collapsed icon rail.
- **large** (>=1024px): the sidebar renders per its persisted
  open/collapsed/closed state.

## 5. What the components expect at runtime

| Component | Requires |
|---|---|
| all `daisyui.*` | compiled Tailwind + daisyUI stylesheet |
| `Glass: true` variants | the `.glass` rule in that stylesheet |
| `shiki.Code` | the shiki bundle at `shiki.ScriptSrc` (plain text renders fine without it) |
| `containers.AppShell` | daisyUI themes listed in your CSS build |
| any `templ.Component` icons | an icon package (e.g. `github.com/iota-uz/icons`) — not required, pass your own SVGs |