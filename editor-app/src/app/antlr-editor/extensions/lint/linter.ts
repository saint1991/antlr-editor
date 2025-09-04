import { type Diagnostic, linter } from '@codemirror/lint';
import type { EditorView } from '@codemirror/view';
import type { Analyzer, Error as AnalyzerError } from '../../../../wasm/analyzer';

// Convert analyzer errors to CodeMirror diagnostics
const errorsToDiagnostics = (errors: AnalyzerError[]): Diagnostic[] => {
  return errors.map((error) => ({
    from: error.start,
    to: error.end,
    severity: 'error' as const,
    message: error.message,
    renderMessage: () => {
      const elem = document.createElement('div');
      elem.className = 'cm-lint-message-error';
      elem.textContent = error.message;

      // Add line and column info
      const posInfo = document.createElement('div');
      posInfo.className = 'cm-lint-message-pos';
      posInfo.style.fontSize = '0.9em';
      posInfo.style.opacity = '0.7';
      posInfo.textContent = `Line ${error.line}, Column ${error.column}`;
      elem.appendChild(posInfo);

      return elem;
    },
  }));
};

// Create the linter extension
export const expressionLinter = (analyzer: Analyzer) => {
  return linter(
    (view: EditorView) => {
      const text = view.state.doc.toString();

      // Skip empty text
      if (!text.trim()) {
        return [];
      }

      // Parse the expression
      const errors = analyzer.lint(text);

      // Convert errors to diagnostics
      return errorsToDiagnostics(errors);
    },
    {
      delay: 500, // Debounce delay in milliseconds
    },
  );
};
