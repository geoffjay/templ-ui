package daisyui

import (
	"testing"

	"github.com/a-h/templ"
)

func TestChatBubbleStart(t *testing.T) {
	got := render(t, ChatBubble(ChatBubbleConfig{
		Message: "It's over Anakin, I have the high ground.",
	}))
	mustContain(t, got, `class="chat chat-start"`)
	mustContain(t, got, `<div class="chat-bubble">`)
	mustContain(t, got, `>It&#39;s over Anakin, I have the high ground.<`)
}

func TestChatBubbleEndWithColor(t *testing.T) {
	got := render(t, ChatBubble(ChatBubbleConfig{
		End:     true,
		Color:   "info",
		Message: "Calm down, Anakin.",
	}))
	mustContain(t, got, `class="chat chat-end"`)
	mustContain(t, got, `class="chat-bubble chat-bubble-info"`)
	mustContain(t, got, `>Calm down, Anakin.<`)
}

func TestChatBubbleHeaderFooter(t *testing.T) {
	got := render(t, ChatBubble(ChatBubbleConfig{
		Author:  "Obi-Wan Kenobi",
		Time:    "12:45",
		Message: "You were the Chosen One!",
		Footer:  "Delivered",
	}))
	mustContain(t, got, `<div class="chat-header">`)
	mustContain(t, got, `>Obi-Wan Kenobi <time`)
	mustContain(t, got, `<time class="text-xs opacity-50">12:45</time>`)
	mustContain(t, got, `<div class="chat-footer opacity-50">Delivered</div>`)
}

func TestChatBubbleAvatar(t *testing.T) {
	got := render(t, ChatBubble(ChatBubbleConfig{
		Avatar:  templ.Raw(`<div class="w-10 rounded-full"><img src="/x.webp" alt="a"/></div>`),
		Message: "hi",
	}))
	mustContain(t, got, `<div class="chat-image avatar">`)
	mustContain(t, got, `<img src="/x.webp" alt="a"/>`)
}
