package ast

import (
	"github.com/Minnozz/gospl/token"
)

type Expression interface {
	Node
}

type BadExpression struct {
	From, To token.Pos
}

func (e *BadExpression) Pos() token.Pos { return e.From }
func (e *BadExpression) End() token.Pos { return e.To }

type Identifier struct {
	NamePos token.Pos
	Name    string
}

func (e *Identifier) Pos() token.Pos { return e.NamePos }
func (e *Identifier) End() token.Pos { return token.Pos(int(e.NamePos) + len(e.Name)) }

type LiteralExpression struct {
	ValuePos token.Pos
	Kind     token.Token
	Value    string
}

func (e *LiteralExpression) Pos() token.Pos { return e.ValuePos }
func (e *LiteralExpression) End() token.Pos { return token.Pos(int(e.ValuePos) + len(e.Value)) }

type UnaryExpression struct {
	OperatorPos token.Pos
	Operator    token.Token
	Expression  Expression
}

func (e *UnaryExpression) Pos() token.Pos { return e.OperatorPos }
func (e *UnaryExpression) End() token.Pos { return e.Expression.End() }

type BinaryExpression struct {
	Left     Expression
	Operator token.Token
	Right    Expression
}

func (e *BinaryExpression) Pos() token.Pos { return e.Left.Pos() }
func (e *BinaryExpression) End() token.Pos { return e.Right.End() }

type FunctionCallExpression struct {
	Name              *Identifier
	Arguments         []Expression
	RoundBracketClose token.Pos
}

func (e *FunctionCallExpression) Pos() token.Pos { return e.Name.Pos() }
func (e *FunctionCallExpression) End() token.Pos { return e.RoundBracketClose + 1 }

type ParenthesizedExpression struct {
	RoundBracketOpen  token.Pos
	Expression        Expression
	RoundBracketClose token.Pos
}

func (e *ParenthesizedExpression) Pos() token.Pos { return e.RoundBracketOpen }
func (e *ParenthesizedExpression) End() token.Pos { return e.RoundBracketClose + 1 }

type TupleExpression struct {
	RoundBracketOpen  token.Pos
	Left              Expression
	Right             Expression
	RoundBracketClose token.Pos
}

func (e *TupleExpression) Pos() token.Pos { return e.RoundBracketOpen }
func (e *TupleExpression) End() token.Pos { return e.RoundBracketClose + 1 }
