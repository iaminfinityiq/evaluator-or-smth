package parser

type BinaryExpression struct {
	Left Expression
	Operator int
	Right Expression
}

func (b BinaryExpression) Kind() int {
	return BinaryExpr
}

type UnaryExpression struct {
	Sign int
	Value Expression
}

func (u UnaryExpression) Kind() int {
	return UnaryExpr
}

type Int struct {
	Value int64
}

func (i Int) Kind() int {
	return IntExpr
}

type Double struct {
	Value float64
}

func (i Double) Kind() int {
	return DoubleExpr
}