# Copyright (C) 2024 Canonical Ltd.
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as
# published by the Free Software Foundation, either version 3 of the
# License, or (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

# Show this help
help:
	@echo 'Usage:'
	@echo 'make <command>'
	@echo ''
	@echo 'Available commands:'
	@awk '/^#/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print substr($$1,1,index($$1,":")),c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t
.PHONY: help

# Pull the latest stable OpenAPI spec and generate Go types
pull-spec:
	./script/generate-resources-from-spec.sh
.PHONY: pull-spec

# Generate test mocks
mocks:
	go install go.uber.org/mock/mockgen@v0.3.0
	go generate ./...
.PHONY: mocks

# Run tests with coverage
test-coverage: mocks
	go test ./... -cover -coverprofile coverage_source.out $(ARGS)
	# this will be cached, just needed to get the test.json
	go test ./... -cover -coverprofile coverage_source.out  $(ARGS) -json > test_source.json
	cat coverage_source.out | grep -v "mock_*" | tee coverage.out
	cat test_source.json | grep -v "mock_*" | tee test.json
.PHONY: test-coverage

# Run tests
test: mocks
	go test ./... $(ARGS)
.PHONY: test

# Clean working directory of generated files
clean:
	find . -name 'mock_*.go' -delete
	rm coverage_source.out coverage.out test_source.json test.json
.PHONY: clean
