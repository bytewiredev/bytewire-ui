package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
	"github.com/bytewiredev/bytewire/pkg/style"
)

// PageShell creates a full-page dark shell with navbar and main content area.
func PageShell(navbar *dom.Node, content *dom.Node) *dom.Node {
	return dom.Div(
		CSS("background-color", ColorBg, "color", ColorText, "min-height", "100vh",
			"font-family", FontSans, "width", "100%"),
		dom.Children(
			navbar,
			dom.Main(
				dom.Class(style.Classes(style.WFull)),
				dom.Children(content),
			),
		),
	)
}

// PageShellWithSidebar creates a page layout with a sidebar and main content.
func PageShellWithSidebar(navbar, sidebar, content *dom.Node) *dom.Node {
	return dom.Div(
		CSS("background-color", ColorBg, "color", ColorText, "min-height", "100vh",
			"font-family", FontSans, "width", "100%"),
		dom.Children(
			navbar,
			dom.Div(
				CSS("display", "flex"),
				dom.Children(
					sidebar,
					dom.Main(
						CSS("flex", "1", "min-width", "0"),
						dom.Children(content),
					),
				),
			),
		),
	)
}

// Hero creates a centered hero section with title, subtitle, and action buttons.
func Hero(title, subtitle string, actions ...*dom.Node) *dom.Node {
	children := []*dom.Node{
		dom.H1(
			dom.Class(style.Classes(style.FontBold)),
			CSS("font-size", "3rem", "line-height", "1.1", "color", ColorTextHeading,
				"margin-bottom", "1.5rem"),
			dom.Children(dom.Text(title)),
		),
		dom.P(
			CSS("color", ColorTextMuted, "font-size", "1.15rem",
				"max-width", "640px", "margin", "0 auto 2rem",
				"line-height", "1.6"),
			dom.Children(dom.Text(subtitle)),
		),
	}
	if len(actions) > 0 {
		children = append(children, dom.Div(
			CSS("display", "flex", "gap", "1rem", "justify-content", "center", "flex-wrap", "wrap"),
			dom.Children(actions...),
		))
	}
	return dom.Section(
		CSS("padding", "5rem 1.5rem 3rem", "text-align", "center"),
		dom.Children(children...),
	)
}

// Footer creates a simple page footer.
func Footer(text string) *dom.Node {
	return dom.El("footer",
		CSS("text-align", "center", "padding", "2rem 1rem",
			"border-top", "1px solid "+ColorBorder,
			"color", ColorTextSubtle, "font-size", "0.8rem"),
		dom.Children(dom.Text(text)),
	)
}

// BaseCSS returns the recommended base CSS to include in engine.WithCSS().
// It sets box-sizing, resets margins, and applies the dark background.
const BaseCSS = `
* { box-sizing: border-box; }
html, body {
	margin: 0;
	padding: 0;
	width: 100%;
	min-height: 100vh;
	font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
	background-color: #0f172a;
	color: #e2e8f0;
	-webkit-font-smoothing: antialiased;
}
a { text-decoration: none; color: inherit; }
pre { margin: 0; }
code { font-family: 'JetBrains Mono', 'Fira Code', 'Cascadia Code', monospace; }
table { border-spacing: 0; }
`
