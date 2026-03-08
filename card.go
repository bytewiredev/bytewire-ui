package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// Card creates a bordered surface container.
func Card(children ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("background-color", ColorBgSurface, "border-radius", RadiusLg,
			"border", "1px solid "+ColorBorder),
		dom.Children(children...),
	)
}

// CardAccent creates a card with a colored top border accent.
func CardAccent(accentColor string, children ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("background-color", ColorBgSurface, "border-radius", RadiusLg,
			"border", "1px solid "+ColorBorder,
			"border-top", "3px solid "+accentColor),
		dom.Children(children...),
	)
}

// CardHeader creates a card header area with bottom border.
func CardHeader(children ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("padding", "1.25rem 1.5rem",
			"border-bottom", "1px solid "+ColorBorder),
		dom.Children(children...),
	)
}

// CardBody creates the main card content area.
func CardBody(children ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("padding", "1.5rem"),
		dom.Children(children...),
	)
}

// CardFooter creates a card footer area with top border.
func CardFooter(children ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("padding", "1rem 1.5rem",
			"border-top", "1px solid "+ColorBorder,
			"background-color", "rgba(0,0,0,0.1)",
			"border-radius", "0 0 "+RadiusLg+" "+RadiusLg),
		dom.Children(children...),
	)
}

// SimpleCard creates a quick card with title and description.
func SimpleCard(title, desc string) *dom.Node {
	return Card(
		CardBody(
			dom.H3(
				CSS("color", ColorTextHeading, "font-weight", "700",
					"font-size", "1rem", "margin-bottom", "0.5rem"),
				dom.Children(dom.Text(title)),
			),
			dom.P(
				CSS("color", ColorTextMuted, "line-height", "1.6", "font-size", "0.875rem"),
				dom.Children(dom.Text(desc)),
			),
		),
	)
}

// FeatureCard creates a card with an icon character, title, and description.
func FeatureCard(icon, title, desc string) *dom.Node {
	return Card(
		CardBody(
			dom.Div(
				CSS("font-size", "2rem", "margin-bottom", "0.75rem"),
				dom.Children(dom.Text(icon)),
			),
			dom.H3(
				CSS("color", ColorTextHeading, "font-weight", "700",
					"font-size", "1.1rem", "margin-bottom", "0.5rem"),
				dom.Children(dom.Text(title)),
			),
			dom.P(
				CSS("color", ColorTextMuted, "line-height", "1.6", "font-size", "0.875rem"),
				dom.Children(dom.Text(desc)),
			),
		),
	)
}

// StatCard creates a large stat display card (value + label).
func StatCard(value, label, color string) *dom.Node {
	return dom.Div(
		CSS("background-color", ColorBgSurface, "border-radius", RadiusLg,
			"padding", "1.5rem", "text-align", "center",
			"border", "1px solid "+ColorBorder),
		dom.Children(
			dom.Div(
				CSS("font-size", "2rem", "font-weight", "800", "color", color,
					"font-family", FontMono, "margin-bottom", "0.25rem"),
				dom.Children(dom.Text(value)),
			),
			dom.Div(
				CSS("color", ColorTextMuted, "font-size", "0.85rem"),
				dom.Children(dom.Text(label)),
			),
		),
	)
}
