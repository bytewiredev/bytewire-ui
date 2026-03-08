package pages

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
	"github.com/bytewiredev/bytewire/pkg/engine"

	ui "github.com/bytewiredev/bytewire-ui"
)

func Feedback(s *engine.Session) *dom.Node {
	content := ui.VStack("2rem",
		// Alerts
		ui.H3("Alerts"),
		ui.VStack("0.75rem",
			ui.Alert("This is an informational message.", ui.AlertInfo),
			ui.Alert("Operation completed successfully.", ui.AlertSuccess),
			ui.Alert("Please review the following warnings.", ui.AlertWarning),
			ui.Alert("An error occurred while processing.", ui.AlertDanger),
		),

		// Alert with title
		ui.H3("Alert with Title"),
		ui.AlertWithTitle("Deployment Complete", "Your application has been deployed to production. All health checks are passing.", ui.AlertSuccess),

		// Badges
		ui.H3("Badges"),
		ui.HStack("0.5rem",
			ui.Badge("Default"),
			ui.Badge("Info", ui.BadgeInfo),
			ui.Badge("Success", ui.BadgeSuccess),
			ui.Badge("Warning", ui.BadgeWarning),
			ui.Badge("Danger", ui.BadgeDanger),
		),

		// Tags
		ui.H3("Tags"),
		ui.HStack("0.5rem",
			ui.Tag("Default"),
			ui.Tag("Info", ui.BadgeInfo),
			ui.Tag("Success", ui.BadgeSuccess),
			ui.Tag("Warning", ui.BadgeWarning),
			ui.Tag("Danger", ui.BadgeDanger),
		),

		// Progress bars
		ui.H3("Progress Bars"),
		ui.VStack("1rem",
			ui.ProgressBar(25),
			ui.ProgressBar(50, ui.ColorInfo),
			ui.ProgressBar(75, ui.ColorWarning),
			ui.ProgressBar(100, ui.ColorSuccess),
		),

		// Labeled progress bar
		ui.H3("Labeled Progress Bar"),
		ui.ProgressBarLabeled("Upload Progress", 67, ui.ColorPrimary),

		// Spinner
		ui.H3("Spinner"),
		ui.HStack("1.5rem",
			ui.Spinner(),
			ui.Spinner("2.5rem"),
		),

		// Empty state
		ui.H3("Empty State"),
		ui.EmptyState(
			"No results found",
			"Try adjusting your search or filters to find what you are looking for.",
			ui.Button("Clear Filters", func(_ []byte) {}),
		),

		// Skeletons
		ui.H3("Skeletons"),
		ui.VStack("0.75rem",
			ui.Skeleton("100%", "1rem"),
			ui.Skeleton("75%", "1rem"),
			ui.Skeleton("50%", "1rem"),
		),
	)

	return showcaseShell("Feedback", content)
}
