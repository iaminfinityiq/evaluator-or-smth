package interpreter

import (
	"evaluator/frontend/parser"
	"evaluator/runtime"
)

func Evaluate(ast_node parser.Statement) runtime.RuntimeResult {
	var rt runtime.RuntimeResult
	switch ast_node.Kind() {
	case parser.BlockStmt:
		rt = EvaluateBlock(ast_node.(parser.Block))
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(rt.Result)
	case parser.IntExpr:
		rt = EvaluateInt(ast_node.(parser.Int))
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(rt.Result)
	case parser.DoubleExpr:
		rt = EvaluateDouble(ast_node.(parser.Double))
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(rt.Result)
	case parser.BinaryExpr:
		rt = EvaluateBinaryExpression(ast_node.(parser.BinaryExpression))
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(rt.Result)
	case parser.UnaryExpr:
		rt = EvaluateUnaryExpression(ast_node.(parser.UnaryExpression))
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(rt.Result)
	default:
		return runtime.Failure(runtime.GoError{"This is probably the rarest error that you have ever met. If you encounter one of these errors, please take a screenshot of that and save it as a part of your memories. This error is occured because there are nodes that are not implemented but for some reason got in your code. Please report to us if you encountered these..."})
	}
}