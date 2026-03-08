package pages

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
	"github.com/bytewiredev/bytewire/pkg/engine"

	ui "github.com/bytewiredev/bytewire-ui"
)

func Buttons(s *engine.Session) *dom.Node {
	noop := func(_ []byte) {}

	content := ui.VStack("2rem",
		// Variants
		ui.H3("Variants"),
		ui.ButtonGroup(
			ui.Button("Primary", noop, ui.WithVariant(ui.ButtonPrimary)),
			ui.Button("Secondary", noop, ui.WithVariant(ui.ButtonSecondary)),
			ui.Button("Outline", noop, ui.WithVariant(ui.ButtonOutline)),
			ui.Button("Ghost", noop, ui.WithVariant(ui.ButtonGhost)),
			ui.Button("Danger", noop, ui.WithVariant(ui.ButtonDanger)),
		),

		// Sizes
		ui.H3("Sizes"),
		ui.ButtonGroup(
			ui.Button("Small", noop, ui.WithSize(ui.ButtonSm)),
			ui.Button("Medium", noop, ui.WithSize(ui.ButtonMd)),
			ui.Button("Large", noop, ui.WithSize(ui.ButtonLg)),
		),

		// Disabled
		ui.H3("Disabled"),
		ui.ButtonGroup(
			ui.Button("Disabled Primary", nil, ui.Disabled()),
			ui.Button("Disabled Outline", nil, ui.WithVariant(ui.ButtonOutline), ui.Disabled()),
		),

		// Full width
		ui.H3("Full Width"),
		ui.Button("Full Width Button", noop, ui.FullWidth()),

		// Button links
		ui.H3("Button Links"),
		ui.ButtonGroup(
			ui.ButtonLink("Primary Link", "/"),
			ui.ButtonLink("Secondary Link", "/", ui.WithVariant(ui.ButtonSecondary)),
			ui.ButtonLink("Small Link", "/", ui.WithSize(ui.ButtonSm)),
		),
	)

	return showcaseShell("Buttons", content)
}
