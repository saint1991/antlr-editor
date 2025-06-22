// Code generated from grammar/Expression.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser
import (
	"fmt"
  	"sync"
	"unicode"
	"github.com/antlr4-go/antlr/v4"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter


type ExpressionLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var ExpressionLexerLexerStaticData struct {
  once                   sync.Once
  serializedATN          []int32
  ChannelNames           []string
  ModeNames              []string
  LiteralNames           []string
  SymbolicNames          []string
  RuleNames              []string
  PredictionContextCache *antlr.PredictionContextCache
  atn                    *antlr.ATN
  decisionToDFA          []*antlr.DFA
}

func expressionlexerLexerInit() {
  staticData := &ExpressionLexerLexerStaticData
  staticData.ChannelNames = []string{
    "DEFAULT_TOKEN_CHANNEL", "HIDDEN",
  }
  staticData.ModeNames = []string{
    "DEFAULT_MODE",
  }
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
    "ADD", "SUB", "MUL", "DIV", "POW", "EQ", "NEQ", "OR", "AND", "LPAREN", 
    "RPAREN", "LBRACKET", "RBRACKET", "COMMA", "STRING_LITERAL", "INTEGER_LITERAL", 
    "FLOAT_LITERAL", "FUNCTION_NAME", "IDENTIFIER", "WS",
  }
  staticData.PredictionContextCache = antlr.NewPredictionContextCache()
  staticData.serializedATN = []int32{
	4, 0, 20, 152, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 
	4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 
	10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 
	7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 1, 0, 1, 
	0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 
	6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 
	10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 
	5, 14, 78, 8, 14, 10, 14, 12, 14, 81, 9, 14, 1, 14, 1, 14, 1, 14, 1, 14, 
	1, 14, 5, 14, 88, 8, 14, 10, 14, 12, 14, 91, 9, 14, 1, 14, 3, 14, 94, 8, 
	14, 1, 15, 4, 15, 97, 8, 15, 11, 15, 12, 15, 98, 1, 16, 4, 16, 102, 8, 
	16, 11, 16, 12, 16, 103, 1, 16, 1, 16, 4, 16, 108, 8, 16, 11, 16, 12, 16, 
	109, 1, 16, 4, 16, 113, 8, 16, 11, 16, 12, 16, 114, 1, 16, 1, 16, 4, 16, 
	119, 8, 16, 11, 16, 12, 16, 120, 3, 16, 123, 8, 16, 1, 16, 1, 16, 3, 16, 
	127, 8, 16, 1, 16, 4, 16, 130, 8, 16, 11, 16, 12, 16, 131, 3, 16, 134, 
	8, 16, 1, 17, 4, 17, 137, 8, 17, 11, 17, 12, 17, 138, 1, 18, 4, 18, 142, 
	8, 18, 11, 18, 12, 18, 143, 1, 19, 4, 19, 147, 8, 19, 11, 19, 12, 19, 148, 
	1, 19, 1, 19, 0, 0, 20, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 
	8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 
	35, 18, 37, 19, 39, 20, 1, 0, 8, 4, 0, 10, 10, 13, 13, 39, 39, 92, 92, 
	4, 0, 10, 10, 13, 13, 34, 34, 92, 92, 1, 0, 48, 57, 2, 0, 69, 69, 101, 
	101, 2, 0, 43, 43, 45, 45, 1, 0, 65, 90, 5, 0, 9, 10, 13, 13, 32, 32, 91, 
	91, 93, 93, 3, 0, 9, 10, 13, 13, 32, 32, 168, 0, 1, 1, 0, 0, 0, 0, 3, 1, 
	0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 
	0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 
	1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 
	27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 
	0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 1, 41, 1, 0, 0, 
	0, 3, 43, 1, 0, 0, 0, 5, 45, 1, 0, 0, 0, 7, 47, 1, 0, 0, 0, 9, 49, 1, 0, 
	0, 0, 11, 51, 1, 0, 0, 0, 13, 54, 1, 0, 0, 0, 15, 57, 1, 0, 0, 0, 17, 60, 
	1, 0, 0, 0, 19, 63, 1, 0, 0, 0, 21, 65, 1, 0, 0, 0, 23, 67, 1, 0, 0, 0, 
	25, 69, 1, 0, 0, 0, 27, 71, 1, 0, 0, 0, 29, 93, 1, 0, 0, 0, 31, 96, 1, 
	0, 0, 0, 33, 133, 1, 0, 0, 0, 35, 136, 1, 0, 0, 0, 37, 141, 1, 0, 0, 0, 
	39, 146, 1, 0, 0, 0, 41, 42, 5, 43, 0, 0, 42, 2, 1, 0, 0, 0, 43, 44, 5, 
	45, 0, 0, 44, 4, 1, 0, 0, 0, 45, 46, 5, 42, 0, 0, 46, 6, 1, 0, 0, 0, 47, 
	48, 5, 47, 0, 0, 48, 8, 1, 0, 0, 0, 49, 50, 5, 94, 0, 0, 50, 10, 1, 0, 
	0, 0, 51, 52, 5, 61, 0, 0, 52, 53, 5, 61, 0, 0, 53, 12, 1, 0, 0, 0, 54, 
	55, 5, 33, 0, 0, 55, 56, 5, 61, 0, 0, 56, 14, 1, 0, 0, 0, 57, 58, 5, 124, 
	0, 0, 58, 59, 5, 124, 0, 0, 59, 16, 1, 0, 0, 0, 60, 61, 5, 38, 0, 0, 61, 
	62, 5, 38, 0, 0, 62, 18, 1, 0, 0, 0, 63, 64, 5, 40, 0, 0, 64, 20, 1, 0, 
	0, 0, 65, 66, 5, 41, 0, 0, 66, 22, 1, 0, 0, 0, 67, 68, 5, 91, 0, 0, 68, 
	24, 1, 0, 0, 0, 69, 70, 5, 93, 0, 0, 70, 26, 1, 0, 0, 0, 71, 72, 5, 44, 
	0, 0, 72, 28, 1, 0, 0, 0, 73, 79, 5, 39, 0, 0, 74, 78, 8, 0, 0, 0, 75, 
	76, 5, 92, 0, 0, 76, 78, 9, 0, 0, 0, 77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 
	0, 0, 78, 81, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 82, 
	1, 0, 0, 0, 81, 79, 1, 0, 0, 0, 82, 94, 5, 39, 0, 0, 83, 89, 5, 34, 0, 
	0, 84, 88, 8, 1, 0, 0, 85, 86, 5, 92, 0, 0, 86, 88, 9, 0, 0, 0, 87, 84, 
	1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 
	89, 90, 1, 0, 0, 0, 90, 92, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 94, 5, 
	34, 0, 0, 93, 73, 1, 0, 0, 0, 93, 83, 1, 0, 0, 0, 94, 30, 1, 0, 0, 0, 95, 
	97, 7, 2, 0, 0, 96, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 
	0, 98, 99, 1, 0, 0, 0, 99, 32, 1, 0, 0, 0, 100, 102, 7, 2, 0, 0, 101, 100, 
	1, 0, 0, 0, 102, 103, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 
	0, 0, 104, 105, 1, 0, 0, 0, 105, 107, 5, 46, 0, 0, 106, 108, 7, 2, 0, 0, 
	107, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 
	110, 1, 0, 0, 0, 110, 134, 1, 0, 0, 0, 111, 113, 7, 2, 0, 0, 112, 111, 
	1, 0, 0, 0, 113, 114, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 
	0, 0, 115, 122, 1, 0, 0, 0, 116, 118, 5, 46, 0, 0, 117, 119, 7, 2, 0, 0, 
	118, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 
	121, 1, 0, 0, 0, 121, 123, 1, 0, 0, 0, 122, 116, 1, 0, 0, 0, 122, 123, 
	1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 126, 7, 3, 0, 0, 125, 127, 7, 4, 
	0, 0, 126, 125, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 129, 1, 0, 0, 0, 
	128, 130, 7, 2, 0, 0, 129, 128, 1, 0, 0, 0, 130, 131, 1, 0, 0, 0, 131, 
	129, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 134, 1, 0, 0, 0, 133, 101, 
	1, 0, 0, 0, 133, 112, 1, 0, 0, 0, 134, 34, 1, 0, 0, 0, 135, 137, 7, 5, 
	0, 0, 136, 135, 1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 
	138, 139, 1, 0, 0, 0, 139, 36, 1, 0, 0, 0, 140, 142, 8, 6, 0, 0, 141, 140, 
	1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 
	0, 0, 144, 38, 1, 0, 0, 0, 145, 147, 7, 7, 0, 0, 146, 145, 1, 0, 0, 0, 
	147, 148, 1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 
	150, 1, 0, 0, 0, 150, 151, 6, 19, 0, 0, 151, 40, 1, 0, 0, 0, 18, 0, 77, 
	79, 87, 89, 93, 98, 103, 109, 114, 120, 122, 126, 131, 133, 138, 143, 148, 
	1, 6, 0, 0,
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

// ExpressionLexerInit initializes any static state used to implement ExpressionLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewExpressionLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func ExpressionLexerInit() {
  staticData := &ExpressionLexerLexerStaticData
  staticData.once.Do(expressionlexerLexerInit)
}

// NewExpressionLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewExpressionLexer(input antlr.CharStream) *ExpressionLexer {
  ExpressionLexerInit()
	l := new(ExpressionLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
  staticData := &ExpressionLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Expression.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExpressionLexer tokens.
const (
	ExpressionLexerADD = 1
	ExpressionLexerSUB = 2
	ExpressionLexerMUL = 3
	ExpressionLexerDIV = 4
	ExpressionLexerPOW = 5
	ExpressionLexerEQ = 6
	ExpressionLexerNEQ = 7
	ExpressionLexerOR = 8
	ExpressionLexerAND = 9
	ExpressionLexerLPAREN = 10
	ExpressionLexerRPAREN = 11
	ExpressionLexerLBRACKET = 12
	ExpressionLexerRBRACKET = 13
	ExpressionLexerCOMMA = 14
	ExpressionLexerSTRING_LITERAL = 15
	ExpressionLexerINTEGER_LITERAL = 16
	ExpressionLexerFLOAT_LITERAL = 17
	ExpressionLexerFUNCTION_NAME = 18
	ExpressionLexerIDENTIFIER = 19
	ExpressionLexerWS = 20
)

