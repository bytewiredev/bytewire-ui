package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
	"github.com/bytewiredev/bytewire/pkg/style"
)

// NavItem defines a navigation link.
type NavItem struct {
	Href  string
	Label string
}

// Navbar creates a top navigation bar with a brand name and links.
func Navbar(brand string, links []NavItem, extraItems ...*dom.Node) *dom.Node {
	items := []*dom.Node{
		dom.A(
			dom.Link("/"),
			dom.Class(style.Classes(style.TextWhite, style.FontBold, style.TextLg, style.Mr2)),
			CSS("padding", "0.25rem 0.75rem", "text-decoration", "none"),
			dom.Children(dom.Text(brand)),
		),
	}

	for _, l := range links {
		items = append(items, dom.A(
			dom.Link(l.Href),
			CSS("color", ColorText, "padding", "0.25rem 0.75rem",
				"text-decoration", "none", "font-weight", "500",
				"font-size", "0.875rem", "border-radius", RadiusMd),
			dom.Children(dom.Text(l.Label)),
		))
	}

	items = append(items, extraItems...)

	return dom.Nav(
		CSS("display", "flex", "align-items", "center", "gap", "0.5rem",
			"padding", "0.75rem 1rem",
			"background-color", ColorBg,
			"border-bottom", "1px solid "+ColorBgSurface),
		dom.Children(items...),
	)
}

// NavbarExtLink creates an external link styled for the navbar.
func NavbarExtLink(label, href string) *dom.Node {
	return dom.A(
		dom.Attr("href", href),
		dom.Attr("target", "_blank"),
		dom.Attr("rel", "noopener noreferrer"),
		CSS("color", ColorText, "padding", "0.25rem 0.75rem",
			"text-decoration", "none", "font-weight", "500",
			"font-size", "0.875rem"),
		dom.Children(dom.Text(label)),
	)
}

// Breadcrumb creates a breadcrumb trail from label/href pairs.
// The last item is rendered as plain text (current page).
func Breadcrumb(items []NavItem) *dom.Node {
	var crumbs []*dom.Node
	for i, item := range items {
		if i > 0 {
			crumbs = append(crumbs, dom.Span(
				CSS("color", ColorTextSubtle, "margin", "0 0.5rem"),
				dom.Children(dom.Text("/")),
			))
		}
		if i == len(items)-1 {
			// Current page (no link)
			crumbs = append(crumbs, dom.Span(
				CSS("color", ColorText, "font-weight", "600"),
				dom.Children(dom.Text(item.Label)),
			))
		} else {
			crumbs = append(crumbs, dom.A(
				dom.Link(item.Href),
				CSS("color", ColorTextMuted, "text-decoration", "none"),
				dom.Children(dom.Text(item.Label)),
			))
		}
	}
	return dom.Nav(
		dom.Attr("aria-label", "breadcrumb"),
		CSS("font-size", "0.8rem", "margin-bottom", "1rem"),
		dom.Children(crumbs...),
	)
}

// Sidebar creates a vertical navigation sidebar.
func Sidebar(title string, links []NavItem) *dom.Node {
	var items []*dom.Node
	if title != "" {
		items = append(items, dom.Div(
			CSS("color", ColorTextMuted, "font-size", "0.7rem", "font-weight", "600",
				"text-transform", "uppercase", "letter-spacing", "0.05em",
				"margin-bottom", "0.75rem", "padding", "0 0.75rem"),
			dom.Children(dom.Text(title)),
		))
	}
	for _, l := range links {
		items = append(items, dom.A(
			dom.Link(l.Href),
			CSS("display", "block", "color", ColorText,
				"padding", "0.5rem 0.75rem", "text-decoration", "none",
				"font-size", "0.875rem", "border-radius", RadiusMd),
			dom.Children(dom.Text(l.Label)),
		))
	}
	return dom.Nav(
		CSS("width", "220px", "padding", "1rem 0",
			"border-right", "1px solid "+ColorBorder),
		dom.Children(items...),
	)
}
