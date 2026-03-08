package ui

import (
	"testing"

	"github.com/bytewiredev/bytewire/pkg/dom"
)

func TestTabDef_StructExists(t *testing.T) {
	_ = TabDef{
		Label: "Tab 1",
		Content: func() *dom.Node {
			return dom.Text("content")
		},
	}
}
