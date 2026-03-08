package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// AvatarSize controls the avatar dimensions.
type AvatarSize int

const (
	AvatarMd AvatarSize = iota
	AvatarSm
	AvatarLg
	AvatarXl
)

// Avatar creates a circular avatar with initials.
func Avatar(initials string, size ...AvatarSize) *dom.Node {
	s := AvatarMd
	if len(size) > 0 {
		s = size[0]
	}
	dim, fontSize := avatarSizing(s)
	return dom.Div(
		CSS("width", dim, "height", dim, "border-radius", RadiusFull,
			"background-color", ColorPrimary, "color", ColorPrimaryText,
			"display", "flex", "align-items", "center", "justify-content", "center",
			"font-weight", "700", "font-size", fontSize, "flex-shrink", "0"),
		dom.Children(dom.Text(initials)),
	)
}

// AvatarImage creates a circular avatar from an image URL.
func AvatarImage(src, alt string, size ...AvatarSize) *dom.Node {
	s := AvatarMd
	if len(size) > 0 {
		s = size[0]
	}
	dim, _ := avatarSizing(s)
	return dom.Img(
		dom.Attr("src", src),
		dom.Attr("alt", alt),
		CSS("width", dim, "height", dim, "border-radius", RadiusFull,
			"object-fit", "cover", "flex-shrink", "0"),
	)
}

func avatarSizing(s AvatarSize) (dim, fontSize string) {
	switch s {
	case AvatarSm:
		return "2rem", "0.7rem"
	case AvatarLg:
		return "3.5rem", "1.25rem"
	case AvatarXl:
		return "5rem", "1.75rem"
	default:
		return "2.5rem", "0.9rem"
	}
}
