package pages

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
	"github.com/bytewiredev/bytewire/pkg/engine"

	ui "github.com/bytewiredev/bytewire-ui"
)

func demoBox(label string) *dom.Node {
	return dom.Div(
		ui.CSS("background-color", ui.ColorBgSurface, "border", "1px solid "+ui.ColorBorder,
			"padding", "1rem", "border-radius", ui.RadiusMd, "text-align", "center",
			"color", ui.ColorTextMuted, "font-size", "0.85rem"),
		dom.Children(dom.Text(label)),
	)
}

func Layout(s *engine.Session) *dom.Node {
	content := ui.VStack("2rem",
		// Container descriptions
		ui.H3("Containers"),
		ui.Muted("Container (max 1100px), ContainerNarrow (max 720px), and ContainerWide (max 1400px) center content horizontally with padding. This page itself is wrapped in a Container."),

		// VStack
		ui.H3("VStack (Vertical Stack)"),
		ui.VStack("0.75rem",
			demoBox("Item 1"),
			demoBox("Item 2"),
			demoBox("Item 3"),
		),

		// HStack
		ui.H3("HStack (Horizontal Stack)"),
		ui.HStack("0.75rem",
			demoBox("Left"),
			demoBox("Center"),
			demoBox("Right"),
		),

		// Grid
		ui.H3("Grid (auto-fill, min 200px)"),
		ui.Grid("200px", "0.75rem",
			demoBox("Cell 1"),
			demoBox("Cell 2"),
			demoBox("Cell 3"),
			demoBox("Cell 4"),
			demoBox("Cell 5"),
			demoBox("Cell 6"),
		),

		// Columns
		ui.H3("Columns (3 fixed)"),
		ui.Columns(3, "0.75rem",
			demoBox("Col 1"),
			demoBox("Col 2"),
			demoBox("Col 3"),
		),

		// Spacer
		ui.H3("Spacer"),
		ui.Muted("A spacer adds vertical space between elements. There is a 3rem spacer below this text."),
		ui.Spacer("3rem"),
		ui.Muted("Content after the spacer."),

		// Divider
		ui.H3("Divider"),
		ui.Muted("Content above the divider."),
		ui.Divider(),
		ui.Muted("Content below the divider."),

		// Section
		ui.H3("Section"),
		ui.Section("Section Title",
			ui.Paragraph("This is a Section component. It wraps content in a Container with an optional title."),
		),

		// SectionAlt
		ui.H3("Section Alt"),
		ui.SectionAlt("Alternate Section",
			ui.Paragraph("SectionAlt uses a contrasting surface background with top and bottom borders."),
		),
	)

	return showcaseShell("Layout", content)
}
