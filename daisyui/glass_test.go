package daisyui

import "testing"

// TestGlassClass asserts that the glass surface treatment is applied to the
// outer element of each supporting component when Glass is true, and absent
// otherwise.
func TestGlassClass(t *testing.T) {
	t.Run("card", func(t *testing.T) {
		got := render(t, Card(CardConfig{Glass: true}))
		mustContain(t, got, `class="card glass"`)
	})
	t.Run("alert", func(t *testing.T) {
		got := render(t, Alert(AlertConfig{Glass: true}))
		mustContain(t, got, `class="alert glass"`)
	})
	t.Run("badge", func(t *testing.T) {
		got := render(t, Badge(BadgeConfig{Glass: true}))
		mustContain(t, got, `class="badge glass"`)
	})
	t.Run("toast", func(t *testing.T) {
		got := render(t, Toast(ToastConfig{Glass: true}))
		mustContain(t, got, `class="toast glass"`)
	})
	t.Run("dock", func(t *testing.T) {
		got := render(t, Dock(DockConfig{Glass: true}))
		mustContain(t, got, `class="dock glass"`)
	})
	t.Run("footer", func(t *testing.T) {
		got := render(t, Footer(FooterConfig{Glass: true}))
		mustContain(t, got, `class="footer glass"`)
	})
	t.Run("table", func(t *testing.T) {
		got := render(t, Table(TableConfig{Glass: true}))
		mustContain(t, got, `class="table glass"`)
	})
	t.Run("stat", func(t *testing.T) {
		got := render(t, Stat(StatConfig{Glass: true}))
		mustContain(t, got, `class="stats glass"`)
	})
	t.Run("list", func(t *testing.T) {
		got := render(t, List(ListConfig{Glass: true}))
		mustContain(t, got, `class="list glass"`)
	})
	t.Run("modal", func(t *testing.T) {
		got := render(t, Modal(ModalConfig{ID: "m", Glass: true}))
		mustContain(t, got, `class="modal-box glass"`)
	})
	t.Run("hero", func(t *testing.T) {
		got := render(t, Hero(HeroConfig{Glass: true}))
		mustContain(t, got, `class="hero-content glass"`)
	})
	t.Run("navbar", func(t *testing.T) {
		got := render(t, Navbar(NavbarData{Glass: true}))
		mustContain(t, got, `class="navbar bg-base-100 shadow-sm glass"`)
	})
	t.Run("chat-bubble", func(t *testing.T) {
		got := render(t, ChatBubble(ChatBubbleConfig{Glass: true}))
		mustContain(t, got, `class="chat-bubble glass"`)
	})
	t.Run("collapse", func(t *testing.T) {
		got := render(t, Collapse(CollapseConfig{Glass: true}, "t", "b"))
		mustContain(t, got, `class="collapse glass"`)
	})
	t.Run("drawer", func(t *testing.T) {
		got := render(t, Drawer(DrawerConfig{Glass: true}))
		mustContain(t, got, `menu p-4 w-60 md:w-80 min-h-full bg-base-200 glass`)
	})
}

// TestGlassAbsent verifies that glass is not added when the flag is false.
func TestGlassAbsent(t *testing.T) {
	got := render(t, Card(CardConfig{}))
	mustNotContain(t, got, "glass")
}

func mustNotContain(t *testing.T, got, want string) {
	t.Helper()
	if contains(got, want) {
		t.Errorf("expected output not to contain %q, got:\n%s", want, got)
	}
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
