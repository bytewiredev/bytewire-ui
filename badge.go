package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// BadgeVariant controls the badge color scheme.
type BadgeVariant int

const (
	BadgeDefault BadgeVariant = iota
	BadgeInfo
	BadgeSuccess
	BadgeWarning
	BadgeDanger
)

// Badge creates a small colored pill label.
func Badge(text string, variant ...BadgeVariant) *dom.Node {
	v := BadgeDefault
	if len(variant) > 0 {
		v = variant[0]
	}
	bg, fg := badgeColors(v)
	return dom.Span(
		CSS("background-color", bg, "color", fg,
			"padding", "0.2rem 0.6rem", "border-radius", RadiusFull,
			"font-size", "0.7rem", "font-weight", "600",
			"display", "inline-block", "line-height", "1.4"),
		dom.Children(dom.Text(text)),
	)
}

// Tag creates a slightly larger tag with optional border.
func Tag(text string, variant ...BadgeVariant) *dom.Node {
	v := BadgeDefault
	if len(variant) > 0 {
		v = variant[0]
	}
	bg, fg := badgeColors(v)
	return dom.Span(
		CSS("background-color", bg, "color", fg,
			"padding", "0.3rem 0.75rem", "border-radius", RadiusMd,
			"font-size", "0.8rem", "font-weight", "500",
			"display", "inline-block"),
		dom.Children(dom.Text(text)),
	)
}

func badgeColors(v BadgeVariant) (bg, fg string) {
	switch v {
	case BadgeInfo:
		return ColorInfoBg, ColorInfoText
	case BadgeSuccess:
		return ColorSuccessBg, ColorSuccessText
	case BadgeWarning:
		return ColorWarningBg, ColorWarningText
	case BadgeDanger:
		return ColorDangerBg, ColorDangerText
	default:
		return ColorBgMuted, ColorText
	}
}
