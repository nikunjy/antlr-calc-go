// Code generated from Calculator.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Calculator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by CalculatorParser.
type CalculatorVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalculatorParser#ToSetVar.
	VisitToSetVar(ctx *ToSetVarContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Calculate.
	VisitCalculate(ctx *CalculateContext) interface{}

	// Visit a parse tree produced by CalculatorParser#SetVariable.
	VisitSetVariable(ctx *SetVariableContext) interface{}

	// Visit a parse tree produced by CalculatorParser#ToMultOrDiv.
	VisitToMultOrDiv(ctx *ToMultOrDivContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Plus.
	VisitPlus(ctx *PlusContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Minus.
	VisitMinus(ctx *MinusContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Multiplication.
	VisitMultiplication(ctx *MultiplicationContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Division.
	VisitDivision(ctx *DivisionContext) interface{}

	// Visit a parse tree produced by CalculatorParser#ToPow.
	VisitToPow(ctx *ToPowContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Power.
	VisitPower(ctx *PowerContext) interface{}

	// Visit a parse tree produced by CalculatorParser#ChangeSign.
	VisitChangeSign(ctx *ChangeSignContext) interface{}

	// Visit a parse tree produced by CalculatorParser#ToAtom.
	VisitToAtom(ctx *ToAtomContext) interface{}

	// Visit a parse tree produced by CalculatorParser#ConstantPI.
	VisitConstantPI(ctx *ConstantPIContext) interface{}

	// Visit a parse tree produced by CalculatorParser#ConstantE.
	VisitConstantE(ctx *ConstantEContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Double.
	VisitDouble(ctx *DoubleContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by CalculatorParser#Braces.
	VisitBraces(ctx *BracesContext) interface{}
}
