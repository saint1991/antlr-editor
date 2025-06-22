# CLAUDE.md - Parser Module

This file provides guidance to Claude Code (claude.ai/code) when working with the parser module code.

## Module Overview

This is the Go-based ANTLR4 parser module for the antlr-editor project. It provides expression parsing capabilities and will support multiple compilation targets (native Go, WASM, Python FFI).

## Rules

- MUST NOT edit .gitignore
- MUST NOT edit content of .git directory
- MUST NOT commit generated code in `gen/` directory
- MUST run parser generation before working with parser code
- MUST follow Go module conventions and idioms

## Development Commands

### Parser Generation

**IMPORTANT: All Docker commands must be executed from the project root directory (parent of parser/).**

```bash
# Generate ANTLR parser code from grammar
docker build --target antlr-generated --output=type=local,dest=parser/gen/parser -f parser/Dockerfile .
```

### Testing

TBD

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
├── examples/           # Usage examples and tests
│   └── main.go        # Basic parser functionality test
├── wasm/              # WebAssembly target (planned)
├── ffi/               # Python FFI target (planned)
└── gen/               # Generated ANTLR parser code (git-ignored)
    └── parser/        # Auto-generated parser implementation
```

## Key Dependencies

- Go 1.24.4+
- ANTLR4 Go Runtime v4.13.1
- Docker (for parser generation)

## Important Notes

1. **Generated Code**: The `gen/` directory contains auto-generated ANTLR parser code. This is excluded from git and must be regenerated when:
   - The grammar file changes
   - Setting up the project for the first time
   - After cleaning the project

2. **Module Path**: The Go module is `antlr-editor/parser`

3. **Import Paths**: When importing generated parser code:
   ```go
   import "antlr-editor/parser/gen/parser"
   ```

4. **Error Handling**: The parser implements custom error listeners for better error reporting

## Common Tasks

### Adding a new parser feature
1. Update the grammar file in `../grammar/Expression.g4`
2. Regenerate parser code using Docker command above
3. Update examples in `examples/main.go` to test new features
4. Run linting and tests

### Debugging parser issues
1. Use the examples program with custom expressions
2. Check the generated parse tree output
3. Verify grammar rules in Expression.g4
