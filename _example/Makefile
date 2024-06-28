# Copyright 2024 Canonical Ltd.

# Show this help
help:
	@echo 'Usage:'
	@echo 'make <command>'
	@echo ''
	@echo 'Available commands:'
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t
.PHONY: help

# Run the server
run:
	@go run ./cmd
.PHONY: run

# Test the server (with cleanup at the end)
test:
	@bash test.sh --bail-on-error --reset --cleanup
	@# At the end, `state.json` and `state.zero.json` must be equal.
	@diff --ignore-trailing-space state.zero.json state.json || echo "unexpected 'state.json' content"
.PHONY: test

# Clean persisted state
clean:
	@rm state.json
.PHONY: clean
