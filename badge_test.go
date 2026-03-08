package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestBadgeDefault(t *testing.T) {
	node := Badge("New")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<span") {
		t.Errorf("expected %q in output, got: %s", "<span", html)
	}
	if !strings.Contains(html, "New") {
		t.Errorf("expected %q in output, got: %s", "New", html)
	}
	if !strings.Contains(html, "border-radius:9999px") {
		t.Errorf("expected %q in output, got: %s", "border-radius:9999px", html)
	}
}

func TestBadgeInfo(t *testing.T) {
	node := Badge("Info", BadgeInfo)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, ColorInfoBg) {
		t.Errorf("expected %q in output, got: %s", ColorInfoBg, html)
	}
}

func TestBadgeSuccess(t *testing.T) {
	node := Badge("OK", BadgeSuccess)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, ColorSuccessBg) {
		t.Errorf("expected %q in output, got: %s", ColorSuccessBg, html)
	}
}

func TestBadgeWarning(t *testing.T) {
	node := Badge("Warn", BadgeWarning)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, ColorWarningBg) {
		t.Errorf("expected %q in output, got: %s", ColorWarningBg, html)
	}
}

func TestBadgeDanger(t *testing.T) {
	node := Badge("Error", BadgeDanger)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, ColorDangerBg) {
		t.Errorf("expected %q in output, got: %s", ColorDangerBg, html)
	}
}

func TestTagDefault(t *testing.T) {
	node := Tag("label")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "padding:0.3rem") {
		t.Errorf("expected %q in output, got: %s", "padding:0.3rem", html)
	}
	if !strings.Contains(html, "border-radius:0.5rem") {
		t.Errorf("expected %q in output, got: %s", "border-radius:0.5rem", html)
	}
}

func TestTagWithVariant(t *testing.T) {
	node := Tag("success", BadgeSuccess)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, ColorSuccessBg) {
		t.Errorf("expected %q in output, got: %s", ColorSuccessBg, html)
	}
}
