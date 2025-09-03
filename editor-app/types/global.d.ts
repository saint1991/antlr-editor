import type { Error as AnalyzerError, TokenizeResult, ParseTreeResult, FormatOptions } from './analyzer';

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

    parseTree: (expression: string) => ParseTreeResult;
    lint: (expression: string) => AnalyzerError[];
    tokenize: (expression: string) => TokenizeResult;
    validate: (expression: string) => boolean;
    format: (expression: string) => string;
    formatWithOptions: (expression: string, options?: FormatOptions) => string;
  }
}
