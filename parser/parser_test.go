package parser

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func getResultFromStream(input string) float64 {
	fs := antlr.NewInputStream(input)
	lex := NewCalculatorLexer(fs)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := NewCalculatorParser(tokens)
	visitor := NewCalculatorVisitorImpl()
	tree := p.Input()
	return visitor.Visit(tree).(float64)
}

func TestSimpleCases(t *testing.T) {
	tests := []struct {
		input  string
		output float64
	}{
		{
			`1+2`,
			3,
		},
		{
			`1*3+2^3`,
			11,
		},
		{
			`(1+2)*3+2^3`,
			17,
		},
	}

	for _, tt := range tests {
		if tt.output != getResultFromStream(tt.input) {
			panic("values dont match")
		}
	}
}

func BenchmarkSimpleCases(b *testing.B) {
	input := `1 + (2*3)*(4+(5*6+(2.0/(3+4*(2*3))))) + 1 + (2*3)*(4+(5*6+(2.0/(3+4*(2*3))))) - (2*3)*(4+(5*6+(2.0/(3+4*(2*3))))) + 1 + (2*3)*(4+(5*6+(2.0/(3+4*(2*3)))))`
	for i := 0; i < b.N; i++ {
		getResultFromStream(input)
	}
}
