// Code generated from grammar/Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Expression
import "github.com/antlr4-go/antlr/v4"


type BaseExpressionVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseExpressionVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitPowerExpr(ctx *PowerExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitFunctionCallExpr(ctx *FunctionCallExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitMulDivExpr(ctx *MulDivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitComparisonExpr(ctx *ComparisonExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitLiteralExpr(ctx *LiteralExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitColumnRefExpr(ctx *ColumnRefExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitAddSubExpr(ctx *AddSubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitColumnReference(ctx *ColumnReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExpressionVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}
