PROTO_DIR ?= proto
TOOL ?= coherence-contracts-tools

.PHONY: check check-contracts check-buf inventory graph cargo-check fmt clean

check: fmt cargo-check check-contracts check-buf

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

inventory:
	cargo run -p $(TOOL) -- inventory --proto-dir $(PROTO_DIR)

graph:
	cargo run -p $(TOOL) -- graph --proto-dir $(PROTO_DIR) > graph/generated/imports.dot

clean:
	cargo clean
