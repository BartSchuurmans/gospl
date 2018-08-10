package parser

import (
	"github.com/Minnozz/gospl/token"
)

type Precedence int

// Precedence groups, from lowest to highest
const (
	binaryBoolean        Precedence = iota // &&, ||
	binaryComparison                       // ==, !=, <, >, <=, >=
	unaryNot                               // !
	binaryColon                            // :
	binaryAddition                         // +, -
	binaryMultiplication                   // *, /, %
	unaryMinus                             // -
	highestPrecedence                      // default
)

type Associativity int

const (
	LeftAssociative Associativity = iota
	RightAssociative
)

func unaryPrecAssoc(op token.Token) (Precedence, Associativity) {
	switch op {
	case token.NOT:
		return unaryNot, RightAssociative
	case token.MINUS:
		return unaryMinus, RightAssociative
	default:
		panic("invalid unary operator: " + op.String())
	}
}

func binaryPrecAssoc(op token.Token) (Precedence, Associativity) {
	switch op {
	case token.PLUS, token.MINUS:
		return binaryAddition, LeftAssociative
	case token.MULTIPLY, token.DIVIDE, token.MODULO:
		return binaryMultiplication, LeftAssociative
	case token.EQUALS, token.LESS_THAN, token.GREATER_THAN, token.LESS_THAN_EQUALS, token.GREATER_THAN_EQUALS, token.NOT_EQUALS:
		return binaryComparison, LeftAssociative
	case token.AND, token.OR:
		return binaryBoolean, LeftAssociative
	case token.COLON:
		return binaryColon, LeftAssociative
	default:
		panic("invalid binary operator: " + op.String())
	}
}
