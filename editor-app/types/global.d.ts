import { type Token, type Error, type TokenizeResult, type FormatOptions } from './analyzer';

declare global {
  // Go WASM runtime class
  class Go {
    constructor();
    argv: string[];
    env: { [key: string]: string };
    exit: (code: number) => void;
    importObject: WebAssembly.Imports;
    run(instance: WebAssembly.Instance): Promise<void>;
  }

  // Extend Window interface
  interface Window {
    readonly Go: typeof Go;

    validateExpression(expression: string): boolean;
    tokenizeExpression(expression: string): TokenizeResult;
    formatExpression(expression: string): string;
    formatExpressionWithOptions(expression: string, options: FormatOptions): string;
  }
}

export {};
