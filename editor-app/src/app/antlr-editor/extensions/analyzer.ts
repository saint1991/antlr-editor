import type { FormatOptions, TokenizeResult } from '@wasm-analyzer';

export type { Error, FormatOptions, Token, TokenizeResult, TokenType } from '@wasm-analyzer';

const wasmModuleUrl = '/analyzer.wasm';
export interface Analyzer {
  validateExpression(expression: string): boolean;
  tokenizeExpression(expression: string): TokenizeResult;
  formatExpression(expression: string): string;
  formatExpressionWithOptions(expression: string, options: FormatOptions): string;
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
    validateExpression: window.validateExpression,
    tokenizeExpression: window.tokenizeExpression,
    formatExpression: window.formatExpression,
    formatExpressionWithOptions: window.formatExpressionWithOptions,
  };

  return instance;
};
