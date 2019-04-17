// Code generated from Calculator.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Calculator

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCalculatorVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalculatorVisitor) VisitToSetVar(ctx *ToSetVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitCalculate(ctx *CalculateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitSetVariable(ctx *SetVariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitToMultOrDiv(ctx *ToMultOrDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitPlus(ctx *PlusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitMinus(ctx *MinusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitMultiplication(ctx *MultiplicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitDivision(ctx *DivisionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitToPow(ctx *ToPowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitPower(ctx *PowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitChangeSign(ctx *ChangeSignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitToAtom(ctx *ToAtomContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitConstantPI(ctx *ConstantPIContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitConstantE(ctx *ConstantEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitDouble(ctx *DoubleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalculatorVisitor) VisitBraces(ctx *BracesContext) interface{} {
	return v.VisitChildren(ctx)
}
