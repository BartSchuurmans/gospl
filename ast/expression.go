package ast

type Expression interface {
	astNode
}

type BadExpression struct {
}

type LiteralExpression struct {
	Type  *Type
	Value string
}
