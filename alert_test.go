package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestAlertInfo(t *testing.T) {
	node := Alert("Info message", AlertInfo)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, `role="alert"`) {
		t.Errorf("expected %q in output, got: %s", `role="alert"`, html)
	}
	if !strings.Contains(html, "border-left") {
		t.Errorf("expected %q in output, got: %s", "border-left", html)
	}
	if !strings.Contains(html, "Info message") {
		t.Errorf("expected %q in output, got: %s", "Info message", html)
	}
}

func TestAlertSuccess(t *testing.T) {
	node := Alert("Success", AlertSuccess)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, ColorSuccessBg) {
		t.Errorf("expected %q in output, got: %s", ColorSuccessBg, html)
	}
}

func TestAlertWarning(t *testing.T) {
	node := Alert("Warning", AlertWarning)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, ColorWarningBg) {
		t.Errorf("expected %q in output, got: %s", ColorWarningBg, html)
	}
}

func TestAlertDanger(t *testing.T) {
	node := Alert("Danger", AlertDanger)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, ColorDangerBg) {
		t.Errorf("expected %q in output, got: %s", ColorDangerBg, html)
	}
}

func TestAlertWithTitle(t *testing.T) {
	node := AlertWithTitle("Error Title", "Something went wrong", AlertDanger)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "Error Title") {
		t.Errorf("expected %q in output, got: %s", "Error Title", html)
	}
	if !strings.Contains(html, "<p") {
		t.Errorf("expected %q in output, got: %s", "<p", html)
	}
	if !strings.Contains(html, "Something went wrong") {
		t.Errorf("expected %q in output, got: %s", "Something went wrong", html)
	}
	if !strings.Contains(html, "font-weight:700") {
		t.Errorf("expected %q in output, got: %s", "font-weight:700", html)
	}
}
