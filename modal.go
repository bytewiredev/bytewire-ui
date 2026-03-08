package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// Modal creates a signal-controlled modal dialog.
// Pass a *dom.Signal[bool] to control visibility.
func Modal(visible *dom.Signal[bool], children ...*dom.Node) *dom.Node {
	return dom.If(visible,
		func(v bool) bool { return v },
		func() *dom.Node {
			return dom.Div(
				// Backdrop
				CSS("position", "fixed", "inset", "0",
					"background-color", "rgba(0,0,0,0.6)",
					"display", "flex", "align-items", "center", "justify-content", "center",
					"z-index", "50", "padding", "1rem"),
				dom.OnClick(func(_ []byte) { visible.Set(false) }),
				dom.Children(
					dom.Div(
						// Dialog panel
						CSS("background-color", ColorBgSurface,
							"border-radius", RadiusLg,
							"border", "1px solid "+ColorBorder,
							"max-width", "560px", "width", "100%",
							"box-shadow", ShadowLg),
						// Stop click propagation by having a separate click handler
						dom.OnClick(func(_ []byte) {}),
						dom.Children(children...),
					),
				),
			)
		},
		nil,
	)
}

// ModalHeader creates a modal header with title and close button.
func ModalHeader(title string, onClose func([]byte)) *dom.Node {
	return dom.Div(
		CSS("display", "flex", "align-items", "center", "justify-content", "space-between",
			"padding", "1.25rem 1.5rem",
			"border-bottom", "1px solid "+ColorBorder),
		dom.Children(
			dom.H3(
				CSS("color", ColorTextHeading, "font-weight", "700", "font-size", "1.1rem"),
				dom.Children(dom.Text(title)),
			),
			dom.Button(
				CSS("background", "none", "border", "none", "color", ColorTextMuted,
					"cursor", "pointer", "font-size", "1.25rem", "padding", "0.25rem"),
				dom.OnClick(onClose),
				dom.Children(dom.Text("x")),
			),
		),
	)
}

// ModalBody creates the modal content area.
func ModalBody(children ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("padding", "1.5rem"),
		dom.Children(children...),
	)
}

// ModalFooter creates a modal footer, typically for action buttons.
func ModalFooter(children ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("display", "flex", "justify-content", "flex-end", "gap", "0.5rem",
			"padding", "1rem 1.5rem",
			"border-top", "1px solid "+ColorBorder,
			"background-color", "rgba(0,0,0,0.1)",
			"border-radius", "0 0 "+RadiusLg+" "+RadiusLg),
		dom.Children(children...),
	)
}
