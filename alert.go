package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// AlertVariant controls the alert color scheme.
type AlertVariant int

const (
	AlertInfo AlertVariant = iota
	AlertSuccess
	AlertWarning
	AlertDanger
)

// Alert creates a colored notification banner.
func Alert(message string, variant AlertVariant) *dom.Node {
	bg, fg, borderColor := alertColors(variant)
	return dom.Div(
		dom.Attr("role", "alert"),
		CSS("background-color", bg, "color", fg,
			"padding", "1rem 1.25rem", "border-radius", RadiusMd,
			"border-left", "4px solid "+borderColor,
			"line-height", "1.6", "font-size", "0.875rem"),
		dom.Children(dom.Text(message)),
	)
}

// AlertWithTitle creates an alert with a bold title and description.
func AlertWithTitle(title, message string, variant AlertVariant) *dom.Node {
	bg, fg, borderColor := alertColors(variant)
	return dom.Div(
		dom.Attr("role", "alert"),
		CSS("background-color", bg, "color", fg,
			"padding", "1rem 1.25rem", "border-radius", RadiusMd,
			"border-left", "4px solid "+borderColor),
		dom.Children(
			dom.Div(
				CSS("font-weight", "700", "margin-bottom", "0.25rem"),
				dom.Children(dom.Text(title)),
			),
			dom.P(
				CSS("line-height", "1.6", "font-size", "0.875rem"),
				dom.Children(dom.Text(message)),
			),
		),
	)
}

func alertColors(v AlertVariant) (bg, fg, border string) {
	switch v {
	case AlertSuccess:
		return ColorSuccessBg, ColorSuccessText, ColorSuccess
	case AlertWarning:
		return ColorWarningBg, ColorWarningText, ColorWarning
	case AlertDanger:
		return ColorDangerBg, ColorDangerText, ColorDanger
	default:
		return ColorInfoBg, ColorInfoText, ColorInfo
	}
}
