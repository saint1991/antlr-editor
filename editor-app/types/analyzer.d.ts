export interface Token {
  readonly type: string;
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

export interface FormatOptions {
  readonly indentSize?: number;
  readonly maxLineLength?: number;
  readonly spaceAroundOps?: boolean;
  readonly breakLongExpressions?: boolean;
}
