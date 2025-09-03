import type { Error as AnalyzerError, FormatOptions, ParseTreeResult, TokenizeResult } from '@wasm-analyzer';

export type { Error, FormatOptions, ParseTreeNode, ParseTreeResult, Token, TokenizeResult, TokenType } from '@wasm-analyzer';

// NodeType constants matching Go analyzer/core/models/node.go
export const NodeType = {
  // Root node
  Expression: 0,

  // Parser Rules - Expression types
  LiteralExpr: 1,
  ColumnRefExpr: 2,
  FunctionCallExpr: 3,
  ParenExpr: 4,
  UnaryMinusExpr: 5,
  PowerExpr: 6,
  MulDivExpr: 7,
  AddSubExpr: 8,
  ComparisonExpr: 9,
  AndExpr: 10,
  OrExpr: 11,

  // Parser Rules - Components
  Literal: 12,
  ColumnReference: 13,
  FunctionCall: 14,
  ArgumentList: 15,

  // Lexer Rules - Operators
  Add: 16,
  Sub: 17,
  Mul: 18,
  Div: 19,
  Pow: 20,
  Lt: 21,
  Le: 22,
  Gt: 23,
  Ge: 24,
  Eq: 25,
  Neq: 26,
  Or: 27,
  And: 28,

  // Lexer Rules - Delimiters
  LParen: 29,
  RParen: 30,
  LBracket: 31,
  RBracket: 32,
  Comma: 33,

  // Lexer Rules - Literals
  BooleanLiteral: 34,
  FloatLiteral: 35,
  IntegerLiteral: 36,
  StringLiteral: 37,

  // Lexer Rules - Identifiers
  FunctionName: 38,
  ColumnRef: 39,

  // Special
  WS: 40,
  ErrorChar: 41,
  Terminal: 42,
  Error: 43,
} as const;

const wasmModuleUrl = '/analyzer.wasm';
export interface Analyzer {
  parseTree: (expression: string) => ParseTreeResult;
  lint: (expression: string) => AnalyzerError[];
  tokenize: (expression: string) => TokenizeResult;
  validate: (expression: string) => boolean;
  format: (expression: string) => string;
  formatWithOptions: (expression: string, options?: FormatOptions) => string;
}

let instance: Analyzer | null = null;

export const loadAnalyzer = async (): Promise<Analyzer> => {
  if (instance) {
    return instance;
  }

  const go = new Go();
  await WebAssembly.instantiateStreaming(fetch(wasmModuleUrl), go.importObject).then((result) => {
    go.run(result.instance);
  });

  instance = {
    parseTree: window.parseTree,
    lint: window.lint,
    tokenize: window.tokenize,
    validate: window.validate,
    format: window.format,
    formatWithOptions: window.formatWithOptions,
  };

  return instance;
};
