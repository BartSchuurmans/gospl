package ast

import (
	"github.com/Minnozz/gospl/token"
)

type Comment struct {
	TextPos token.Pos
	Text    string
}

func (c *Comment) Pos() token.Pos { return c.TextPos }
func (c *Comment) End() token.Pos { return token.Pos(int(c.TextPos) + len(c.Text)) }
