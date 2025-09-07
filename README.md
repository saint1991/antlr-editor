# ANTLR Editor

A comprehensive expression editor built with ANTLR4, featuring real-time syntax highlighting, error detection, and multi-platform support. The project consists of a modern web editor and a powerful analyzer engine that supports WebAssembly and Python FFI.

## 🌟 Features

### Web Editor
- **Real-time syntax highlighting** with CodeMirror 6
- **Error detection and display** with inline diagnostics
- **Intelligent autocompletion** with context-aware suggestions
- **Code formatting** with proper indentation
- **Hover tooltips** for function documentation
- **Dark/light theme support**
- **Smart bracket matching and folding**

### Analyzer Engine
- **ANTLR4-based parsing** for mathematical and logical expressions
- **WebAssembly support** for browser integration
- **Python FFI bindings** for native integration
- **Optimized WASM builds** with TinyGo for minimal binary sizes
- **Multiple compilation targets** (native Go, WASM, Python FFI)

### Supported Expression Grammar
- **Arithmetic operators**: `+`, `-`, `*`, `/`, `^` (with proper precedence)
- **Comparison operators**: `<`, `<=`, `>`, `>=`, `==`, `!=`
- **Logical operators**: `&&`, `||`
- **Literals**: integers, floats, strings, booleans
- **Column references**: `[column_name]` format
- **Function calls**: `FUNCTION_NAME(arg1, arg2, ...)` format
- **Parentheses**: for expression grouping

## 🚀 Quick Start

### Web Editor

Start the development server:

```bash
cd editor-app
npm install
npm start
```

Navigate to `http://localhost:4200` to use the editor.

### Analyzer - WASM Build

Build optimized WASM modules for browser usage:

```bash
# From project root
docker build --target wasm-output --output=type=local,dest=editor-app/public -f analyzer/Dockerfile .
```

### Analyzer - Python FFI

Install and use the Python bindings:

```bash
cd analyzer/ffi/python
pip install .
```

Example usage:

```python
from antlr_analyzer import Analyzer

analyzer = Analyzer()
result = analyzer.tokenize("age > 18 AND name = 'John'")

if result.is_valid:
    for token in result.tokens:
        print(f"{token.text} -> {token.token_type.name}")
else:
    for error in result.errors:
        print(f"Error: {error.message}")
```

## 🏗️ Project Structure

```
antlr-editor/
├── grammar/              # ANTLR4 grammar definitions
│   ├── Expression.g4        # Main expression grammar file
│   └── grammar.md           # Grammar documentation
├── analyzer/             # Go-based ANTLR4 analyzer engine
│   ├── core/                # Core analyzer logic
│   ├── wasm/                # WebAssembly target
│   ├── ffi/                 # Python FFI and C shared library
│   ├── gen/                 # Generated parser code (git-ignored)
│   ├── codegen.sh           # Parser generation script
│   └── Dockerfile           # Multi-stage build for ANTLR generation
├── editor-app/           # Angular-based web editor
│   ├── src/app/antlr-editor/   # Main editor component
│   │   └── extensions/         # CodeMirror extensions
│   │       ├── completion/        # Autocompletion
│   │       ├── format/            # Code formatting
│   │       ├── lint/              # Error linting
│   │       ├── syntax-highlight/  # Syntax highlighting
│   │       └── tooltip/           # Hover tooltips
│   └── src/wasm/             # WASM integration
├── .github/workflows/    # CI/CD pipelines
└── .devcontainer/        # Development container configuration
```

## 🔧 Development

### Prerequisites

- **Node.js** 18+ and npm
- **Go** 1.24.6+
- **Docker** (for ANTLR code generation)
- **Angular CLI** (`npm install -g @angular/cli`)

### Setup

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd antlr-editor
   ```

2. **Generate ANTLR parser code**:
   ```bash
   docker build --target antlr-generated --output=type=local,dest=analyzer/gen/parser -f analyzer/Dockerfile .
   ```

3. **Install editor dependencies**:
   ```bash
   cd editor-app
   npm install
   ```

4. **Build WASM module**:
   ```bash
   npm run build:wasm
   ```

### Development Commands

#### Web Editor
```bash
cd editor-app
npm start          # Start development server
npm run build      # Build production bundle
npm run test       # Run unit tests
npm run lint       # Run linter
npm run format     # Format code
```

#### Analyzer
```bash
cd analyzer
./codegen.sh       # Generate ANTLR parser
go test ./...      # Run tests
go build ./...     # Build all targets
golangci-lint run  # Run linter
```

### Testing Expressions

Try these sample expressions in the editor:

```javascript
// Arithmetic
2 + 3 * 4
(10 + 5) / 3

// Comparisons and logic
age >= 18 && status == 'active'
[price] * [quantity] > 100

// Functions
MAX(a, b, c)
ROUND([value], 2)

// Complex expressions
([salary] + [bonus]) * 0.8 > MIN_THRESHOLD && [department] == 'Engineering'
```

## 📚 Documentation

- [Analyzer Documentation](./analyzer/README.md) - Detailed analyzer module guide
- [Python FFI Documentation](./analyzer/ffi/python/README.md) - Python bindings usage
- [Grammar Definition](./grammar/Expression.g4) - ANTLR4 grammar specification
- [Development Guidelines](./CLAUDE.md) - Contributing and development rules

## 🔄 CI/CD

The project includes GitHub Actions workflows for:

- **Analyzer CI**: Go testing, linting, and building
- **Editor App CI**: Angular testing, building, and deployment
- **WASM Build**: Automated WASM generation and optimization
- **Python FFI**: Testing and packaging Python bindings

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
