package pages

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
	"github.com/bytewiredev/bytewire/pkg/engine"

	ui "github.com/bytewiredev/bytewire-ui"
)

var navLinks = []ui.NavItem{
	{Href: "/buttons", Label: "Buttons"},
	{Href: "/forms", Label: "Forms"},
	{Href: "/data", Label: "Data"},
	{Href: "/feedback", Label: "Feedback"},
	{Href: "/navigation", Label: "Navigation"},
	{Href: "/typography", Label: "Typography"},
	{Href: "/layout", Label: "Layout"},
}

func showcaseShell(title string, content *dom.Node) *dom.Node {
	navbar := ui.Navbar("bytewire-ui", navLinks)
	body := ui.Container(
		ui.H1(title),
		ui.Spacer("1rem"),
		content,
		ui.Spacer("3rem"),
	)
	return ui.PageShell(navbar, body)
}

func Home(s *engine.Session) *dom.Node {
	navbar := ui.Navbar("bytewire-ui", navLinks)

	hero := ui.Hero(
		"bytewire-ui",
		"A server-driven component library for Bytewire. Dark-themed, composable, and zero-JS on the client.",
		ui.ButtonLink("Get Started", "/buttons"),
		ui.ButtonLink("GitHub", "https://github.com/bytewiredev/bytewire-ui", ui.WithVariant(ui.ButtonSecondary)),
	)

	grid := ui.Grid("250px", "1.5rem",
		featureLink("Buttons", "Primary, secondary, outline, ghost, danger, sizes, groups, and link buttons.", "/buttons"),
		featureLink("Forms", "Inputs, password fields, textareas, selects, checkboxes, form fields, and groups.", "/forms"),
		featureLink("Data Display", "Tables, key-value tables, cards, stat cards, lists, and description lists.", "/data"),
		featureLink("Feedback", "Alerts, badges, tags, progress bars, spinners, empty states, and skeletons.", "/feedback"),
		featureLink("Navigation", "Navbar, breadcrumbs, and sidebar navigation components.", "/navigation"),
		featureLink("Typography", "Headings, paragraphs, muted text, inline code, links, and code blocks.", "/typography"),
		featureLink("Layout", "Containers, stacks, grids, columns, spacers, dividers, and sections.", "/layout"),
		featureLink("Interactive", "Modals, tabs, and accordions with signal-driven state.", "/feedback"),
	)

	content := dom.Div(
		dom.Children(
			hero,
			ui.Container(grid, ui.Spacer("3rem")),
		),
	)

	return ui.PageShell(navbar, content)
}

func featureLink(title, desc, href string) *dom.Node {
	return dom.A(
		dom.Link(href),
		ui.CSS("text-decoration", "none"),
		dom.Children(
			ui.FeatureCard("", title, desc),
		),
	)
}
