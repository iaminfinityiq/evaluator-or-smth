package interpreter

import (
	"evaluator/frontend/parser"
	"evaluator/runtime"
)

func EvaluateBlock(ast_node parser.Block) runtime.RuntimeResult {
	var last_returned any = nil
	for _, statement := range ast_node.Body {
		var rt runtime.RuntimeResult = Evaluate(statement)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		last_returned = rt.Result
	}

	return runtime.Success(last_returned)
}