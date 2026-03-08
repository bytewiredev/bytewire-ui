package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestList_UlTag(t *testing.T) {
	html := dom.RenderHTML(List(dom.Text("item1")))
	if !strings.Contains(html, "<ul") {
		t.Errorf("expected ul tag in output, got: %s", html)
	}
}

func TestList_LiTags(t *testing.T) {
	html := dom.RenderHTML(List(dom.Text("item1"), dom.Text("item2")))
	if !strings.Contains(html, "<li") {
		t.Errorf("expected li tag in output, got: %s", html)
	}
}

func TestList_BorderBottom(t *testing.T) {
	html := dom.RenderHTML(List(dom.Text("item1")))
	if !strings.Contains(html, "border-bottom") {
		t.Errorf("expected %q in output, got: %s", "border-bottom", html)
	}
}

func TestListItem_PrimaryOnly(t *testing.T) {
	node := ListItem("Primary text")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "Primary text") {
		t.Errorf("expected %q in output, got: %s", "Primary text", html)
	}
}

func TestListItem_WithSecondary(t *testing.T) {
	node := ListItem("Primary", "Secondary info")
	html := dom.RenderHTML(node)
	if !strings.Contains(html, "Secondary info") {
		t.Errorf("expected %q in output, got: %s", "Secondary info", html)
	}
}

func TestDescriptionList_DlTag(t *testing.T) {
	items := []KV{{Key: "Name", Value: "Alice"}}
	html := dom.RenderHTML(DescriptionList(items))
	if !strings.Contains(html, "<dl") {
		t.Errorf("expected dl tag in output, got: %s", html)
	}
}

func TestDescriptionList_DtAndDd(t *testing.T) {
	items := []KV{{Key: "Name", Value: "Alice"}}
	html := dom.RenderHTML(DescriptionList(items))
	if !strings.Contains(html, "<dt") {
		t.Errorf("expected dt tag in output, got: %s", html)
	}
	if !strings.Contains(html, "<dd") {
		t.Errorf("expected dd tag in output, got: %s", html)
	}
}

func TestDescriptionList_KeyAndValue(t *testing.T) {
	items := []KV{{Key: "Status", Value: "Active"}}
	html := dom.RenderHTML(DescriptionList(items))
	if !strings.Contains(html, "Status") {
		t.Errorf("expected %q in output, got: %s", "Status", html)
	}
	if !strings.Contains(html, "Active") {
		t.Errorf("expected %q in output, got: %s", "Active", html)
	}
}
