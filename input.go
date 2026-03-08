package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// Input creates a styled text input.
func Input(placeholder string, onInput func([]byte), opts ...dom.Option) *dom.Node {
	options := []dom.Option{
		dom.Attr("type", "text"),
		dom.Attr("placeholder", placeholder),
		CSS("background-color", ColorBgInput, "color", ColorText,
			"border", "1px solid "+ColorBorder, "border-radius", RadiusMd,
			"padding", "0.625rem 0.875rem", "font-size", "0.875rem",
			"width", "100%", "outline", "none",
			"font-family", "inherit"),
	}
	if onInput != nil {
		options = append(options, dom.OnInput(onInput))
	}
	options = append(options, opts...)
	return dom.Input(options...)
}

// PasswordInput creates a password input field.
func PasswordInput(placeholder string, onInput func([]byte), opts ...dom.Option) *dom.Node {
	options := []dom.Option{
		dom.Attr("type", "password"),
		dom.Attr("placeholder", placeholder),
		CSS("background-color", ColorBgInput, "color", ColorText,
			"border", "1px solid "+ColorBorder, "border-radius", RadiusMd,
			"padding", "0.625rem 0.875rem", "font-size", "0.875rem",
			"width", "100%", "outline", "none",
			"font-family", "inherit"),
	}
	if onInput != nil {
		options = append(options, dom.OnInput(onInput))
	}
	options = append(options, opts...)
	return dom.Input(options...)
}

// Textarea creates a styled multi-line text input.
func Textarea(placeholder string, rows int, onInput func([]byte), opts ...dom.Option) *dom.Node {
	options := []dom.Option{
		dom.Attr("placeholder", placeholder),
		dom.Attr("rows", itoa(rows)),
		CSS("background-color", ColorBgInput, "color", ColorText,
			"border", "1px solid "+ColorBorder, "border-radius", RadiusMd,
			"padding", "0.625rem 0.875rem", "font-size", "0.875rem",
			"width", "100%", "outline", "none", "resize", "vertical",
			"font-family", "inherit", "line-height", "1.6"),
	}
	if onInput != nil {
		options = append(options, dom.OnInput(onInput))
	}
	options = append(options, opts...)
	return dom.Textarea(options...)
}

// SelectOption defines an option for a select dropdown.
type SelectOption struct {
	Value string
	Label string
}

// Select creates a styled dropdown select.
func Select(options []SelectOption, onInput func([]byte), opts ...dom.Option) *dom.Node {
	var items []*dom.Node
	for _, o := range options {
		items = append(items, dom.OptionEl(
			dom.Attr("value", o.Value),
			dom.Children(dom.Text(o.Label)),
		))
	}

	selectOpts := []dom.Option{
		CSS("background-color", ColorBgInput, "color", ColorText,
			"border", "1px solid "+ColorBorder, "border-radius", RadiusMd,
			"padding", "0.625rem 0.875rem", "font-size", "0.875rem",
			"width", "100%", "outline", "none",
			"font-family", "inherit", "cursor", "pointer"),
		dom.Children(items...),
	}
	if onInput != nil {
		selectOpts = append(selectOpts, dom.OnInput(onInput))
	}
	selectOpts = append(selectOpts, opts...)
	return dom.SelectEl(selectOpts...)
}

// Checkbox creates a styled checkbox with label.
func Checkbox(label string, onClick func([]byte)) *dom.Node {
	return dom.Label(
		CSS("display", "flex", "align-items", "center", "gap", "0.5rem",
			"cursor", "pointer", "color", ColorText, "font-size", "0.875rem"),
		dom.Children(
			dom.Input(
				dom.Attr("type", "checkbox"),
				CSS("width", "1rem", "height", "1rem", "cursor", "pointer",
					"accent-color", ColorPrimary),
				dom.OnClick(onClick),
			),
			dom.Span(dom.Children(dom.Text(label))),
		),
	)
}

// FormField wraps an input with a label and optional help text.
func FormField(label string, input *dom.Node, help ...string) *dom.Node {
	children := []*dom.Node{
		dom.Label(
			CSS("color", ColorTextMuted, "font-size", "0.8rem", "font-weight", "600",
				"display", "block", "margin-bottom", "0.375rem"),
			dom.Children(dom.Text(label)),
		),
		input,
	}
	if len(help) > 0 && help[0] != "" {
		children = append(children, dom.Div(
			CSS("color", ColorTextSubtle, "font-size", "0.75rem", "margin-top", "0.375rem"),
			dom.Children(dom.Text(help[0])),
		))
	}
	return dom.Div(
		CSS("margin-bottom", "1.25rem"),
		dom.Children(children...),
	)
}

// FormGroup arranges form fields in a horizontal row.
func FormGroup(children ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("display", "grid", "grid-template-columns", "repeat(auto-fill, minmax(200px, 1fr))",
			"gap", "1rem"),
		dom.Children(children...),
	)
}
