package ast

type Type interface {
	astNode
}

type BadType struct {
}

type NamedType struct {
	Name *Identifier
}

type TupleType struct {
	Left  Type
	Right Type
}

type ListType struct {
	ElementType Type
}
