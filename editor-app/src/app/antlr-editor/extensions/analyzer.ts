import { type AnalyzeResult, type FormatOptions } from '../../../../types/analyzer';

const wasmModuleUrl = '/analyzer.wasm';

export interface Analyzer {
  validateExpression(expression: string): boolean;
  analyzeExpression(expression: string): AnalyzeResult;
  formatExpression(expression: string): string;
  formatExpressionWithOptions(expression: string, options: FormatOptions): string;
}


let instance: Analyzer | null = null;

export const loadAnalyzer = async (): Promise<Analyzer> => {
  if (instance) {
    return instance;
  }

  const go = new Go();
  await WebAssembly.instantiateStreaming(fetch(wasmModuleUrl), go.importObject).then(result => {
    go.run(result.instance)
  });

  return {
    validateExpression: window.validateExpression,
    analyzeExpression: window.analyzeExpression,
    formatExpression: window.formatExpression,
    formatExpressionWithOptions: window.formatExpressionWithOptions
  }
};
