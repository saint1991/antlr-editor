import { delimitedIndent, getIndentUnit, indentNodeProp, indentService, LanguageSupport, LRLanguage } from '@codemirror/language';
import { NodeProp } from '@lezer/common';
import { styleTags, tags } from '@lezer/highlight';
import type { Analyzer } from '../../../../wasm/analyzer';
import { ExpressionParser, nodeSet } from './parser';

// Create the expression language with proper syntax highlighting
export const expressionLanguageSupport = (analyzer: Analyzer): LanguageSupport => {
  // Configure the parser with syntax highlighting and indentation
  const parserWithMetadata = new ExpressionParser(analyzer).configure({
    props: [
      styleTags({
        StringLiteral: tags.string,
        IntegerLiteral: tags.number,
        FloatLiteral: tags.number,
        BooleanLiteral: tags.bool,
        FunctionName: tags.function(tags.variableName),
        ColumnReference: tags.variableName,
        Error: tags.invalid,
        Literal: tags.literal,
        FunctionCall: tags.function(tags.name),
        FunctionCallExpr: tags.function(tags.name),
        Expression: tags.content,
      }),
      // Add indentation rules for function calls
      indentNodeProp.add({
        FunctionCall: delimitedIndent({ closing: ')', align: false }),
        FunctionCallExpr: delimitedIndent({ closing: ')', align: false }),
      }),
    ],
  });

  // Create LRLanguage with the configured parser
  const language = LRLanguage.define({
    name: 'expression',
    parser: parserWithMetadata,
    languageData: {
      commentTokens: { line: '//' },
      // Re-indent when typing closing brackets
      indentOnInput: /^\s*\)$/,
    },
  });

  // Custom indentation service
  const customIndentService = indentService.of((context, pos) => {
    const prevLine = pos > 0 ? context.lineAt(pos - 1) : null;

    if (prevLine) {
      const prevText = prevLine.text.trimEnd();
      const prevIndent = prevLine.text.match(/^\s*/)?.[0] || '';

      // If previous line ends with opening bracket, add indentation
      if (prevText.endsWith('(')) {
        const indentUnit = getIndentUnit(context.state);
        return prevIndent.length + indentUnit;
      }

      // Otherwise maintain current indentation
      return prevIndent.length;
    }

    return 0;
  });

  return new LanguageSupport(language, [customIndentService]);
};
