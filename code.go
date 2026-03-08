package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// CodeBlock creates a monospace code block with dark background.
func CodeBlock(code string) *dom.Node {
	return dom.El("pre",
		CSS("background-color", ColorBgSurface, "border-radius", RadiusMd,
			"padding", "1.5rem", "overflow-x", "auto",
			"font-family", FontMono, "font-size", "0.875rem",
			"line-height", "1.6", "color", ColorText,
			"border", "1px solid "+ColorBorder),
		dom.Children(
			dom.El("code",
				dom.Children(dom.Text(code)),
			),
		),
	)
}

// CodeBlockWithTitle creates a code block with a filename or language header.
func CodeBlockWithTitle(title, code string) *dom.Node {
	return dom.Div(
		CSS("border-radius", RadiusMd, "border", "1px solid "+ColorBorder,
			"overflow", "hidden"),
		dom.Children(
			dom.Div(
				CSS("background-color", ColorBgMuted, "padding", "0.5rem 1rem",
					"font-size", "0.75rem", "color", ColorTextMuted,
					"font-family", FontMono, "border-bottom", "1px solid "+ColorBorder),
				dom.Children(dom.Text(title)),
			),
			dom.El("pre",
				CSS("background-color", ColorBgSurface, "padding", "1.5rem",
					"overflow-x", "auto", "margin", "0",
					"font-family", FontMono, "font-size", "0.875rem",
					"line-height", "1.6", "color", ColorText),
				dom.Children(
					dom.El("code",
						dom.Children(dom.Text(code)),
					),
				),
			),
		),
	)
}
