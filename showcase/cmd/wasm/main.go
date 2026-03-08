//go:build js && wasm

package main

import "github.com/bytewiredev/bytewire/pkg/wasm"

func main() {
	wasm.Start()
}
