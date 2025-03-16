package parser

const (
	BlockStmt = iota
	BinaryExpr
	UnaryExpr
	IntExpr
	DoubleExpr
)

type Statement interface {
	Kind() int
}

type Expression interface {
	Statement
}