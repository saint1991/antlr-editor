# Expression Language Grammar Specification

## Overview

This document defines the grammar specification for an expression language. This language is designed for describing expressions that combine column references, function calls, arithmetic and logical operations in data processing and computation contexts.

## Lexical Elements

### 1. Literals

#### 1.1 String Literals
- **Syntax**: `'string'` or `"string"`
- **Description**: Strings enclosed in single quotes (') or double quotes (")
- **Escaping**: Quote characters within strings are escaped with backslash (\)
- **Examples**: 
  - `'Hello World'`
  - `"She said \"Hello\""`
  - `'It\'s a beautiful day'`

#### 1.2 Integer Literals
- **Syntax**: `[0-9]+`
- **Description**: Decimal numeric values
- **Examples**: `123`, `0`, `999`

#### 1.3 Float Literals
- **Syntax**: `[0-9]+\.[0-9]+` or scientific notation `[0-9]+(\.[0-9]+)?[eE][+-]?[0-9]+`
- **Description**: Numbers with decimal points, scientific notation is also supported
- **Examples**: 
  - `3.14`
  - `0.5`
  - `1.23e-4`
  - `2.5E+3`

### 2. Column References
- **Syntax**: `[identifier]`
- **Description**: Identifiers enclosed in square brackets
- **Constraints**: 
  - Can contain non-ASCII characters
  - Whitespace is not allowed
- **Examples**: 
  - `[name]`
  - `[年齢]`
  - `[user_id]`

### 3. Functions
- **Syntax**: `FUNCTION_NAME(arguments)`
- **Function Name Constraints**:
  - Must be uppercase alphabetic characters only (A-Z)
  - Numbers and symbols are not allowed
- **Arguments**: 
  - Specified within parentheses ()
  - Multiple arguments are separated by commas (,)
- **Examples**: 
  - `SUM([price])`
  - `MAX([score1], [score2])`
  - `CONCAT([first_name], [last_name])`

### 4. Operators

#### 4.1 Arithmetic Operators
- `+` : Addition
- `-` : Subtraction
- `*` : Multiplication
- `/` : Division
- `^` : Exponentiation

#### 4.2 Comparison Operators
- `==` : Equality
- `!=` : Inequality

#### 4.3 Logical Operators
- `||` : Logical OR
- `&&` : Logical AND

### 5. Grouping
- **Syntax**: `(expression)`
- **Description**: Expression grouping with parentheses
- **Purpose**: Explicit control of operation precedence

## Syntax Rules

### Expression
Expressions are composed of combinations of the following elements:

```
expression := literal
           | column_reference
           | function_call
           | '(' expression ')'
           | expression operator expression

literal := string_literal
        | integer_literal
        | float_literal

column_reference := '[' identifier ']'

function_call := FUNCTION_NAME '(' argument_list? ')'

argument_list := expression (',' expression)*

operator := '+' | '-' | '*' | '/' | '^'
         | '==' | '!='
         | '||' | '&&'
```

### Operator Precedence
1. `()` - Parentheses (highest precedence)
2. `^` - Exponentiation
3. `*`, `/` - Multiplication, Division
4. `+`, `-` - Addition, Subtraction
5. `==`, `!=` - Comparison operators
6. `&&` - Logical AND
7. `||` - Logical OR (lowest precedence)

## Usage Examples

### Basic Arithmetic Expressions
```
[price] * [quantity]
[total] / ([count] + 1)
[base] ^ 2
```

### Function Calls
```
SUM([sales])
MAX([score1], [score2], [score3])
CONCAT([first_name], ' ', [last_name])
```

### Complex Expressions
```
([unit_price] * [quantity]) * (1 + [tax_rate])
SUM([revenue]) / COUNT([transactions])
[status] == 'active' && [score] > 80
```

### Logical Expressions
```
[age] >= 18 && [verified] == true
[category] == 'premium' || [points] > 1000
```

## Notes and Constraints

1. **Case Sensitivity**: Function names must be uppercase only
2. **Whitespace**: Whitespace characters are not allowed within column references
3. **Escape Characters**: Quote characters within strings must be escaped with backslash
4. **Operator Associativity**: Operators of the same precedence are left-associative

## Future Extensions

The following features are under consideration for future additions:
- Conditional operator (ternary operator)
- Array/list operations
- Regular expression pattern matching
- Date/time functions
- More built-in functions

## Next Steps

Based on this grammar specification, the following work is planned:

1. Create ANTLR grammar file (.g4)
2. Generate parser and lexer
3. Implement expression editor UI
4. Create unit tests
5. Implement integration tests

## Related Files

- `grammar/grammar.md` - Current simplified grammar definition
- (Planned) `grammar/Expression.g4` - ANTLR grammar file