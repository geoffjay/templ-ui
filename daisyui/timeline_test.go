package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

const timelineCheckSVG = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" class="h-5 w-5"><path/></svg>`

func TestTimeline(t *testing.T) {
	got := render(t, Timeline(TimelineConfig{
		Items: []TimelineItem{
			{
				Side:    TimelineSideStart,
				Icon:    templ.Raw(timelineCheckSVG),
				Content: templ.Raw("1984"),
				HRAfter: true,
			},
			{
				Side:     TimelineSideEnd,
				Icon:     templ.Raw(timelineCheckSVG),
				Content:  templ.Raw("iMac"),
				Box:      true,
				HRBefore: true,
			},
		},
	}))
	mustContain(t, got, `<ul class="timeline">`)
	mustContain(t, got, `<div class="timeline-start">1984</div>`)
	mustContain(t, got, `<div class="timeline-middle">`)
	mustContain(t, got, `<div class="timeline-end timeline-box">iMac</div>`)
	mustContain(t, got, `<hr>`)
}

func TestTimelineDirections(t *testing.T) {
	for _, d := range []struct {
		dir, expect string
	}{
		{"horizontal", "timeline timeline-horizontal"},
		{"vertical", "timeline timeline-vertical"},
	} {
		got := render(t, Timeline(TimelineConfig{Direction: d.dir}))
		mustContain(t, got, `class="`+d.expect+`"`)
	}
}

func TestTimelineModifiers(t *testing.T) {
	for _, m := range []string{"snap-icon", "compact"} {
		got := render(t, Timeline(TimelineConfig{Modifier: m}))
		mustContain(t, got, `class="timeline timeline-`+m+`"`)
	}
}

func TestTimelineColorHR(t *testing.T) {
	got := render(t, Timeline(TimelineConfig{
		Items: []TimelineItem{
			{Side: TimelineSideStart, Content: templ.Raw("A"), HRAfter: true, HRClass: "bg-primary"},
		},
	}))
	mustContain(t, got, `<hr class="bg-primary">`)
}
