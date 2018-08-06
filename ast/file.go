package ast

import (
	"github.com/Minnozz/gospl/token"
)

type File struct {
	Declarations []Declaration
	Comments     []Comment
}

func (f *File) Pos() token.Pos {
	if len(f.Declarations) == 0 {
		return token.NoPos
	}
	return f.Declarations[0].Pos()
}

func (f *File) End() token.Pos {
	if len(f.Declarations) == 0 {
		return token.NoPos
	}
	return f.Declarations[len(f.Declarations)-1].End()
}
