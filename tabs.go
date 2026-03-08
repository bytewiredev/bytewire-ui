package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// TabDef defines a tab with label and content builder.
type TabDef struct {
	Label   string
	Content func() *dom.Node
}

// Tabs creates a tabbed interface controlled by a signal.
// The signal holds the active tab index.
func Tabs(active *dom.Signal[int], tabs []TabDef) *dom.Node {
	// Tab buttons
	var buttons []*dom.Node
	for i, tab := range tabs {
		idx := i
		label := tab.Label
		buttons = append(buttons, tabButton(active, idx, label))
	}

	// Tab content
	content := dom.If(active,
		func(v int) bool { return v >= 0 },
		func() *dom.Node {
			var panels []*dom.Node
			for i, tab := range tabs {
				idx := i
				builder := tab.Content
				panels = append(panels,
					dom.If(active,
						func(v int) bool { return v == idx },
						func() *dom.Node { return builder() },
						nil,
					),
				)
			}
			return dom.Div(dom.Children(panels...))
		},
		nil,
	)

	return dom.Div(
		dom.Children(
			dom.Div(
				CSS("display", "flex", "gap", "0",
					"border-bottom", "2px solid "+ColorBorder,
					"margin-bottom", "1.5rem"),
				dom.Children(buttons...),
			),
			content,
		),
	)
}

func tabButton(active *dom.Signal[int], idx int, label string) *dom.Node {
	return dom.Button(
		CSS("background", "none", "border", "none",
			"padding", "0.75rem 1.25rem",
			"font-size", "0.875rem", "cursor", "pointer",
			"color", ColorTextMuted, "font-weight", "500",
			"border-bottom", "2px solid transparent",
			"margin-bottom", "-2px", "transition", "color 0.15s"),
		dom.ClassF(active, func(v int) string {
			if v == idx {
				return "bw-tab-active"
			}
			return ""
		}),
		dom.StyleF(active, "color", func(v int) string {
			if v == idx {
				return ColorTextHeading
			}
			return ColorTextMuted
		}),
		dom.StyleF(active, "border-bottom-color", func(v int) string {
			if v == idx {
				return ColorPrimary
			}
			return "transparent"
		}),
		dom.OnClick(func(_ []byte) { active.Set(idx) }),
		dom.Children(dom.Text(label)),
	)
}
