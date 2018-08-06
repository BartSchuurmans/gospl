package ast

import (
	"github.com/Minnozz/gospl/token"
)

type Node interface {
	Pos() token.Pos // First character of the node
	End() token.Pos // First character after the node
}
