package tree

import (
	"github.com/antlr4-go/antlr/v4"

	"antlr-editor/analyzer/core/models"
	"antlr-editor/analyzer/gen/parser"
)

// buildParseTreeFromContext recursively builds a models.ParseTreeNode from ANTLR parse tree
// Visitor implements the ANTLR visitor pattern for building parse trees
type Visitor struct {
	parser.BaseExpressionVisitor
	input string
}

// NewParseTreeVisitor creates a new parse tree visitor
func NewParseTreeVisitor(input string) *Visitor {
	return &Visitor{
		input: input,
	}
}

// Visit is the main entry point for visiting nodes
func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	result := tree.Accept(v)
	if node, ok := result.(*models.ParseTreeNode); ok {
		return node
	}
	return nil
}

// VisitChildren visits all children and combines results
func (v *Visitor) VisitChildren(ctx antlr.RuleNode) interface{} {
	// Default implementation - should not be called directly in most cases
	return nil
}

// VisitTerminal handles terminal nodes
func (v *Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	token := node.GetSymbol()
	return &models.ParseTreeNode{
		Type:     models.NodeTypeTerminal,
		Text:     token.GetText(),
		Start:    token.GetStart(),
		End:      token.GetStop() + 1,
		Children: []models.ParseTreeNode{},
	}
}

// VisitErrorNode handles error nodes
func (v *Visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	token := node.GetSymbol()
	return &models.ParseTreeNode{
		Type:     models.NodeTypeError,
		Text:     token.GetText(),
		Start:    token.GetStart(),
		End:      token.GetStop() + 1,
		Children: []models.ParseTreeNode{},
	}
}

// VisitLiteralExpr handles literal expressions
func (v *Visitor) VisitLiteralExpr(ctx *parser.LiteralExprContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	children := []models.ParseTreeNode{}
	if literalNode := v.Visit(ctx.Literal()); literalNode != nil {
		if node, ok := literalNode.(*models.ParseTreeNode); ok {
			children = append(children, *node)
		}
	}

	return &models.ParseTreeNode{
		Type:     models.NodeTypeLiteralExpr,
		Text:     text,
		Start:    start,
		End:      end,
		Children: children,
	}
}

// VisitColumnRefExpr handles column reference expressions
func (v *Visitor) VisitColumnRefExpr(ctx *parser.ColumnRefExprContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	return &models.ParseTreeNode{
		Type:     models.NodeTypeColumnRefExpr,
		Text:     text,
		Start:    start,
		End:      end,
		Children: []models.ParseTreeNode{},
	}
}

// VisitFunctionCallExpr handles function call expressions
func (v *Visitor) VisitFunctionCallExpr(ctx *parser.FunctionCallExprContext) interface{} {
	if funcCall := ctx.FunctionCall(); funcCall != nil {
		return v.Visit(funcCall)
	}
	return nil
}

// VisitParenExpr handles parenthesized expressions
func (v *Visitor) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	children := []models.ParseTreeNode{}
	if expr := ctx.Expression(); expr != nil {
		if innerExpr := v.Visit(expr); innerExpr != nil {
			if node, ok := innerExpr.(*models.ParseTreeNode); ok {
				children = append(children, *node)
			}
		}
	}

	return &models.ParseTreeNode{
		Type:     models.NodeTypeParenExpr,
		Text:     text,
		Start:    start,
		End:      end,
		Children: children,
	}
}

// VisitUnaryMinusExpr handles unary minus expressions
func (v *Visitor) VisitUnaryMinusExpr(ctx *parser.UnaryMinusExprContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	children := []models.ParseTreeNode{}
	if expr := ctx.Expression(); expr != nil {
		if innerExpr := v.Visit(expr); innerExpr != nil {
			if node, ok := innerExpr.(*models.ParseTreeNode); ok {
				children = append(children, *node)
			}
		}
	}

	return &models.ParseTreeNode{
		Type:     models.NodeTypeUnaryMinusExpr,
		Text:     text,
		Start:    start,
		End:      end,
		Children: children,
	}
}

// VisitPowerExpr handles power expressions
func (v *Visitor) VisitPowerExpr(ctx *parser.PowerExprContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	children := []models.ParseTreeNode{}
	for _, expr := range ctx.AllExpression() {
		if child := v.Visit(expr); child != nil {
			if node, ok := child.(*models.ParseTreeNode); ok {
				children = append(children, *node)
			}
		}
	}

	return &models.ParseTreeNode{
		Type:     models.NodeTypePowerExpr,
		Text:     text,
		Start:    start,
		End:      end,
		Children: children,
	}
}

// VisitMulDivExpr handles multiplication/division expressions
func (v *Visitor) VisitMulDivExpr(ctx *parser.MulDivExprContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	children := []models.ParseTreeNode{}
	for _, expr := range ctx.AllExpression() {
		if child := v.Visit(expr); child != nil {
			if node, ok := child.(*models.ParseTreeNode); ok {
				children = append(children, *node)
			}
		}
	}

	return &models.ParseTreeNode{
		Type:     models.NodeTypeMulDivExpr,
		Text:     text,
		Start:    start,
		End:      end,
		Children: children,
	}
}

// VisitAddSubExpr handles addition/subtraction expressions
func (v *Visitor) VisitAddSubExpr(ctx *parser.AddSubExprContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	children := []models.ParseTreeNode{}
	for _, expr := range ctx.AllExpression() {
		if child := v.Visit(expr); child != nil {
			if node, ok := child.(*models.ParseTreeNode); ok {
				children = append(children, *node)
			}
		}
	}

	return &models.ParseTreeNode{
		Type:     models.NodeTypeAddSubExpr,
		Text:     text,
		Start:    start,
		End:      end,
		Children: children,
	}
}

// VisitComparisonExpr handles comparison expressions
func (v *Visitor) VisitComparisonExpr(ctx *parser.ComparisonExprContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	children := []models.ParseTreeNode{}
	for _, expr := range ctx.AllExpression() {
		if child := v.Visit(expr); child != nil {
			if node, ok := child.(*models.ParseTreeNode); ok {
				children = append(children, *node)
			}
		}
	}

	return &models.ParseTreeNode{
		Type:     models.NodeTypeComparisonExpr,
		Text:     text,
		Start:    start,
		End:      end,
		Children: children,
	}
}

// VisitAndExpr handles AND expressions
func (v *Visitor) VisitAndExpr(ctx *parser.AndExprContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	children := []models.ParseTreeNode{}
	for _, expr := range ctx.AllExpression() {
		if child := v.Visit(expr); child != nil {
			if node, ok := child.(*models.ParseTreeNode); ok {
				children = append(children, *node)
			}
		}
	}

	return &models.ParseTreeNode{
		Type:     models.NodeTypeAndExpr,
		Text:     text,
		Start:    start,
		End:      end,
		Children: children,
	}
}

// VisitOrExpr handles OR expressions
func (v *Visitor) VisitOrExpr(ctx *parser.OrExprContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	children := []models.ParseTreeNode{}
	for _, expr := range ctx.AllExpression() {
		if child := v.Visit(expr); child != nil {
			if node, ok := child.(*models.ParseTreeNode); ok {
				children = append(children, *node)
			}
		}
	}

	return &models.ParseTreeNode{
		Type:     models.NodeTypeOrExpr,
		Text:     text,
		Start:    start,
		End:      end,
		Children: children,
	}
}

// VisitLiteral handles literal nodes
func (v *Visitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	var nodeType models.NodeType
	if ctx.STRING_LITERAL() != nil {
		nodeType = models.NodeTypeStringLiteral
	} else if ctx.INTEGER_LITERAL() != nil {
		nodeType = models.NodeTypeIntegerLiteral
	} else if ctx.FLOAT_LITERAL() != nil {
		nodeType = models.NodeTypeFloatLiteral
	} else if ctx.BOOLEAN_LITERAL() != nil {
		nodeType = models.NodeTypeBooleanLiteral
	} else {
		nodeType = models.NodeTypeLiteral
	}

	return &models.ParseTreeNode{
		Type:     nodeType,
		Text:     text,
		Start:    start,
		End:      end,
		Children: []models.ParseTreeNode{},
	}
}

// VisitColumnReference handles column reference nodes
func (v *Visitor) VisitColumnReference(ctx *parser.ColumnReferenceContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	return &models.ParseTreeNode{
		Type:     models.NodeTypeColumnReference,
		Text:     text,
		Start:    start,
		End:      end,
		Children: []models.ParseTreeNode{},
	}
}

// VisitFunctionCall handles function call nodes
func (v *Visitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	children := []models.ParseTreeNode{}

	// Add function name as a child
	if ctx.FUNCTION_NAME() != nil {
		fnToken := ctx.FUNCTION_NAME().GetSymbol()
		children = append(children, models.ParseTreeNode{
			Type:     models.NodeTypeFunctionName,
			Text:     fnToken.GetText(),
			Start:    fnToken.GetStart(),
			End:      fnToken.GetStop() + 1,
			Children: []models.ParseTreeNode{},
		})
	}

	// Add argument list if present
	if argList := ctx.ArgumentList(); argList != nil {
		if argListNode := v.Visit(argList); argListNode != nil {
			if node, ok := argListNode.(*models.ParseTreeNode); ok {
				children = append(children, *node)
			}
		}
	}

	return &models.ParseTreeNode{
		Type:     models.NodeTypeFunctionCall,
		Text:     text,
		Start:    start,
		End:      end,
		Children: children,
	}
}

// VisitArgumentList handles argument list nodes
func (v *Visitor) VisitArgumentList(ctx *parser.ArgumentListContext) interface{} {
	start := ctx.GetStart().GetStart()
	end := ctx.GetStop().GetStop() + 1
	text := v.input[start:end]

	children := []models.ParseTreeNode{}
	for _, expr := range ctx.AllExpression() {
		if child := v.Visit(expr); child != nil {
			if node, ok := child.(*models.ParseTreeNode); ok {
				children = append(children, *node)
			}
		}
	}

	return &models.ParseTreeNode{
		Type:     models.NodeTypeArgumentList,
		Text:     text,
		Start:    start,
		End:      end,
		Children: children,
	}
}
