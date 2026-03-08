package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
	"github.com/bytewiredev/bytewire/pkg/style"
)

// Container centers content with a max width and horizontal padding.
func Container(children ...*dom.Node) *dom.Node {
	return dom.Div(
		dom.Class(style.Classes(style.MxAuto, style.Px4)),
		CSS("max-width", "1100px"),
		dom.Children(children...),
	)
}

// ContainerNarrow is a narrower container for text-heavy content.
func ContainerNarrow(children ...*dom.Node) *dom.Node {
	return dom.Div(
		dom.Class(style.Classes(style.MxAuto, style.Px4)),
		CSS("max-width", "720px"),
		dom.Children(children...),
	)
}

// ContainerWide is a wider container for dashboards and tables.
func ContainerWide(children ...*dom.Node) *dom.Node {
	return dom.Div(
		dom.Class(style.Classes(style.MxAuto, style.Px4)),
		CSS("max-width", "1400px"),
		dom.Children(children...),
	)
}

// Section wraps content in a padded section with an optional title.
func Section(title string, children ...*dom.Node) *dom.Node {
	var inner []*dom.Node
	if title != "" {
		inner = append(inner, SectionTitle(title))
	}
	inner = append(inner, children...)
	return dom.Section(
		dom.Children(
			Container(inner...),
		),
	)
}

// SectionAlt creates a section with a contrasting background.
func SectionAlt(title string, children ...*dom.Node) *dom.Node {
	var inner []*dom.Node
	if title != "" {
		inner = append(inner, SectionTitle(title))
	}
	inner = append(inner, children...)
	return dom.Section(
		CSS("background-color", ColorBgSurface,
			"border-top", "1px solid "+ColorBorder,
			"border-bottom", "1px solid "+ColorBorder),
		dom.Children(
			Container(inner...),
		),
	)
}

// SectionTitle creates a styled section heading.
func SectionTitle(text string) *dom.Node {
	return dom.H2(
		dom.Class(style.Classes(style.Text2Xl, style.FontBold, style.Mb4)),
		CSS("color", ColorTextHeading),
		dom.Children(dom.Text(text)),
	)
}

// VStack arranges children vertically with a gap.
func VStack(gap string, children ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("display", "flex", "flex-direction", "column", "gap", gap),
		dom.Children(children...),
	)
}

// HStack arranges children horizontally with a gap.
func HStack(gap string, children ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("display", "flex", "flex-direction", "row", "align-items", "center", "gap", gap),
		dom.Children(children...),
	)
}

// Center horizontally and vertically centers its children.
func Center(children ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("display", "flex", "align-items", "center", "justify-content", "center"),
		dom.Children(children...),
	)
}

// Grid creates a responsive grid. minWidth is the minimum column width (e.g. "280px").
func Grid(minWidth string, gap string, children ...*dom.Node) *dom.Node {
	return dom.Div(
		CSS("display", "grid",
			"grid-template-columns", "repeat(auto-fill, minmax("+minWidth+", 1fr))",
			"gap", gap),
		dom.Children(children...),
	)
}

// Columns creates a fixed-column grid.
func Columns(cols int, gap string, children ...*dom.Node) *dom.Node {
	colTemplate := "repeat(" + itoa(cols) + ", 1fr)"
	return dom.Div(
		CSS("display", "grid", "grid-template-columns", colTemplate, "gap", gap),
		dom.Children(children...),
	)
}

// Spacer adds vertical space.
func Spacer(height string) *dom.Node {
	return dom.Div(CSS("height", height))
}

// Divider creates a horizontal line.
func Divider() *dom.Node {
	return dom.El("hr",
		CSS("border", "none", "border-top", "1px solid "+ColorBorder, "margin", "1.5rem 0"),
	)
}

// itoa converts small ints to string without importing strconv.
func itoa(n int) string {
	if n < 0 {
		return "-" + itoa(-n)
	}
	if n < 10 {
		return string(rune('0' + n))
	}
	return itoa(n/10) + string(rune('0'+n%10))
}
