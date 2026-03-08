package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestButtonDefault(t *testing.T) {
	node := Button("Click", nil)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<button") {
		t.Errorf("expected %q in output, got: %s", "<button", html)
	}
	if !strings.Contains(html, "Click") {
		t.Errorf("expected %q in output, got: %s", "Click", html)
	}
	if !strings.Contains(html, "background-color:#22c55e") {
		t.Errorf("expected %q in output, got: %s", "background-color:#22c55e", html)
	}
}

func TestButtonSecondary(t *testing.T) {
	node := Button("Secondary", nil, WithVariant(ButtonSecondary))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "background-color:#1e293b") {
		t.Errorf("expected %q in output, got: %s", "background-color:#1e293b", html)
	}
}

func TestButtonOutline(t *testing.T) {
	node := Button("Outline", nil, WithVariant(ButtonOutline))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "background-color:transparent") {
		t.Errorf("expected %q in output, got: %s", "background-color:transparent", html)
	}
}

func TestButtonGhost(t *testing.T) {
	node := Button("Ghost", nil, WithVariant(ButtonGhost))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "background-color:transparent") {
		t.Errorf("expected %q in output, got: %s", "background-color:transparent", html)
	}
}

func TestButtonDanger(t *testing.T) {
	node := Button("Delete", nil, WithVariant(ButtonDanger))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "#ef4444") {
		t.Errorf("expected %q in output, got: %s", "#ef4444", html)
	}
}

func TestButtonSmall(t *testing.T) {
	node := Button("Small", nil, WithSize(ButtonSm))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "0.375rem") {
		t.Errorf("expected %q in output, got: %s", "0.375rem", html)
	}
}

func TestButtonLarge(t *testing.T) {
	node := Button("Large", nil, WithSize(ButtonLg))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "0.875rem") {
		t.Errorf("expected %q in output, got: %s", "0.875rem", html)
	}
}

func TestButtonDisabled(t *testing.T) {
	node := Button("Disabled", nil, Disabled())
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "opacity:0.5") {
		t.Errorf("expected %q in output, got: %s", "opacity:0.5", html)
	}
	if !strings.Contains(html, `disabled="true"`) {
		t.Errorf("expected %q in output, got: %s", `disabled="true"`, html)
	}
}

func TestButtonFullWidth(t *testing.T) {
	node := Button("Full", nil, FullWidth())
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "width:100%") {
		t.Errorf("expected %q in output, got: %s", "width:100%", html)
	}
}

func TestButtonLink(t *testing.T) {
	node := ButtonLink("Go", "/page")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<a") {
		t.Errorf("expected %q in output, got: %s", "<a", html)
	}
	if !strings.Contains(html, "/page") {
		t.Errorf("expected %q in output, got: %s", "/page", html)
	}
	if !strings.Contains(html, "inline-block") {
		t.Errorf("expected %q in output, got: %s", "inline-block", html)
	}
}

func TestButtonGroup(t *testing.T) {
	node := ButtonGroup(Button("A", nil), Button("B", nil))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "flex") {
		t.Errorf("expected %q in output, got: %s", "flex", html)
	}
	if !strings.Contains(html, "gap:0.5rem") {
		t.Errorf("expected %q in output, got: %s", "gap:0.5rem", html)
	}
}
