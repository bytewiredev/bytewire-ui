package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// List creates an unordered list.
func List(items ...*dom.Node) *dom.Node {
	var lis []*dom.Node
	for _, item := range items {
		lis = append(lis, dom.Li(
			CSS("padding", "0.5rem 0", "border-bottom", "1px solid "+ColorBorder),
			dom.Children(item),
		))
	}
	return dom.Ul(
		CSS("list-style", "none", "padding", "0", "margin", "0"),
		dom.Children(lis...),
	)
}

// ListItem creates a list item with primary text and optional secondary text.
func ListItem(primary string, secondary ...string) *dom.Node {
	children := []*dom.Node{
		dom.Div(
			CSS("color", ColorText, "font-size", "0.875rem"),
			dom.Children(dom.Text(primary)),
		),
	}
	if len(secondary) > 0 && secondary[0] != "" {
		children = append(children, dom.Div(
			CSS("color", ColorTextMuted, "font-size", "0.75rem", "margin-top", "0.125rem"),
			dom.Children(dom.Text(secondary[0])),
		))
	}
	return dom.Div(dom.Children(children...))
}

// DescriptionList creates a description list from key-value pairs.
func DescriptionList(items []KV) *dom.Node {
	var children []*dom.Node
	for _, item := range items {
		children = append(children,
			dom.El("dt",
				CSS("color", ColorTextMuted, "font-size", "0.75rem", "font-weight", "600",
					"text-transform", "uppercase", "letter-spacing", "0.03em",
					"margin-bottom", "0.25rem"),
				dom.Children(dom.Text(item.Key)),
			),
			dom.El("dd",
				CSS("color", ColorText, "margin", "0 0 1.25rem 0",
					"font-size", "0.875rem", "line-height", "1.6"),
				dom.Children(dom.Text(item.Value)),
			),
		)
	}
	return dom.El("dl",
		CSS("margin", "0"),
		dom.Children(children...),
	)
}
