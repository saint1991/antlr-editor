# Makefile for Python FFI distribution

# Configuration
UV := uv
PYTHON := python3
PACKAGE_NAME := analyzer
GO_MODULE_DIR := ../../
DIST_DIR := dist
BUILD_DIR := build
WHEEL_DIR := $(DIST_DIR)

# Platform detection
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
    LIB_EXT := .dylib
    LIB_NAME := libanalyzer.dylib
endif
ifeq ($(UNAME_S),Linux)
    LIB_EXT := .so
    LIB_NAME := libanalyzer.so
endif

# Targets
.PHONY: all shared-lib codegen clean wheel lint format info check-deps

# Default target
all: wheel

# Generate ANTLR parser (prerequisite for Go build)
codegen:
	@cd $(GO_MODULE_DIR) && ./codegen.sh


# Build Go FFI library
shared-lib: codegen
	@echo "Building Go FFI library..."
	@cd $(GO_MODULE_DIR) && \
		CGO_ENABLED=1 go build -buildmode=c-shared \
		-o ffi/python/src/analyzer/$(LIB_NAME) \
		./ffi/analyzer.go
	@echo "Go FFI library built successfully"


# Build Python wheel
wheel: shared-lib
	@echo "Building Python wheel..."
	@$(UV) build --wheel --out-dir $(WHEEL_DIR)
	@echo "Wheel built successfully in $(WHEEL_DIR)/"

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)/
	@rm -rf $(DIST_DIR)/
	@rm -rf *.egg-info/
	@rm -rf src/analyzer.egg-info
	@rm -rf src/analyzer/$(LIB_NAME)
	@rm -rf src/analyzer/libanalyzer.h
	@find . -name "*.pyc" -delete
	@find . -name "__pycache__" -type d -exec rm -rf {} + 2>/dev/null || true
	@echo "Clean completed"

# Run linting
lint:
	@echo "Running linting..."
	@$(UV) run ruff check src/
	@echo "Linting completed"

# Format code
format:
	@echo "Formatting code..."
	@$(UV) run ruff format src/
	@echo "Code formatted"

# Show package info
info:
	@echo "Package: $(PACKAGE_NAME)"
	@echo "Python: $(shell $(PYTHON) --version)"
	@echo "Go: $(shell go version 2>/dev/null || echo 'Go not found')"
	@echo "Platform: $(UNAME_S)"
	@echo "Library: $(LIB_NAME)"
	@if [ -f "$(WHEEL_DIR)"/*.whl ]; then \
		echo "Built wheels:"; \
		ls -la $(WHEEL_DIR)/*.whl; \
	fi
