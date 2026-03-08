package ui

import (
	"strings"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

// CSS builds a single style attribute from property:value pairs.
// This survives SSR hydration (unlike dom.Style which produces OpSetStyle opcodes
// that are lost when the WASM client reconnects).
//
// Usage: CSS("display", "grid", "gap", "1rem")
func CSS(pairs ...string) dom.Option {
	var b strings.Builder
	for i := 0; i+1 < len(pairs); i += 2 {
		if b.Len() > 0 {
			b.WriteByte(';')
		}
		b.WriteString(pairs[i])
		b.WriteByte(':')
		b.WriteString(pairs[i+1])
	}
	return dom.Attr("style", b.String())
}

// mergeCSS combines multiple CSS pair slices into one.
func mergeCSS(base []string, overrides ...string) []string {
	return append(base, overrides...)
}
