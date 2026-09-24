.PHONY: all templ generate build test vet fmt lint assets css shiki safelist css-audit run export demo tidy clean ci

TEMPL_PKG := github.com/a-h/templ/cmd/templ
TEMPL_VERSION := v0.3.1020
BUN := bun

# Find all directories containing .templ files
TEMPL_SOURCES := $(shell find . -type f -name '*.templ' -not -path './node_modules/*' 2>/dev/null)

## Install templ CLI if not present
$(HOME)/go/bin/templ:
	go install $(TEMPL_PKG)@$(TEMPL_VERSION)

## Generate Go code from .templ files
templ: $(HOME)/go/bin/templ
	@ if [ -n "$(TEMPL_SOURCES)" ]; then templ generate; else echo "no .templ files yet"; fi
	@ # Detach the "templ: version:" comment from the package clause so it does
	@ # not become package documentation on pkg.go.dev (repeats once per file).
	@ find . -name '*_templ.go' -not -path './node_modules/*' -exec perl -0pi -e 's{(^// templ: version:.*\n)(package )}{$$1\n$$2}m' {} +

generate: templ

## Build all Go packages
build:
	go build ./...

## Run tests
test:
	go test ./...

## Run go vet
vet:
	go vet ./...

## Check gofmt (check only)
fmt:
	@ out=$$(gofmt -l . 2>/dev/null | grep -v '/node_modules/' | grep -v '/.git/'); if [ -n "$$out" ]; then echo "gofmt needs to format:"; echo "$$out"; exit 1; else echo "gofmt: ok"; fi

## Lint (vet + fmt)
lint: vet fmt

## Build frontend assets: safelist + shiki bundle + compiled stylesheet
assets: safelist shiki css

## Build the shiki highlighter bundle into examples/gallery/static/
shiki:
	$(BUN) run build:shiki

## Compile the Tailwind v4 + daisyUI stylesheet into examples/gallery/static/
css:
	$(BUN) run css

## Regenerate daisyui/safelist.css (every class the components can render)
safelist:
	$(BUN) run safelist

## Export the gallery and fail if any rendered class is missing from the stylesheet
css-audit: assets
	rm -rf /tmp/templ-ui-audit
	go run ./examples/gallery/cmd/export /tmp/templ-ui-audit
	$(BUN) run css:audit /tmp/templ-ui-audit examples/gallery/static/styles.css

## Run the gallery demo server (http://localhost:8280)
run-demo:
	go run ./examples/gallery/cmd/server

## Export the gallery to static HTML for GitHub Pages (writes examples/gallery/dist/)
export:
	go run ./examples/gallery/cmd/export

## CI entry point: lint + test
ci: lint test
	@ echo "ci: ok"

## Tidy modules
tidy:
	go mod tidy

## Clean generated artifacts
clean:
	find . -type f -name 'templ_*.go' -not -path './node_modules/*' -delete 2>/dev/null || true
