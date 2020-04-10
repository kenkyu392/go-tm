OUTPUT_PATH   := _output
COVERAGE_FILE := $(OUTPUT_PATH)/coverage.out
COVERAGE_HTML := $(OUTPUT_PATH)/coverage.html

.PHONY: clean
clean:
	@go clean -r -cache -testcache -modcache

.PHONY: test
test:
	@go test -count=1 -v ./...

.PHONY: cover
cover:
	@mkdir -p $(OUTPUT_PATH)
	@go test -coverprofile=$(COVERAGE_FILE) -covermode=count ./...
	@go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@go tool cover -func=$(COVERAGE_FILE)

.PHONY: generate
generate:
	@go generate ./...
