import { LanguageSupport, LRLanguage } from '@codemirror/language';
import { NodeProp } from '@lezer/common';
import { styleTags, tags } from '@lezer/highlight';
import type { Analyzer } from '../../../../wasm/analyzer';
import { ExpressionParser, nodeSet } from './parser';

// Create the expression language with proper syntax highlighting
export const expressionLanguageSupport = (analyzer: Analyzer): LanguageSupport => {
  // Configure the parser with syntax highlighting
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
    ],
  });

  // Create LRLanguage with the configured parser
  const language = LRLanguage.define({
    name: 'expression',
    parser: parserWithMetadata,
    languageData: {
      commentTokens: { line: '//' },
      indentOnInput: /^\s*\)$/,
    },
  });

  return new LanguageSupport(language, []);
};
