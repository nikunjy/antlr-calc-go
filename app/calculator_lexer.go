// Code generated from Calculator.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 17, 86, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 6, 2, 35, 10,
	2, 13, 2, 14, 2, 36, 3, 3, 6, 3, 40, 10, 3, 13, 3, 14, 3, 41, 3, 3, 3,
	3, 6, 3, 46, 10, 3, 13, 3, 14, 3, 47, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 6, 8, 60, 10, 8, 13, 8, 14, 8, 61, 3, 8, 3,
	8, 3, 9, 3, 9, 7, 9, 68, 10, 9, 12, 9, 14, 9, 71, 11, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	16, 3, 16, 2, 2, 17, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10,
	19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 3, 2, 6, 3, 2,
	50, 59, 5, 2, 11, 11, 15, 15, 34, 34, 5, 2, 67, 92, 97, 97, 99, 124, 6,
	2, 50, 59, 67, 92, 97, 97, 99, 124, 2, 90, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 3, 34, 3, 2, 2, 2, 5, 39, 3, 2, 2, 2,
	7, 49, 3, 2, 2, 2, 9, 52, 3, 2, 2, 2, 11, 54, 3, 2, 2, 2, 13, 56, 3, 2,
	2, 2, 15, 59, 3, 2, 2, 2, 17, 65, 3, 2, 2, 2, 19, 72, 3, 2, 2, 2, 21, 74,
	3, 2, 2, 2, 23, 76, 3, 2, 2, 2, 25, 78, 3, 2, 2, 2, 27, 80, 3, 2, 2, 2,
	29, 82, 3, 2, 2, 2, 31, 84, 3, 2, 2, 2, 33, 35, 9, 2, 2, 2, 34, 33, 3,
	2, 2, 2, 35, 36, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 37, 3, 2, 2, 2, 37,
	4, 3, 2, 2, 2, 38, 40, 9, 2, 2, 2, 39, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2,
	2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 45,
	7, 48, 2, 2, 44, 46, 9, 2, 2, 2, 45, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2,
	47, 45, 3, 2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 6, 3, 2, 2, 2, 49, 50, 7, 114,
	2, 2, 50, 51, 7, 107, 2, 2, 51, 8, 3, 2, 2, 2, 52, 53, 7, 103, 2, 2, 53,
	10, 3, 2, 2, 2, 54, 55, 7, 96, 2, 2, 55, 12, 3, 2, 2, 2, 56, 57, 7, 12,
	2, 2, 57, 14, 3, 2, 2, 2, 58, 60, 9, 3, 2, 2, 59, 58, 3, 2, 2, 2, 60, 61,
	3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2,
	63, 64, 8, 8, 2, 2, 64, 16, 3, 2, 2, 2, 65, 69, 9, 4, 2, 2, 66, 68, 9,
	5, 2, 2, 67, 66, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69,
	70, 3, 2, 2, 2, 70, 18, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 73, 7, 45,
	2, 2, 73, 20, 3, 2, 2, 2, 74, 75, 7, 63, 2, 2, 75, 22, 3, 2, 2, 2, 76,
	77, 7, 47, 2, 2, 77, 24, 3, 2, 2, 2, 78, 79, 7, 44, 2, 2, 79, 26, 3, 2,
	2, 2, 80, 81, 7, 49, 2, 2, 81, 28, 3, 2, 2, 2, 82, 83, 7, 42, 2, 2, 83,
	30, 3, 2, 2, 2, 84, 85, 7, 43, 2, 2, 85, 32, 3, 2, 2, 2, 8, 2, 36, 41,
	47, 61, 69, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "'pi'", "'e'", "'^'", "'\n'", "", "", "'+'", "'='", "'-'",
	"'*'", "'/'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "INT", "DOUBLE", "PI", "E", "POW", "NL", "WS", "ID", "PLUS", "EQUAL",
	"MINUS", "MULT", "DIV", "LPAR", "RPAR",
}

var lexerRuleNames = []string{
	"INT", "DOUBLE", "PI", "E", "POW", "NL", "WS", "ID", "PLUS", "EQUAL", "MINUS",
	"MULT", "DIV", "LPAR", "RPAR",
}

type CalculatorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCalculatorLexer(input antlr.CharStream) *CalculatorLexer {

	l := new(CalculatorLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Calculator.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalculatorLexer tokens.
const (
	CalculatorLexerINT    = 1
	CalculatorLexerDOUBLE = 2
	CalculatorLexerPI     = 3
	CalculatorLexerE      = 4
	CalculatorLexerPOW    = 5
	CalculatorLexerNL     = 6
	CalculatorLexerWS     = 7
	CalculatorLexerID     = 8
	CalculatorLexerPLUS   = 9
	CalculatorLexerEQUAL  = 10
	CalculatorLexerMINUS  = 11
	CalculatorLexerMULT   = 12
	CalculatorLexerDIV    = 13
	CalculatorLexerLPAR   = 14
	CalculatorLexerRPAR   = 15
)
