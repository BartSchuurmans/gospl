package scanner

import (
	"github.com/Minnozz/gospl/token"
)

type Error struct {
	Pos token.Position
	Msg string
}

func (e Error) Error() string {
	return e.Pos.String() + ": " + e.Msg
}

type ErrorList []*Error

func (el *ErrorList) Add(pos token.Position, msg string) {
	*el = append(*el, &Error{pos, msg})
}
