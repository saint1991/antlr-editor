# ANTLR Expression Analyzer - Python Bindings

Python bindings for the ANTLR Expression Analyzer, providing syntax analysis and token classification for expressions.

## Project Structure

```
analyzer/ffi/python/
├── Makefile           # Build automation
├── pyproject.toml     # Package configuration
├── README.md          # This file
├── src/
│   └── analyzer/
│       ├── __init__.py     # Package exports
│       ├── analyzer.py     # Main Analyzer class
│       └── models/         # Data models
│           ├── __init__.py
│           ├── error.py    # ErrorInfo model
│           ├── result.py   # TokenizeResult model
│           └── token.py    # TokenInfo and TokenType models
└── examples/
    └── example.py     # Usage examples
```

## Installation

### From Source with Make

1. Build the wheel package:
```bash
cd analyzer/ffi/python
make wheel
```

2. Install the built wheel:
```bash
pip install dist/*.whl
```

### From Source with Docker

1. Build the wheel using Docker:
```bash
docker build --target python-wheel-output --output=type=local,dest=. -f analyzer/Dockerfile .
```

2. Install the built wheel:
```bash
pip install analyzer-*.whl
```

## Usage

### Basic Example

```python
from analyzer import Analyzer

# Create analyzer instance
analyzer = Analyzer()

# Validate an expression
expression = "user.age > 18 AND user.name = 'John'"
is_valid = analyzer.validate(expression)
print(f"Expression is valid: {is_valid}")

# Tokenize expression for detailed information
result = analyzer.tokenize(expression)

# Check if valid
if result.is_valid:
    print("Expression is valid!")
else:
    print("Expression has errors:")
    for error in result.errors:
        print(f"  - {error.message} at line {error.line}, column {error.column}")

# Examine tokens
for token in result.tokens:
    print(f"Token: {token.text} (type: {token.token_type.name})")
```

### Token Types

The analyzer recognizes the following token types:

- `STRING` - String literals (e.g., 'hello', "world")
- `INTEGER` - Integer numbers (e.g., 42, -10)
- `FLOAT` - Floating point numbers (e.g., 3.14, -0.5)
- `BOOLEAN` - Boolean values (true, false)
- `COLUMN_REFERENCE` - Column references (e.g., user.name, table.column)
- `FUNCTION` - Function calls (e.g., MAX, MIN, COUNT)
- `OPERATOR` - Operators (e.g., +, -, *, /, =, >, <, AND, OR)
- `COMMA` - Comma separator
- `LEFT_PAREN` / `RIGHT_PAREN` - Parentheses
- `LEFT_BRACKET` / `RIGHT_BRACKET` - Square brackets
- `WHITESPACE` - Spaces, tabs, newlines
- `ERROR` - Invalid tokens
- `EOF` - End of file
