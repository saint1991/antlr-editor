import { autocompletion, type CompletionContext, type CompletionResult } from '@codemirror/autocomplete';

// Define function completions with descriptions
const functionCompletions = [
  {
    label: 'UPPER',
    type: 'function',
    detail: '(text) → text',
    info: 'Converts text to uppercase.\nExample: UPPER("hello") → "HELLO"',
  },
  {
    label: 'LOWER',
    type: 'function',
    detail: '(text) → text',
    info: 'Converts text to lowercase.\nExample: LOWER("HELLO") → "hello"',
  },
  {
    label: 'TRIM',
    type: 'function',
    detail: '(text) → text',
    info: 'Removes leading and trailing whitespace.\nExample: TRIM("  hello  ") → "hello"',
  },
  {
    label: 'LENGTH',
    type: 'function',
    detail: '(text) → number',
    info: 'Returns the length of a string.\nExample: LENGTH("hello") → 5',
  },
  {
    label: 'CONCAT',
    type: 'function',
    detail: '(text1, text2, ...) → text',
    info: 'Concatenates multiple strings.\nExample: CONCAT("hello", " ", "world") → "hello world"',
  },
  {
    label: 'SUBSTRING',
    type: 'function',
    detail: '(text, start, length) → text',
    info: 'Extracts a substring.\nExample: SUBSTRING("hello", 2, 3) → "llo"',
  },
  {
    label: 'REPLACE',
    type: 'function',
    detail: '(text, search, replace) → text',
    info: 'Replaces occurrences of search with replace.\nExample: REPLACE("hello", "l", "r") → "herro"',
  },
  {
    label: 'IF',
    type: 'function',
    detail: '(condition, true_value, false_value) → any',
    info: 'Conditional expression.\nExample: IF([age] > 18, "adult", "minor")',
  },
  {
    label: 'CASE',
    type: 'function',
    detail: 'CASE WHEN condition THEN value ... END',
    info: 'Multi-condition expression.\nExample: CASE WHEN [score] > 90 THEN "A" WHEN [score] > 80 THEN "B" ELSE "C" END',
  },
  {
    label: 'COALESCE',
    type: 'function',
    detail: '(value1, value2, ...) → any',
    info: 'Returns the first non-null value.\nExample: COALESCE([field1], [field2], "default")',
  },
  {
    label: 'ROUND',
    type: 'function',
    detail: '(number, decimals?) → number',
    info: 'Rounds a number to specified decimals.\nExample: ROUND(3.14159, 2) → 3.14',
  },
  {
    label: 'FLOOR',
    type: 'function',
    detail: '(number) → number',
    info: 'Rounds down to the nearest integer.\nExample: FLOOR(3.7) → 3',
  },
  {
    label: 'CEIL',
    type: 'function',
    detail: '(number) → number',
    info: 'Rounds up to the nearest integer.\nExample: CEIL(3.2) → 4',
  },
  {
    label: 'ABS',
    type: 'function',
    detail: '(number) → number',
    info: 'Returns the absolute value.\nExample: ABS(-5) → 5',
  },
  {
    label: 'MIN',
    type: 'function',
    detail: '(value1, value2, ...) → any',
    info: 'Returns the minimum value.\nExample: MIN(1, 2, 3) → 1',
  },
  {
    label: 'MAX',
    type: 'function',
    detail: '(value1, value2, ...) → any',
    info: 'Returns the maximum value.\nExample: MAX(1, 2, 3) → 3',
  },
  {
    label: 'SUM',
    type: 'function',
    detail: '(number1, number2, ...) → number',
    info: 'Returns the sum of values.\nExample: SUM(1, 2, 3) → 6',
  },
  {
    label: 'AVG',
    type: 'function',
    detail: '(number1, number2, ...) → number',
    info: 'Returns the average of values.\nExample: AVG(1, 2, 3) → 2',
  },
  {
    label: 'COUNT',
    type: 'function',
    detail: '(value1, value2, ...) → number',
    info: 'Counts non-null values.\nExample: COUNT([field1], [field2])',
  },
  {
    label: 'NOW',
    type: 'function',
    detail: '() → datetime',
    info: 'Returns the current date and time.\nExample: NOW()',
  },
  {
    label: 'DATE',
    type: 'function',
    detail: '(datetime) → date',
    info: 'Extracts the date part.\nExample: DATE(NOW())',
  },
  {
    label: 'YEAR',
    type: 'function',
    detail: '(datetime) → number',
    info: 'Extracts the year.\nExample: YEAR(NOW()) → 2024',
  },
  {
    label: 'MONTH',
    type: 'function',
    detail: '(datetime) → number',
    info: 'Extracts the month (1-12).\nExample: MONTH(NOW())',
  },
  {
    label: 'DAY',
    type: 'function',
    detail: '(datetime) → number',
    info: 'Extracts the day of month.\nExample: DAY(NOW())',
  },
];

// Boolean literals
const booleanCompletions = [
  {
    label: 'true',
    type: 'constant',
    detail: 'boolean',
    info: 'Boolean true value',
  },
  {
    label: 'false',
    type: 'constant',
    detail: 'boolean',
    info: 'Boolean false value',
  },
];

// Operator completions
const operatorCompletions = [
  {
    label: 'AND',
    type: 'keyword',
    detail: 'logical operator',
    info: 'Logical AND operator',
  },
  {
    label: 'OR',
    type: 'keyword',
    detail: 'logical operator',
    info: 'Logical OR operator',
  },
];

// Completion function
function expressionCompletions(context: CompletionContext): CompletionResult | null {
  const word = context.matchBefore(/\w*/);
  if (!word) return null;

  // Don't show completions in the middle of a word
  if (word.from === word.to && !context.explicit) return null;

  // Get the text before cursor
  const textBefore = context.state.doc.sliceString(0, context.pos);

  // Determine what kind of completions to show
  let options = [];

  // If we're likely at the start of a function name
  if (/[^a-zA-Z0-9_]$/.test(textBefore) || textBefore === '') {
    options = [...functionCompletions, ...booleanCompletions, ...operatorCompletions];
  } else {
    // Show all completions
    options = [...functionCompletions, ...booleanCompletions, ...operatorCompletions];
  }

  return {
    from: word.from,
    options: options.map((completion) => ({
      label: completion.label,
      type: completion.type,
      detail: completion.detail,
      info: completion.info,
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
    })),
    validFor: /^\w*$/,
  };
}

// Create the autocompletion extension
export function expressionAutocompletion() {
  return autocompletion({
    override: [expressionCompletions],
    defaultKeymap: true,
    closeOnBlur: true,
    icons: true,
    optionClass: (completion) => `cm-completion-${completion.type}`,
    tooltipClass: () => 'cm-autocomplete-tooltip',
  });
}
