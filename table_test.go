package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestTable(t *testing.T) {
	cols := []TableColumn{
		{Label: "Name"},
		{Label: "Age"},
	}
	rows := [][]string{
		{"Alice", "30"},
		{"Bob", "25"},
	}
	node := Table(cols, rows)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<table") {
		t.Errorf("expected %q in output, got: %s", "<table", html)
	}
	if !strings.Contains(html, "<thead") {
		t.Errorf("expected %q in output, got: %s", "<thead", html)
	}
	if !strings.Contains(html, "<tbody") {
		t.Errorf("expected %q in output, got: %s", "<tbody", html)
	}
	if !strings.Contains(html, "Name") {
		t.Errorf("expected %q in output, got: %s", "Name", html)
	}
	if !strings.Contains(html, "Alice") {
		t.Errorf("expected %q in output, got: %s", "Alice", html)
	}
	if !strings.Contains(html, "text-transform:uppercase") {
		t.Errorf("expected %q in output, got: %s", "text-transform:uppercase", html)
	}
}

func TestKeyValueTable(t *testing.T) {
	pairs := []KV{
		{Key: "Host", Value: "localhost"},
		{Key: "Port", Value: "8080"},
	}
	node := KeyValueTable(pairs)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "Host") {
		t.Errorf("expected %q in output, got: %s", "Host", html)
	}
	if !strings.Contains(html, "localhost") {
		t.Errorf("expected %q in output, got: %s", "localhost", html)
	}
	if !strings.Contains(html, "Key") {
		t.Errorf("expected %q in output, got: %s", "Key", html)
	}
	if !strings.Contains(html, "Value") {
		t.Errorf("expected %q in output, got: %s", "Value", html)
	}
}
