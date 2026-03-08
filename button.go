package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// ButtonVariant controls the visual style of a button.
type ButtonVariant int

const (
	ButtonPrimary ButtonVariant = iota
	ButtonSecondary
	ButtonOutline
	ButtonGhost
	ButtonDanger
)

// ButtonSize controls the padding and font size.
type ButtonSize int

const (
	ButtonMd ButtonSize = iota
	ButtonSm
	ButtonLg
)

// ButtonOpts configures a button.
type ButtonOpts struct {
	Variant  ButtonVariant
	Size     ButtonSize
	Disabled bool
	FullWidth bool
}

// ButtonOpt is a functional option for Button.
type ButtonOpt func(*ButtonOpts)

// WithVariant sets the button variant.
func WithVariant(v ButtonVariant) ButtonOpt {
	return func(o *ButtonOpts) { o.Variant = v }
}

// WithSize sets the button size.
func WithSize(s ButtonSize) ButtonOpt {
	return func(o *ButtonOpts) { o.Size = s }
}

// Disabled marks the button as disabled.
func Disabled() ButtonOpt {
	return func(o *ButtonOpts) { o.Disabled = true }
}

// FullWidth makes the button take full container width.
func FullWidth() ButtonOpt {
	return func(o *ButtonOpts) { o.FullWidth = true }
}

// Button creates a styled button.
func Button(label string, onClick func([]byte), opts ...ButtonOpt) *dom.Node {
	cfg := &ButtonOpts{Variant: ButtonPrimary, Size: ButtonMd}
	for _, o := range opts {
		o(cfg)
	}

	bg, fg, border := buttonColors(cfg.Variant)
	pad, fontSize := buttonSizing(cfg.Size)

	pairs := []string{
		"background-color", bg,
		"color", fg,
		"padding", pad,
		"border-radius", RadiusMd,
		"font-weight", "600",
		"font-size", fontSize,
		"border", border,
		"cursor", "pointer",
		"transition", "opacity 0.15s",
		"line-height", "1",
	}
	if cfg.FullWidth {
		pairs = append(pairs, "width", "100%")
	}
	if cfg.Disabled {
		pairs = append(pairs, "opacity", "0.5", "pointer-events", "none")
	}

	options := []dom.Option{
		CSS(pairs...),
		dom.Children(dom.Text(label)),
	}
	if onClick != nil && !cfg.Disabled {
		options = append(options, dom.OnClick(onClick))
	}
	if cfg.Disabled {
		options = append(options, dom.Attr("disabled", "true"))
	}
	return dom.Button(options...)
}

// ButtonLink creates a button-styled anchor link.
func ButtonLink(label, href string, opts ...ButtonOpt) *dom.Node {
	cfg := &ButtonOpts{Variant: ButtonPrimary, Size: ButtonMd}
	for _, o := range opts {
		o(cfg)
	}

	bg, fg, border := buttonColors(cfg.Variant)
	pad, fontSize := buttonSizing(cfg.Size)

	pairs := []string{
		"background-color", bg,
		"color", fg,
		"padding", pad,
		"border-radius", RadiusMd,
		"font-weight", "600",
		"font-size", fontSize,
		"border", border,
		"text-decoration", "none",
		"display", "inline-block",
		"line-height", "1",
	}
	return dom.A(
		dom.Link(href),
		CSS(pairs...),
		dom.Children(dom.Text(label)),
	)
}

// ButtonGroup arranges buttons horizontally.
func ButtonGroup(buttons ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("display", "flex", "gap", "0.5rem", "flex-wrap", "wrap"),
		dom.Children(buttons...),
	)
}

func buttonColors(v ButtonVariant) (bg, fg, border string) {
	switch v {
	case ButtonPrimary:
		return ColorPrimary, ColorPrimaryText, "none"
	case ButtonSecondary:
		return ColorBgSurface, ColorText, "1px solid " + ColorBorder
	case ButtonOutline:
		return "transparent", ColorText, "1px solid " + ColorBorder
	case ButtonGhost:
		return "transparent", ColorTextMuted, "none"
	case ButtonDanger:
		return ColorDanger, "#ffffff", "none"
	default:
		return ColorPrimary, ColorPrimaryText, "none"
	}
}

func buttonSizing(s ButtonSize) (padding, fontSize string) {
	switch s {
	case ButtonSm:
		return "0.375rem 0.75rem", "0.8rem"
	case ButtonLg:
		return "0.875rem 2rem", "1.1rem"
	default:
		return "0.625rem 1.25rem", "0.875rem"
	}
}
