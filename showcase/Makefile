GOROOT_WASM := $(shell go env GOROOT)/lib/wasm

.DEFAULT_GOAL := help

.PHONY: help build wasm run dev stop clean

help: ## Show available commands
	@echo "bytewire-ui showcase commands:"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  make %-12s %s\n", $$1, $$2}'
	@echo ""

build: ## Build the server binary
	go build -o server .

wasm: ## Compile the WASM client
	GOOS=js GOARCH=wasm go build -o static/bytewire.wasm ./cmd/wasm/
	@cp "$(GOROOT_WASM)/wasm_exec.js" static/wasm_exec.js 2>/dev/null || true

dev: stop wasm build run ## Full rebuild and run (stops existing first)

run: ## Run the server (http://localhost:9091)
	@./server &
	@echo "showcase running at http://localhost:9091"
	@echo "Run 'make stop' to stop"

stop: ## Stop the running server
	@lsof -ti :9091 | xargs kill 2>/dev/null || true
	@lsof -ti :4553 | xargs kill 2>/dev/null || true
	@lsof -ti udp:4553 | xargs kill 2>/dev/null || true
	@echo "showcase stopped"

clean: stop ## Stop server and remove build artifacts
	rm -f server static/bytewire.wasm
