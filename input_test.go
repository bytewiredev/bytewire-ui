package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestInput(t *testing.T) {
	node := Input("Enter name", nil)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<input") {
		t.Errorf("expected %q in output, got: %s", "<input", html)
	}
	if !strings.Contains(html, `type="text"`) {
		t.Errorf("expected %q in output, got: %s", `type="text"`, html)
	}
	if !strings.Contains(html, `placeholder="Enter name"`) {
		t.Errorf("expected %q in output, got: %s", `placeholder="Enter name"`, html)
	}
}

func TestPasswordInput(t *testing.T) {
	node := PasswordInput("Enter password", nil)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, `type="password"`) {
		t.Errorf("expected %q in output, got: %s", `type="password"`, html)
	}
}

func TestTextarea(t *testing.T) {
	node := Textarea("Write here", 5, nil)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<textarea") {
		t.Errorf("expected %q in output, got: %s", "<textarea", html)
	}
	if !strings.Contains(html, `rows="5"`) {
		t.Errorf("expected %q in output, got: %s", `rows="5"`, html)
	}
	if !strings.Contains(html, `placeholder="Write here"`) {
		t.Errorf("expected %q in output, got: %s", `placeholder="Write here"`, html)
	}
}

func TestSelect(t *testing.T) {
	opts := []SelectOption{
		{Value: "a", Label: "Option A"},
		{Value: "b", Label: "Option B"},
	}
	node := Select(opts, nil)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<select") {
		t.Errorf("expected %q in output, got: %s", "<select", html)
	}
	if !strings.Contains(html, "<option") {
		t.Errorf("expected %q in output, got: %s", "<option", html)
	}
	if !strings.Contains(html, `value="a"`) {
		t.Errorf("expected %q in output, got: %s", `value="a"`, html)
	}
	if !strings.Contains(html, "Option A") {
		t.Errorf("expected %q in output, got: %s", "Option A", html)
	}
}

func TestCheckbox(t *testing.T) {
	node := Checkbox("Accept terms", nil)
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "<label") {
		t.Errorf("expected %q in output, got: %s", "<label", html)
	}
	if !strings.Contains(html, `type="checkbox"`) {
		t.Errorf("expected %q in output, got: %s", `type="checkbox"`, html)
	}
	if !strings.Contains(html, "Accept terms") {
		t.Errorf("expected %q in output, got: %s", "Accept terms", html)
	}
}

func TestFormField(t *testing.T) {
	node := FormField("Email", Input("you@example.com", nil), "Your email is safe")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "Email") {
		t.Errorf("expected %q in output, got: %s", "Email", html)
	}
	if !strings.Contains(html, "Your email is safe") {
		t.Errorf("expected %q in output, got: %s", "Your email is safe", html)
	}
}

func TestFormGroup(t *testing.T) {
	node := FormGroup(Input("first", nil), Input("last", nil))
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "grid-template-columns") {
		t.Errorf("expected %q in output, got: %s", "grid-template-columns", html)
	}
}
