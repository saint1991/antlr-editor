# antlr-editor

Syntax highlighted editor. Syntax is defined using ANTLR.

## Features

- ANTLR4-based expression parsing and analysis
- WebAssembly (WASM) support for browser integration
- Python FFI bindings for native integration
- Optimized builds with TinyGo for smaller binary sizes
- Real-time syntax highlighting and error detection

## Quick Start

### WASM Build

Build optimized WASM modules:

```bash
cd analyzer
./build-wasm.sh
```

Compare WASM binary sizes:

```bash
cd analyzer  
./benchmark-wasm.sh
```

### Python FFI

Install the Python bindings:

```bash
cd analyzer/ffi/python
pip install .
```

Example usage:

```python
from antlr_analyzer import Analyzer

analyzer = Analyzer()
result = analyzer.analyze("age > 18 AND name = 'John'")

if result.is_valid:
    for token in result.tokens:
        print(f"{token.text} -> {token.token_type.name}")
else:
    for error in result.errors:
        print(f"Error: {error.message}")
```

For detailed development instructions, see [analyzer/CLAUDE.md](./analyzer/CLAUDE.md) and [analyzer/ffi/python/README.md](./analyzer/ffi/python/README.md).
