package ast

import (
	"github.com/Minnozz/gospl/token"
)

type Statement interface {
	Node
}

type BadStatement struct {
	From, To token.Pos
}

func (s *BadStatement) Pos() token.Pos { return s.From }
func (s *BadStatement) End() token.Pos { return s.To }

type BlockStatement struct {
	CurlyBracketOpen  token.Pos
	List              []Statement
	CurlyBracketClose token.Pos
}

func (s *BlockStatement) Pos() token.Pos { return s.CurlyBracketOpen }
func (s *BlockStatement) End() token.Pos { return s.CurlyBracketClose + 1 }

type ReturnStatement struct {
	Return    token.Pos
	Value     Expression
	Semicolon token.Pos
}

func (s *ReturnStatement) Pos() token.Pos { return s.Return }
func (s *ReturnStatement) End() token.Pos { return s.Semicolon + 1 }

type IfStatement struct {
	If        token.Pos
	Condition Expression
	Body      Statement
	Else      Statement // Can be nil
}

func (s *IfStatement) Pos() token.Pos { return s.If }
func (s *IfStatement) End() token.Pos {
	if s.Else != nil {
		return s.Else.End()
	}
	return s.Body.End()
}

type WhileStatement struct {
	While     token.Pos
	Condition Expression
	Body      Statement
}

func (s *WhileStatement) Pos() token.Pos { return s.While }
func (s *WhileStatement) End() token.Pos { return s.Body.End() }

type AssignmentStatement struct {
	Name      *Identifier
	Value     Expression
	Semicolon token.Pos
}

func (s *AssignmentStatement) Pos() token.Pos { return s.Name.Pos() }
func (s *AssignmentStatement) End() token.Pos { return s.Semicolon + 1 }

type FunctionCallStatement struct {
	FunctionCall *FunctionCallExpression
	Semicolon    token.Pos
}

func (s *FunctionCallStatement) Pos() token.Pos { return s.FunctionCall.Pos() }
func (s *FunctionCallStatement) End() token.Pos { return s.Semicolon + 1 }
