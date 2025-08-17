# CLAUDE.md - Analyzer Module

This file provides guidance to Claude Code (claude.ai/code) when working with the analyzer module.

## Module Overview

This is the Go-based ANTLR4 analyzer module for the antlr-editor project. It provides expression parsing and analysis capabilities with multiple compilation targets (native Go, WASM, Python FFI).

## Rules

- DO NOT edit .gitignore
- DO NOT edit the contents of the .git directory
- DO NOT commit generated code in the `gen/` directory
- Follow Go module conventions and idioms
- MUST run parser generation before building or testing

## Development Commands

**IMPORTANT: Parser generation is required before any build or test operations. The generated code is not committed to the repository.**

### Parser Generation

```bash
# From the analyzer directory
./codegen.sh
```

### Testing

```bash
go vet ./core/... ./ffi/... ./wasm/... 
go test ./...
```

### Linting and Type Checking

```bash
# Run golangci-lint (from analyzer directory)
golangci-lint run

# Check for compilation errors
go build ./...
```

### Building WASM

```bash
# Build WASM module with standard Go (from analyzer directory)
GOOS=js GOARCH=wasm go build -o analyzer.wasm ./wasm/analyzer.go

# Build optimized WASM module with TinyGo and wasm-opt
./build-wasm.sh

# Run WASM size benchmark and comparison
./benchmark-wasm.sh
```

## Project Structure

```
analyzer/
├── go.mod              # Go module definition
├── go.sum              # Dependency checksums
├── main.go             # Example usage and analyzer testing
├── codegen.sh          # Script to generate ANTLR parser code
├── Dockerfile          # Docker configuration for ANTLR code generation
├── .golangci.toml      # Linting configuration
├── README.md           # Analyzer module documentation
├── core/               # Core analyzer logic
│   ├── app/            # Application layer
│   │   ├── app.go          # Application coordinator
│   │   ├── analyzer.go     # Expression analyzer for syntax highlighting
│   │   └── analyzer_test.go # Analyzer tests
│   ├── infrastructure/ # Infrastructure layer
│   │   ├── error_listener.go # Error listener implementations
│   │   └── parser.go       # Parser helper utilities
│   └── models/         # Shared data structures
│       ├── error.go    # ErrorInfo types
│       └── tokens.go   # TokenInfo, TokenType types
├── wasm/               # WebAssembly target
│   └── analyzer.go     # WASM-compatible analyzer
├── ffi/                # Python FFI target
│   └── analyzer.go     # FFI-compatible analyzer
└── gen/                # Generated ANTLR parser code (git-ignored)
    └── parser/         # Auto-generated parser implementation
```

## Key Dependencies

- Go 1.24.6 or higher
- ANTLR4 Go Runtime v4.13.1
- Docker (for parser generation)

## Common Tasks

### Adding a New Analyzer Feature
1. Update the grammar file in `../grammar/Expression.g4`
2. Regenerate the parser code using `./codegen.sh`
3. Update the analyzer logic in `core/app/analyzer.go`
4. Update the examples in `main.go` to test the new features
5. Run linting and tests

### Debugging Analyzer Issues
1. Use the main program with custom expressions
2. Check the generated parse tree output
3. Verify the grammar rules in Expression.g4
4. Check token classification in `core/models/tokens.go`
