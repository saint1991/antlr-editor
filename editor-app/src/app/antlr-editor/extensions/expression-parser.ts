import { type Input, NodeSet, NodeType, Parser, type PartialParse, Tree, type TreeFragment } from '@lezer/common';
import { styleTags, tags } from '@lezer/highlight';
import type { Analyzer, ParseTreeNode, ParseTreeResult } from './analyzer';

// Define node types for our expression language (matching Go NodeType constants)
const nodeTypeIds = {
  // Root node
  Expression: 0,

  // Parser Rules - Expression types
  LiteralExpr: 1,
  ColumnRefExpr: 2,
  FunctionCallExpr: 3,
  ParenExpr: 4,
  UnaryMinusExpr: 5,
  PowerExpr: 6,
  MulDivExpr: 7,
  AddSubExpr: 8,
  ComparisonExpr: 9,
  AndExpr: 10,
  OrExpr: 11,

  // Parser Rules - Components
  Literal: 12,
  ColumnReference: 13,
  FunctionCall: 14,
  ArgumentList: 15,

  // Lexer Rules - Literals
  BooleanLiteral: 34,
  FloatLiteral: 35,
  IntegerLiteral: 36,
  StringLiteral: 37,

  // Lexer Rules - Identifiers
  FunctionName: 38,

  // Special
  Error: 43,
} as const;

// Create node types (dense array with sequential IDs for CodeMirror)
const nodeTypes: NodeType[] = [
  NodeType.define({ id: 0, name: 'Expression', top: true }), // 0
  NodeType.define({ id: 1, name: 'LiteralExpr' }), // 1
  NodeType.define({ id: 2, name: 'ColumnRefExpr' }), // 2
  NodeType.define({ id: 3, name: 'FunctionCallExpr' }), // 3
  NodeType.define({ id: 4, name: 'ParenExpr' }), // 4
  NodeType.define({ id: 5, name: 'UnaryMinusExpr' }), // 5
  NodeType.define({ id: 6, name: 'PowerExpr' }), // 6
  NodeType.define({ id: 7, name: 'MulDivExpr' }), // 7
  NodeType.define({ id: 8, name: 'AddSubExpr' }), // 8
  NodeType.define({ id: 9, name: 'ComparisonExpr' }), // 9
  NodeType.define({ id: 10, name: 'AndExpr' }), // 10
  NodeType.define({ id: 11, name: 'OrExpr' }), // 11
  NodeType.define({ id: 12, name: 'Literal' }), // 12
  NodeType.define({ id: 13, name: 'ColumnReference' }), // 13
  NodeType.define({ id: 14, name: 'FunctionCall' }), // 14
  NodeType.define({ id: 15, name: 'ArgumentList' }), // 15
  NodeType.define({ id: 16, name: 'BooleanLiteral' }), // 16
  NodeType.define({ id: 17, name: 'FloatLiteral' }), // 17
  NodeType.define({ id: 18, name: 'IntegerLiteral' }), // 18
  NodeType.define({ id: 19, name: 'StringLiteral' }), // 19
  NodeType.define({ id: 20, name: 'FunctionName' }), // 20
  NodeType.define({ id: 21, name: 'Error', error: true }), // 21
];

// Map from Go NodeType constants to CodeMirror array indices
const goNodeTypeToIndex: Map<number, number> = new Map([
  [nodeTypeIds.Expression, 0],
  [nodeTypeIds.LiteralExpr, 1],
  [nodeTypeIds.ColumnRefExpr, 2],
  [nodeTypeIds.FunctionCallExpr, 3],
  [nodeTypeIds.ParenExpr, 4],
  [nodeTypeIds.UnaryMinusExpr, 5],
  [nodeTypeIds.PowerExpr, 6],
  [nodeTypeIds.MulDivExpr, 7],
  [nodeTypeIds.AddSubExpr, 8],
  [nodeTypeIds.ComparisonExpr, 9],
  [nodeTypeIds.AndExpr, 10],
  [nodeTypeIds.OrExpr, 11],
  [nodeTypeIds.Literal, 12],
  [nodeTypeIds.ColumnReference, 13],
  [nodeTypeIds.FunctionCall, 14],
  [nodeTypeIds.ArgumentList, 15],
  [nodeTypeIds.BooleanLiteral, 16],
  [nodeTypeIds.FloatLiteral, 17],
  [nodeTypeIds.IntegerLiteral, 18],
  [nodeTypeIds.StringLiteral, 19],
  [nodeTypeIds.FunctionName, 20],
  [nodeTypeIds.Error, 21],
]);

// Create node set with style tags
export const nodeSet = new NodeSet(nodeTypes).extend(
  styleTags({
    StringLiteral: tags.string,
    IntegerLiteral: tags.number,
    FloatLiteral: tags.number,
    BooleanLiteral: tags.bool,
    FunctionName: tags.function(tags.variableName),
    ColumnReference: tags.variableName,
    Error: tags.invalid,
    // Add fallback styles for common node types
    Literal: tags.literal,
    FunctionCall: tags.function(tags.name),
    FunctionCallExpr: tags.function(tags.name),
  }),
);

// Convert Go NodeType constants to CodeMirror array indices
const getNodeTypeId = (goNodeType: number): number => {
  // Convert Go NodeType to CodeMirror array index
  const index = goNodeTypeToIndex.get(goNodeType);
  if (index !== undefined) {
    return index;
  }

  // Fallback to LiteralExpr index for unknown types
  return 0;
};

// Custom parser implementation that mimics LRParser interface
export class ExpressionParser extends Parser {
  readonly nodeSet = nodeSet;

  constructor(private readonly analyzer: Analyzer) {
    super();
  }

  createParse(input: Input, fragments: readonly TreeFragment[], ranges: readonly { from: number; to: number }[]): PartialParse {
    return new ExpressionPartialParse(this.analyzer, input, fragments, ranges);
  }

  configure(config: any): ExpressionParser {
    // Create a new instance with the configuration applied to nodeSet
    const newParser = new ExpressionParser(this.analyzer);

    if (config.props && config.props.length > 0) {
      // Apply style tags and other properties to the nodeSet
      const styledNodeSet = nodeSet.extend(...config.props);
      (newParser as any).nodeSet = styledNodeSet;
    }

    return newParser;
  }

  hasWrappers(): boolean {
    return false;
  }

  get topNode() {
    return this.nodeSet.types[0]; // Expression
  }

  getName(id: number): string {
    const nodeType = this.nodeSet.types[id];
    return nodeType ? nodeType.name : 'Unknown';
  }
}

// Partial parse implementation
class ExpressionPartialParse implements PartialParse {
  private tree: Tree | null = null;
  parsedPos = 0;
  stoppedAt: number | null = null;

  constructor(
    private analyzer: Analyzer,
    private input: Input,
    _fragments: readonly TreeFragment[],
    _ranges: readonly { from: number; to: number }[],
  ) {}

  advance(): Tree | null {
    if (this.tree) {
      return this.tree;
    }
    // Get the full text
    const text = this.input.read(0, this.input.length);

    // Parse with analyzer
    const result: ParseTreeResult = this.analyzer.parseTree(text);

    // Build tree from parse result
    if (result.tree) {
      // Use the parse tree even if there are errors (error nodes are embedded in the tree)
      this.tree = this.buildTree(result.tree);
    } else {
      // Empty tree
      this.tree = Tree.build({
        buffer: [0, 0, 0, 4], // 0 = Expression index
        nodeSet,
        topID: 0, // 0 = Expression index
        maxBufferLength: 1024,
        reused: [],
        minRepeatType: 22, // Error + 1
      });
    }

    this.parsedPos = this.input.length;
    return this.tree;
  }

  stopAt(pos: number): void {
    this.stoppedAt = pos;
  }

  private buildTree(rootNode: ParseTreeNode): Tree {
    const buffer = this.nodeToBuffer(rootNode);

    return Tree.build({
      buffer,
      nodeSet,
      topID: 0, // 0 = Expression index
      maxBufferLength: 1024,
      reused: [],
      minRepeatType: 22, // Error + 1
    });
  }

  private nodeToBuffer(node: ParseTreeNode): number[] {
    const buffer: number[] = [];

    // Helper to add nodes in postfix order
    const addNode = (n: ParseTreeNode) => {
      const startIndex = buffer.length;

      // First, add all children (if they exist)
      if (n.children && n.children.length > 0) {
        for (const child of n.children) {
          addNode(child);
        }
      }

      // Calculate size (includes this node + all descendants)
      const endIndex = buffer.length;
      const size = 4 + (endIndex - startIndex); // 4 for this node + size of children

      // Add this node
      const nodeId = getNodeTypeId(n.type);
      buffer.push(nodeId, n.start, n.end, size);
    };

    addNode(node);
    return buffer;
  }
}
