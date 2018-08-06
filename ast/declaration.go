package ast

import (
	"github.com/Minnozz/gospl/token"
)

type Declaration interface {
	Node
}

type BadDeclaration struct {
	From, To token.Pos
}

func (d *BadDeclaration) Pos() token.Pos { return d.From }
func (d *BadDeclaration) End() token.Pos { return d.To }

type VariableDeclaration struct {
	Type        Type
	Name        *Identifier
	Initializer Expression
	Semicolon   token.Pos
}

func (d *VariableDeclaration) Pos() token.Pos { return d.Type.Pos() }
func (d *VariableDeclaration) End() token.Pos { return d.Semicolon + 1 }

type FunctionDeclaration struct {
	ReturnType        Type
	Name              *Identifier
	Parameters        *FunctionParameters
	Variables         []*VariableDeclaration
	Statements        []Statement
	CurlyBracketClose token.Pos
}

func (d *FunctionDeclaration) Pos() token.Pos { return d.ReturnType.Pos() }
func (d *FunctionDeclaration) End() token.Pos { return d.CurlyBracketClose + 1 }

type FunctionParameters struct {
	RoundBracketOpen  token.Pos
	Parameters        []*FunctionParameter
	RoundBracketClose token.Pos
}

func (d *FunctionParameters) Pos() token.Pos { return d.RoundBracketOpen }
func (d *FunctionParameters) End() token.Pos { return d.RoundBracketClose + 1 }

type FunctionParameter struct {
	Type Type
	Name *Identifier
}

func (d *FunctionParameter) Pos() token.Pos { return d.Type.Pos() }
func (d *FunctionParameter) End() token.Pos { return d.Name.End() }
