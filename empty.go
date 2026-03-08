package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// EmptyState creates a centered placeholder for empty views.
func EmptyState(title, description string, action ...*dom.Node) *dom.Node {
	children := []*dom.Node{
		dom.Div(
			CSS("font-size", "3rem", "margin-bottom", "1rem", "opacity", "0.3"),
			dom.Children(dom.Text("--")),
		),
		dom.H3(
			CSS("color", ColorTextHeading, "font-weight", "700",
				"font-size", "1.1rem", "margin-bottom", "0.5rem"),
			dom.Children(dom.Text(title)),
		),
		dom.P(
			CSS("color", ColorTextMuted, "font-size", "0.875rem",
				"line-height", "1.6", "max-width", "360px", "margin", "0 auto"),
			dom.Children(dom.Text(description)),
		),
	}
	if len(action) > 0 {
		children = append(children, dom.Div(
			CSS("margin-top", "1.5rem"),
			dom.Children(action...),
		))
	}
	return dom.Div(
		CSS("text-align", "center", "padding", "4rem 1rem"),
		dom.Children(children...),
	)
}

// Skeleton creates a loading placeholder bar.
func Skeleton(width, height string) *dom.Node {
	return dom.Div(
		CSS("background-color", ColorBgMuted, "border-radius", RadiusMd,
			"width", width, "height", height,
			"animation", "pulse 1.5s ease-in-out infinite"),
	)
}

// SkeletonCSS returns the @keyframes rule needed for Skeleton.
// Add this to your engine.WithCSS() call.
const SkeletonCSS = `@keyframes pulse { 0%, 100% { opacity: 1; } 50% { opacity: 0.4; } }`
