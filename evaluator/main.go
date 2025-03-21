package main

import (
	"evaluator/backend/interpreter"
	"evaluator/frontend/lexer"
	"evaluator/frontend/parser"
	"evaluator/runtime"
	"fmt"
)

func main() {
	var code string = "1.5 / 5"
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

	rt = interpreter.Evaluate(rt.Result.(parser.Statement))
	if rt.Error != nil {
		e := *rt.Error
		runtime.DisplayError(e)
		return
	}

	fmt.Println(rt.Result)
}
