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

type IfStatement struct {
	Condition Expression
	Body      Statement
	Else      Statement
}

type WhileStatement struct {
	Condition Expression
	Body      Statement
}

type AssignmentStatement struct {
	Name  *Identifier
	Value Expression
}

type FunctionCallStatement struct {
	FunctionCall *FunctionCallExpression
}
