package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// ProgressBar creates a horizontal progress bar.
// percent should be 0-100.
func ProgressBar(percent int, color ...string) *dom.Node {
	c := ColorPrimary
	if len(color) > 0 {
		c = color[0]
	}
	pct := itoa(percent)
	if percent < 0 {
		pct = "0"
	}
	if percent > 100 {
		pct = "100"
	}
	return dom.Div(
		CSS("background-color", ColorBgMuted, "border-radius", RadiusFull,
			"height", "0.5rem", "overflow", "hidden"),
		dom.Children(
			dom.Div(
				CSS("background-color", c, "height", "100%",
					"width", pct+"%", "border-radius", RadiusFull,
					"transition", "width 0.3s ease"),
			),
		),
	)
}

// ProgressBarLabeled creates a progress bar with a label above it.
func ProgressBarLabeled(label string, percent int, color ...string) *dom.Node {
	pct := itoa(percent)
	return dom.Div(
		dom.Children(
			dom.Div(
				CSS("display", "flex", "justify-content", "space-between",
					"margin-bottom", "0.375rem"),
				dom.Children(
					dom.Span(
						CSS("color", ColorText, "font-size", "0.8rem", "font-weight", "600"),
						dom.Children(dom.Text(label)),
					),
					dom.Span(
						CSS("color", ColorTextMuted, "font-size", "0.8rem"),
						dom.Children(dom.Text(pct+"%")),
					),
				),
			),
			ProgressBar(percent, color...),
		),
	)
}

// Spinner creates a simple CSS loading spinner.
func Spinner(size ...string) *dom.Node {
	s := "1.5rem"
	if len(size) > 0 {
		s = size[0]
	}
	return dom.Div(
		CSS("width", s, "height", s,
			"border", "2px solid "+ColorBgMuted,
			"border-top-color", ColorPrimary,
			"border-radius", RadiusFull,
			"animation", "spin 0.6s linear infinite"),
	)
}

// SpinnerCSS returns the @keyframes rule needed for Spinner.
// Add this to your engine.WithCSS() call.
const SpinnerCSS = `@keyframes spin { to { transform: rotate(360deg); } }`
