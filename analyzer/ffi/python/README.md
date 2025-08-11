# ANTLR Expression Analyzer - Python Bindings

Python bindings for the ANTLR Expression Analyzer, providing syntax analysis and validation for expressions.

## Installation

### From Source

1. Clone the repository:
```bash
git clone https://github.com/saint1991/antlr-editor.git
cd antlr-editor/analyzer/ffi/python
```

2. Build and install the package:
```bash
pip install .
```

This will automatically:
- Generate the ANTLR parser (if needed)
- Build the Go shared library
- Install the Python package

### From Wheel

Pre-built wheels can be installed directly:

```bash
pip install antlr-analyzer
```

## Usage

### Basic Example

```python
from antlr_analyzer import Analyzer

# Create analyzer instance
analyzer = Analyzer()

# Validate an expression
expression = "user.age > 18 AND user.name = 'John'"
is_valid = analyzer.validate(expression)
print(f"Expression is valid: {is_valid}")

# Analyze expression for detailed information
result = analyzer.analyze(expression)

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

### API Reference

#### `Analyzer`

Main class for expression analysis.

##### Methods

- `__init__(lib_path: Optional[Path] = None)` - Initialize analyzer with optional custom library path
- `validate(expression: str) -> bool` - Check if an expression is syntactically valid
- `analyze(expression: str) -> AnalysisResult` - Perform detailed analysis of an expression

#### `AnalysisResult`

Result of expression analysis.

##### Properties

- `tokens: List[TokenInfo]` - List of tokens found in the expression
- `errors: List[ErrorInfo]` - List of syntax errors (empty if valid)
- `is_valid: bool` - True if the expression has no errors

#### `TokenInfo`

Information about a single token.

##### Properties

- `token_type: TokenType` - Type of the token
- `text: str` - Text content of the token
- `start: int` - Start position in the expression
- `end: int` - End position in the expression
- `line: int` - Line number (1-based)
- `column: int` - Column number (1-based)
- `is_valid: bool` - Whether the token is valid

#### `ErrorInfo`

Information about a syntax error.

##### Properties

- `message: str` - Error description
- `line: int` - Line number where error occurred
- `column: int` - Column number where error occurred
- `start: int` - Start position of the error
- `end: int` - End position of the error

## Examples

### SQL-like Expression
```python
analyzer = Analyzer()
result = analyzer.analyze("SELECT name, age FROM users WHERE age > 21 AND status = 'active'")

for token in result.tokens:
    if token.token_type == TokenType.FUNCTION:
        print(f"Function: {token.text}")
    elif token.token_type == TokenType.COLUMN_REFERENCE:
        print(f"Column: {token.text}")
```

### Mathematical Expression
```python
result = analyzer.analyze("(price * quantity) - (discount / 100 * price * quantity)")

# Find all operators
operators = [t for t in result.tokens if t.token_type == TokenType.OPERATOR]
for op in operators:
    print(f"Operator: {op.text} at position {op.start}")
```

### Error Detection
```python
# Expression with syntax error
result = analyzer.analyze("user.name = 'John AND age > 18")

if not result.is_valid:
    for error in result.errors:
        print(f"Syntax error: {error.message}")
        print(f"Location: line {error.line}, column {error.column}")
```

## Development

### Building from Source

1. Ensure you have Go 1.24+ installed
2. Install build dependencies:
   ```bash
   pip install setuptools wheel
   ```

3. Build the wheel:
   ```bash
   cd analyzer/ffi
   ./build-ffi.sh --wheel
   ```

### Running Tests

```python
# Example test script
import unittest
from antlr_analyzer import Analyzer, TokenType

class TestAnalyzer(unittest.TestCase):
    def setUp(self):
        self.analyzer = Analyzer()
    
    def test_valid_expression(self):
        self.assertTrue(self.analyzer.validate("age > 18"))
    
    def test_invalid_expression(self):
        self.assertFalse(self.analyzer.validate("age > "))
    
    def test_token_analysis(self):
        result = self.analyzer.analyze("name = 'John'")
        self.assertEqual(len(result.tokens), 5)  # name, =, 'John', whitespaces
        self.assertEqual(result.tokens[0].token_type, TokenType.COLUMN_REFERENCE)

if __name__ == '__main__':
    unittest.main()
```

## License

MIT License - see LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.