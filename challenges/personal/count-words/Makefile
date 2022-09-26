.PHONY: all
all: build

-include config.mk

COMPILERS ?= go
BIN ?= solution
TEST_CASE ?= 00
MODE ?=
CLEAN_TARGET ?=

_input  := input/input$(TEST_CASE).txt
_output := output/output$(TEST_CASE).txt

.PHONY: bench
bench:
	@echo "  Comp  |      Size      |   Time   | CPU  |  Memory  "
	@echo "--------+----------------+----------+------+----------"
	@$(MAKE) -s $(COMPILERS:%=bench-%)

TIME ?= time
TIME_FORMAT ?= %E |  %P |  %Mkb

bench-%: $(BIN)-%
	@_comp="$$(printf "%7s" "$(@:bench-%=%)")" && \
		_size="$$(printf "%6s" "$$(du -h "$<" | cut -f 1)")" && \
		_mode="$$(printf "%-7s" "$$( \
			file "$<" | grep -q "statically linked" && \
			echo "Static" || (file "$<" | grep -q "ASCII text" && \
				echo "Interp." || echo "Dynamic") \
		)")" && \
		_fmt="$$_comp | $$_size $$_mode |  $(TIME_FORMAT)" && \
	cat "$(_input)" | $(TIME) -f "$$_fmt" "./$<" > "/dev/null"

.PHONY: build
build: $(COMPILERS:%=build-%)

build-%:
	@$(MAKE) -s $(@:build-%=$(BIN)-%)

.PHONY: check
check: $(COMPILERS:%=check-%)

check-%:
	@_got="$$(mktemp -u)" && _want="$$(mktemp -u)" && \
		$(MAKE) -s $(@:check-%=run-%) > "$$_got" && \
		([ "$(MODE)" = "unordered" ] && sort -o "$$_got" "$$_got" || true) && \
		cat "$(_output)" > "$$_want" && \
		([ "$(MODE)" = "unordered" ] && sort -o "$$_want" "$$_want" || true) && \
		busybox diff "$$_want" "$$_got"

.PHONY: clean
clean:
	rm -rf solution solution-* *.prof profile.txt target $(CLEAN_TARGET)

.PHONY: run
run: $(COMPILERS:%=run-%)

run-%: $(BIN)-%
	@cat "$(_input)" | "./$<"

# C.

CLANG ?= clang
GCC ?= gcc
TCC ?= tcc
CFLAGS ?= -O2
CSOURCE ?= main.c

$(BIN)-clang: $(CSOURCE)
	$(CLANG) -o "$@" $(CFLAGS) $^

$(BIN)-gcc: $(CSOURCE)
	$(GCC) -o "$@" $(CFLAGS) $^

$(BIN)-tcc: $(CSOURCE)
	$(TCC) -o "$@" $(CFLAGS) $^

# Go.

GO ?= go
GOTIP ?= gotip
DLV ?= dlv
GOFLAGS ?=
GOSOURCE ?= main.go

.PHONY: debug-go
debug-go: $(GOSOURCE)
	$(DLV) -r "$(_input)" debug "$<"

$(BIN)-go: $(GOSOURCE)
	$(GO) build -o "$@" $(GOFLAGS) $^

$(BIN)-gotip: $(GOSOURCE)
	$(GOTIP) build -o "$@" $(GOFLAGS) $^

TINYGO ?= docker run --rm -i -t -u "$$UID" \
	-v "$$HOME/.cache:/.cache" -v "$$HOME/go/pkg:/go/pkg" \
	-v "/dev/bus/usb:/dev/bus/usb" -v "$$PWD:/src" -w "/src" \
	tinygo/tinygo:0.24.0 tinygo

TINYGOFLAGS ?=
TINYGOSOURCE ?= main_tinygo.go

$(BIN)-tinygo: $(TINYGOSOURCE)
	$(TINYGO) build -o "$@" $(TINYGOFLAGS) $^

# Node.js.

NODE ?= node
NODEFLAGS ?= 
NODESOURCE ?= main.js

$(BIN)-node: $(NODESOURCE)
	@echo "#!/usr/bin/env $(NODE) $(NODEFLAGS)" > "$@"
	@cat "$<" >> "$@"
	@chmod +x "$@"

# Python.

PYTHON ?= python
PYTHONFLAGS ?= 
PYTHONSOURCE ?= main.py

$(BIN)-python: $(PYTHONSOURCE)
	@echo "#!/usr/bin/env $(PYTHON) $(PYTHONFLAGS)" > "$@"
	@cat "$<" >> "$@"
	@chmod +x "$@"

# Rust.

CARGO ?= cargo
RUSTFLAGS ?= --release -q
RUSTSOURCE ?= main.rs

$(BIN)-rust: $(RUSTSOURCE)
	$(CARGO) build $(RUSTFLAGS) && cp "target/release/$(BIN)" "$@"

# V.

V ?= v
VFLAGS ?= -skip-unused -prod
VSOURCE ?= main.v

$(BIN)-v: $(VSOURCE)
	$(V) -o "$@" $(VFLAGS) $^
