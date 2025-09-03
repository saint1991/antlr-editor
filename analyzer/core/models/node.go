package models

type NodeType int

const (
	// Root node
	NodeTypeExpression NodeType = 0

	// Parser Rules - Expression types
	NodeTypeLiteralExpr      NodeType = 1
	NodeTypeColumnRefExpr    NodeType = 2
	NodeTypeFunctionCallExpr NodeType = 3
	NodeTypeParenExpr        NodeType = 4
	NodeTypeUnaryMinusExpr   NodeType = 5
	NodeTypePowerExpr        NodeType = 6
	NodeTypeMulDivExpr       NodeType = 7
	NodeTypeAddSubExpr       NodeType = 8
	NodeTypeComparisonExpr   NodeType = 9
	NodeTypeAndExpr          NodeType = 10
	NodeTypeOrExpr           NodeType = 11

	// Parser Rules - Components
	NodeTypeLiteral         NodeType = 12
	NodeTypeColumnReference NodeType = 13
	NodeTypeFunctionCall    NodeType = 14
	NodeTypeArgumentList    NodeType = 15

	// Lexer Rules - Operators
	NodeTypeAdd NodeType = 16
	NodeTypeSub NodeType = 17
	NodeTypeMul NodeType = 18
	NodeTypeDiv NodeType = 19
	NodeTypePow NodeType = 20
	NodeTypeLt  NodeType = 21
	NodeTypeLe  NodeType = 22
	NodeTypeGt  NodeType = 23
	NodeTypeGe  NodeType = 24
	NodeTypeEq  NodeType = 25
	NodeTypeNeq NodeType = 26
	NodeTypeOr  NodeType = 27
	NodeTypeAnd NodeType = 28

	// Lexer Rules - Delimiters
	NodeTypeLParen   NodeType = 29
	NodeTypeRParen   NodeType = 30
	NodeTypeLBracket NodeType = 31
	NodeTypeRBracket NodeType = 32
	NodeTypeComma    NodeType = 33

	// Lexer Rules - Literals
	NodeTypeBooleanLiteral NodeType = 34
	NodeTypeFloatLiteral   NodeType = 35
	NodeTypeIntegerLiteral NodeType = 36
	NodeTypeStringLiteral  NodeType = 37

	// Lexer Rules - Identifiers
	NodeTypeFunctionName NodeType = 38
	NodeTypeColumnRef    NodeType = 39

	// Special
	NodeTypeWS        NodeType = 40
	NodeTypeErrorChar NodeType = 41
	NodeTypeTerminal  NodeType = 42
	NodeTypeError     NodeType = 43
)

// ParseTreeNode represents a node in the parse tree hierarchy
type ParseTreeNode struct {
	Type     NodeType        `json:"type"`     // Node type
	Text     string          `json:"text"`     // Text content of this node
	Start    int             `json:"start"`    // Start position in the input
	End      int             `json:"end"`      // End position in the input
	Children []ParseTreeNode `json:"children"` // Child nodes
}

// AsMap converts ParseTreeNode to a map for JSON serialization
func (n *ParseTreeNode) AsMap() map[string]any {
	children := make([]any, len(n.Children))
	for i, child := range n.Children {
		children[i] = child.AsMap()
	}

	return map[string]any{
		"type":     int(n.Type),
		"text":     n.Text,
		"start":    n.Start,
		"end":      n.End,
		"children": children,
	}
}
