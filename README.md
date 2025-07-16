# antlr-editor

Syntax highlighted editor. Syntax is defined using ANTLR.

## Features

- ANTLR4-based expression parsing and analysis
- WebAssembly (WASM) support for browser integration
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

For detailed development instructions, see [analyzer/CLAUDE.md](./analyzer/CLAUDE.md).
