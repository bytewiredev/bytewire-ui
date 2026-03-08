package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestCard(t *testing.T) {
	node := Card(dom.Text("content"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "background-color:#1e293b") {
		t.Errorf("expected %q in output, got: %s", "background-color:#1e293b", html)
	}
	if !strings.Contains(html, "border-radius:0.75rem") {
		t.Errorf("expected %q in output, got: %s", "border-radius:0.75rem", html)
	}
}

func TestCardAccent(t *testing.T) {
	node := CardAccent("#22c55e", dom.Text("accent"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "border-top:3px solid") {
		t.Errorf("expected %q in output, got: %s", "border-top:3px solid", html)
	}
}

func TestCardHeader(t *testing.T) {
	node := CardHeader(dom.Text("header"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "border-bottom") {
		t.Errorf("expected %q in output, got: %s", "border-bottom", html)
	}
}

func TestCardBody(t *testing.T) {
	node := CardBody(dom.Text("body"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "padding:1.5rem") {
		t.Errorf("expected %q in output, got: %s", "padding:1.5rem", html)
	}
}

func TestCardFooter(t *testing.T) {
	node := CardFooter(dom.Text("footer"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "border-top") {
		t.Errorf("expected %q in output, got: %s", "border-top", html)
	}
}

func TestSimpleCard(t *testing.T) {
	node := SimpleCard("My Title", "My description")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<h3") {
		t.Errorf("expected %q in output, got: %s", "<h3", html)
	}
	if !strings.Contains(html, "My Title") {
		t.Errorf("expected %q in output, got: %s", "My Title", html)
	}
	if !strings.Contains(html, "<p") {
		t.Errorf("expected %q in output, got: %s", "<p", html)
	}
	if !strings.Contains(html, "My description") {
		t.Errorf("expected %q in output, got: %s", "My description", html)
	}
}

func TestFeatureCard(t *testing.T) {
	node := FeatureCard("🚀", "Feature", "A great feature")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "🚀") {
		t.Errorf("expected %q in output, got: %s", "🚀", html)
	}
	if !strings.Contains(html, "Feature") {
		t.Errorf("expected %q in output, got: %s", "Feature", html)
	}
	if !strings.Contains(html, "A great feature") {
		t.Errorf("expected %q in output, got: %s", "A great feature", html)
	}
}

func TestStatCard(t *testing.T) {
	node := StatCard("42", "Users", "#22c55e")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "42") {
		t.Errorf("expected %q in output, got: %s", "42", html)
	}
	if !strings.Contains(html, "Users") {
		t.Errorf("expected %q in output, got: %s", "Users", html)
	}
	if !strings.Contains(html, "font-family") {
		t.Errorf("expected %q in output, got: %s", "font-family", html)
	}
}
