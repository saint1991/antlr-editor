# Analyzer Module

ANTLR4-based expression analyzer for the antlr-editor project, implemented in Go with support for multiple target platforms.

## Overview

This module provides parsing capabilities for mathematical and logical expressions using ANTLR4 Go runtime. It supports compilation to multiple targets including native Go binaries, WebAssembly (WASM), and Python FFI.

## Project Structure

```
analyzer/
├── go.mod              # Go module definition
├── go.sum              # Dependency checksums
├── examples/
│   └── main.go         # Basic usage examples and tests
├── wasm/               # WebAssembly target (planned)
├── ffi/                # Python FFI target (planned)
└── gen/                # Generated ANTLR parser code (git-ignored)
    └── parser/
```

## Dependencies

- **Go**: 1.24.4 or later
- **ANTLR4 Go Runtime**: v4.13.1
- **Docker**: Required for parser code generation

## Quick Start

### 1. Generate Parser Code

From the project root directory:

```bash
# Generate ANTLR parser code
docker build --target antlr-generated --output=type=local,dest=analyzer/gen/parser -f analyzer/Dockerfile .
```

### 2. Run Examples

```bash
cd analyzer
go run ./examples
```

Expected output:
```
Parse tree: (expression (expression (literal 1)) + (expression (expression (literal 2)) * (expression (literal 3))))
```

## Supported Grammar

The parser supports the Expression.g4 grammar located in `../grammar/Expression.g4`, which includes:

- **Arithmetic operators**: `+`, `-`, `*`, `/`, `^` (with proper precedence)
- **Comparison operators**: `==`, `!=`
- **Logical operators**: `||`, `&&`
- **Literals**: integers, floats, strings
- **Function calls**: `FUNCTION_NAME(arg1, arg2, ...)`
- **Column references**: `[column_name]`
- **Parentheses**: for grouping expressions

## Development

### Code Generation

The parser code is generated from ANTLR grammar files and is not committed to git. To regenerate:

```bash
# From project root
docker build --target antlr-generated --output=type=local,dest=parser/gen/parser -f parser/Dockerfile .
```

### Testing

```bash
# Run basic functionality test
go run ./examples

# Run with different expressions
go run ./examples "2 + 3 * 4"
go run ./examples "MAX(1, 2, 3)"
go run ./examples "[column_a] > 5"
```

## WebAssembly (WASM) Support

The analyzer can be compiled to WebAssembly for browser usage:


## Python FFI Support

### Python FFI Target (`ffi/`)
- CGO-based shared library for Python integration
- Native performance with Python convenience
- Support for complex data structures

## Architecture

The parser module follows Go project conventions:

- **Generated code** in `gen/` (excluded from git)
- **Target-specific implementations** in dedicated directories
- **Shared core logic** accessible to all targets
- **Examples and tests** in `examples/`

## Contributing

1. Ensure Docker is installed for parser generation
2. Follow Go coding conventions
3. Test changes with `go run ./examples`
4. Do not commit generated code in `gen/`

## License

MIT License - see the main project LICENSE file.