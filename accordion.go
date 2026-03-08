package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// AccordionItem defines a collapsible section.
type AccordionItem struct {
	Title   string
	Content func() *dom.Node
}

// Accordion creates a collapsible accordion.
// Each item gets its own signal to track open/closed state.
func Accordion(items []AccordionItem) *dom.Node {
	var sections []*dom.Node
	for _, item := range items {
		open := dom.NewSignal(false)
		sections = append(sections, accordionSection(open, item.Title, item.Content))
	}
	return dom.Div(
		CSS("border", "1px solid "+ColorBorder, "border-radius", RadiusMd, "overflow", "hidden"),
		dom.Children(sections...),
	)
}

func accordionSection(open *dom.Signal[bool], title string, content func() *dom.Node) *dom.Node {
	return dom.Div(
		CSS("border-bottom", "1px solid "+ColorBorder),
		dom.Children(
			// Header (always visible)
			dom.Button(
				CSS("display", "flex", "align-items", "center", "justify-content", "space-between",
					"width", "100%", "padding", "1rem 1.25rem",
					"background", "none", "border", "none",
					"color", ColorTextHeading, "font-size", "0.9rem", "font-weight", "600",
					"cursor", "pointer", "text-align", "left"),
				dom.OnClick(func(_ []byte) {
					open.Update(func(v bool) bool { return !v })
				}),
				dom.Children(
					dom.Span(dom.Children(dom.Text(title))),
					dom.TextF(open, func(v bool) string {
						if v {
							return "-"
						}
						return "+"
					}),
				),
			),
			// Content (collapsible)
			dom.If(open,
				func(v bool) bool { return v },
				func() *dom.Node {
					return dom.Div(
						CSS("padding", "0 1.25rem 1.25rem",
							"color", ColorTextMuted, "font-size", "0.875rem", "line-height", "1.6"),
						dom.Children(content()),
					)
				},
				nil,
			),
		),
	)
}
