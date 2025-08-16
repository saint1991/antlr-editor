# Expression Formatter Rules

This document describes the formatting rules applied by the analyzer's expression formatter.

## Spacing Rules

### Operators
All operators have spaces before and after them:
- Arithmetic: `+`, `-`, `*`, `/`, `^`
- Comparison: `<`, `<=`, `>`, `>=`, `==`, `!=`
- Logical: `&&`, `||`

**Example:**
- Input: `[a]+[b]*2`
- Output: `[a] + [b] * 2`

### Commas
- One space after commas, no space before
- **Example:** `FUNC([a], [b], [c])`

### Parentheses and Brackets
- No spaces inside parentheses: `(expression)` not `( expression )`
- No spaces inside brackets for column references: `[column]` not `[ column ]`
- Function calls have no space before parentheses: `FUNC(...)` not `FUNC (...)`

### Function Call Formatting (Python-style)

#### Basic Rules
- No space between function name and opening parenthesis: `FUNC(` not `FUNC (`
- Space after each comma in arguments: `FUNC(arg1, arg2, arg3)`
- No trailing comma after the last argument (unless multi-line)
- Closing parenthesis immediately follows last argument or is on its own line

#### Single-line Function Calls
- Keep on one line if under the maximum line length
- **Example:** `FUNC([column1], [column2], 123, "text")`

#### Multi-line Function Calls
When function calls exceed the maximum line length or have many arguments:
All arguments on separate lines:
```
FUNC(
  [column1],
  [column2],
  [column3],
  "long string value",
  123
)
```

#### Nested Function Calls
- Format inner functions first, then outer
- Break at the outermost level first when line is too long
- **Example:**
  ```
  OUTER(
    INNER1([a], [b]),
    INNER2([c], [d], [e]),
    [f]
  )
  ```

#### Alignment Rules
- Arguments aligned with the opening parenthesis or indented one level
- Hanging indent
  ```
  VERY_LONG_FUNCTION_NAME(
      first_argument, second_argument,
      third_argument, fourth_argument
  )
  ```

## Operator Precedence and Parentheses

### Precedence Order (highest to lowest)
1. Unary minus: `-expression`
2. Power: `^` (right-associative)
3. Multiplication/Division: `*`, `/`
4. Addition/Subtraction: `+`, `-`
5. Comparison: `<`, `<=`, `>`, `>=`, `==`, `!=`
6. Logical AND: `&&`
7. Logical OR: `||`

### Parentheses Rules
- **Remove unnecessary parentheses** when precedence is clear
- **Keep parentheses** that change evaluation order or improve readability
- **Add parentheses** to clarify complex expressions

**Examples:**
- Input: `((([a] + [b])) * ([c]))`
- Output: `([a] + [b]) * [c]`

- Input: `[a] + [b] * [c]`
- Output: `[a] + [b] * [c]` (no parentheses needed due to precedence)

## Line Breaking and Indentation

### Line Length
- Default maximum line length: 80 characters
- Break at lower-precedence operators first
- Logical operators (`&&`, `||`) are preferred break points

### Break Rules
1. **All binary operators except power operator**: Can break before the operator when line is too long
2. **Function arguments**: Can break after each comma for long argument lists

### Indentation
- Default indent size: 2 spaces (configurable to 4)
- Continuation lines are indented one level
- Nested expressions maintain proper indentation hierarchy

**Example:**
```
[column1] + [column2] * 3 > 10
  && FUNC([col3], [col4], [col5])
  && [column6] == "value"
```

## Special Formatting

### Function Names
- Function names remain in uppercase: `FUNC`, `MAX`, `MIN`
- No changes to function name casing

### Literals
- **Strings**: Preserve original quotes (single or double)
- **Numbers**: Keep as-is (future: option to normalize decimals)
- **Booleans**: Preserve `true`/`false` exactly

### Column References
- Always wrapped in brackets: `[column_name]`
- No spaces inside brackets
- Preserve original column name casing

## Complex Expression Examples

### Example 1: Simple Expression
**Input:**
```
[col1]+[col2]*3>10&&FUNC([col3],[col4])
```
**Output:**
```
[col1] + [col2] * 3 > 10 && FUNC([col3], [col4])
```

### Example 2: Multi-line Expression
**Input:**
```
([column1]+[column2])*3>10&&(FUNC([col3],[col4],[col5])||[column6]=="value")&&[column7]<100
```
**Output:**
```
([column1] + [column2]) * 3 > 10
  && (FUNC([col3], [col4], [col5]) || [column6] == "value")
  && [column7] < 100
```

### Example 3: Nested Functions
**Input:**
```
FUNC1(FUNC2([a],[b]),FUNC3([c],[d]),[e])
```
**Output (single-line):**
```
FUNC1(FUNC2([a], [b]), FUNC3([c], [d]), [e])
```

**Output (multi-line - when exceeds max length):**
```
FUNC1(
  FUNC2([a], [b]),
  FUNC3([c], [d]),
  [e]
)
```

### Example 3b: Complex Nested Functions
**Input:**
```
CALCULATE(SUM([sales],[tax]),AVERAGE([price],[discount],[quantity]),FILTER([region],"APAC"))
```
**Output (multi-line with grouped arguments):**
```
CALCULATE(
  SUM([sales], [tax]),
  AVERAGE([price], [discount], [quantity]),
  FILTER([region], "APAC")
)
```

### Example 3c: Deeply Nested Functions
**Input:**
```
OUTER(MIDDLE1(INNER([a],[b]),INNER([c],[d])),MIDDLE2(INNER([e],[f]),INNER([g],[h])),[i])
```
**Output:**
```
OUTER(
  MIDDLE1(
    INNER([a], [b]),
    INNER([c], [d])
  ),
  MIDDLE2(
    INNER([e], [f]),
    INNER([g], [h])
  ),
  [i]
)
```

### Example 4: Mathematical Expression
**Input:**
```
-[a]+[b]^2/([c]-[d])*[e]
```
**Output:**
```
-[a] + [b] ^ 2 / ([c] - [d]) * [e]
```

## Configuration Options

The formatter supports the following configuration options:

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `IndentSize` | int | 2 | Number of spaces per indent level |
| `MaxLineLength` | int | 80 | Maximum line length before breaking |
| `SpaceAroundOps` | bool | true | Add spaces around operators |
| `UppercaseFunctions` | bool | true | Keep function names uppercase |
| `RemoveUnnecessaryParens` | bool | true | Remove redundant parentheses |
| `BreakLongExpressions` | bool | true | Auto-break long expressions |
| `AlignOperators` | bool | false | Vertically align operators (future) |

## Edge Cases

### Empty Expression
- Input: `""`
- Output: `""` (unchanged)

### Single Token
- Input: `[column]`
- Output: `[column]` (unchanged)

### Invalid Expression
- Expressions with syntax errors are not formatted
- Original expression is returned unchanged

## Future Enhancements

1. **Comment preservation**: Maintain comments in formatted output
2. **Vertical alignment**: Align operators in multi-line expressions
3. **Custom break rules**: User-defined line break preferences
4. **Style presets**: SQL-style, Math-style, Compact-style options