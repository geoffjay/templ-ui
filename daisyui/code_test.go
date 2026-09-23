package daisyui

import (
	"strings"
	"testing"
)

func TestCode(t *testing.T) {
	got := render(t, Code(CodeConfig{
		Class: "w-full",
		Lines: []CodeLine{
			{Prefix: "$", Text: "npm i daisyui"},
			{Prefix: ">", Text: "installing...", Class: "text-warning"},
			{Prefix: ">", Text: "Done!", Class: "text-success"},
		},
	}))
	mustContain(t, got, `<div class="mockup-code w-full">`)
	mustContain(t, got, `data-prefix="$"`)
	mustContain(t, got, `<code>npm i daisyui</code>`)
	mustContain(t, got, `class="text-warning"`)
	mustContain(t, got, `<code>installing...</code>`)
	mustContain(t, got, `class="text-success"`)
}

func TestCodeWithoutPrefix(t *testing.T) {
	got := render(t, Code(CodeConfig{
		Lines: []CodeLine{{Text: "without prefix"}},
	}))
	mustContain(t, got, `<pre><code>without prefix</code></pre>`)
	if strings.Contains(got, "data-prefix") {
		t.Errorf("did not expect data-prefix:\n%s", got)
	}
}

func TestCodeWithColor(t *testing.T) {
	got := render(t, Code(CodeConfig{
		Class: "bg-primary text-primary-content w-full",
		Lines: []CodeLine{{Text: "can be any color!"}},
	}))
	mustContain(t, got, `class="mockup-code bg-primary text-primary-content w-full"`)
}
