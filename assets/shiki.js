// Progressive enhancement for every [data-shiki] element: replace the plain
// <code> text with shiki's syntax-highlighted HTML. Uses the JavaScript regex
// engine (no WebAssembly) so it stays within the page's script-src 'self'
// CSP and avoids a separate .wasm fetch.
//
// Built against shiki/core with explicit lang/theme imports so the bundle
// pulls in only the grammars we use (and never imports shiki/wasm, which the
// full index does). Bundled with esbuild because Parcel 2.16's resolver
// mishandles the ESM subpath exports of @shikijs/langs and @shikijs/themes.
//
// The script is an ES module loaded once per page via a <script type="module"
// src from shiki.ScriptSrc, default "/static/shiki.js"> tag emitted by the shiki.Code templ component; the
// browser deduplicates module scripts by src.

import bash from "@shikijs/langs/bash";
import css from "@shikijs/langs/css";

import go from "@shikijs/langs/go";
import html from "@shikijs/langs/html";
import javascript from "@shikijs/langs/javascript";
import templ from "@shikijs/langs/templ";
import typescript from "@shikijs/langs/typescript";
import githubDark from "@shikijs/themes/github-dark";
import nord from "@shikijs/themes/nord";
import vitesseDark from "@shikijs/themes/vitesse-dark";
import vitesseLight from "@shikijs/themes/vitesse-light";
import { createHighlighterCore } from "shiki/dist/core.mjs";
import { createJavaScriptRegexEngine } from "shiki/dist/engine-javascript.mjs";

const defaultTheme = "nord";

const langs = { go, html, bash, templ, javascript, typescript, css };
const themes = {
  nord,
  "github-dark": githubDark,
  "vitesse-dark": vitesseDark,
  "vitesse-light": vitesseLight,
};

function applyHighlighter(highlighter) {
  const nodes = document.querySelectorAll("pre[data-shiki]");
  for (const node of nodes) {
    const code = node.querySelector("code");
    if (!code) continue;
    const lang = node.dataset.lang || "text";
    const theme = node.dataset.theme || defaultTheme;
    const extraClass = node.className.replace(/\bshiki-code\b/g, "").trim();
    try {
      const html = highlighter.codeToHtml(code.textContent, {
        lang,
        theme,
      });
      const wrapper = document.createElement("template");
      wrapper.innerHTML = html.trim();
      const pre = wrapper.content.firstElementChild;
      if (pre) {
        const merged =
          (pre.className || "") + (extraClass ? " " + extraClass : "");
        if (merged) pre.className = merged.trim();
        node.outerHTML = pre.outerHTML;
      } else {
        node.outerHTML = html;
      }
    } catch (err) {
      // Unknown lang/theme (not preloaded): leave the plain <pre><code> as-is.
      console.warn("[shiki] skipped node:", err);
    }
  }
}

async function main() {
  const highlighter = await createHighlighterCore({
    themes: Object.values(themes),
    langs: Object.values(langs),
    engine: createJavaScriptRegexEngine(),
  });
  applyHighlighter(highlighter);
}

main().catch((err) => {
  console.error("[shiki] failed to initialize:", err);
});
