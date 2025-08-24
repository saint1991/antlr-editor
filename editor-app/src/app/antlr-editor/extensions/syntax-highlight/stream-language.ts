import { HighlightStyle, LanguageSupport, StreamLanguage, syntaxHighlighting } from '@codemirror/language';
import { tags } from '@lezer/highlight';
import type { Analyzer, Token, TokenType } from '../analyzer';

const tokenTypeMapping: Record<TokenType, string | null> = {
  string: 'string',
  integer: 'number',
  float: 'number',
  boolean: 'bool',
  columnReference: 'variableName',
  function: 'keyword',
  operator: 'operator',
  comma: 'punctuation',
  leftParen: 'paren',
  rightParen: 'paren',
  leftBracket: 'variableName',
  rightBracket: 'variableName',
  whitespace: null,
  error: 'invalid',
  eof: null,
};

// Create the language definition using StreamLanguage
export const createExpressionLanguage = (analyzer: Analyzer) => {
  return StreamLanguage.define<{
    tokens: Token[];
    errors: Array<{ start: number; end: number }>;
    currentIndex: number;
    text: string;
  }>({
    name: 'expression',
    startState: () => ({
      tokens: [],
      errors: [],
      currentIndex: 0,
      text: '',
    }),
    token: (stream, state) => {
      // If at start of line or text changed, tokenize the entire line
      if (stream.pos === 0 || state.text !== stream.string) {
        state.text = stream.string;
        const result = analyzer.tokenizeExpression(stream.string);
        state.tokens = result.tokens;
        state.errors = result.errors || [];
        state.currentIndex = 0;
      }

      // Skip whitespace
      if (stream.eatSpace()) {
        return null;
      }

      const pos = stream.pos;

      // Check if current position is within an error range
      for (const error of state.errors) {
        if (pos >= error.start && pos < error.end) {
          // Move stream position forward by one character within error range
          stream.next();
          return 'invalid';
        }
      }

      // Find token at current position
      for (const token of state.tokens) {
        if (token.start === pos) {
          // Move stream position to end of token
          stream.pos = token.end;

          return tokenTypeMapping[token.type];
        }
      }

      // No token found, advance one character
      stream.next();
      return null;
    },
  });
};

// Define custom highlight style for the Expression language
export const expressionHighlightStyle = HighlightStyle.define([
  { tag: tags.string, color: '#a31515' }, // Red for strings
  { tag: tags.number, color: '#098658' }, // Green for numbers
  { tag: tags.keyword, color: '#795e26', fontWeight: 'bold' }, // Brown bold for function names
  { tag: tags.variableName, color: '#001080' }, // Dark blue for identifiers
  { tag: tags.bracket, color: '#000000' }, // Black for brackets
  { tag: tags.operator, color: '#000000' }, // Black for operators
  { tag: tags.punctuation, color: '#000000' }, // Black for punctuation
  { tag: tags.invalid, color: '#ff0000', textDecoration: 'underline wavy' }, // Red with wavy underline for errors
]);

// Create the complete language support extension
export const expressionLanguage = (analyzer: Analyzer): LanguageSupport => {
  const language = createExpressionLanguage(analyzer);
  return new LanguageSupport(language, [syntaxHighlighting(expressionHighlightStyle)]);
};
