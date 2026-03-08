package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
	"github.com/bytewiredev/bytewire/pkg/style"
)

// H1 creates a large page heading.
func H1(text string) *dom.Node {
	return dom.H1(
		dom.Class(style.Classes(style.FontBold)),
		CSS("font-size", "2.5rem", "color", ColorTextHeading, "line-height", "1.2", "margin-bottom", "1rem"),
		dom.Children(dom.Text(text)),
	)
}

// H2 creates a section heading.
func H2(text string) *dom.Node {
	return dom.H2(
		dom.Class(style.Classes(style.FontBold)),
		CSS("font-size", "1.75rem", "color", ColorTextHeading, "line-height", "1.3", "margin-bottom", "0.75rem"),
		dom.Children(dom.Text(text)),
	)
}

// H3 creates a subsection heading.
func H3(text string) *dom.Node {
	return dom.H3(
		dom.Class(style.Classes(style.FontBold)),
		CSS("font-size", "1.25rem", "color", ColorTextHeading, "line-height", "1.4", "margin-bottom", "0.5rem"),
		dom.Children(dom.Text(text)),
	)
}

// Paragraph creates a body text paragraph.
func Paragraph(text string) *dom.Node {
	return dom.P(
		CSS("color", ColorText, "line-height", "1.6", "margin-bottom", "1rem"),
		dom.Children(dom.Text(text)),
	)
}

// Muted creates dimmed secondary text.
func Muted(text string) *dom.Node {
	return dom.P(
		CSS("color", ColorTextMuted, "line-height", "1.6", "font-size", "0.875rem"),
		dom.Children(dom.Text(text)),
	)
}

// Small creates small text.
func Small(text string) *dom.Node {
	return dom.Span(
		CSS("color", ColorTextSubtle, "font-size", "0.75rem"),
		dom.Children(dom.Text(text)),
	)
}

// Strong creates bold inline text.
func Strong(text string) *dom.Node {
	return dom.El("strong",
		CSS("color", ColorTextHeading, "font-weight", "700"),
		dom.Children(dom.Text(text)),
	)
}

// InlineCode creates inline monospace code text.
func InlineCode(text string) *dom.Node {
	return dom.El("code",
		CSS("background-color", ColorBgSurface, "color", ColorPrimary,
			"padding", "0.15rem 0.4rem", "border-radius", RadiusSm,
			"font-family", FontMono, "font-size", "0.85em"),
		dom.Children(dom.Text(text)),
	)
}

// TextLink creates a styled anchor link for inline text.
func TextLink(text, href string) *dom.Node {
	return dom.A(
		dom.Link(href),
		CSS("color", ColorPrimary, "text-decoration", "none", "font-weight", "600"),
		dom.Children(dom.Text(text)),
	)
}

// ExternalLink creates a link that opens in a new tab.
func ExternalLink(text, href string) *dom.Node {
	return dom.A(
		dom.Attr("href", href),
		dom.Attr("target", "_blank"),
		dom.Attr("rel", "noopener noreferrer"),
		CSS("color", ColorPrimary, "text-decoration", "none", "font-weight", "600"),
		dom.Children(dom.Text(text)),
	)
}
