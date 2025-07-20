# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Rules

- DO NOT edit .gitignore
- DO NOT edit the contents of the .git directory
- DO NOT edit generated files
- Prefer functional style without side effect.

## Project Overview

Please refer to [README.md](./README.md) for detailed information.

## Development Commands

### Analyzer Build Commands

To generate the Go ANTLR parser from grammar/Expression.g4, run the following command from the project root:

```bash
# Generate ANTLR parser from Expression.g4
docker build --target antlr-generated --output=type=local,dest=analyzer/gen/parser -f analyzer/Dockerfile .
```

To build the WASM binary using Docker:

```bash
# Build optimized WASM binary with TinyGo and wasm-opt
docker build --target wasm-output --output=type=local,dest=. -f analyzer/Dockerfile .
```

For additional analyzer build commands, please see [analyzer/CLAUDE.md](./analyzer/CLAUDE.md).

## Project Structure

```
├── grammar/   # Grammar definitions in ANTLR4 format
├── analyzer/    # ANTLR4 analyzer implementation in Go
└── README.md
```

## Architecture

*Architecture details will be documented as the codebase grows*