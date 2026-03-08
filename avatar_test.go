package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestAvatar_DefaultBorderRadius(t *testing.T) {
	html := dom.RenderHTML(Avatar("AB"))
	if !strings.Contains(html, "border-radius:9999px") {
		t.Errorf("expected %q in output, got: %s", "border-radius:9999px", html)
	}
}

func TestAvatar_DefaultInitials(t *testing.T) {
	html := dom.RenderHTML(Avatar("JD"))
	if !strings.Contains(html, "JD") {
		t.Errorf("expected %q in output, got: %s", "JD", html)
	}
}

func TestAvatar_DefaultWidth(t *testing.T) {
	html := dom.RenderHTML(Avatar("AB"))
	if !strings.Contains(html, "width:2.5rem") {
		t.Errorf("expected %q in output, got: %s", "width:2.5rem", html)
	}
}

func TestAvatar_SmallWidth(t *testing.T) {
	html := dom.RenderHTML(Avatar("AB", AvatarSm))
	if !strings.Contains(html, "width:2rem") {
		t.Errorf("expected %q in output, got: %s", "width:2rem", html)
	}
}

func TestAvatar_LargeWidth(t *testing.T) {
	html := dom.RenderHTML(Avatar("AB", AvatarLg))
	if !strings.Contains(html, "width:3.5rem") {
		t.Errorf("expected %q in output, got: %s", "width:3.5rem", html)
	}
}

func TestAvatar_XlWidth(t *testing.T) {
	html := dom.RenderHTML(Avatar("AB", AvatarXl))
	if !strings.Contains(html, "width:5rem") {
		t.Errorf("expected %q in output, got: %s", "width:5rem", html)
	}
}

func TestAvatarImage_ImgTag(t *testing.T) {
	html := dom.RenderHTML(AvatarImage("https://example.com/pic.jpg", "John"))
	if !strings.Contains(html, "<img") {
		t.Errorf("expected img tag in output, got: %s", html)
	}
}

func TestAvatarImage_Src(t *testing.T) {
	html := dom.RenderHTML(AvatarImage("https://example.com/pic.jpg", "John"))
	if !strings.Contains(html, `src="https://example.com/pic.jpg"`) {
		t.Errorf("expected src attribute in output, got: %s", html)
	}
}

func TestAvatarImage_Alt(t *testing.T) {
	html := dom.RenderHTML(AvatarImage("https://example.com/pic.jpg", "John"))
	if !strings.Contains(html, `alt="John"`) {
		t.Errorf("expected alt attribute in output, got: %s", html)
	}
}

func TestAvatarImage_BorderRadius(t *testing.T) {
	html := dom.RenderHTML(AvatarImage("https://example.com/pic.jpg", "John"))
	if !strings.Contains(html, "border-radius:9999px") {
		t.Errorf("expected %q in output, got: %s", "border-radius:9999px", html)
	}
}
