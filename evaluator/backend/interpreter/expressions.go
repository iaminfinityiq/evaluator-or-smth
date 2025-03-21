package interpreter

import (
	"evaluator/backend/value_types"
	"evaluator/frontend/lexer"
	"evaluator/frontend/parser"
	"evaluator/helpers"
	"evaluator/runtime"
	"math/big"
)

func EvaluateInt(ast_node parser.Int) runtime.RuntimeResult {
	return runtime.Success(value_types.Fraction{ast_node.Value, 1})
}

func EvaluateDouble(ast_node parser.Double) runtime.RuntimeResult {
	var big big.Rat = *new(big.Rat).SetFloat64(ast_node.Value)
	return runtime.Success(value_types.Fraction{
		big.Num().Int64(),
		big.Denom().Int64(),
	})
}

func EvaluateBinaryExpression(ast_node parser.BinaryExpression) runtime.RuntimeResult {
	var rt runtime.RuntimeResult = Evaluate(ast_node.Left)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var left value_types.Fraction = rt.Result.(value_types.Fraction)

	rt = Evaluate(ast_node.Right)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var right value_types.Fraction = rt.Result.(value_types.Fraction)

	switch ast_node.Operator {
	case lexer.Plus:
		var denominator int64 = helpers.LCM(left.Denominator, right.Denominator)
		var numerator int64 = left.Numerator * denominator / left.Denominator + right.Numerator * denominator / right.Denominator

		var gcd = helpers.GCD(numerator, denominator)
		return runtime.Success(value_types.Fraction{
			numerator / gcd,
			denominator / gcd,
		})
	case lexer.Minus:
		var denominator int64 = helpers.LCM(left.Denominator, right.Denominator)
		var numerator int64 = left.Numerator * denominator / left.Denominator - right.Numerator * denominator / right.Denominator

		var gcd = helpers.GCD(numerator, denominator)
		return runtime.Success(value_types.Fraction{
			numerator / gcd,
			denominator / gcd,
		})
	case lexer.Multiply:
		var numerator int64 = left.Numerator * right.Numerator
		var denominator int64 = left.Denominator * right.Denominator

		var gcd = helpers.GCD(numerator, denominator)
		return runtime.Success(value_types.Fraction{
			numerator / gcd,
			denominator / gcd,
		})
	case lexer.Divide:
		var numerator int64 = left.Numerator * right.Denominator
		var denominator int64 = left.Denominator * right.Numerator

		var gcd = helpers.GCD(numerator, denominator)
		return runtime.Success(value_types.Fraction{
			numerator / gcd,
			denominator / gcd,
		})
	default:
		return runtime.Failure(runtime.GoError{"This is probably the rarest error that you have ever met. If you encounter one of these errors, please take a screenshot of that and save it as a part of your memories. This error is occured because there are operators that you are trying to use but yet not implemented. Please report to us if you encountered these..."})
	}
}

func EvaluateUnaryExpression(ast_node parser.UnaryExpression) runtime.RuntimeResult {
	var rt runtime.RuntimeResult = Evaluate(ast_node.Value)
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	if ast_node.Sign == lexer.Plus {
		return runtime.Success(rt.Result)
	}

	var base value_types.Fraction = rt.Result.(value_types.Fraction)
	return runtime.Success(value_types.Fraction{
		-base.Numerator,
		base.Denominator,
	})
}