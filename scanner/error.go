package scanner

import (
	"github.com/Minnozz/gompiler/token"
)

type Error struct {
	Pos token.Position
	Msg string
}

func (e Error) Error() string {
	return e.Pos.String() + ": " + e.Msg
}
