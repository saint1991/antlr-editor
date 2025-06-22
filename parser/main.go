package main

import (
	"fmt"
	"antlr-editor/parser/src/gen/parser"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	input := antlr.NewInputStream("1 + 2 * 3")
	lexer := parser.NewExpressionLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewExpressionParser(stream)

	tree := p.Expression()
	fmt.Println("Parse tree:", tree.ToStringTree(nil, p))
}