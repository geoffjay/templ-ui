package daisyui

import "testing"

func TestDiffViewRendersAddedAndRemoved(t *testing.T) {
	got := render(t, DiffView(DiffViewConfig{
		Files: []DiffViewFile{
			{
				Path: "handlers.go", Added: 1, Removed: 1,
				Hunks: []DiffViewHunk{
					{
						Lines: []DiffViewLine{
							{Kind: DiffViewContext, OldLineNo: 1, NewLineNo: 1, Text: "ctx"},
							{Kind: DiffViewRemoved, OldLineNo: 2, NewLineNo: 0, Text: "old"},
							{Kind: DiffViewAdded, OldLineNo: 0, NewLineNo: 2, Text: "new"},
						},
					},
				},
			},
		},
	}))
	mustContain(t, got, "handlers.go")
	mustContain(t, got, `badge-success`)
	mustContain(t, got, `badge-error`)
	mustContain(t, got, `bg-error/20 text-error-content`)
	mustContain(t, got, `bg-success/20 text-success-content`)
	mustContain(t, got, ">old<")
	mustContain(t, got, ">new<")
	mustContain(t, got, ">ctx<")
}
