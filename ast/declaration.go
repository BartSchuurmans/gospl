package ast

type Declaration interface {
	astNode
}

type BadDeclaration struct {
}

type VariableDeclaration struct {
	Type        Type
	Name        *Identifier
	Initializer Expression
}

type FunctionDeclaration struct {
	ReturnType Type
	Name       *Identifier
	Parameters *FunctionParameters
	Variables  []*VariableDeclaration
	Statements []Statement
}

type FunctionParameters struct {
	Parameters []*FunctionParameter
}

type FunctionParameter struct {
	Type Type
	Name *Identifier
}
