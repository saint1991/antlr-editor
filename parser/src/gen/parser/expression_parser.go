// Code generated from grammar/Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Expression
import (
	"fmt"
	"strconv"
  	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}


type ExpressionParser struct {
	*antlr.BaseParser
}

var ExpressionParserStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func expressionParserInit() {
  staticData := &ExpressionParserStaticData
  staticData.LiteralNames = []string{
    "", "'+'", "'-'", "'*'", "'/'", "'^'", "'=='", "'!='", "'||'", "'&&'", 
    "'('", "')'", "'['", "']'", "','",
  }
  staticData.SymbolicNames = []string{
    "", "ADD", "SUB", "MUL", "DIV", "POW", "EQ", "NEQ", "OR", "AND", "LPAREN", 
    "RPAREN", "LBRACKET", "RBRACKET", "COMMA", "STRING_LITERAL", "INTEGER_LITERAL", 
    "FLOAT_LITERAL", "FUNCTION_NAME", "IDENTIFIER", "WS",
  }
  staticData.RuleNames = []string{
    "expression", "literal", "columnReference", "functionCall", "argumentList",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 1, 20, 65, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 
	4, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 19, 8, 0, 1, 0, 
	1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 
	1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 39, 8, 0, 10, 0, 12, 0, 42, 9, 0, 1, 
	1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 53, 8, 3, 1, 3, 
	1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 60, 8, 4, 10, 4, 12, 4, 63, 9, 4, 1, 4, 0, 
	1, 0, 5, 0, 2, 4, 6, 8, 0, 4, 1, 0, 3, 4, 1, 0, 1, 2, 1, 0, 6, 7, 1, 0, 
	15, 17, 70, 0, 18, 1, 0, 0, 0, 2, 43, 1, 0, 0, 0, 4, 45, 1, 0, 0, 0, 6, 
	49, 1, 0, 0, 0, 8, 56, 1, 0, 0, 0, 10, 11, 6, 0, -1, 0, 11, 19, 3, 2, 1, 
	0, 12, 19, 3, 4, 2, 0, 13, 19, 3, 6, 3, 0, 14, 15, 5, 10, 0, 0, 15, 16, 
	3, 0, 0, 0, 16, 17, 5, 11, 0, 0, 17, 19, 1, 0, 0, 0, 18, 10, 1, 0, 0, 0, 
	18, 12, 1, 0, 0, 0, 18, 13, 1, 0, 0, 0, 18, 14, 1, 0, 0, 0, 19, 40, 1, 
	0, 0, 0, 20, 21, 10, 6, 0, 0, 21, 22, 5, 5, 0, 0, 22, 39, 3, 0, 0, 6, 23, 
	24, 10, 5, 0, 0, 24, 25, 7, 0, 0, 0, 25, 39, 3, 0, 0, 6, 26, 27, 10, 4, 
	0, 0, 27, 28, 7, 1, 0, 0, 28, 39, 3, 0, 0, 5, 29, 30, 10, 3, 0, 0, 30, 
	31, 7, 2, 0, 0, 31, 39, 3, 0, 0, 4, 32, 33, 10, 2, 0, 0, 33, 34, 5, 9, 
	0, 0, 34, 39, 3, 0, 0, 3, 35, 36, 10, 1, 0, 0, 36, 37, 5, 8, 0, 0, 37, 
	39, 3, 0, 0, 2, 38, 20, 1, 0, 0, 0, 38, 23, 1, 0, 0, 0, 38, 26, 1, 0, 0, 
	0, 38, 29, 1, 0, 0, 0, 38, 32, 1, 0, 0, 0, 38, 35, 1, 0, 0, 0, 39, 42, 
	1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 1, 1, 0, 0, 0, 
	42, 40, 1, 0, 0, 0, 43, 44, 7, 3, 0, 0, 44, 3, 1, 0, 0, 0, 45, 46, 5, 12, 
	0, 0, 46, 47, 5, 19, 0, 0, 47, 48, 5, 13, 0, 0, 48, 5, 1, 0, 0, 0, 49, 
	50, 5, 18, 0, 0, 50, 52, 5, 10, 0, 0, 51, 53, 3, 8, 4, 0, 52, 51, 1, 0, 
	0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 55, 5, 11, 0, 0, 55, 
	7, 1, 0, 0, 0, 56, 61, 3, 0, 0, 0, 57, 58, 5, 14, 0, 0, 58, 60, 3, 0, 0, 
	0, 59, 57, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61, 59, 1, 0, 0, 0, 61, 62, 
	1, 0, 0, 0, 62, 9, 1, 0, 0, 0, 63, 61, 1, 0, 0, 0, 5, 18, 38, 40, 52, 61,
}
  deserializer := antlr.NewATNDeserializer(nil)
  staticData.atn = deserializer.Deserialize(staticData.serializedATN)
  atn := staticData.atn
  staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
  decisionToDFA := staticData.decisionToDFA
  for index, state := range atn.DecisionToState {
    decisionToDFA[index] = antlr.NewDFA(state, index)
  }
}

// ExpressionParserInit initializes any static state used to implement ExpressionParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewExpressionParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExpressionParserInit() {
  staticData := &ExpressionParserStaticData
  staticData.once.Do(expressionParserInit)
}

// NewExpressionParser produces a new parser instance for the optional input antlr.TokenStream.
func NewExpressionParser(input antlr.TokenStream) *ExpressionParser {
	ExpressionParserInit()
	this := new(ExpressionParser)
	this.BaseParser = antlr.NewBaseParser(input)
  staticData := &ExpressionParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Expression.g4"

	return this
}


// ExpressionParser tokens.
const (
	ExpressionParserEOF = antlr.TokenEOF
	ExpressionParserADD = 1
	ExpressionParserSUB = 2
	ExpressionParserMUL = 3
	ExpressionParserDIV = 4
	ExpressionParserPOW = 5
	ExpressionParserEQ = 6
	ExpressionParserNEQ = 7
	ExpressionParserOR = 8
	ExpressionParserAND = 9
	ExpressionParserLPAREN = 10
	ExpressionParserRPAREN = 11
	ExpressionParserLBRACKET = 12
	ExpressionParserRBRACKET = 13
	ExpressionParserCOMMA = 14
	ExpressionParserSTRING_LITERAL = 15
	ExpressionParserINTEGER_LITERAL = 16
	ExpressionParserFLOAT_LITERAL = 17
	ExpressionParserFUNCTION_NAME = 18
	ExpressionParserIDENTIFIER = 19
	ExpressionParserWS = 20
)

// ExpressionParser rules.
const (
	ExpressionParserRULE_expression = 0
	ExpressionParserRULE_literal = 1
	ExpressionParserRULE_columnReference = 2
	ExpressionParserRULE_functionCall = 3
	ExpressionParserRULE_argumentList = 4
)

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type AndExprContext struct {
	ExpressionContext
}

func NewAndExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndExprContext {
	var p = new(AndExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AndExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndExprContext) AND() antlr.TerminalNode {
	return s.GetToken(ExpressionParserAND, 0)
}


func (s *AndExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitAndExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type PowerExprContext struct {
	ExpressionContext
}

func NewPowerExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerExprContext {
	var p = new(PowerExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *PowerExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PowerExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowerExprContext) POW() antlr.TerminalNode {
	return s.GetToken(ExpressionParserPOW, 0)
}


func (s *PowerExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitPowerExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type FunctionCallExprContext struct {
	ExpressionContext
}

func NewFunctionCallExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExprContext {
	var p = new(FunctionCallExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FunctionCallExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExprContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}


func (s *FunctionCallExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitFunctionCallExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type MulDivExprContext struct {
	ExpressionContext
}

func NewMulDivExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivExprContext {
	var p = new(MulDivExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulDivExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(ExpressionParserMUL, 0)
}

func (s *MulDivExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(ExpressionParserDIV, 0)
}


func (s *MulDivExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitMulDivExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type ComparisonExprContext struct {
	ExpressionContext
}

func NewComparisonExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComparisonExprContext {
	var p = new(ComparisonExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ComparisonExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ComparisonExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ComparisonExprContext) EQ() antlr.TerminalNode {
	return s.GetToken(ExpressionParserEQ, 0)
}

func (s *ComparisonExprContext) NEQ() antlr.TerminalNode {
	return s.GetToken(ExpressionParserNEQ, 0)
}


func (s *ComparisonExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitComparisonExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type LiteralExprContext struct {
	ExpressionContext
}

func NewLiteralExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExprContext {
	var p = new(LiteralExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}


func (s *LiteralExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitLiteralExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type ColumnRefExprContext struct {
	ExpressionContext
}

func NewColumnRefExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnRefExprContext {
	var p = new(ColumnRefExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ColumnRefExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnRefExprContext) ColumnReference() IColumnReferenceContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnReferenceContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnReferenceContext)
}


func (s *ColumnRefExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitColumnRefExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type ParenExprContext struct {
	ExpressionContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLPAREN, 0)
}

func (s *ParenExprContext) Expression() IExpressionContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRPAREN, 0)
}


func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type AddSubExprContext struct {
	ExpressionContext
}

func NewAddSubExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExprContext {
	var p = new(AddSubExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddSubExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(ExpressionParserADD, 0)
}

func (s *AddSubExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(ExpressionParserSUB, 0)
}


func (s *AddSubExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitAddSubExpr(s)

	default:
		return t.VisitChildren(s)
	}
}


type OrExprContext struct {
	ExpressionContext
}

func NewOrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrExprContext {
	var p = new(OrExprContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *OrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *OrExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(ExpressionParserOR, 0)
}


func (s *OrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitOrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *ExpressionParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ExpressionParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, ExpressionParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ExpressionParserSTRING_LITERAL, ExpressionParserINTEGER_LITERAL, ExpressionParserFLOAT_LITERAL:
		localctx = NewLiteralExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(11)
			p.Literal()
		}


	case ExpressionParserLBRACKET:
		localctx = NewColumnRefExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(12)
			p.ColumnReference()
		}


	case ExpressionParserFUNCTION_NAME:
		localctx = NewFunctionCallExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(13)
			p.FunctionCall()
		}


	case ExpressionParserLPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(14)
			p.Match(ExpressionParserLPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(15)
			p.expression(0)
		}
		{
			p.SetState(16)
			p.Match(ExpressionParserRPAREN)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}



	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPowerExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(20)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(21)
					p.Match(ExpressionParserPOW)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(22)
					p.expression(6)
				}


			case 2:
				localctx = NewMulDivExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(23)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(24)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ExpressionParserMUL || _la == ExpressionParserDIV) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(25)
					p.expression(6)
				}


			case 3:
				localctx = NewAddSubExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(26)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(27)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ExpressionParserADD || _la == ExpressionParserSUB) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(28)
					p.expression(5)
				}


			case 4:
				localctx = NewComparisonExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(30)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ExpressionParserEQ || _la == ExpressionParserNEQ) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(31)
					p.expression(4)
				}


			case 5:
				localctx = NewAndExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(32)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(33)
					p.Match(ExpressionParserAND)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(34)
					p.expression(3)
				}


			case 6:
				localctx = NewOrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ExpressionParserRULE_expression)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(36)
					p.Match(ExpressionParserOR)
					if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
					}
				}
				{
					p.SetState(37)
					p.expression(2)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}



	errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL() antlr.TerminalNode
	INTEGER_LITERAL() antlr.TerminalNode
	FLOAT_LITERAL() antlr.TerminalNode

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_literal
	return p
}

func InitEmptyLiteralContext(p *LiteralContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_literal
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(ExpressionParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) INTEGER_LITERAL() antlr.TerminalNode {
	return s.GetToken(ExpressionParserINTEGER_LITERAL, 0)
}

func (s *LiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(ExpressionParserFLOAT_LITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExpressionParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ExpressionParserRULE_literal)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		_la = p.GetTokenStream().LA(1)

		if !(((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 229376) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IColumnReferenceContext is an interface to support dynamic dispatch.
type IColumnReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACKET() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode

	// IsColumnReferenceContext differentiates from other interfaces.
	IsColumnReferenceContext()
}

type ColumnReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnReferenceContext() *ColumnReferenceContext {
	var p = new(ColumnReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_columnReference
	return p
}

func InitEmptyColumnReferenceContext(p *ColumnReferenceContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_columnReference
}

func (*ColumnReferenceContext) IsColumnReferenceContext() {}

func NewColumnReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnReferenceContext {
	var p = new(ColumnReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_columnReference

	return p
}

func (s *ColumnReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnReferenceContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLBRACKET, 0)
}

func (s *ColumnReferenceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ExpressionParserIDENTIFIER, 0)
}

func (s *ColumnReferenceContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRBRACKET, 0)
}

func (s *ColumnReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ColumnReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitColumnReference(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExpressionParser) ColumnReference() (localctx IColumnReferenceContext) {
	localctx = NewColumnReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ExpressionParserRULE_columnReference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(ExpressionParserLBRACKET)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(ExpressionParserIDENTIFIER)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(47)
		p.Match(ExpressionParserRBRACKET)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNCTION_NAME() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgumentList() IArgumentListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) FUNCTION_NAME() antlr.TerminalNode {
	return s.GetToken(ExpressionParserFUNCTION_NAME, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ExpressionParserRPAREN, 0)
}

func (s *FunctionCallContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext;
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext);
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExpressionParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ExpressionParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(ExpressionParserFUNCTION_NAME)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(ExpressionParserLPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	if ((int64(_la) & ^0x3f) == 0 && ((int64(1) << _la) & 496640) != 0) {
		{
			p.SetState(51)
			p.ArgumentList()
		}

	}
	{
		p.SetState(54)
		p.Match(ExpressionParserRPAREN)
		if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
		}
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext)  {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ExpressionParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ExpressionParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext;
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext);
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ExpressionParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ExpressionParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ExpressionVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *ExpressionParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ExpressionParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.expression(0)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)


	for _la == ExpressionParserCOMMA {
		{
			p.SetState(57)
			p.Match(ExpressionParserCOMMA)
			if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
			}
		}
		{
			p.SetState(58)
			p.expression(0)
		}


		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
	    	goto errorExit
	    }
		_la = p.GetTokenStream().LA(1)
	}



errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}


func (p *ExpressionParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
			var t *ExpressionContext = nil
			if localctx != nil { t = localctx.(*ExpressionContext) }
			return p.Expression_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ExpressionParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
			return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
			return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
			return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

