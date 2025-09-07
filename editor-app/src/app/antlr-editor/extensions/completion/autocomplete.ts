import { autocompletion, type Completion, type CompletionContext, type CompletionResult } from '@codemirror/autocomplete';
import type { FunctionDescription } from '../function';

// Completion function factory
const createExpressionCompletions = (functionDescriptions: Record<string, FunctionDescription>) => {
  return (context: CompletionContext): CompletionResult | null => {
    const word = context.matchBefore(/\w*/);
    if (!word) {
      return null;
    }

    // Don't show completions in the middle of a word
    if (word.from === word.to && !context.explicit) {
      return null;
    }

    // Convert function descriptions to completion options
    const functionCompletions: Completion[] = Object.values(functionDescriptions).map((desc) => ({
      label: desc.name,
      type: desc.type || 'function',
      detail: desc.detail || desc.syntax,
      info: desc.info || desc.description,
      apply: (view, completion, from, to) => {
        let insert = completion.label;

        // Add parentheses for functions
        if (completion.type === 'function') {
          insert += '()';
          // Position cursor between parentheses
          view.dispatch({
            changes: { from, to, insert },
            selection: { anchor: from + insert.length - 1 },
          });
        } else {
          view.dispatch({
            changes: { from, to, insert },
          });
        }
      },
    }));

    return {
      from: word.from,
      options: functionCompletions,
      validFor: /^\w*$/,
    };
  };
};

// Apply styles to match hover tooltip design
const applyAutocompleteStyles = () => {
  // Check if styles are already applied
  if (document.getElementById('autocomplete-custom-styles')) {
    return;
  }

  const styleSheet = document.createElement('style');
  styleSheet.id = 'autocomplete-custom-styles';
  styleSheet.textContent = `
    .cm-tooltip.cm-tooltip-autocomplete {
      background: #2d3748 !important;
      border: 1px solid rgba(99, 179, 237, 0.2) !important;
      border-radius: 8px !important;
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3) !important;
      padding: 4px !important;
      max-height: 300px !important;
      overflow-y: auto !important;
    }

    .cm-completionIcon {
      width: 14px !important;
      height: 14px !important;
      margin-right: 8px !important;
      border-radius: 50% !important;
      display: inline-flex !important;
      align-items: center !important;
      justify-content: center !important;
      font-size: 8px !important;
      font-weight: 500 !important;
    }

    .cm-completionIcon-function {
      background: rgba(99, 179, 237, 0.2) !important;
      color: #63b3ed !important;
      border: 1px solid rgba(99, 179, 237, 0.3) !important;
    }

    .cm-completionIcon-constant {
      background: rgba(104, 211, 145, 0.2) !important;
      color: #68d391 !important;
      border: 1px solid rgba(104, 211, 145, 0.3) !important;
    }

    .cm-completionIcon-keyword {
      background: rgba(237, 137, 54, 0.2) !important;
      color: #ed8936 !important;
      border: 1px solid rgba(237, 137, 54, 0.3) !important;
    }

    .cm-completionIcon-function::before {
      content: "ƒ" !important;
    }

    .cm-completionIcon-constant::before {
      content: "•" !important;
    }

    .cm-completionIcon-keyword::before {
      content: "⚬" !important;
    }

    .cm-completion {
      padding: 8px 12px !important;
      margin: 2px !important;
      border-radius: 6px !important;
      display: flex !important;
      align-items: center !important;
      color: #e2e8f0 !important;
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', sans-serif !important;
      font-size: 13px !important;
      line-height: 1.5 !important;
      cursor: pointer !important;
      transition: all 0.15s ease !important;
      background: transparent !important;
      border: 1px solid transparent !important;
    }

    .cm-completion:hover,
    .cm-completion[aria-selected] {
      background: rgba(99, 179, 237, 0.15) !important;
      border: 1px solid rgba(99, 179, 237, 0.3) !important;
      color: #ffffff !important;
      transform: translateX(1px) !important;
    }

    .cm-completionLabel {
      flex-grow: 1 !important;
      font-weight: 400 !important;
      color: inherit !important;
    }

    .cm-completionDetail {
      margin-left: 8px !important;
      opacity: 0.8 !important;
      font-family: 'SF Mono', Monaco, 'Cascadia Code', 'Roboto Mono', Consolas, 'Courier New', monospace !important;
      font-size: 11px !important;
      color: #68d391 !important;
      background: rgba(104, 211, 145, 0.1) !important;
      padding: 2px 6px !important;
      border-radius: 3px !important;
      border: 1px solid rgba(104, 211, 145, 0.3) !important;
    }

    .cm-tooltip.cm-completionInfo {
      background: #2d3748 !important;
      border: 1px solid rgba(99, 179, 237, 0.2) !important;
      border-radius: 8px !important;
      padding: 16px !important;
      max-width: 450px !important;
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3) !important;
      color: #e2e8f0 !important;
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', sans-serif !important;
      font-size: 13px !important;
      line-height: 1.5 !important;
    }

    .cm-completionInfo p {
      margin: 0 !important;
      color: #e2e8f0 !important;
    }

    /* Custom scrollbar for autocomplete */
    .cm-tooltip.cm-tooltip-autocomplete::-webkit-scrollbar {
      width: 6px !important;
    }

    .cm-tooltip.cm-tooltip-autocomplete::-webkit-scrollbar-track {
      background: rgba(0, 0, 0, 0.2) !important;
      border-radius: 3px !important;
    }

    .cm-tooltip.cm-tooltip-autocomplete::-webkit-scrollbar-thumb {
      background: rgba(255, 255, 255, 0.2) !important;
      border-radius: 3px !important;
    }

    .cm-tooltip.cm-tooltip-autocomplete::-webkit-scrollbar-thumb:hover {
      background: rgba(255, 255, 255, 0.3) !important;
    }
  `;
  document.head.appendChild(styleSheet);
};

// Create the autocompletion extension
export const expressionAutocompletion = (functionDescriptions: Record<string, FunctionDescription>) => {
  // Apply custom styles
  applyAutocompleteStyles();

  return autocompletion({
    override: [createExpressionCompletions(functionDescriptions)],
    defaultKeymap: true,
    closeOnBlur: true,
    icons: true,
    optionClass: (completion) => `cm-completion-${completion.type}`,
    tooltipClass: () => 'cm-autocomplete-tooltip',
  });
};
