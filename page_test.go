package ui

import (
	"strings"
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestHero_SectionTag(t *testing.T) {
	html := dom.RenderHTML(Hero("Welcome", "A subtitle"))
	if !strings.Contains(html, "<section") {
		t.Errorf("expected section tag in output, got: %s", html)
	}
}

func TestHero_H1Title(t *testing.T) {
	html := dom.RenderHTML(Hero("Welcome", "A subtitle"))
	if !strings.Contains(html, "<h1") {
		t.Errorf("expected h1 tag in output, got: %s", html)
	}
	if !strings.Contains(html, "Welcome") {
		t.Errorf("expected %q in output, got: %s", "Welcome", html)
	}
}

func TestHero_Subtitle(t *testing.T) {
	html := dom.RenderHTML(Hero("Title", "My subtitle text"))
	if !strings.Contains(html, "<p") {
		t.Errorf("expected p tag in output, got: %s", html)
	}
	if !strings.Contains(html, "My subtitle text") {
		t.Errorf("expected %q in output, got: %s", "My subtitle text", html)
	}
}

func TestHero_ActionButtons(t *testing.T) {
	btn := dom.Button(dom.Children(dom.Text("Click")))
	html := dom.RenderHTML(Hero("Title", "Sub", btn))
	if !strings.Contains(html, "Click") {
		t.Errorf("expected action button text %q in output, got: %s", "Click", html)
	}
	if !strings.Contains(html, "display:flex") {
		t.Errorf("expected actions div with %q in output, got: %s", "display:flex", html)
	}
}

func TestFooter_FooterTag(t *testing.T) {
	html := dom.RenderHTML(Footer("Copyright 2024"))
	if !strings.Contains(html, "<footer") {
		t.Errorf("expected footer tag in output, got: %s", html)
	}
}

func TestFooter_Text(t *testing.T) {
	html := dom.RenderHTML(Footer("Copyright 2024"))
	if !strings.Contains(html, "Copyright 2024") {
		t.Errorf("expected %q in output, got: %s", "Copyright 2024", html)
	}
}

func TestFooter_BorderTop(t *testing.T) {
	html := dom.RenderHTML(Footer("text"))
	if !strings.Contains(html, "border-top") {
		t.Errorf("expected %q in output, got: %s", "border-top", html)
	}
}

func TestPageShell_Navbar(t *testing.T) {
	nav := Navbar("App", nil)
	content := dom.Div(dom.Children(dom.Text("content")))
	html := dom.RenderHTML(PageShell(nav, content))
	if !strings.Contains(html, "<nav") {
		t.Errorf("expected navbar nav tag in output, got: %s", html)
	}
}

func TestPageShell_MainTag(t *testing.T) {
	nav := Navbar("App", nil)
	content := dom.Div(dom.Children(dom.Text("content")))
	html := dom.RenderHTML(PageShell(nav, content))
	if !strings.Contains(html, "<main") {
		t.Errorf("expected main tag in output, got: %s", html)
	}
}

func TestPageShell_MinHeight(t *testing.T) {
	nav := Navbar("App", nil)
	content := dom.Div(dom.Children(dom.Text("content")))
	html := dom.RenderHTML(PageShell(nav, content))
	if !strings.Contains(html, "min-height:100vh") {
		t.Errorf("expected %q in output, got: %s", "min-height:100vh", html)
	}
}

func TestPageShellWithSidebar_Flex(t *testing.T) {
	nav := Navbar("App", nil)
	side := Sidebar("Menu", nil)
	content := dom.Div(dom.Children(dom.Text("content")))
	html := dom.RenderHTML(PageShellWithSidebar(nav, side, content))
	if !strings.Contains(html, "display:flex") {
		t.Errorf("expected %q in output, got: %s", "display:flex", html)
	}
}

func TestPageShellWithSidebar_Sidebar(t *testing.T) {
	nav := Navbar("App", nil)
	side := Sidebar("Menu", nil)
	content := dom.Div(dom.Children(dom.Text("content")))
	html := dom.RenderHTML(PageShellWithSidebar(nav, side, content))
	if !strings.Contains(html, "width:220px") {
		t.Errorf("expected sidebar %q in output, got: %s", "width:220px", html)
	}
}

func TestPageShellWithSidebar_MainTag(t *testing.T) {
	nav := Navbar("App", nil)
	side := Sidebar("Menu", nil)
	content := dom.Div(dom.Children(dom.Text("content")))
	html := dom.RenderHTML(PageShellWithSidebar(nav, side, content))
	if !strings.Contains(html, "<main") {
		t.Errorf("expected main tag in output, got: %s", html)
	}
}

func TestBaseCSS_BoxSizing(t *testing.T) {
	if !strings.Contains(BaseCSS, "box-sizing: border-box") {
		t.Errorf("expected %q in BaseCSS, got: %s", "box-sizing: border-box", BaseCSS)
	}
}
