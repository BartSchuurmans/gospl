package ast

import (
	"github.com/Minnozz/gospl/token"
)

type Expression interface {
	astNode
}

type BadExpression struct {
}

type Identifier struct {
	Name string
}

type LiteralExpression struct {
	Kind  token.Token
	Value string
}

type BinaryExpression struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

type FunctionCallExpression struct {
	Name      *Identifier
	Arguments []Expression
}

type ParenthesizedExpression struct {
	Expression Expression
}

type TupleExpression struct {
	Left  Expression
	Right Expression
}
