package pages

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
	"github.com/bytewiredev/bytewire/pkg/engine"

	ui "github.com/bytewiredev/bytewire-ui"
)

func Data(s *engine.Session) *dom.Node {
	content := ui.VStack("2rem",
		// Table
		ui.H3("Table"),
		ui.Table(
			[]ui.TableColumn{
				{Label: "Name"},
				{Label: "Role"},
				{Label: "Status", Align: "center"},
			},
			[][]string{
				{"Alice Johnson", "Admin", "Active"},
				{"Bob Smith", "Editor", "Active"},
				{"Carol Williams", "Viewer", "Inactive"},
				{"Dave Brown", "Editor", "Active"},
			},
		),

		// Key-value table
		ui.H3("Key-Value Table"),
		ui.KeyValueTable([]ui.KV{
			{Key: "Framework", Value: "Bytewire"},
			{Key: "Language", Value: "Go"},
			{Key: "Rendering", Value: "Server-driven"},
			{Key: "Transport", Value: "WebTransport / WebSocket"},
			{Key: "License", Value: "MIT"},
		}),

		// Simple card
		ui.H3("Simple Card"),
		ui.SimpleCard("Simple Card", "A minimal card with a title and description. Useful for quick content blocks."),

		// Feature card
		ui.H3("Feature Card"),
		ui.FeatureCard("~", "Server-Driven UI", "All rendering happens on the server. The client is a thin WASM binary that applies DOM patches."),

		// Stat cards
		ui.H3("Stat Cards"),
		ui.Columns(3, "1rem",
			ui.StatCard("1.2ms", "Avg Latency", ui.ColorSuccess),
			ui.StatCard("42", "Components", ui.ColorPrimary),
			ui.StatCard("99.9%", "Uptime", ui.ColorInfo),
		),

		// List
		ui.H3("List"),
		ui.List(
			ui.ListItem("Getting Started", "Install the library and run your first app"),
			ui.ListItem("Components", "Explore all available UI components"),
			ui.ListItem("Theming", "Customize colors, fonts, and spacing"),
			ui.ListItem("Deployment", "Deploy your Bytewire app to production"),
		),

		// Description list
		ui.H3("Description List"),
		ui.DescriptionList([]ui.KV{
			{Key: "Version", Value: "0.1.0"},
			{Key: "Author", Value: "Bytewire Team"},
			{Key: "Repository", Value: "github.com/bytewiredev/bytewire-ui"},
		}),
	)

	return showcaseShell("Data Display", content)
}
