package ast

type Statement interface {
	astNode
}

type BadStatement struct {
}

type VariableDeclarationStatement struct {
	Declaration *VariableDeclaration
}

type BlockStatement struct {
	List []Statement
}

type ReturnStatement struct {
	Value Expression
}
