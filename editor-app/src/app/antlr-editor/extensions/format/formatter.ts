import { Transaction } from '@codemirror/state';
import { type EditorView, keymap } from '@codemirror/view';
import type { Analyzer } from '../../../../wasm/analyzer';

export const formatExpression = (analyzer: Analyzer) => (view: EditorView) => {
  const currentDoc = view.state.doc.toString();

  try {
    const formatted = analyzer.format(currentDoc);

    if (formatted === currentDoc) {
      return false;
    }

    const transaction = view.state.update({
      changes: {
        from: 0,
        to: view.state.doc.length,
        insert: formatted,
      },
      annotations: Transaction.userEvent.of('format'),
    });

    view.dispatch(transaction);
    return true;
  } catch (error) {
    console.error('Failed to format expression:', error);
    return false;
  }
};

export const formatKeymap = (analyzer: Analyzer) =>
  keymap.of([
    { key: 'Shift-Alt-f', run: formatExpression(analyzer) },
    { key: 'Shift-Cmd-f', mac: 'Shift-Cmd-f', run: formatExpression(analyzer) },
  ]);
