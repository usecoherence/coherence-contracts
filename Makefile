PROTO_DIR ?= proto
TOOL ?= coherence-contracts-tools
BUF_PLUGIN_DIR ?= buf-plugins/coherence

.PHONY: check check-contracts check-buf inventory graph cargo-check fmt clean
.PHONY: go-check go-test go-lint go-build

check: fmt cargo-check go-check check-contracts check-buf

fmt:
	cargo fmt --all --check

cargo-check:
	cargo check --workspace

check-contracts:
	cargo run -p $(TOOL) -- check --proto-dir $(PROTO_DIR)

check-buf:
	@if command -v buf >/dev/null 2>&1; then \
		buf lint; \
	else \
		echo "buf not installed; skipping buf lint"; \
	fi

go-check: go-build go-test go-lint

go-build:
	cd $(BUF_PLUGIN_DIR) && go build ./...

go-test:
	cd $(BUF_PLUGIN_DIR) && go test ./...

go-lint:
	@lint=$$(command -v golangci-lint 2>/dev/null || echo "$(HOME)/.local/share/nvim/mason/bin/golangci-lint"); \
	if [ -x "$$lint" ]; then \
		"$$lint" run ./$(BUF_PLUGIN_DIR)/...; \
	else \
		echo "golangci-lint not installed; skipping"; \
	fi

inventory:
	cargo run -p $(TOOL) -- inventory --proto-dir $(PROTO_DIR)

graph:
	cargo run -p $(TOOL) -- graph --proto-dir $(PROTO_DIR) > graph/generated/imports.dot

clean:
	cargo clean
