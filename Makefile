PROTO_DIR ?= proto
TOOL ?= coherence-contracts-tools
BUF_PLUGIN_DIR ?= buf-plugins/coherence
BUF_PLUGIN_NAME ?= buf-plugin-coherence
SITE_DIR ?= site
HUGO ?= hugo

.PHONY: check check-contracts check-buf inventory graph cargo-check fmt clean
.PHONY: go-check go-test go-lint go-build go-install
.PHONY: site-build site-serve

check: fmt cargo-check go-check check-contracts check-buf site-build

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
	cd $(BUF_PLUGIN_DIR) && go build -o $(BUF_PLUGIN_NAME) .

go-install: go-build
	mv $(BUF_PLUGIN_DIR)/$(BUF_PLUGIN_NAME) $(HOME)/.local/bin/

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

site-build:
	@if command -v $(HUGO) >/dev/null 2>&1; then \
		cd $(SITE_DIR) && $(HUGO) --minify; \
	else \
		echo "hugo not installed; skipping site build"; \
	fi

site-serve:
	cd $(SITE_DIR) && $(HUGO) server

clean:
	cargo clean
