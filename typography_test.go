package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestH1(t *testing.T) {
	node := H1("Main Title")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<h1") {
		t.Errorf("expected %q in output, got: %s", "<h1", html)
	}
	if !strings.Contains(html, "font-size:2.5rem") {
		t.Errorf("expected %q in output, got: %s", "font-size:2.5rem", html)
	}
	if !strings.Contains(html, "Main Title") {
		t.Errorf("expected %q in output, got: %s", "Main Title", html)
	}
}

func TestH2(t *testing.T) {
	node := H2("Section Title")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<h2") {
		t.Errorf("expected %q in output, got: %s", "<h2", html)
	}
	if !strings.Contains(html, "font-size:1.75rem") {
		t.Errorf("expected %q in output, got: %s", "font-size:1.75rem", html)
	}
}

func TestH3(t *testing.T) {
	node := H3("Subsection")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<h3") {
		t.Errorf("expected %q in output, got: %s", "<h3", html)
	}
	if !strings.Contains(html, "font-size:1.25rem") {
		t.Errorf("expected %q in output, got: %s", "font-size:1.25rem", html)
	}
}

func TestParagraph(t *testing.T) {
	node := Paragraph("Hello world")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<p") {
		t.Errorf("expected %q in output, got: %s", "<p", html)
	}
	if !strings.Contains(html, "Hello world") {
		t.Errorf("expected %q in output, got: %s", "Hello world", html)
	}
}

func TestMuted(t *testing.T) {
	node := Muted("secondary text")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "0.875rem") {
		t.Errorf("expected %q in output, got: %s", "0.875rem", html)
	}
	if !strings.Contains(html, "#94a3b8") {
		t.Errorf("expected %q in output, got: %s", "#94a3b8", html)
	}
}

func TestSmall(t *testing.T) {
	node := Small("tiny text")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "0.75rem") {
		t.Errorf("expected %q in output, got: %s", "0.75rem", html)
	}
	if !strings.Contains(html, "<span") {
		t.Errorf("expected %q in output, got: %s", "<span", html)
	}
}

func TestStrong(t *testing.T) {
	node := Strong("bold text")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<strong") {
		t.Errorf("expected %q in output, got: %s", "<strong", html)
	}
	if !strings.Contains(html, "font-weight:700") {
		t.Errorf("expected %q in output, got: %s", "font-weight:700", html)
	}
}

func TestInlineCode(t *testing.T) {
	node := InlineCode("fmt.Println")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<code") {
		t.Errorf("expected %q in output, got: %s", "<code", html)
	}
	if !strings.Contains(html, "font-family") {
		t.Errorf("expected %q in output, got: %s", "font-family", html)
	}
}

func TestTextLink(t *testing.T) {
	node := TextLink("click me", "https://example.com")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<a") {
		t.Errorf("expected %q in output, got: %s", "<a", html)
	}
	if !strings.Contains(html, "https://example.com") {
		t.Errorf("expected %q in output, got: %s", "https://example.com", html)
	}
}

func TestExternalLink(t *testing.T) {
	node := ExternalLink("ext link", "https://example.com")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, `target="_blank"`) {
		t.Errorf("expected %q in output, got: %s", `target="_blank"`, html)
	}
	if !strings.Contains(html, `rel="noopener noreferrer"`) {
		t.Errorf("expected %q in output, got: %s", `rel="noopener noreferrer"`, html)
	}
}
