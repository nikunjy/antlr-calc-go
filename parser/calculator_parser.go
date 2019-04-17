// Code generated from Calculator.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Calculator

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 17, 82, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 23, 10, 2, 3, 2, 3,
	2, 5, 2, 27, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 42, 10, 4, 12, 4, 14, 4, 45, 11, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 56, 10, 5, 12,
	5, 14, 5, 59, 11, 5, 3, 6, 3, 6, 3, 6, 5, 6, 64, 10, 6, 3, 7, 3, 7, 3,
	7, 5, 7, 69, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 5, 8, 80, 10, 8, 3, 8, 2, 4, 6, 8, 9, 2, 4, 6, 8, 10, 12, 14, 2, 2,
	2, 87, 2, 26, 3, 2, 2, 2, 4, 28, 3, 2, 2, 2, 6, 32, 3, 2, 2, 2, 8, 46,
	3, 2, 2, 2, 10, 60, 3, 2, 2, 2, 12, 68, 3, 2, 2, 2, 14, 79, 3, 2, 2, 2,
	16, 17, 5, 4, 3, 2, 17, 18, 7, 8, 2, 2, 18, 19, 5, 2, 2, 2, 19, 27, 3,
	2, 2, 2, 20, 22, 5, 6, 4, 2, 21, 23, 7, 8, 2, 2, 22, 21, 3, 2, 2, 2, 22,
	23, 3, 2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 25, 7, 2, 2, 3, 25, 27, 3, 2, 2,
	2, 26, 16, 3, 2, 2, 2, 26, 20, 3, 2, 2, 2, 27, 3, 3, 2, 2, 2, 28, 29, 7,
	10, 2, 2, 29, 30, 7, 12, 2, 2, 30, 31, 5, 6, 4, 2, 31, 5, 3, 2, 2, 2, 32,
	33, 8, 4, 1, 2, 33, 34, 5, 8, 5, 2, 34, 43, 3, 2, 2, 2, 35, 36, 12, 5,
	2, 2, 36, 37, 7, 11, 2, 2, 37, 42, 5, 8, 5, 2, 38, 39, 12, 4, 2, 2, 39,
	40, 7, 13, 2, 2, 40, 42, 5, 8, 5, 2, 41, 35, 3, 2, 2, 2, 41, 38, 3, 2,
	2, 2, 42, 45, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 7,
	3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 46, 47, 8, 5, 1, 2, 47, 48, 5, 10, 6, 2,
	48, 57, 3, 2, 2, 2, 49, 50, 12, 5, 2, 2, 50, 51, 7, 14, 2, 2, 51, 56, 5,
	10, 6, 2, 52, 53, 12, 4, 2, 2, 53, 54, 7, 15, 2, 2, 54, 56, 5, 10, 6, 2,
	55, 49, 3, 2, 2, 2, 55, 52, 3, 2, 2, 2, 56, 59, 3, 2, 2, 2, 57, 55, 3,
	2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 9, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 60,
	63, 5, 12, 7, 2, 61, 62, 7, 7, 2, 2, 62, 64, 5, 10, 6, 2, 63, 61, 3, 2,
	2, 2, 63, 64, 3, 2, 2, 2, 64, 11, 3, 2, 2, 2, 65, 66, 7, 13, 2, 2, 66,
	69, 5, 12, 7, 2, 67, 69, 5, 14, 8, 2, 68, 65, 3, 2, 2, 2, 68, 67, 3, 2,
	2, 2, 69, 13, 3, 2, 2, 2, 70, 80, 7, 5, 2, 2, 71, 80, 7, 6, 2, 2, 72, 80,
	7, 4, 2, 2, 73, 80, 7, 3, 2, 2, 74, 80, 7, 10, 2, 2, 75, 76, 7, 16, 2,
	2, 76, 77, 5, 6, 4, 2, 77, 78, 7, 17, 2, 2, 78, 80, 3, 2, 2, 2, 79, 70,
	3, 2, 2, 2, 79, 71, 3, 2, 2, 2, 79, 72, 3, 2, 2, 2, 79, 73, 3, 2, 2, 2,
	79, 74, 3, 2, 2, 2, 79, 75, 3, 2, 2, 2, 80, 15, 3, 2, 2, 2, 11, 22, 26,
	41, 43, 55, 57, 63, 68, 79,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "'pi'", "'e'", "'^'", "'\n'", "", "", "'+'", "'='", "'-'",
	"'*'", "'/'", "'('", "')'",
}
var symbolicNames = []string{
	"", "INT", "DOUBLE", "PI", "E", "POW", "NL", "WS", "ID", "PLUS", "EQUAL",
	"MINUS", "MULT", "DIV", "LPAR", "RPAR",
}

var ruleNames = []string{
	"input", "setVar", "plusOrMinus", "multOrDiv", "pow", "unaryMinus", "atom",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CalculatorParser struct {
	*antlr.BaseParser
}

func NewCalculatorParser(input antlr.TokenStream) *CalculatorParser {
	this := new(CalculatorParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Calculator.g4"

	return this
}

// CalculatorParser tokens.
const (
	CalculatorParserEOF    = antlr.TokenEOF
	CalculatorParserINT    = 1
	CalculatorParserDOUBLE = 2
	CalculatorParserPI     = 3
	CalculatorParserE      = 4
	CalculatorParserPOW    = 5
	CalculatorParserNL     = 6
	CalculatorParserWS     = 7
	CalculatorParserID     = 8
	CalculatorParserPLUS   = 9
	CalculatorParserEQUAL  = 10
	CalculatorParserMINUS  = 11
	CalculatorParserMULT   = 12
	CalculatorParserDIV    = 13
	CalculatorParserLPAR   = 14
	CalculatorParserRPAR   = 15
)

// CalculatorParser rules.
const (
	CalculatorParserRULE_input       = 0
	CalculatorParserRULE_setVar      = 1
	CalculatorParserRULE_plusOrMinus = 2
	CalculatorParserRULE_multOrDiv   = 3
	CalculatorParserRULE_pow         = 4
	CalculatorParserRULE_unaryMinus  = 5
	CalculatorParserRULE_atom        = 6
)

// IInputContext is an interface to support dynamic dispatch.
type IInputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputContext differentiates from other interfaces.
	IsInputContext()
}

type InputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputContext() *InputContext {
	var p = new(InputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_input
	return p
}

func (*InputContext) IsInputContext() {}

func NewInputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputContext {
	var p = new(InputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_input

	return p
}

func (s *InputContext) GetParser() antlr.Parser { return s.parser }

func (s *InputContext) CopyFrom(ctx *InputContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CalculateContext struct {
	*InputContext
}

func NewCalculateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CalculateContext {
	var p = new(CalculateContext)

	p.InputContext = NewEmptyInputContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InputContext))

	return p
}

func (s *CalculateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculateContext) PlusOrMinus() IPlusOrMinusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlusOrMinusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlusOrMinusContext)
}

func (s *CalculateContext) EOF() antlr.TerminalNode {
	return s.GetToken(CalculatorParserEOF, 0)
}

func (s *CalculateContext) NL() antlr.TerminalNode {
	return s.GetToken(CalculatorParserNL, 0)
}

func (s *CalculateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitCalculate(s)

	default:
		return t.VisitChildren(s)
	}
}

type ToSetVarContext struct {
	*InputContext
}

func NewToSetVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToSetVarContext {
	var p = new(ToSetVarContext)

	p.InputContext = NewEmptyInputContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InputContext))

	return p
}

func (s *ToSetVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToSetVarContext) SetVar() ISetVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetVarContext)
}

func (s *ToSetVarContext) NL() antlr.TerminalNode {
	return s.GetToken(CalculatorParserNL, 0)
}

func (s *ToSetVarContext) Input() IInputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputContext)
}

func (s *ToSetVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitToSetVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Input() (localctx IInputContext) {
	localctx = NewInputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalculatorParserRULE_input)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewToSetVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(14)
			p.SetVar()
		}
		{
			p.SetState(15)
			p.Match(CalculatorParserNL)
		}
		{
			p.SetState(16)
			p.Input()
		}

	case 2:
		localctx = NewCalculateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(18)
			p.plusOrMinus(0)
		}
		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CalculatorParserNL {
			{
				p.SetState(19)
				p.Match(CalculatorParserNL)
			}

		}
		{
			p.SetState(22)
			p.Match(CalculatorParserEOF)
		}

	}

	return localctx
}

// ISetVarContext is an interface to support dynamic dispatch.
type ISetVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetVarContext differentiates from other interfaces.
	IsSetVarContext()
}

type SetVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetVarContext() *SetVarContext {
	var p = new(SetVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_setVar
	return p
}

func (*SetVarContext) IsSetVarContext() {}

func NewSetVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetVarContext {
	var p = new(SetVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_setVar

	return p
}

func (s *SetVarContext) GetParser() antlr.Parser { return s.parser }

func (s *SetVarContext) CopyFrom(ctx *SetVarContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SetVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SetVariableContext struct {
	*SetVarContext
}

func NewSetVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetVariableContext {
	var p = new(SetVariableContext)

	p.SetVarContext = NewEmptySetVarContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SetVarContext))

	return p
}

func (s *SetVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetVariableContext) ID() antlr.TerminalNode {
	return s.GetToken(CalculatorParserID, 0)
}

func (s *SetVariableContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(CalculatorParserEQUAL, 0)
}

func (s *SetVariableContext) PlusOrMinus() IPlusOrMinusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlusOrMinusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlusOrMinusContext)
}

func (s *SetVariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitSetVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) SetVar() (localctx ISetVarContext) {
	localctx = NewSetVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CalculatorParserRULE_setVar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewSetVariableContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(CalculatorParserID)
	}
	{
		p.SetState(27)
		p.Match(CalculatorParserEQUAL)
	}
	{
		p.SetState(28)
		p.plusOrMinus(0)
	}

	return localctx
}

// IPlusOrMinusContext is an interface to support dynamic dispatch.
type IPlusOrMinusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPlusOrMinusContext differentiates from other interfaces.
	IsPlusOrMinusContext()
}

type PlusOrMinusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPlusOrMinusContext() *PlusOrMinusContext {
	var p = new(PlusOrMinusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_plusOrMinus
	return p
}

func (*PlusOrMinusContext) IsPlusOrMinusContext() {}

func NewPlusOrMinusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PlusOrMinusContext {
	var p = new(PlusOrMinusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_plusOrMinus

	return p
}

func (s *PlusOrMinusContext) GetParser() antlr.Parser { return s.parser }

func (s *PlusOrMinusContext) CopyFrom(ctx *PlusOrMinusContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PlusOrMinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusOrMinusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ToMultOrDivContext struct {
	*PlusOrMinusContext
}

func NewToMultOrDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToMultOrDivContext {
	var p = new(ToMultOrDivContext)

	p.PlusOrMinusContext = NewEmptyPlusOrMinusContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PlusOrMinusContext))

	return p
}

func (s *ToMultOrDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToMultOrDivContext) MultOrDiv() IMultOrDivContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultOrDivContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultOrDivContext)
}

func (s *ToMultOrDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitToMultOrDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type PlusContext struct {
	*PlusOrMinusContext
}

func NewPlusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlusContext {
	var p = new(PlusContext)

	p.PlusOrMinusContext = NewEmptyPlusOrMinusContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PlusOrMinusContext))

	return p
}

func (s *PlusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlusContext) PlusOrMinus() IPlusOrMinusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlusOrMinusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlusOrMinusContext)
}

func (s *PlusContext) PLUS() antlr.TerminalNode {
	return s.GetToken(CalculatorParserPLUS, 0)
}

func (s *PlusContext) MultOrDiv() IMultOrDivContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultOrDivContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultOrDivContext)
}

func (s *PlusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitPlus(s)

	default:
		return t.VisitChildren(s)
	}
}

type MinusContext struct {
	*PlusOrMinusContext
}

func NewMinusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinusContext {
	var p = new(MinusContext)

	p.PlusOrMinusContext = NewEmptyPlusOrMinusContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PlusOrMinusContext))

	return p
}

func (s *MinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinusContext) PlusOrMinus() IPlusOrMinusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlusOrMinusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlusOrMinusContext)
}

func (s *MinusContext) MINUS() antlr.TerminalNode {
	return s.GetToken(CalculatorParserMINUS, 0)
}

func (s *MinusContext) MultOrDiv() IMultOrDivContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultOrDivContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultOrDivContext)
}

func (s *MinusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitMinus(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) PlusOrMinus() (localctx IPlusOrMinusContext) {
	return p.plusOrMinus(0)
}

func (p *CalculatorParser) plusOrMinus(_p int) (localctx IPlusOrMinusContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPlusOrMinusContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPlusOrMinusContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, CalculatorParserRULE_plusOrMinus, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewToMultOrDivContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(31)
		p.multOrDiv(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPlusContext(p, NewPlusOrMinusContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_plusOrMinus)
				p.SetState(33)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(34)
					p.Match(CalculatorParserPLUS)
				}
				{
					p.SetState(35)
					p.multOrDiv(0)
				}

			case 2:
				localctx = NewMinusContext(p, NewPlusOrMinusContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_plusOrMinus)
				p.SetState(36)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(37)
					p.Match(CalculatorParserMINUS)
				}
				{
					p.SetState(38)
					p.multOrDiv(0)
				}

			}

		}
		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IMultOrDivContext is an interface to support dynamic dispatch.
type IMultOrDivContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultOrDivContext differentiates from other interfaces.
	IsMultOrDivContext()
}

type MultOrDivContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultOrDivContext() *MultOrDivContext {
	var p = new(MultOrDivContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_multOrDiv
	return p
}

func (*MultOrDivContext) IsMultOrDivContext() {}

func NewMultOrDivContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultOrDivContext {
	var p = new(MultOrDivContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_multOrDiv

	return p
}

func (s *MultOrDivContext) GetParser() antlr.Parser { return s.parser }

func (s *MultOrDivContext) CopyFrom(ctx *MultOrDivContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MultOrDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultOrDivContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultiplicationContext struct {
	*MultOrDivContext
}

func NewMultiplicationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicationContext {
	var p = new(MultiplicationContext)

	p.MultOrDivContext = NewEmptyMultOrDivContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MultOrDivContext))

	return p
}

func (s *MultiplicationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicationContext) MultOrDiv() IMultOrDivContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultOrDivContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultOrDivContext)
}

func (s *MultiplicationContext) MULT() antlr.TerminalNode {
	return s.GetToken(CalculatorParserMULT, 0)
}

func (s *MultiplicationContext) Pow() IPowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPowContext)
}

func (s *MultiplicationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitMultiplication(s)

	default:
		return t.VisitChildren(s)
	}
}

type DivisionContext struct {
	*MultOrDivContext
}

func NewDivisionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivisionContext {
	var p = new(DivisionContext)

	p.MultOrDivContext = NewEmptyMultOrDivContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MultOrDivContext))

	return p
}

func (s *DivisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivisionContext) MultOrDiv() IMultOrDivContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultOrDivContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultOrDivContext)
}

func (s *DivisionContext) DIV() antlr.TerminalNode {
	return s.GetToken(CalculatorParserDIV, 0)
}

func (s *DivisionContext) Pow() IPowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPowContext)
}

func (s *DivisionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitDivision(s)

	default:
		return t.VisitChildren(s)
	}
}

type ToPowContext struct {
	*MultOrDivContext
}

func NewToPowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToPowContext {
	var p = new(ToPowContext)

	p.MultOrDivContext = NewEmptyMultOrDivContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MultOrDivContext))

	return p
}

func (s *ToPowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToPowContext) Pow() IPowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPowContext)
}

func (s *ToPowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitToPow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) MultOrDiv() (localctx IMultOrDivContext) {
	return p.multOrDiv(0)
}

func (p *CalculatorParser) multOrDiv(_p int) (localctx IMultOrDivContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMultOrDivContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMultOrDivContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, CalculatorParserRULE_multOrDiv, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewToPowContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(45)
		p.Pow()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(53)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicationContext(p, NewMultOrDivContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_multOrDiv)
				p.SetState(47)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(48)
					p.Match(CalculatorParserMULT)
				}
				{
					p.SetState(49)
					p.Pow()
				}

			case 2:
				localctx = NewDivisionContext(p, NewMultOrDivContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, CalculatorParserRULE_multOrDiv)
				p.SetState(50)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(51)
					p.Match(CalculatorParserDIV)
				}
				{
					p.SetState(52)
					p.Pow()
				}

			}

		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IPowContext is an interface to support dynamic dispatch.
type IPowContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPowContext differentiates from other interfaces.
	IsPowContext()
}

type PowContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowContext() *PowContext {
	var p = new(PowContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_pow
	return p
}

func (*PowContext) IsPowContext() {}

func NewPowContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowContext {
	var p = new(PowContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_pow

	return p
}

func (s *PowContext) GetParser() antlr.Parser { return s.parser }

func (s *PowContext) CopyFrom(ctx *PowContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PowerContext struct {
	*PowContext
}

func NewPowerContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowerContext {
	var p = new(PowerContext)

	p.PowContext = NewEmptyPowContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PowContext))

	return p
}

func (s *PowerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowerContext) UnaryMinus() IUnaryMinusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryMinusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryMinusContext)
}

func (s *PowerContext) POW() antlr.TerminalNode {
	return s.GetToken(CalculatorParserPOW, 0)
}

func (s *PowerContext) Pow() IPowContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPowContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPowContext)
}

func (s *PowerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitPower(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Pow() (localctx IPowContext) {
	localctx = NewPowContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CalculatorParserRULE_pow)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewPowerContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.UnaryMinus()
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(59)
			p.Match(CalculatorParserPOW)
		}
		{
			p.SetState(60)
			p.Pow()
		}

	}

	return localctx
}

// IUnaryMinusContext is an interface to support dynamic dispatch.
type IUnaryMinusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryMinusContext differentiates from other interfaces.
	IsUnaryMinusContext()
}

type UnaryMinusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryMinusContext() *UnaryMinusContext {
	var p = new(UnaryMinusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_unaryMinus
	return p
}

func (*UnaryMinusContext) IsUnaryMinusContext() {}

func NewUnaryMinusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryMinusContext {
	var p = new(UnaryMinusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_unaryMinus

	return p
}

func (s *UnaryMinusContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryMinusContext) CopyFrom(ctx *UnaryMinusContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *UnaryMinusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ChangeSignContext struct {
	*UnaryMinusContext
}

func NewChangeSignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ChangeSignContext {
	var p = new(ChangeSignContext)

	p.UnaryMinusContext = NewEmptyUnaryMinusContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnaryMinusContext))

	return p
}

func (s *ChangeSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChangeSignContext) MINUS() antlr.TerminalNode {
	return s.GetToken(CalculatorParserMINUS, 0)
}

func (s *ChangeSignContext) UnaryMinus() IUnaryMinusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryMinusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryMinusContext)
}

func (s *ChangeSignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitChangeSign(s)

	default:
		return t.VisitChildren(s)
	}
}

type ToAtomContext struct {
	*UnaryMinusContext
}

func NewToAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ToAtomContext {
	var p = new(ToAtomContext)

	p.UnaryMinusContext = NewEmptyUnaryMinusContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UnaryMinusContext))

	return p
}

func (s *ToAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToAtomContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *ToAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitToAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) UnaryMinus() (localctx IUnaryMinusContext) {
	localctx = NewUnaryMinusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CalculatorParserRULE_unaryMinus)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalculatorParserMINUS:
		localctx = NewChangeSignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.Match(CalculatorParserMINUS)
		}
		{
			p.SetState(64)
			p.UnaryMinus()
		}

	case CalculatorParserINT, CalculatorParserDOUBLE, CalculatorParserPI, CalculatorParserE, CalculatorParserID, CalculatorParserLPAR:
		localctx = NewToAtomContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)
			p.Atom()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalculatorParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalculatorParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) CopyFrom(ctx *AtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConstantPIContext struct {
	*AtomContext
}

func NewConstantPIContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantPIContext {
	var p = new(ConstantPIContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ConstantPIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantPIContext) PI() antlr.TerminalNode {
	return s.GetToken(CalculatorParserPI, 0)
}

func (s *ConstantPIContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitConstantPI(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableContext struct {
	*AtomContext
}

func NewVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableContext {
	var p = new(VariableContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ID() antlr.TerminalNode {
	return s.GetToken(CalculatorParserID, 0)
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

type BracesContext struct {
	*AtomContext
}

func NewBracesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracesContext {
	var p = new(BracesContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *BracesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracesContext) LPAR() antlr.TerminalNode {
	return s.GetToken(CalculatorParserLPAR, 0)
}

func (s *BracesContext) PlusOrMinus() IPlusOrMinusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPlusOrMinusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPlusOrMinusContext)
}

func (s *BracesContext) RPAR() antlr.TerminalNode {
	return s.GetToken(CalculatorParserRPAR, 0)
}

func (s *BracesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitBraces(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstantEContext struct {
	*AtomContext
}

func NewConstantEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantEContext {
	var p = new(ConstantEContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *ConstantEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantEContext) E() antlr.TerminalNode {
	return s.GetToken(CalculatorParserE, 0)
}

func (s *ConstantEContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitConstantE(s)

	default:
		return t.VisitChildren(s)
	}
}

type DoubleContext struct {
	*AtomContext
}

func NewDoubleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleContext {
	var p = new(DoubleContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *DoubleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(CalculatorParserDOUBLE, 0)
}

func (s *DoubleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitDouble(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	*AtomContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.AtomContext = NewEmptyAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AtomContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(CalculatorParserINT, 0)
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalculatorVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalculatorParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CalculatorParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalculatorParserPI:
		localctx = NewConstantPIContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Match(CalculatorParserPI)
		}

	case CalculatorParserE:
		localctx = NewConstantEContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.Match(CalculatorParserE)
		}

	case CalculatorParserDOUBLE:
		localctx = NewDoubleContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(70)
			p.Match(CalculatorParserDOUBLE)
		}

	case CalculatorParserINT:
		localctx = NewIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(71)
			p.Match(CalculatorParserINT)
		}

	case CalculatorParserID:
		localctx = NewVariableContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(72)
			p.Match(CalculatorParserID)
		}

	case CalculatorParserLPAR:
		localctx = NewBracesContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(73)
			p.Match(CalculatorParserLPAR)
		}
		{
			p.SetState(74)
			p.plusOrMinus(0)
		}
		{
			p.SetState(75)
			p.Match(CalculatorParserRPAR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *CalculatorParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *PlusOrMinusContext = nil
		if localctx != nil {
			t = localctx.(*PlusOrMinusContext)
		}
		return p.PlusOrMinus_Sempred(t, predIndex)

	case 3:
		var t *MultOrDivContext = nil
		if localctx != nil {
			t = localctx.(*MultOrDivContext)
		}
		return p.MultOrDiv_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CalculatorParser) PlusOrMinus_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CalculatorParser) MultOrDiv_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
