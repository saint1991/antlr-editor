// Code generated from grammar/Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Expression
import "github.com/antlr4-go/antlr/v4"


// A complete Visitor for a parse tree produced by ExpressionParser.
type ExpressionVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExpressionParser#AndExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#PowerExpr.
	VisitPowerExpr(ctx *PowerExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#FunctionCallExpr.
	VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#MulDivExpr.
	VisitMulDivExpr(ctx *MulDivExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#ComparisonExpr.
	VisitComparisonExpr(ctx *ComparisonExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#LiteralExpr.
	VisitLiteralExpr(ctx *LiteralExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#ColumnRefExpr.
	VisitColumnRefExpr(ctx *ColumnRefExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#AddSubExpr.
	VisitAddSubExpr(ctx *AddSubExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#OrExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by ExpressionParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by ExpressionParser#columnReference.
	VisitColumnReference(ctx *ColumnReferenceContext) interface{}

	// Visit a parse tree produced by ExpressionParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by ExpressionParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

}