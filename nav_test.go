package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestNavbar_NavTag(t *testing.T) {
	html := dom.RenderHTML(Navbar("MyApp", nil))
	if !strings.Contains(html, "<nav") {
		t.Errorf("expected nav tag in output, got: %s", html)
	}
}

func TestNavbar_BrandLink(t *testing.T) {
	html := dom.RenderHTML(Navbar("MyApp", nil))
	if !strings.Contains(html, `href="/"`) {
		t.Errorf("expected brand link href=\"/\" in output, got: %s", html)
	}
	if !strings.Contains(html, "MyApp") {
		t.Errorf("expected brand text %q in output, got: %s", "MyApp", html)
	}
}

func TestNavbar_LinkItems(t *testing.T) {
	links := []NavItem{
		{Href: "/about", Label: "About"},
		{Href: "/docs", Label: "Docs"},
	}
	html := dom.RenderHTML(Navbar("Brand", links))
	if !strings.Contains(html, "About") {
		t.Errorf("expected %q in output, got: %s", "About", html)
	}
	if !strings.Contains(html, "/docs") {
		t.Errorf("expected %q in output, got: %s", "/docs", html)
	}
}

func TestNavbarExtLink_TargetBlank(t *testing.T) {
	html := dom.RenderHTML(NavbarExtLink("GitHub", "https://github.com"))
	if !strings.Contains(html, `target="_blank"`) {
		t.Errorf("expected %q in output, got: %s", `target="_blank"`, html)
	}
}

func TestNavbarExtLink_RelNoopener(t *testing.T) {
	html := dom.RenderHTML(NavbarExtLink("GitHub", "https://github.com"))
	if !strings.Contains(html, `rel="noopener noreferrer"`) {
		t.Errorf("expected %q in output, got: %s", `rel="noopener noreferrer"`, html)
	}
}

func TestBreadcrumb_NavTag(t *testing.T) {
	items := []NavItem{{Href: "/", Label: "Home"}, {Href: "", Label: "Page"}}
	html := dom.RenderHTML(Breadcrumb(items))
	if !strings.Contains(html, "<nav") {
		t.Errorf("expected nav tag in output, got: %s", html)
	}
}

func TestBreadcrumb_AriaLabel(t *testing.T) {
	items := []NavItem{{Href: "/", Label: "Home"}, {Href: "", Label: "Page"}}
	html := dom.RenderHTML(Breadcrumb(items))
	if !strings.Contains(html, `aria-label="breadcrumb"`) {
		t.Errorf("expected %q in output, got: %s", `aria-label="breadcrumb"`, html)
	}
}

func TestBreadcrumb_Separator(t *testing.T) {
	items := []NavItem{{Href: "/", Label: "Home"}, {Href: "", Label: "Page"}}
	html := dom.RenderHTML(Breadcrumb(items))
	if !strings.Contains(html, "/") {
		t.Errorf("expected separator %q in output, got: %s", "/", html)
	}
}

func TestBreadcrumb_LastItemNoLink(t *testing.T) {
	items := []NavItem{{Href: "/", Label: "Home"}, {Href: "", Label: "Current"}}
	html := dom.RenderHTML(Breadcrumb(items))
	// The last item should be in a span, not an anchor
	idx := strings.LastIndex(html, "Current")
	if idx == -1 {
		t.Fatalf("expected %q in output, got: %s", "Current", html)
	}
	// Check there is no <a wrapping the last item
	before := html[:idx]
	lastA := strings.LastIndex(before, "<a")
	lastSpan := strings.LastIndex(before, "<span")
	if lastA > lastSpan {
		t.Errorf("expected last breadcrumb item to not be a link, got: %s", html)
	}
}

func TestSidebar_NavTag(t *testing.T) {
	html := dom.RenderHTML(Sidebar("Menu", []NavItem{{Href: "/a", Label: "A"}}))
	if !strings.Contains(html, "<nav") {
		t.Errorf("expected nav tag in output, got: %s", html)
	}
}

func TestSidebar_Width(t *testing.T) {
	html := dom.RenderHTML(Sidebar("Menu", []NavItem{{Href: "/a", Label: "A"}}))
	if !strings.Contains(html, "width:220px") {
		t.Errorf("expected %q in output, got: %s", "width:220px", html)
	}
}

func TestSidebar_Title(t *testing.T) {
	html := dom.RenderHTML(Sidebar("Menu", []NavItem{{Href: "/a", Label: "A"}}))
	if !strings.Contains(html, "Menu") {
		t.Errorf("expected %q in output, got: %s", "Menu", html)
	}
}

func TestSidebar_Links(t *testing.T) {
	links := []NavItem{{Href: "/a", Label: "Link A"}, {Href: "/b", Label: "Link B"}}
	html := dom.RenderHTML(Sidebar("Nav", links))
	if !strings.Contains(html, "Link A") {
		t.Errorf("expected %q in output, got: %s", "Link A", html)
	}
	if !strings.Contains(html, "Link B") {
		t.Errorf("expected %q in output, got: %s", "Link B", html)
	}
}
