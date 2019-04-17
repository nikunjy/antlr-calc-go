package main

import (
	"fmt"
	"log"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/nikunjy/antlr-calc-go/parser"
)

func main() {
	fs, err := antlr.NewFileStream("input_file")
	if err != nil {
		log.Fatal(err)
	}
	lex := parser.NewCalculatorLexer(fs)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewCalculatorParser(tokens)
	visitor := parser.NewCalculatorVisitorImpl()
	tree := p.Input()
	fmt.Println(visitor.Visit(tree))
}
