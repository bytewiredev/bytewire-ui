package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestCodeBlock_PreTag(t *testing.T) {
	html := dom.RenderHTML(CodeBlock("fmt.Println()"))
	if !strings.Contains(html, "<pre") {
		t.Errorf("expected pre tag in output, got: %s", html)
	}
}

func TestCodeBlock_CodeTag(t *testing.T) {
	html := dom.RenderHTML(CodeBlock("fmt.Println()"))
	if !strings.Contains(html, "<code") {
		t.Errorf("expected code tag in output, got: %s", html)
	}
}

func TestCodeBlock_FontFamily(t *testing.T) {
	html := dom.RenderHTML(CodeBlock("x := 1"))
	if !strings.Contains(html, "font-family") {
		t.Errorf("expected %q in output, got: %s", "font-family", html)
	}
}

func TestCodeBlock_Content(t *testing.T) {
	html := dom.RenderHTML(CodeBlock("hello world"))
	if !strings.Contains(html, "hello world") {
		t.Errorf("expected %q in output, got: %s", "hello world", html)
	}
}

func TestCodeBlockWithTitle_TitleDiv(t *testing.T) {
	html := dom.RenderHTML(CodeBlockWithTitle("main.go", "package main"))
	if !strings.Contains(html, "main.go") {
		t.Errorf("expected title %q in output, got: %s", "main.go", html)
	}
}

func TestCodeBlockWithTitle_PreAndCode(t *testing.T) {
	html := dom.RenderHTML(CodeBlockWithTitle("app.js", "const x = 1"))
	if !strings.Contains(html, "<pre") {
		t.Errorf("expected pre tag in output, got: %s", html)
	}
	if !strings.Contains(html, "<code") {
		t.Errorf("expected code tag in output, got: %s", html)
	}
}

func TestCodeBlockWithTitle_CodeContent(t *testing.T) {
	html := dom.RenderHTML(CodeBlockWithTitle("test.go", "func Test()"))
	if !strings.Contains(html, "func Test()") {
		t.Errorf("expected %q in output, got: %s", "func Test()", html)
	}
}
