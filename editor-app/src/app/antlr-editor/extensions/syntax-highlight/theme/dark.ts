import { HighlightStyle } from '@codemirror/language';
import { tags } from '@lezer/highlight';

// Custom highlight style optimized for dark themes
export const darkHighlightStyle = HighlightStyle.define([
  // Strings - Soft green
  { tag: tags.string, color: '#98c379' },

  // Numbers - Light blue
  { tag: tags.number, color: '#61afef' },

  // Booleans - Orange
  { tag: tags.bool, color: '#d19a66' },

  // Functions - Yellow
  { tag: tags.function(tags.variableName), color: '#e5c07b' },

  // Variables/Column references - Light red
  { tag: tags.variableName, color: '#e06c75' },

  // Operators - Cyan
  { tag: tags.operator, color: '#56b6c2' },

  // Punctuation - Gray
  { tag: tags.punctuation, color: '#abb2bf' },

  // Comments - Muted gray (if we add them later)
  { tag: tags.comment, color: '#5c6370', fontStyle: 'italic' },

  // Keywords - Purple
  { tag: tags.keyword, color: '#c678dd' },

  // Errors - Red with underline
  { tag: tags.invalid, color: '#ff6b6b', textDecoration: 'underline wavy #ff6b6b' },
]);
