package app

import (
	"testing"

	"antlr-editor/analyzer/core/app/formatter"

	"github.com/stretchr/testify/assert"
)

func createFormatterWithOptions(options *formatter.FormatOptions) *Formatter {
	if options == nil {
		return newFormatter()
	}
	return NewFormatterWithOptions(options)
}

func TestFormatter_BasicSpacing(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		options  *formatter.FormatOptions
	}{
		{
			name:     "arithmetic operators spacing",
			input:    "[a]+[b]*2",
			expected: "[a] + [b] * 2",
		},
		{
			name:     "comparison operators spacing",
			input:    "[a]<[b]&&[c]>=[d]",
			expected: "[a] < [b] && [c] >= [d]",
		},
		{
			name:     "logical operators spacing",
			input:    "[a]||[b]&&[c]",
			expected: "[a] || [b] && [c]",
		},
		{
			name:     "power operator spacing",
			input:    "[a]^2+[b]^3",
			expected: "[a] ^ 2 + [b] ^ 3",
		},
		{
			name:     "unary minus",
			input:    "-[a]+[b]",
			expected: "-[a] + [b]",
		},
		{
			name:     "equality and inequality",
			input:    "[a]==[b]&&[c]!=[d]",
			expected: "[a] == [b] && [c] != [d]",
		},
		{
			name:     "less than and less than or equal",
			input:    "[a]<[b]&&[c]<=[d]",
			expected: "[a] < [b] && [c] <= [d]",
		},
		{
			name:     "greater than and greater than or equal",
			input:    "[a]>[b]&&[c]>=[d]",
			expected: "[a] > [b] && [c] >= [d]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := createFormatterWithOptions(tt.options)
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatter_FunctionCalls_SingleLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple function call",
			input:    "FUNC([a])",
			expected: "FUNC([a])",
		},
		{
			name:     "function with multiple arguments",
			input:    "FUNC([a],[b],[c])",
			expected: "FUNC([a], [b], [c])",
		},
		{
			name:     "function with mixed argument types",
			input:    "FUNC([column1],[column2],123,\"text\")",
			expected: `FUNC([column1], [column2], 123, "text")`,
		},
		{
			name:     "no space before parenthesis",
			input:    "FUNC  ([a],[b])",
			expected: "FUNC([a], [b])",
		},
		{
			name:     "nested function calls inline",
			input:    "SUM(MIN([a],[b]),MAX([c],[d]),[e])",
			expected: "SUM(MIN([a], [b]), MAX([c], [d]), [e])",
		},
		{
			name:     "function in expression",
			input:    "[a]+FUNC([b],[c])*[d]",
			expected: "[a] + FUNC([b], [c]) * [d]",
		},
		{
			name:     "multiple functions in expression",
			input:    "MAX([a],[b])+MIN([c],[d])",
			expected: "MAX([a], [b]) + MIN([c], [d])",
		},
		{
			name:     "function with no arguments",
			input:    "NOW()",
			expected: "NOW()",
		},
		{
			name:     "function with boolean literals",
			input:    "FUNC(true,false,[column])",
			expected: "FUNC(true, false, [column])",
		},
		{
			name:     "function with negative numbers",
			input:    "FUNC(-1,-2.5,[a])",
			expected: "FUNC(-1, -2.5, [a])",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newFormatter()
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatter_FunctionCalls_MultiLine(t *testing.T) {
	shortLineOpts := &formatter.FormatOptions{
		IndentSize:           2,
		MaxLineLength:        30,
		SpaceAroundOps:       true,
		BreakLongExpressions: true,
	}

	tests := []struct {
		name     string
		input    string
		expected string
		options  *formatter.FormatOptions
	}{
		{
			name:    "long function call with many arguments",
			input:   "FUNC([column1],[column2],[column3],\"long string value\",123)",
			options: shortLineOpts,
			expected: `FUNC(
  [column1],
  [column2],
  [column3],
  "long string value",
  123
)`,
		},
		{
			name:    "nested functions multi-line",
			input:   "FUNC(SUM([a],[b]),MAX([c],[d],[e]),[f])",
			options: shortLineOpts,
			expected: `FUNC(
  SUM([a], [b]),
  MAX([c], [d], [e]),
  [f]
)`,
		},
		{
			name:    "complex nested functions - CALCULATE example",
			input:   "CALC(SUM([sales],[tax]),AVG([price],[discount],[quantity]),FILTER([region],\"APAC\"))",
			options: shortLineOpts,
			expected: `CALC(
  SUM([sales], [tax]),
  AVG(
    [price],
    [discount],
    [quantity]
  ),
  FILTER([region], "APAC")
)`,
		},
		{
			name:    "deeply nested functions",
			input:   "CALC(SUM(MIN([a],[b]),MIN([c],[d])),MAX(MIN([e],[f]),MIN([g],[h])),[i])",
			options: shortLineOpts,
			expected: `CALC(
  SUM(
    MIN([a], [b]),
    MIN([c], [d])
  ),
  MAX(
    MIN([e], [f]),
    MIN([g], [h])
  ),
  [i]
)`,
		},
		{
			name:  "very long function name hanging indent",
			input: "CALCULATE([sales],[tax],[discount],[commission])",
			options: &formatter.FormatOptions{
				IndentSize:           2,
				MaxLineLength:        50,
				SpaceAroundOps:       true,
				BreakLongExpressions: true,
			},
			expected: `CALCULATE(
  [sales],
  [tax],
  [discount],
  [commission]
)`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := createFormatterWithOptions(tt.options)
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatter_FunctionCalls_CommaSpacing(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "one space after commas",
			input:    "FUNC([a],[b],[c])",
			expected: "FUNC([a], [b], [c])",
		},
		{
			name:     "extra spaces after commas removed",
			input:    "FUNC([a],  [b],   [c])",
			expected: "FUNC([a], [b], [c])",
		},
		{
			name:     "no spaces before commas",
			input:    "FUNC([a] ,[b] ,[c])",
			expected: "FUNC([a], [b], [c])",
		},
		{
			name:     "mixed spacing issues",
			input:    "FUNC([a] ,  [b]  ,   [c] )",
			expected: "FUNC([a], [b], [c])",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newFormatter()
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatter_Parentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "remove unnecessary parentheses",
			input:    "((([a]+[b]))*([c]))",
			expected: "((([a] + [b])) * ([c]))",
		},
		{
			name:     "keep necessary parentheses for precedence",
			input:    "([a]+[b])*[c]",
			expected: "([a] + [b]) * [c]",
		},
		{
			name:     "no parentheses needed due to precedence",
			input:    "[a]+[b]*[c]",
			expected: "[a] + [b] * [c]",
		},
		{
			name:     "no spaces inside parentheses",
			input:    "( [a] + [b] )",
			expected: "([a] + [b])",
		},
		{
			name:     "no spaces inside function parentheses",
			input:    "FUNC( [a], [b] )",
			expected: "FUNC([a], [b])",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newFormatter()
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatter_OperatorPrecedence(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "power is highest precedence",
			input:    "[a]+[b]^2*[c]",
			expected: "[a] + [b] ^ 2 * [c]",
		},
		{
			name:     "multiplication before addition",
			input:    "[a]+[b]*[c]",
			expected: "[a] + [b] * [c]",
		},
		{
			name:     "comparison after arithmetic",
			input:    "[a]+[b]>[c]*[d]",
			expected: "[a] + [b] > [c] * [d]",
		},
		{
			name:     "logical AND before OR",
			input:    "[a]||[b]&&[c]",
			expected: "[a] || [b] && [c]",
		},
		{
			name:     "right associative exponentiation",
			input:    "[a]^[b]^[c]",
			expected: "[a] ^ [b] ^ [c]",
		},
		{
			name:     "unary minus highest precedence",
			input:    "-[a]^2",
			expected: "-[a] ^ 2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newFormatter()
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatter_LineBreaking(t *testing.T) {
	shortLineOpts := &formatter.FormatOptions{
		IndentSize:           2,
		MaxLineLength:        10,
		SpaceAroundOps:       true,
		BreakLongExpressions: true,
	}

	tests := []struct {
		name     string
		input    string
		expected string
		options  *formatter.FormatOptions
	}{
		{
			name:  "break at logical operators",
			input: "[column1]+[column2]*3>10&&FUNC([col3],[col4])&&[column6]==\"value\"",
			expected: `[column1] + [column2] * 3 > 10
  && FUNC([col3], [col4])
  && [column6] == "value"`,
		},
		{
			name:  "complex multi-line expression from docs",
			input: "([column1]+[column2])*3>10&&(FUNC([col3],[col4],[col5])||[column6]==\"value\")&&[column7]<100",
			expected: `([column1] + [column2]) * 3 > 10
  && (FUNC([col3], [col4], [col5])
  || [column6] == "value")
  && [column7] < 100`,
		},
		{
			name:    "break before logical operators",
			input:   "[a]>[b]&&[c]<[d]||[e]==[f]",
			options: shortLineOpts,
			expected: `[a]
  > [b]
  && [c]
  < [d]
  || [e]
  == [f]`,
		},
		{
			name:    "break before arithmetic operators",
			input:   "[a]+[b]*[c]-[d]/[e]",
			options: shortLineOpts,
			expected: `[a]
  + [b]
  * [c]
  - [d]
  / [e]`,
		},
		{
			name:    "break before comparison operators",
			input:   "[a]>[b]&&[c]<=[d]",
			options: shortLineOpts,
			expected: `[a]
  > [b]
  && [c]
  <= [d]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := createFormatterWithOptions(tt.options)
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatter_ComplexExpressions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		options  *formatter.FormatOptions
	}{
		{
			name:  "simple expression from docs",
			input: "[col1]+[col2]*3>10&&FUNC([col3],[col4])",
			expected: `[col1] + [col2] * 3 > 10
  && FUNC([col3], [col4])`,
		},
		{
			name:     "mathematical expression from docs",
			input:    "-[a]+[b]^2/([c]-[d])*[e]",
			expected: "-[a] + [b] ^ 2 / ([c] - [d]) * [e]",
		},
		{
			name:  "mixed operators and functions",
			input: "[a]+FUNC([b]*2,[c]/3)>[d]&&[e]!=[f]",
			expected: `[a] + FUNC([b] * 2, [c] / 3) > [d]
  && [e] != [f]`,
		},
		{
			name:  "complex nested with various operators",
			input: "SUM([a]^2,MIN([b]+[c],[d]-[e]))*[f]<=100||[g]==\"test\"",
			expected: `SUM([a] ^ 2, MIN([b] + [c], [d] - [e]))
  * [f] <= 100 || [g] == "test"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := createFormatterWithOptions(tt.options)
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatter_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty expression",
			input:    "",
			expected: "",
		},
		{
			name:     "single column reference",
			input:    "[column]",
			expected: "[column]",
		},
		{
			name:     "single number",
			input:    "123",
			expected: "123",
		},
		{
			name:     "single string",
			input:    "\"text\"",
			expected: "\"text\"",
		},
		{
			name:     "invalid expression - unmatched parenthesis returns original",
			input:    "([a]+[b]",
			expected: "([a]+[b]",
		},
		{
			name:     "invalid expression - syntax error returns original",
			input:    "[a]++[b]",
			expected: "[a]++[b]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newFormatter()
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatter_ConfigurationOptions(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		options  *formatter.FormatOptions
	}{
		{
			name:  "custom indent size 4",
			input: "FUNC([a],[b],[c])",
			options: &formatter.FormatOptions{
				IndentSize:           4,
				MaxLineLength:        20,
				SpaceAroundOps:       true,
				BreakLongExpressions: true,
			},
			expected: `FUNC([a], [b], [c])`,
		},
		{
			name:  "no space around operators",
			input: "[a]+[b]*[c]",
			options: &formatter.FormatOptions{
				IndentSize:           2,
				MaxLineLength:        80,
				SpaceAroundOps:       false,
				BreakLongExpressions: true,
			},
			expected: "[a]+[b]*[c]",
		},
		{
			name:  "no line breaking",
			input: "[column1]+[column2]*3>10&&FUNC([col3],[col4])&&[column6]==\"value\"",
			options: &formatter.FormatOptions{
				IndentSize:           2,
				MaxLineLength:        40,
				SpaceAroundOps:       true,
				BreakLongExpressions: false,
			},
			expected: "[column1] + [column2] * 3 > 10 && FUNC([col3], [col4]) && [column6] == \"value\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := createFormatterWithOptions(tt.options)
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatter_StringAndBooleanLiterals(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "preserve single quotes",
			input:    "'single'==[column]",
			expected: "'single' == [column]",
		},
		{
			name:     "preserve double quotes",
			input:    "\"double\"==[column]",
			expected: "\"double\" == [column]",
		},
		{
			name:     "preserve true/false",
			input:    "true&&false||[column]",
			expected: "true && false || [column]",
		},
		{
			name:     "mixed quotes in function",
			input:    "FUNC('single',\"double\",[column])",
			expected: "FUNC('single', \"double\", [column])",
		},
		{
			name:     "empty strings",
			input:    "FUNC(\"\",'',[column])",
			expected: "FUNC(\"\", '', [column])",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newFormatter()
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatter_ColumnReferences(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "preserve column name casing",
			input:    "[Column_Name]+[another_column]",
			expected: "[Column_Name] + [another_column]",
		},
		{
			name:     "complex column names",
			input:    "[column1]+[column_2]+[ColumnThree]",
			expected: "[column1] + [column_2] + [ColumnThree]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newFormatter()
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestFormatter_NumberLiterals(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "integers",
			input:    "123+456",
			expected: "123 + 456",
		},
		{
			name:     "decimals",
			input:    "1.23+4.56",
			expected: "1.23 + 4.56",
		},
		{
			name:     "scientific notation",
			input:    "1e5+2E-3",
			expected: "1e5 + 2E-3",
		},
		{
			name:     "negative numbers",
			input:    "-123+-4.56",
			expected: "-123 + -4.56",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := newFormatter()
			result := f.Format(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
