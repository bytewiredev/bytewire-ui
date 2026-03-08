package pages

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
	"github.com/bytewiredev/bytewire/pkg/engine"

	ui "github.com/bytewiredev/bytewire-ui"
)

func Navigation(s *engine.Session) *dom.Node {
	content := ui.VStack("2rem",
		// Breadcrumb
		ui.H3("Breadcrumb"),
		ui.Breadcrumb([]ui.NavItem{
			{Href: "/", Label: "Home"},
			{Href: "/navigation", Label: "Navigation"},
			{Href: "", Label: "Breadcrumb"},
		}),

		// Sidebar
		ui.H3("Sidebar"),
		ui.Muted("The sidebar component renders a vertical navigation panel. Below is an inline example."),
		ui.Spacer("0.5rem"),
		dom.Div(
			ui.CSS("border", "1px solid "+ui.ColorBorder, "border-radius", ui.RadiusMd,
				"overflow", "hidden", "display", "inline-block"),
			dom.Children(
				ui.Sidebar("Documentation", []ui.NavItem{
					{Href: "/", Label: "Getting Started"},
					{Href: "/buttons", Label: "Buttons"},
					{Href: "/forms", Label: "Forms"},
					{Href: "/data", Label: "Data Display"},
					{Href: "/feedback", Label: "Feedback"},
				}),
			),
		),

		// Navbar note
		ui.H3("Navbar"),
		ui.Muted("The Navbar component is used in the page shell at the top of every page in this showcase. It accepts a brand name, navigation links, and optional extra items."),
	)

	return showcaseShell("Navigation", content)
}
