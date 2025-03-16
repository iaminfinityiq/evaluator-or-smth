package parser

type Block struct {
	Body []Statement
}

func (b Block) Kind() int {
	return BlockStmt
}