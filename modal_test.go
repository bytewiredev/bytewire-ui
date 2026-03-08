package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestModalHeader_TitleInH3(t *testing.T) {
	node := ModalHeader("Confirm", func(_ []byte) {})
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<h3") {
		t.Errorf("expected h3 tag in output, got: %s", html)
	}
	if !strings.Contains(html, "Confirm") {
		t.Errorf("expected %q in output, got: %s", "Confirm", html)
	}
}

func TestModalHeader_CloseButton(t *testing.T) {
	node := ModalHeader("Title", func(_ []byte) {})
	html := dom.RenderHTML(node)
	if !strings.Contains(html, ">x<") {
		t.Errorf("expected close button with %q in output, got: %s", "x", html)
	}
}

func TestModalBody_Padding(t *testing.T) {
	node := ModalBody(dom.Text("body content"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "padding:1.5rem") {
		t.Errorf("expected %q in output, got: %s", "padding:1.5rem", html)
	}
}

func TestModalFooter_JustifyContent(t *testing.T) {
	node := ModalFooter(dom.Text("ok"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "justify-content:flex-end") {
		t.Errorf("expected %q in output, got: %s", "justify-content:flex-end", html)
	}
}

func TestModalFooter_BorderTop(t *testing.T) {
	node := ModalFooter(dom.Text("ok"))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "border-top") {
		t.Errorf("expected %q in output, got: %s", "border-top", html)
	}
}
