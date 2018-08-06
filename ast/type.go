package ast

import (
	"github.com/Minnozz/gospl/token"
)

type Type interface {
	Node
}

type BadType struct {
	From, To token.Pos
}

func (t *BadType) Pos() token.Pos { return t.From }
func (t *BadType) End() token.Pos { return t.To }

type NamedType struct {
	Name *Identifier
}

func (t *NamedType) Pos() token.Pos { return t.Name.Pos() }
func (t *NamedType) End() token.Pos { return t.Name.End() }

type TupleType struct {
	RoundBracketOpen  token.Pos
	Left              Type
	Right             Type
	RoundBracketClose token.Pos
}

func (t *TupleType) Pos() token.Pos { return t.RoundBracketOpen }
func (t *TupleType) End() token.Pos { return t.RoundBracketClose + 1 }

type ListType struct {
	SquareBracketOpen  token.Pos
	ElementType        Type
	SquareBracketClose token.Pos
}

func (t *ListType) Pos() token.Pos { return t.SquareBracketOpen }
func (t *ListType) End() token.Pos { return t.SquareBracketClose + 1 }
