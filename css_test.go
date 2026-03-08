package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestCSS_MultipleProperties(t *testing.T) {
	node := dom.Div(CSS("display", "grid", "gap", "1rem"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "display:grid;gap:1rem") {
		t.Errorf("expected %q in output, got: %s", "display:grid;gap:1rem", html)
	}
}

func TestCSS_SinglePair(t *testing.T) {
	node := dom.Div(CSS("color", "red"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "color:red") {
		t.Errorf("expected %q in output, got: %s", "color:red", html)
	}
}

func TestCSS_EmptyPairs(t *testing.T) {
	node := dom.Div(CSS())
	html := dom.RenderHTML(node)
	if strings.Contains(html, "style") {
		// If it does produce a style attr, it should be empty
		if !strings.Contains(html, `style=""`) {
			t.Errorf("expected empty or no style attribute, got: %s", html)
		}
	}
}
