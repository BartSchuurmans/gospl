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
	Name       *Identifier
	Type       *FunctionType
	Variables  []*VariableDeclaration
	Statements []Statement
}

type FunctionType struct {
	Parameters *FunctionParameters
	Return     Type
}

type FunctionParameters struct {
	Parameters []*FunctionParameter
}

type FunctionParameter struct {
	Type Type
	Name *Identifier
}
