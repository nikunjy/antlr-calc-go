package parser

import (
	"fmt"
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
		ParseTreeVisitor: &BaseCalculatorVisitor{},
		vars:             make(map[string]float64),
	}
}

func (c *CalculatorVisitorImpl) VisitPlus(ctx *PlusContext) interface{} {
	fmt.Println("Came here to visit plus")
	first := ctx.PlusOrMinus().Accept(c).(float64)
	fmt.Println("first done")
	second := ctx.MultOrDiv().Accept(c).(float64)
	fmt.Println("second done", first, second)
	return first + second
}

func (c *CalculatorVisitorImpl) VisitMinus(ctx *MinusContext) interface{} {
	fmt.Println("Came here to visit minus")
	return ctx.PlusOrMinus().Accept(c).(float64) - ctx.MultOrDiv().Accept(c).(float64)
}

func (c *CalculatorVisitorImpl) VisitMultiplication(ctx *MultiplicationContext) interface{} {
	fmt.Println("Came here to multi")
	return ctx.MultOrDiv().Accept(c).(float64) * ctx.Pow().Accept(c).(float64)
}

func (c *CalculatorVisitorImpl) VisitDivision(ctx *DivisionContext) interface{} {
	return ctx.MultOrDiv().Accept(c).(float64) / ctx.Pow().Accept(c).(float64)
}

func (c *CalculatorVisitorImpl) VisitSetVariable(ctx *SetVariableContext) interface{} {
	value := ctx.PlusOrMinus().Accept(c).(float64)
	c.vars[ctx.ID().GetText()] = value
	return value
}

func (c *CalculatorVisitorImpl) VisitPower(ctx *PowerContext) interface{} {
	if ctx.Pow() != nil {
		return math.Pow(ctx.UnaryMinus().Accept(c).(float64), ctx.Pow().Accept(c).(float64))
	}
	return ctx.UnaryMinus().Accept(c)
}

func (c *CalculatorVisitorImpl) VisitChangeSign(ctx *ChangeSignContext) interface{} {
	return -1 * ctx.UnaryMinus().Accept(c).(float64)
}

func (c *CalculatorVisitorImpl) VisitBraces(ctx *BracesContext) interface{} {
	return ctx.PlusOrMinus().Accept(c)
}

func (c *CalculatorVisitorImpl) VisitConstantPI(ctx *ConstantPIContext) interface{} {
	return float64(math.Pi)
}

func (c *CalculatorVisitorImpl) VisitConstantE(ctx *ConstantEContext) interface{} {
	return float64(math.E)
}

func (c *CalculatorVisitorImpl) VisitDouble(ctx *DoubleContext) interface{} {
	fmt.Println("Came here to visit double")
	val, err := strconv.ParseFloat(ctx.DOUBLE().GetText(), 10)
	if err != nil {
		panic(err)
	}
	return val
}

func (c *CalculatorVisitorImpl) VisitInt(ctx *IntContext) interface{} {
	fmt.Println("Came here to visit int")
	val, err := strconv.ParseFloat(ctx.INT().GetText(), 10)
	if err != nil {
		panic(err)
	}
	return val
}

func (c *CalculatorVisitorImpl) VisitVariable(ctx *VariableContext) interface{} {
	fmt.Println("Got to visit variable")
	return c.vars[ctx.ID().GetText()]
}

func (c *CalculatorVisitorImpl) VisitCalculate(ctx *CalculateContext) interface{} {
	fmt.Println("Came here to visit calculate")
	return ctx.PlusOrMinus().Accept(c)
}

func (c *CalculatorVisitorImpl) VisitToAtom(ctx *ToAtomContext) interface{} {
	fmt.Println("Came here to visit atom")
	return ctx.Atom().Accept(c)
}

func (c *CalculatorVisitorImpl) VisitToMultOrDiv(ctx *ToMultOrDivContext) interface{} {
	fmt.Println("visit to mult or div")
	return ctx.MultOrDiv().Accept(c)
}

func (c *CalculatorVisitorImpl) VisitToPow(ctx *ToPowContext) interface{} {
	return ctx.Pow().Accept(c)
}

func (c *CalculatorVisitorImpl) VisitToSetVar(ctx *ToSetVarContext) interface{} {
	return ctx.SetVar().Accept(c)
}

/*


    @Override
    public Double visitCalculate(CalculatorParser.CalculateContext ctx) {
        return visit(ctx.plusOrMinus());
    }
}
*/
