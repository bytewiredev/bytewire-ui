package pages

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
	"github.com/bytewiredev/bytewire/pkg/engine"

	ui "github.com/bytewiredev/bytewire-ui"
)

func Forms(s *engine.Session) *dom.Node {
	content := ui.VStack("2rem",
		// Text input
		ui.H3("Text Input"),
		ui.FormField("Username", ui.Input("Enter your username", nil), "Choose a unique username"),

		// Password input
		ui.H3("Password Input"),
		ui.FormField("Password", ui.PasswordInput("Enter your password", nil), "Must be at least 8 characters"),

		// Textarea
		ui.H3("Textarea"),
		ui.FormField("Bio", ui.Textarea("Tell us about yourself...", 4, nil), "Brief description"),

		// Select
		ui.H3("Select"),
		ui.FormField("Role", ui.Select([]ui.SelectOption{
			{Value: "admin", Label: "Admin"},
			{Value: "editor", Label: "Editor"},
			{Value: "viewer", Label: "Viewer"},
		}, nil), "Select your role"),

		// Checkbox
		ui.H3("Checkbox"),
		ui.Checkbox("I agree to the terms and conditions", nil),

		// Form group
		ui.H3("Form Group (Side by Side)"),
		ui.FormGroup(
			ui.FormField("First Name", ui.Input("First name", nil)),
			ui.FormField("Last Name", ui.Input("Last name", nil)),
		),
	)

	return showcaseShell("Forms", content)
}
