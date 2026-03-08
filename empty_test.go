package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestEmptyState_TitleInH3(t *testing.T) {
	html := dom.RenderHTML(EmptyState("No items", "Nothing to show"))
	if !strings.Contains(html, "<h3") {
		t.Errorf("expected h3 tag in output, got: %s", html)
	}
	if !strings.Contains(html, "No items") {
		t.Errorf("expected %q in output, got: %s", "No items", html)
	}
}

func TestEmptyState_DescriptionInP(t *testing.T) {
	html := dom.RenderHTML(EmptyState("No items", "Nothing to show"))
	if !strings.Contains(html, "<p") {
		t.Errorf("expected p tag in output, got: %s", html)
	}
	if !strings.Contains(html, "Nothing to show") {
		t.Errorf("expected %q in output, got: %s", "Nothing to show", html)
	}
}

func TestEmptyState_Placeholder(t *testing.T) {
	html := dom.RenderHTML(EmptyState("Title", "Desc"))
	if !strings.Contains(html, "--") {
		t.Errorf("expected placeholder %q in output, got: %s", "--", html)
	}
}

func TestEmptyState_ActionDiv(t *testing.T) {
	btn := dom.Button(dom.Children(dom.Text("Add")))
	html := dom.RenderHTML(EmptyState("Empty", "Nothing here", btn))
	if !strings.Contains(html, "Add") {
		t.Errorf("expected action button %q in output, got: %s", "Add", html)
	}
	if !strings.Contains(html, "margin-top:1.5rem") {
		t.Errorf("expected action div with %q in output, got: %s", "margin-top:1.5rem", html)
	}
}

func TestEmptyState_WithoutAction(t *testing.T) {
	html := dom.RenderHTML(EmptyState("Empty", "Nothing here"))
	// Should still render title and description without error
	if !strings.Contains(html, "Empty") {
		t.Errorf("expected %q in output, got: %s", "Empty", html)
	}
	// Should not contain the action wrapper margin
	if strings.Contains(html, "margin-top:1.5rem") {
		t.Errorf("expected no action div in output, got: %s", html)
	}
}

func TestSkeleton_Width(t *testing.T) {
	html := dom.RenderHTML(Skeleton("200px", "20px"))
	if !strings.Contains(html, "width:200px") {
		t.Errorf("expected %q in output, got: %s", "width:200px", html)
	}
}

func TestSkeleton_Height(t *testing.T) {
	html := dom.RenderHTML(Skeleton("200px", "20px"))
	if !strings.Contains(html, "height:20px") {
		t.Errorf("expected %q in output, got: %s", "height:20px", html)
	}
}

func TestSkeleton_PulseAnimation(t *testing.T) {
	html := dom.RenderHTML(Skeleton("100px", "16px"))
	if !strings.Contains(html, "animation:pulse") {
		t.Errorf("expected %q in output, got: %s", "animation:pulse", html)
	}
}

func TestSkeletonCSS_KeyframesPulse(t *testing.T) {
	if !strings.Contains(SkeletonCSS, "@keyframes pulse") {
		t.Errorf("expected %q in SkeletonCSS, got: %s", "@keyframes pulse", SkeletonCSS)
	}
}
