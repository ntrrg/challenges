GO ?= go
module := $(shell $(GO) list -m)
PACKAGE ?= $(notdir $(module))

GODOC_PORT ?= 6060
HUGO_PORT ?= 1313

goFiles := $(shell find . -iname "*.go" -type f | grep -v "/_" | grep -v "^\./vendor")
goFilesSrc := $(shell $(GO) list -f '{{ range .GoFiles }}{{ $$.Dir }}/{{ . }} {{ end }}' ./...)
goFilesTest := $(shell $(GO) list -f "{{ range .TestGoFiles }}{{ $$.Dir }}/{{ . }} {{ end }}{{ range .XTestGoFiles }}{{ $$.Dir }}/{{ . }} {{ end }}" ./...)

.PHONY: all
all: build

.PHONY: build
build:
	$(GO) build ./...

.PHONY: clean
clean:

# Development

COVERAGE_FILE ?= coverage.txt
TARGET_FUNC ?= .
TARGET_PKG ?= ./...
WATCH_TARGET ?= test

.PHONY: benchmark
benchmark:
	$(GO) test -run none -bench "$(TARGET_FUNC)" -benchmem -v $(TARGET_PKG)

.PHONY: ca
ca:
	golangci-lint run

.PHONY: ca-fast
ca-fast:
	golangci-lint run --fast

.PHONY: ci
ci: build test lint ca

.PHONY: ci-race
ci-race: build test-race lint ca

.PHONY: clean-dev
clean-dev: clean
	rm -rf $(COVERAGE_FILE)

.PHONY: coverage
coverage:
	$(GO) tool cover -func $(COVERAGE_FILE)

.PHONY: coverage-web
coverage-web:
	$(GO) tool cover -html $(COVERAGE_FILE)

.PHONY: doc
doc:
	@echo "Go to http://localhost:$(HUGO_PORT)/en/projects/$(PACKAGE)/"
	@docker run --rm -it \
		-e PORT=$(HUGO_PORT) \
		-p $(HUGO_PORT):$(HUGO_PORT) \
		-v "$$PWD/.ntweb":/site/content/projects/$(PACKAGE)/ \
		ntrrg/ntweb:editing --port $(HUGO_PORT)

.PHONY: format
format:
	gofmt -s -w -l $(goFiles)

.PHONE: godoc
godoc:
	@echo "Go to http://localhost:$(GODOC_PORT)/pkg/$(module)/"
	godoc -http :$(GODOC_PORT) -play

.PHONY: lint
lint:
	gofmt -d -e -s $(goFiles)

.PHONY: test
test:
	$(GO) test \
		-run "$(TARGET_FUNC)" \
		-coverprofile $(COVERAGE_FILE) \
		-v $(TARGET_PKG)

.PHONY: test-race
test-race:
	$(GO) test -race \
		-run "$(TARGET_FUNC)" \
		-coverprofile $(COVERAGE_FILE) \
		-v $(TARGET_PKG)

.PHONY: watch
watch:
	reflex -d "none" -r '\.go$$' -- $(MAKE) -s $(WATCH_TARGET)
