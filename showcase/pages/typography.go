package pages

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
	"github.com/bytewiredev/bytewire/pkg/engine"

	ui "github.com/bytewiredev/bytewire-ui"
)

func Typography(s *engine.Session) *dom.Node {
	content := ui.VStack("2rem",
		// Headings
		ui.H3("Headings"),
		ui.H1("Heading 1"),
		ui.H2("Heading 2"),
		ui.H3("Heading 3"),

		// Paragraph
		ui.H3("Paragraph"),
		ui.Paragraph("This is a body text paragraph. It uses the default text color and comfortable line spacing for readability. Bytewire renders all UI on the server and sends efficient DOM patches to the client."),

		// Muted
		ui.H3("Muted Text"),
		ui.Muted("This is muted secondary text, useful for descriptions and helper text."),

		// Small
		ui.H3("Small Text"),
		ui.Small("This is small text, typically used for captions or footnotes."),

		// Strong
		ui.H3("Strong Text"),
		dom.P(
			dom.Children(
				dom.Text("This sentence has "),
				ui.Strong("bold emphasis"),
				dom.Text(" on certain words."),
			),
		),

		// Inline code
		ui.H3("Inline Code"),
		dom.P(
			dom.Children(
				dom.Text("Use "),
				ui.InlineCode("ui.Button()"),
				dom.Text(" to create a styled button, or "),
				ui.InlineCode("ui.Card()"),
				dom.Text(" for a card container."),
			),
		),

		// Links
		ui.H3("Links"),
		ui.VStack("0.5rem",
			ui.TextLink("Internal text link", "/buttons"),
			ui.ExternalLink("External link (opens in new tab)", "https://github.com/bytewiredev/bytewire-ui"),
		),

		// Code block
		ui.H3("Code Block"),
		ui.CodeBlock(`package main

import ui "github.com/bytewiredev/bytewire-ui"

func example() *dom.Node {
    return ui.Button("Click me", onClick)
}`),

		// Code block with title
		ui.H3("Code Block with Title"),
		ui.CodeBlockWithTitle("main.go", `srv := engine.NewServer(":4553", nil, r.Mount,
    engine.WithSSR(),
    engine.WithCSS(ui.BaseCSS),
)`),
	)

	return showcaseShell("Typography", content)
}
