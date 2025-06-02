# MyLang Interpreter Makefile

# Set the Go binary path
GO := go

# Set the ANTLR tool path - adjust as needed for your system
ANTLR_JAR := antlr-4.13.1-complete.jar

# Project directories
GRAMMAR_DIR := ./grammar
PARSER_DIR := ./parser
CMD_DIR := ./cmd

# Binary output
BIN_NAME := mylang
BIN_PATH := ./bin/$(BIN_NAME)

.PHONY: all
all: build

# Download ANTLR4 jar if needed
download-antlr:
	@if [ ! -f $(ANTLR_JAR) ]; then \
		echo "Downloading ANTLR4 JAR..." && \
		curl -O https://www.antlr.org/download/$(ANTLR_JAR); \
	fi

# Generate parser code from grammar
.PHONY: generate
generate: download-antlr
	@echo "Generating parser from grammar..."
	java -jar $(ANTLR_JAR) -Dlanguage=Go -visitor -package parser -o $(PARSER_DIR) $(GRAMMAR_DIR)/MyLang.g4

# Run all tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test -v ./...

# Run tests with coverage
.PHONY: coverage
coverage:
	@echo "Running tests with coverage..."
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out

# Build the interpreter
.PHONY: build
build: generate
	@echo "Building interpreter..."
	$(GO) build -o $(BIN_PATH) $(CMD_DIR)/main.go

# Clean generated files
.PHONY: clean
clean:
	@echo "Cleaning generated files..."
	rm -rf $(PARSER_DIR)/*.go
	rm -f $(BIN_PATH)
	rm -f coverage.out

# Format the code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Vet the code
.PHONY: vet
vet:
	@echo "Vetting code..."
	$(GO) vet ./...

# Run the interpreter
.PHONY: run
run: build
	@echo "Running interpreter..."
	$(BIN_PATH)

# Run a specific example
.PHONY: run-example
run-example: build
	@echo "Running example $(EXAMPLE)..."
	$(BIN_PATH) ./examples/$(EXAMPLE).mylang
