package parser

import (
	"evaluator/frontend/lexer"
	"evaluator/helpers"
	"evaluator/runtime"
	"strconv"
)

type Parser struct {
	Tokens []lexer.Token
}

func (p *Parser) At() lexer.Token {
	return p.Tokens[0]
}

func (p *Parser) Eat() lexer.Token {
	var returned lexer.Token = p.At()
	p.Tokens = p.Tokens[1:]

	return returned
}

func (p *Parser) Expect(expected helpers.IntSet) runtime.RuntimeResult {
	var returned lexer.Token = p.Eat()
	if !expected.Contains(returned.TokenType) {
		return runtime.Failure(runtime.SyntaxError{"Invalid syntax!"})
	}

	return runtime.Success(returned)
}

func (p *Parser) IsEOF() bool {
	return p.At().TokenType == lexer.EOF
}

func (p *Parser) NotEOF() bool {
	return !p.IsEOF()
}

func (p *Parser) ParseBlock() runtime.RuntimeResult {
	var block Block = Block{[]Statement{}}
	for p.NotEOF() {
		var rt runtime.RuntimeResult = p.ParseStatement()
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		block.Body = append(block.Body, rt.Result.(Statement))
	}

	return runtime.Success(block)
}

func (p *Parser) ParseStatement() runtime.RuntimeResult {
	var rt runtime.RuntimeResult
	switch p.At().TokenType {
	default:
		rt = p.ParseExpression()
	}

	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	return runtime.Success(rt.Result)
}

func (p *Parser) ParseExpression() runtime.RuntimeResult {
	var rt runtime.RuntimeResult = p.ParseAdditiveExpression()
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	return runtime.Success(rt.Result)
}

func (p *Parser) ParseAdditiveExpression() runtime.RuntimeResult {
	var rt runtime.RuntimeResult = p.ParseMultiplicativeExpression()
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var left Expression = rt.Result.(Expression)
	var operators helpers.IntSet = helpers.IntSet{make(map[int]bool)}
	operators.Add(lexer.Plus)
	operators.Add(lexer.Minus)

	for operators.Contains(p.At().TokenType) {
		var operator int = p.Eat().TokenType
		rt = p.ParseMultiplicativeExpression()
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		left = BinaryExpression{
			left,
			operator,
			rt.Result.(Expression),
		}
	}

	return runtime.Success(left)
}

func (p *Parser) ParseMultiplicativeExpression() runtime.RuntimeResult {
	var rt runtime.RuntimeResult = p.ParseUnaryExpression()
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	var left Expression = rt.Result.(Expression)
	var operators helpers.IntSet = helpers.IntSet{make(map[int]bool)}
	operators.Add(lexer.Multiply)
	operators.Add(lexer.Divide)

	for operators.Contains(p.At().TokenType) {
		var operator int = p.Eat().TokenType
		rt = p.ParseUnaryExpression()
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		left = BinaryExpression{
			left,
			operator,
			rt.Result.(Expression),
		}
	}

	return runtime.Success(left)
}

func (p *Parser) ParseUnaryExpression() runtime.RuntimeResult {
	var signs helpers.IntSet = helpers.IntSet{make(map[int]bool)}
	signs.Add(lexer.Plus)
	signs.Add(lexer.Minus)

	var rt runtime.RuntimeResult
	if !signs.Contains(p.At().TokenType) {
		rt = p.ParsePrimaryExpression()
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(rt.Result)
	}

	var sign int = lexer.Plus
	for signs.Contains(p.At().TokenType) {
		var operator int = p.Eat().TokenType
		if operator == lexer.Plus {
			continue
		}

		if sign == lexer.Plus {
			sign = lexer.Minus
		} else {
			sign = lexer.Plus
		}
	}

	rt = p.ParsePrimaryExpression()
	if rt.Error != nil {
		return runtime.Failure(*rt.Error)
	}

	return runtime.Success(UnaryExpression{
		sign,
		rt.Result.(Expression),
	})
}

func (p *Parser) ParsePrimaryExpression() runtime.RuntimeResult {
	var token lexer.Token = p.Eat()
	switch token.TokenType {
	case lexer.Int:
		value, _ := strconv.ParseInt(token.Value, 10, 64)
		return runtime.Success(Int{value})
	case lexer.Double:
		value, _ := strconv.ParseFloat(token.Value, 64)
		return runtime.Success(Double{value})
	case lexer.LeftParentheses:
		var rt runtime.RuntimeResult = p.ParseExpression()
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		var value Expression = rt.Result.(Expression)
		var expected helpers.IntSet = helpers.IntSet{make(map[int]bool)}
		expected.Add(lexer.RightParentheses)
		rt = p.Expect(expected)
		if rt.Error != nil {
			return runtime.Failure(*rt.Error)
		}

		return runtime.Success(value)
	default:
		return runtime.Failure(runtime.SyntaxError{"Invalid syntax!"})
	}
}
