package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestAccordionItem_StructExists(t *testing.T) {
	_ = AccordionItem{
		Title: "Section 1",
		Content: func() *dom.Node {
			return dom.Text("content")
		},
	}
}

func TestAccordion_Border(t *testing.T) {
	items := []AccordionItem{
		{Title: "First", Content: func() *dom.Node { return dom.Text("body") }},
	}
	html := dom.RenderHTML(Accordion(items))
	if !strings.Contains(html, "border:1px solid") {
		t.Errorf("expected %q in output, got: %s", "border:1px solid", html)
	}
}

func TestAccordion_ButtonWithTitle(t *testing.T) {
	items := []AccordionItem{
		{Title: "My Section", Content: func() *dom.Node { return dom.Text("body") }},
	}
	html := dom.RenderHTML(Accordion(items))
	if !strings.Contains(html, "<button") {
		t.Errorf("expected button tag in output, got: %s", html)
	}
	if !strings.Contains(html, "My Section") {
		t.Errorf("expected %q in output, got: %s", "My Section", html)
	}
}

func TestAccordion_PlusIndicator(t *testing.T) {
	items := []AccordionItem{
		{Title: "Item", Content: func() *dom.Node { return dom.Text("body") }},
	}
	html := dom.RenderHTML(Accordion(items))
	if !strings.Contains(html, "+") {
		t.Errorf("expected %q indicator in output, got: %s", "+", html)
	}
}
