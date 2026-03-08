package ui

import (
	"github.com/bytewiredev/bytewire/pkg/dom"
)

// TableColumn defines a column header.
type TableColumn struct {
	Label string
	Align string // "left" (default), "center", "right"
}

// Table creates a styled data table from columns and row data.
// Each row is a slice of strings matching the column order.
func Table(columns []TableColumn, rows [][]string) *dom.Node {
	// Header
	var headerCells []*dom.Node
	for _, col := range columns {
		align := col.Align
		if align == "" {
			align = "left"
		}
		headerCells = append(headerCells, dom.Th(
			CSS("padding", "0.75rem 1rem", "text-align", align,
				"font-weight", "600", "color", ColorTextMuted,
				"font-size", "0.75rem", "text-transform", "uppercase",
				"letter-spacing", "0.05em"),
			dom.Children(dom.Text(col.Label)),
		))
	}

	// Body rows
	var bodyRows []*dom.Node
	for i, row := range rows {
		var cells []*dom.Node
		for j, cell := range row {
			align := "left"
			if j < len(columns) && columns[j].Align != "" {
				align = columns[j].Align
			}
			cells = append(cells, dom.Td(
				CSS("padding", "0.75rem 1rem", "color", ColorText,
					"border-top", "1px solid "+ColorBorder, "text-align", align),
				dom.Children(dom.Text(cell)),
			))
		}
		bg := "transparent"
		if i%2 == 1 {
			bg = "rgba(255,255,255,0.02)"
		}
		bodyRows = append(bodyRows, dom.Tr(
			CSS("background-color", bg),
			dom.Children(cells...),
		))
	}

	return dom.Div(
		CSS("overflow-x", "auto", "border-radius", RadiusLg,
			"border", "1px solid "+ColorBorder),
		dom.Children(
			dom.Table(
				CSS("width", "100%", "border-collapse", "collapse", "font-size", "0.875rem"),
				dom.Children(
					dom.Thead(
						CSS("background-color", ColorBgSurface),
						dom.Children(dom.Tr(dom.Children(headerCells...))),
					),
					dom.Tbody(dom.Children(bodyRows...)),
				),
			),
		),
	)
}

// KeyValueTable creates a two-column table from key-value pairs.
func KeyValueTable(pairs []KV) *dom.Node {
	var rows []*dom.Node
	for _, kv := range pairs {
		rows = append(rows, dom.Tr(
			dom.Children(
				dom.Td(
					CSS("padding", "0.5rem 1rem", "border-top", "1px solid "+ColorBorder,
						"color", ColorPrimary, "font-family", FontMono,
						"font-size", "0.8rem", "white-space", "nowrap"),
					dom.Children(dom.Text(kv.Key)),
				),
				dom.Td(
					CSS("padding", "0.5rem 1rem", "border-top", "1px solid "+ColorBorder,
						"color", ColorTextMuted),
					dom.Children(dom.Text(kv.Value)),
				),
			),
		))
	}

	return dom.Div(
		CSS("overflow-x", "auto", "border-radius", RadiusMd,
			"border", "1px solid "+ColorBorder),
		dom.Children(
			dom.Table(
				CSS("width", "100%", "border-collapse", "collapse", "font-size", "0.875rem"),
				dom.Children(
					dom.Thead(
						CSS("background-color", ColorBgSurface),
						dom.Children(dom.Tr(dom.Children(
							tableHeaderCell("Key"),
							tableHeaderCell("Value"),
						))),
					),
					dom.Tbody(dom.Children(rows...)),
				),
			),
		),
	)
}

// KV is a key-value pair for KeyValueTable.
type KV struct {
	Key   string
	Value string
}

func tableHeaderCell(text string) *dom.Node {
	return dom.Th(
		CSS("padding", "0.5rem 1rem", "text-align", "left", "font-weight", "600",
			"color", ColorTextMuted, "font-size", "0.75rem", "text-transform", "uppercase"),
		dom.Children(dom.Text(text)),
	)
}
