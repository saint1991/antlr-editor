type TokenType =
  | 'string'
  | 'integer'
  | 'float'
  | 'boolean'
  | 'columnReference'
  | 'function'
  | 'operator'
  | 'comma'
  | 'leftParen'
  | 'rightParen'
  | 'leftBracket'
  | 'rightBracket'
  | 'whitespace'
  | 'error'
  | 'eof';

// NodeType values matching Go analyzer/core/models/node.go
export type NodeType =
  // Root node
  | 0 // Expression
  // Parser Rules - Expression types
  | 1 // LiteralExpr
  | 2 // ColumnRefExpr
  | 3 // FunctionCallExpr
  | 4 // ParenExpr
  | 5 // UnaryMinusExpr
  | 6 // PowerExpr
  | 7 // MulDivExpr
  | 8 // AddSubExpr
  | 9 // ComparisonExpr
  | 10 // AndExpr
  | 11 // OrExpr
  // Parser Rules - Components
  | 12 // Literal
  | 13 // ColumnReference
  | 14 // FunctionCall
  | 15 // ArgumentList
  // Lexer Rules - Operators
  | 16 // Add
  | 17 // Sub
  | 18 // Mul
  | 19 // Div
  | 20 // Pow
  | 21 // Lt
  | 22 // Le
  | 23 // Gt
  | 24 // Ge
  | 25 // Eq
  | 26 // Neq
  | 27 // Or
  | 28 // And
  // Lexer Rules - Delimiters
  | 29 // LParen
  | 30 // RParen
  | 31 // LBracket
  | 32 // RBracket
  | 33 // Comma
  // Lexer Rules - Literals
  | 34 // BooleanLiteral
  | 35 // FloatLiteral
  | 36 // IntegerLiteral
  | 37 // StringLiteral
  // Lexer Rules - Identifiers
  | 38 // FunctionName
  | 39 // ColumnRef
  // Special
  | 40 // WS
  | 41 // ErrorChar
  | 42 // Terminal
  | 43; // Error

export interface Token {
  readonly type: TokenType;
  readonly text: string;
  readonly start: number;
  readonly end: number;
  readonly line: number;
  readonly column: number;
  readonly isValid: boolean;
}

export interface Error {
  readonly message: string;
  readonly line: number;
  readonly column: number;
  readonly start: number;
  readonly end: number;
}

export interface TokenizeResult {
  readonly tokens: Token[];
  readonly errors: Error[];
}

export interface ParseTreeNode {
  readonly type: NodeType;
  readonly text: string;
  readonly start: number;
  readonly end: number;
  readonly children: ParseTreeNode[];
}

export interface ParseTreeResult {
  readonly tree: ParseTreeNode | null;
  readonly errors: Error[];
}

export interface FormatOptions {
  readonly indentSize?: number;
  readonly maxLineLength?: number;
  readonly spaceAroundOps?: boolean;
  readonly breakLongExpressions?: boolean;
}
