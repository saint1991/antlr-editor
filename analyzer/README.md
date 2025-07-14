# Analyzer Module

ANTLR4-based expression analyzer for the antlr-editor project, implemented in Go with support for multiple target platforms.

## Overview

This module provides parsing capabilities for mathematical and logical expressions using ANTLR4 Go runtime. It supports compilation to multiple targets including native Go binaries, WebAssembly (WASM), and Python FFI.

## Project Structure

```
analyzer/
├── go.mod              # Go module definition
├── go.sum              # Dependency checksums
├── main.go             # Basic usage examples and tests
├── core/               # Core analyzer logic
│   ├── app/            # Application layer
│   ├── infrastructure/ # Infrastructure layer  
│   └── models/         # Shared data structures
├── wasm/               # WebAssembly target
├── ffi/                # Python FFI target
├── python/             # Python bindings and shared libraries
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

```bash
# Build WASM module (from analyzer directory)
docker build --target wasm-output --output=type=local,dest=dist -f Dockerfile .
```

## Python FFI Support

The analyzer provides Python bindings through a C shared library interface.

### Building the Python FFI Library

```bash
# Build FFI shared library using Docker
docker build --target ffi-output --output=type=local,dest=python -f Dockerfile .

# Or build directly with Go (requires ANTLR parser generation first)
go build -buildmode=c-shared -o python/libanalyzer.dylib ./ffi/analyzer.go  # macOS
go build -buildmode=c-shared -o python/libanalyzer.so ./ffi/analyzer.go    # Linux
```

### Using the Python Bindings

```python
# Basic usage
from python.analyzer_ffi import AnalyzerFFI

analyzer = AnalyzerFFI()

# Validate expression syntax
is_valid = analyzer.validate("2 + 3 * 4")
print(f"Valid: {is_valid}")

# Get detailed analysis
result = analyzer.analyze("func(a, b)")
print(f"Tokens: {len(result['tokens'])}")
print(f"Errors: {len(result['errors'])}")

# Convenience functions
from python.analyzer_ffi import validate, analyze, get_tokens

print(validate("x + y"))           # True
tokens = get_tokens("2 + 3")       # List of token information
result = analyze("invalid syntax") # Full analysis with errors
```

### Python FFI Features

- **Expression validation**: Check syntax correctness
- **Token analysis**: Get detailed token information for syntax highlighting  
- **Error reporting**: Detailed error messages with line/column information
- **Memory management**: Automatic cleanup of allocated strings
- **Cross-platform**: Supports Linux, macOS, and Windows
- **Zero dependencies**: Uses only Python standard library

### Python Package Structure

```
python/
├── __init__.py          # Package initialization
├── analyzer_ffi.py      # Main FFI wrapper
├── example.py           # Usage examples
├── setup.py             # Python package setup
├── requirements.txt     # Dependencies (none required)
├── libanalyzer.so       # Shared library (Linux)
├── libanalyzer.dylib    # Shared library (macOS)
└── libanalyzer.h        # C header file
```

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