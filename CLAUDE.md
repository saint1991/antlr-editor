# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Rules

- DO NOT edit .gitignore
- DO NOT edit the contents of the .git directory
- DO NOT edit generated files

## Project Overview

Please refer to [README.md](./README.md) for detailed information.

## Development Commands

### Parser Build Commands

To generate the Go ANTLR parser from grammar/Expression.g4, run the following command from the project root:

```bash
# Generate ANTLR parser from Expression.g4
docker build --target antlr-generated --output=type=local,dest=parser/gen/parser -f parser/Dockerfile .
```

For additional parser build commands, please see [parser/CLAUDE.md](./parser/CLAUDE.md).

## Project Structure

```
├── grammar/   # Grammar definitions in ANTLR4 format
├── parser/    # ANTLR4 parser implementation in Go
└── README.md
```

## Architecture

*Architecture details will be documented as the codebase grows*