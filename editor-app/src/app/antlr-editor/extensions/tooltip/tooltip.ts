import { syntaxTree } from '@codemirror/language';
import { hoverTooltip } from '@codemirror/view';
import type { FunctionDescription } from '../function';

export const expressionHoverTooltip = (functionDescriptions: Record<string, FunctionDescription>) => {
  return hoverTooltip((view, pos, side) => {
    // Get the syntax tree from the current state
    const tree = syntaxTree(view.state);
    if (!tree) {
      return null;
    }

    // Use cursor to find the node at the current position
    const node = tree.resolveInner(pos, side);

    // Check if the node is a FunctionName node
    if (node && node.type.name === 'FunctionName') {
      // Get the text of the function name from the document
      const functionName = view.state.doc.sliceString(node.from, node.to).toUpperCase();
      const description = functionDescriptions[functionName];

      if (description) {
        return {
          pos: node.from,
          end: node.to,
          above: true,
          create: () => ({
            dom: createTooltipDOM(description),
          }),
        };
      }
    }

    return null;
  });
};

const createTooltipDOM = (description: FunctionDescription): HTMLElement => {
  const tooltip = document.createElement('div');
  tooltip.className = 'cm-tooltip-hover';

  // Apply inline styles directly
  tooltip.style.cssText = `
    background: #2d3748;
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 8px;
    padding: 16px;
    max-width: 450px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.3);
    color: #e2e8f0;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', sans-serif;
    font-size: 13px;
    line-height: 1.5;
  `;

  // Function name and syntax
  const header = document.createElement('div');
  header.className = 'cm-tooltip-header';
  header.style.cssText = `
    display: flex;
    align-items: baseline;
    gap: 8px;
    margin-bottom: 12px;
    padding-bottom: 10px;
    border-bottom: 1px solid rgba(99, 179, 237, 0.2);
  `;

  const functionName = document.createElement('strong');
  functionName.textContent = description.name;
  functionName.style.cssText = `
    color: #63b3ed;
    font-weight: 600;
    font-size: 13px;
    letter-spacing: 0.3px;
  `;
  header.appendChild(functionName);

  const separator = document.createElement('span');
  separator.textContent = ' ';
  header.appendChild(separator);

  const syntax = document.createElement('code');
  syntax.textContent = description.syntax;
  syntax.style.cssText = `
    background: rgba(104, 211, 145, 0.1);
    color: #68d391;
    padding: 4px 8px;
    border-radius: 4px;
    border: 1px solid rgba(104, 211, 145, 0.3);
    font-family: 'SF Mono', Monaco, 'Cascadia Code', 'Roboto Mono', Consolas, 'Courier New', monospace;
    font-size: 12px;
    font-weight: 500;
  `;
  header.appendChild(syntax);

  tooltip.appendChild(header);

  // Description
  const desc = document.createElement('div');
  desc.className = 'cm-tooltip-description';
  desc.textContent = description.description;
  desc.style.cssText = `
    margin-bottom: 14px;
    color: #e2e8f0;
    font-size: 13px;
    line-height: 1.6;
    padding: 8px;
    background: rgba(0, 0, 0, 0.2);
    border-radius: 4px;
    border-left: 3px solid #63b3ed;
  `;
  tooltip.appendChild(desc);

  // Examples (if available)
  if (description.examples && description.examples.length > 0) {
    const examplesTitle = document.createElement('div');
    examplesTitle.className = 'cm-tooltip-examples-title';
    examplesTitle.textContent = '▸ Examples';
    examplesTitle.style.cssText = `
      font-weight: 600;
      margin-bottom: 8px;
      color: #63b3ed;
      font-size: 12px;
      text-transform: uppercase;
      letter-spacing: 1px;
      display: flex;
      align-items: center;
      gap: 6px;
    `;
    tooltip.appendChild(examplesTitle);

    const examplesList = document.createElement('ul');
    examplesList.className = 'cm-tooltip-examples';
    examplesList.style.cssText = `
      margin: 0;
      padding: 0;
      list-style: none;
      background: rgba(0, 0, 0, 0.15);
      border-radius: 4px;
      padding: 8px 12px;
    `;

    description.examples.forEach((example, index) => {
      const li = document.createElement('li');
      li.style.cssText = `
        margin-bottom: ${index === (description.examples?.length || 0) - 1 ? '0' : '8px'};
        padding: 6px 8px;
        background: rgba(26, 32, 44, 0.5);
        border-radius: 4px;
        position: relative;
        padding-left: 24px;
        transition: background 0.2s ease;
      `;

      // Add arrow
      const arrow = document.createElement('span');
      arrow.textContent = '→';
      arrow.style.cssText = `
        color: #68d391;
        position: absolute;
        left: 8px;
        font-weight: bold;
      `;
      li.appendChild(arrow);

      const code = document.createElement('code');
      code.textContent = example;
      code.style.cssText = `
        background: rgba(26, 32, 44, 0.7);
        color: #f7fafc;
        padding: 3px 6px;
        border-radius: 3px;
        font-family: 'SF Mono', Monaco, 'Cascadia Code', 'Roboto Mono', Consolas, 'Courier New', monospace;
        font-size: 12px;
        font-weight: 500;
        display: inline-block;
        border: 1px solid rgba(255, 255, 255, 0.1);
      `;
      li.appendChild(code);

      examplesList.appendChild(li);
    });
    tooltip.appendChild(examplesList);
  }

  return tooltip;
};
