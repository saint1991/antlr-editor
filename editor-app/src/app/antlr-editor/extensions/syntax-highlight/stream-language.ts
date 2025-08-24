import { HighlightStyle, LanguageSupport, StreamLanguage, type StringStream, syntaxHighlighting } from '@codemirror/language';
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

interface State {
  tokens: Token[];
  errors: Array<{ start: number; end: number }>;
  currentIndex: number;
  text: string;
}

// Create the language definition using StreamLanguage
export const createExpressionLanguage = (analyzer: Analyzer) => {
  return StreamLanguage.define<State>({
    name: 'expression',
    startState: () => ({
      tokens: [],
      errors: [],
      currentIndex: 0,
      text: '',
    }),
    token: (stream: StringStream, state: State) => {
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

// Define light theme highlight style
export const lightThemeHighlightStyle = HighlightStyle.define([
  { tag: tags.string, color: '#a31515' }, // Red for strings
  { tag: tags.number, color: '#098658' }, // Green for numbers
  { tag: tags.keyword, color: '#795e26', fontWeight: 'bold' }, // Brown bold for function names
  { tag: tags.variableName, color: '#001080' }, // Dark blue for identifiers
  { tag: tags.bracket, color: '#000000' }, // Black for brackets
  { tag: tags.operator, color: '#000000' }, // Black for operators
  { tag: tags.punctuation, color: '#000000' }, // Black for punctuation
  { tag: tags.invalid, color: '#ff0000', textDecoration: 'underline wavy' }, // Red with wavy underline for errors
]);

// Define dark theme highlight style
export const darkThemeHighlightStyle = HighlightStyle.define([
  { tag: tags.string, color: '#ce9178' }, // Light orange for strings
  { tag: tags.number, color: '#b5cea8' }, // Light green for numbers
  { tag: tags.keyword, color: '#c586c0', fontWeight: 'bold' }, // Pink bold for function names
  { tag: tags.variableName, color: '#9cdcfe' }, // Light blue for identifiers
  { tag: tags.bracket, color: '#d4d4d4' }, // Light gray for brackets
  { tag: tags.operator, color: '#d4d4d4' }, // Light gray for operators
  { tag: tags.punctuation, color: '#d4d4d4' }, // Light gray for punctuation
  { tag: tags.invalid, color: '#f48771', textDecoration: 'underline wavy' }, // Light red with wavy underline for errors
]);

// Create the complete language support extension
export const expressionLanguage = (analyzer: Analyzer, theme: 'light' | 'dark' = 'light'): LanguageSupport => {
  const language = createExpressionLanguage(analyzer);
  const highlightStyle = theme === 'dark' ? darkThemeHighlightStyle : lightThemeHighlightStyle;
  return new LanguageSupport(language, [syntaxHighlighting(highlightStyle)]);
};
