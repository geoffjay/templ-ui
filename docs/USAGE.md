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
- Tailwind v4 scans templates for class names automatically; if your templ
  files live outside the source root, add `@source` directives.

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

## 5. What the components expect at runtime

| Component | Requires |
|---|---|
| all `daisyui.*` | compiled Tailwind + daisyUI stylesheet |
| `Glass: true` variants | the `.glass` rule in that stylesheet |
| `shiki.Code` | the shiki bundle at `shiki.ScriptSrc` (plain text renders fine without it) |
| `containers.AppShell` | daisyUI themes listed in your CSS build |
| any `templ.Component` icons | an icon package (e.g. `github.com/iota-uz/icons`) — not required, pass your own SVGs |