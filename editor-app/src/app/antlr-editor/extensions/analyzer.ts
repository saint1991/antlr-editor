import type { FormatOptions, TokenizeResult } from '../../../../types/analyzer';

const wasmModuleUrl = '/analyzer.wasm';

export interface Analyzer {
  validateExpression(expression: string): boolean;
  tokenizeExpression(expression: string): TokenizeResult;
  formatExpression(expression: string): string;
  formatExpressionWithOptions(expression: string, options: FormatOptions): string;
}

const instance: Analyzer | null = null;

export const loadAnalyzer = async (): Promise<Analyzer> => {
  if (instance) {
    return instance;
  }

  const go = new Go();
  await WebAssembly.instantiateStreaming(fetch(wasmModuleUrl), go.importObject).then((result) => {
    go.run(result.instance);
  });

  return {
    validateExpression: window.validateExpression,
    tokenizeExpression: window.tokenizeExpression,
    formatExpression: window.formatExpression,
    formatExpressionWithOptions: window.formatExpressionWithOptions,
  };
};
