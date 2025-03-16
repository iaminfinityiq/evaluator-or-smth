package lexer

import (
	"evaluator/helpers"
	"evaluator/runtime"
)

const (
	EOF = iota
	Plus
	Minus
	Multiply
	Divide
	LeftParentheses
	RightParentheses
	Int
	Double
)

func Tokenize(snippet string) runtime.RuntimeResult {
	var tokens []Token = []Token{}
	
	var characters []rune = []rune{}
	for _, c := range snippet {
		characters = append(characters, c)
	}

	const DIGITS string = "0123456789."
	const WHITESPACES string = " \n\t"

	for len(characters) > 0 {
		switch characters[0] {
		case '+':
			tokens = append(tokens, Token{Plus, "+"})
			characters = characters[1:]
		case '-':
			tokens = append(tokens, Token{Minus, "-"})
			characters = characters[1:]
		case '*':
			tokens = append(tokens, Token{Multiply, "*"})
			characters = characters[1:]
		case '/':
			tokens = append(tokens, Token{Divide, "/"})
			characters = characters[1:]
		case '(':
			tokens = append(tokens, Token{LeftParentheses, "("})
			characters = characters[1:]
		case ')':
			tokens = append(tokens, Token{RightParentheses, ")"})
			characters = characters[1:]
		default:
			if helpers.RuneInStringChecker(characters[0], WHITESPACES) {
				characters = characters[1:]
				continue
			}

			if helpers.RuneInStringChecker(characters[0], DIGITS) {
				var dot_count int = 0
				var number string = ""
				for helpers.RuneInStringChecker(characters[0], DIGITS) {
					if characters[0] == '.' {
						dot_count++
					}

					number += string(characters[0])
					characters = characters[1:]

					if len(characters) == 0 {
						break
					}
				}

				switch dot_count {
				case 0:
					tokens = append(tokens, Token{Int, number})
				case 1:
					tokens = append(tokens, Token{Double, number})
				default:
					return runtime.Failure(runtime.SyntaxError{"Expected 0 or 1 '.' in a number, got " + string(dot_count) + "/1"})
				}

				continue
			}

			return runtime.Failure(runtime.SyntaxError{"Unexpected character: '" + string(characters[0]) + "'"})
		}
	}

	return runtime.Success(append(tokens, Token{EOF, "EOF"}))
}