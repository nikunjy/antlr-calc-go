package parser

import (
	"math"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CalculatorVisitorImpl struct {
	antlr.ParseTreeVisitor
	vars map[string]float64
}

func NewCalculatorVisitorImpl() CalculatorVisitor {
	return &CalculatorVisitorImpl{
		vars: make(map[string]float64),
	}
}

func (c *CalculatorVisitorImpl) visit(tree antlr.ParseTree) float64 {
	return c.Visit(tree).(float64)
}

func (c *CalculatorVisitorImpl) VisitPlus(ctx *PlusContext) interface{} {
	return c.visit(ctx.PlusOrMinus()) + c.visit(ctx.MultOrDiv())
}

func (c *CalculatorVisitorImpl) VisitMinus(ctx *MinusContext) interface{} {
	return c.visit(ctx.PlusOrMinus()) - c.visit(ctx.MultOrDiv())
}

func (c *CalculatorVisitorImpl) VisitMultiplication(ctx *MultiplicationContext) interface{} {
	return c.visit(ctx.MultOrDiv()) * c.visit(ctx.Pow())
}

func (c *CalculatorVisitorImpl) VisitDivision(ctx *DivisionContext) interface{} {
	return c.visit(ctx.MultOrDiv()) / c.visit(ctx.Pow())
}

func (c *CalculatorVisitorImpl) VisitSetVariable(ctx *SetVariableContext) interface{} {
	value := c.visit(ctx.PlusOrMinus())
	c.vars[ctx.ID().GetText()] = value
	return value
}

func (c *CalculatorVisitorImpl) VisitPower(ctx *PowerContext) interface{} {
	if ctx.Pow() != nil {
		return math.Pow(c.visit(ctx.UnaryMinus()), c.visit(ctx.Pow()))
	}
	return c.visit(ctx.UnaryMinus())
}

func (c *CalculatorVisitorImpl) VisitChangeSign(ctx *ChangeSignContext) interface{} {
	return -1 * c.visit(ctx.UnaryMinus())
}

func (c *CalculatorVisitorImpl) VisitBraces(ctx *BracesContext) interface{} {
	return c.visit(ctx.PlusOrMinus())
}

func (c *CalculatorVisitorImpl) VisitConstantPI(ctx *ConstantPIContext) interface{} {
	return float64(math.Pi)
}

func (c *CalculatorVisitorImpl) VisitConstantE(ctx *ConstantEContext) interface{} {
	return float64(math.E)
}

func (c *CalculatorVisitorImpl) VisitDouble(ctx *DoubleContext) interface{} {
	val, err := strconv.ParseFloat(ctx.DOUBLE().GetText(), 10)
	if err != nil {
		panic(err)
	}
	return val
}

func (c *CalculatorVisitorImpl) VisitInt(ctx *IntContext) interface{} {
	val, err := strconv.ParseFloat(ctx.INT().GetText(), 10)
	if err != nil {
		panic(err)
	}
	return val
}

func (c *CalculatorVisitorImpl) VisitVariable(ctx *VariableContext) interface{} {
	return c.vars[ctx.ID().GetText()]
}

func (c *CalculatorVisitorImpl) VisitCalculate(ctx *CalculateContext) interface{} {
	return c.visit(ctx.PlusOrMinus())
}

func (c *CalculatorVisitorImpl) VisitToAtom(ctx *ToAtomContext) interface{} {
	return c.visit(ctx)
}

func (c *CalculatorVisitorImpl) VisitToMultOrDiv(ctx *ToMultOrDivContext) interface{} {
	return c.visit(ctx)
}

func (c *CalculatorVisitorImpl) VisitToPow(ctx *ToPowContext) interface{} {
	return c.visit(ctx)
}

func (c *CalculatorVisitorImpl) VisitToSetVar(ctx *ToSetVarContext) interface{} {
	return c.visit(ctx)
}

/*


    @Override
    public Double visitCalculate(CalculatorParser.CalculateContext ctx) {
        return visit(ctx.plusOrMinus());
    }
}
*/
