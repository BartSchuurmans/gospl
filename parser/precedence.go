package parser

import (
	"github.com/Minnozz/gospl/token"
)

// Precedence groups, from lowest to highest
const (
	binaryBoolean        = iota // &&, ||
	binaryComparison            // ==, !=, <, >, <=, >=
	unaryNot                    // !
	binaryColon                 // :
	binaryAddition              // +, -
	binaryMultiplication        // *, /, %
	unaryMinus                  // -
	highestPrecedence           // default
)

func unaryOperatorPrecedence(op token.Token) int {
	switch op {
	case token.NOT:
		return unaryNot
	case token.MINUS:
		return unaryMinus
	default:
		panic("invalid unary operator: " + op.String())
	}
}

func binaryOperatorPrecedence(op token.Token) int {
	switch op {
	case token.PLUS, token.MINUS:
		return binaryAddition
	case token.MULTIPLY, token.DIVIDE, token.MODULO:
		return binaryMultiplication
	case token.EQUALS, token.LESS_THAN, token.GREATER_THAN, token.LESS_THAN_EQUALS, token.GREATER_THAN_EQUALS, token.NOT_EQUALS:
		return binaryComparison
	case token.AND, token.OR:
		return binaryBoolean
	case token.COLON:
		return binaryColon
	default:
		panic("invalid binary operator: " + op.String())
	}
}
