all: lint format tests
.PHONY: lint
lint:
	@echo "Running golangci-lint";
	@golangci-lint run --color=always;

.PHONY: format 
format:
	@echo "Running gofumpt formatter";
	@gofumpt -l -w -d;

.PHONY: test
test:
	@echo "Running tests";
# See https://github.com/gotestyourself/gotestsum
	@gotestsum --format-hide-empty-pkg; 
