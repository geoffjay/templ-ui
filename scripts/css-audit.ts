// Fails when the exported gallery renders a class that the compiled
// stylesheet does not define. Catches the two ways templ-ui pages end up
// silently unstyled: a stale stylesheet (built before a component/demo
// changed) and runtime-assembled daisyUI modifiers missing from the
// generated safelist.
//
// Usage: bun scripts/css-audit.ts <export-dir> <styles.css>
// (CI runs it after `go run ./examples/gallery/cmd/export`.)

import { readdirSync, readFileSync, statSync } from "node:fs";
import { join, relative } from "node:path";

const [exportDir, cssPath] = process.argv.slice(2);
if (!exportDir || !cssPath) {
  console.error("usage: bun scripts/css-audit.ts <export-dir> <styles.css>");
  process.exit(2);
}

// Classes rendered on purpose with no stylesheet rule.
const NO_CSS: Record<string, true> = {
  // Marker read by the shiki client bundle, never styled.
  "shiki-code": true,
};

const defined = new Set<string>();
for (const m of readFileSync(cssPath, "utf8").matchAll(/\.((?:\\.|[A-Za-z0-9_-])+)/g)) {
  defined.add(m[1].replace(/\\(.)/g, "$1"));
}

const firstSeen = new Map<string, string>();
const walk = (dir: string) => {
  for (const name of readdirSync(dir)) {
    const path = join(dir, name);
    if (statSync(path).isDirectory()) walk(path);
    else if (path.endsWith(".html")) {
      for (const m of readFileSync(path, "utf8").matchAll(/class="([^"]*)"/g)) {
        for (const cls of m[1].split(/\s+/)) {
          if (cls && !firstSeen.has(cls)) firstSeen.set(cls, relative(exportDir, path));
        }
      }
    }
  }
};
walk(exportDir);

const missing = [...firstSeen].filter(([cls]) => !defined.has(cls) && !NO_CSS[cls]).sort();
if (missing.length > 0) {
  console.error(`css-audit: ${missing.length} rendered class(es) have no CSS rule:`);
  for (const [cls, page] of missing) console.error(`  ${cls.padEnd(32)} ${page}`);
  console.error("Rebuild the stylesheet (bun run css); if a component builds the class at runtime, run bun run safelist.");
  process.exit(1);
}
console.log(`css-audit: all ${firstSeen.size} rendered classes are defined`);
