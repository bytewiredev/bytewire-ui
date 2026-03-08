package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestContainer(t *testing.T) {
	node := Container(dom.Text("hello"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "max-width:1100px") {
		t.Errorf("expected %q in output, got: %s", "max-width:1100px", html)
	}
	if !strings.Contains(html, "mx-auto") {
		t.Errorf("expected %q in output, got: %s", "mx-auto", html)
	}
}

func TestContainerNarrow(t *testing.T) {
	node := ContainerNarrow(dom.Text("narrow"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "max-width:720px") {
		t.Errorf("expected %q in output, got: %s", "max-width:720px", html)
	}
}

func TestContainerWide(t *testing.T) {
	node := ContainerWide(dom.Text("wide"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "max-width:1400px") {
		t.Errorf("expected %q in output, got: %s", "max-width:1400px", html)
	}
}

func TestVStack(t *testing.T) {
	node := VStack("1rem", dom.Text("a"), dom.Text("b"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "flex-direction:column") {
		t.Errorf("expected %q in output, got: %s", "flex-direction:column", html)
	}
	if !strings.Contains(html, "gap:1rem") {
		t.Errorf("expected %q in output, got: %s", "gap:1rem", html)
	}
}

func TestHStack(t *testing.T) {
	node := HStack("0.5rem", dom.Text("a"), dom.Text("b"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "flex-direction:row") {
		t.Errorf("expected %q in output, got: %s", "flex-direction:row", html)
	}
	if !strings.Contains(html, "gap:0.5rem") {
		t.Errorf("expected %q in output, got: %s", "gap:0.5rem", html)
	}
}

func TestCenter(t *testing.T) {
	node := Center(dom.Text("centered"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "justify-content:center") {
		t.Errorf("expected %q in output, got: %s", "justify-content:center", html)
	}
}

func TestGrid(t *testing.T) {
	node := Grid("280px", "1rem", dom.Text("item"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "grid-template-columns") {
		t.Errorf("expected %q in output, got: %s", "grid-template-columns", html)
	}
	if !strings.Contains(html, "gap:1rem") {
		t.Errorf("expected %q in output, got: %s", "gap:1rem", html)
	}
}

func TestColumns(t *testing.T) {
	node := Columns(3, "1rem", dom.Text("col"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "repeat(3, 1fr)") {
		t.Errorf("expected %q in output, got: %s", "repeat(3, 1fr)", html)
	}
}

func TestSpacer(t *testing.T) {
	node := Spacer("2rem")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "height:2rem") {
		t.Errorf("expected %q in output, got: %s", "height:2rem", html)
	}
}

func TestDivider(t *testing.T) {
	node := Divider()
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<hr") {
		t.Errorf("expected %q in output, got: %s", "<hr", html)
	}
	if !strings.Contains(html, "border-top") {
		t.Errorf("expected %q in output, got: %s", "border-top", html)
	}
}

func TestSection(t *testing.T) {
	node := Section("My Section", dom.Text("content"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<section") {
		t.Errorf("expected %q in output, got: %s", "<section", html)
	}
	if !strings.Contains(html, "My Section") {
		t.Errorf("expected %q in output, got: %s", "My Section", html)
	}
}

func TestSectionAlt(t *testing.T) {
	node := SectionAlt("Alt", dom.Text("alt content"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "background-color") {
		t.Errorf("expected %q in output, got: %s", "background-color", html)
	}
}

func TestSectionTitle(t *testing.T) {
	node := SectionTitle("Title")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<h2") {
		t.Errorf("expected %q in output, got: %s", "<h2", html)
	}
	if !strings.Contains(html, "Title") {
		t.Errorf("expected %q in output, got: %s", "Title", html)
	}
}
