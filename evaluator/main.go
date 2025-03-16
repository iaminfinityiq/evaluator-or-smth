package main

import (
	"evaluator/frontend/lexer"
	"evaluator/runtime"
	"fmt"
	"evaluator/frontend/parser"
)

func main() {
	var code string = "1 + 1 * 2 / (5 - 3)"
	var rt runtime.RuntimeResult = lexer.Tokenize(code)

	if rt.Error != nil {
		e := *rt.Error
		runtime.DisplayError(e)
		return
	}

	var p parser.Parser = parser.Parser{rt.Result.([]lexer.Token)}
	rt = p.ParseBlock()
	if rt.Error != nil {
		e := *rt.Error
		runtime.DisplayError(e)
		return
	}

	fmt.Println(rt.Result)
}