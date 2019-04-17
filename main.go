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
	count := 0
	for {

		switch val := tree.(type) {
		case *parser.CalculateContext:
			fmt.Println(val.Accept(visitor))
			fmt.Println(val.EOF())
		case *parser.ToSetVarContext:
			val.Accept(visitor)
			tree = val.Input()
		default:
			return
		}
		count++
		if count > 5 {
			return
		}
	}
}
