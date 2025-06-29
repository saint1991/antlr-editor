# CLAUDE.md - Parser Module

This file provides guidance to Claude Code (claude.ai/code) when working with the parser module.

## Module Overview

This is the Go-based ANTLR4 parser module for the antlr-editor project. It provides expression parsing capabilities and will support multiple compilation targets (native Go, WASM, Python FFI).

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
# From the parser directory
./codegen.sh
```

### Testing

```bash
cd parser
go vet ./core/... ./ffi/... ./wasm/... 
go test ./...
```

### Linting and Type Checking

```bash
# Run golangci-lint (from parser directory)
cd parser
golangci-lint run

# Check for compilation errors
go build ./...
```

## Project Structure

```
parser/
├── go.mod              # Go module definition
├── go.sum              # Dependency checksums
├── main.go             # Example usage and parser testing
├── codegen.sh          # Script to generate ANTLR parser code
├── Dockerfile          # Docker configuration for ANTLR code generation
├── .golangci.toml      # Linting configuration
├── core/               # Core parser logic
│   ├── analyzer.go     # Expression analyzer for syntax highlighting
│   ├── analyzer_test.go # Analyzer tests
│   ├── validator.go    # Expression validator implementation
│   ├── validator_test.go # Validator tests
│   └── models/         # Shared data structures
│       └── analysis.go # TokenInfo, ErrorInfo, AnalysisResult types
├── wasm/               # WebAssembly target
│   └── validator.go    # WASM-compatible validator and analyzer
├── ffi/                # Python FFI target
│   └── validator.go    # FFI-compatible validator and analyzer
└── gen/                # Generated ANTLR parser code (git-ignored)
    └── parser/         # Auto-generated parser implementation
```

## Key Dependencies

- Go 1.24.4 or higher
- ANTLR4 Go Runtime v4.13.1
- Docker (for parser generation)

## Common Tasks

### Adding a New Parser Feature
1. Update the grammar file in `../grammar/Expression.g4`
2. Regenerate the parser code using the Docker command above
3. Update the examples in `examples/main.go` to test the new features
4. Run linting and tests

### Debugging Parser Issues
1. Use the examples program with custom expressions
2. Check the generated parse tree output
3. Verify the grammar rules in Expression.g4
