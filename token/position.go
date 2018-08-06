package token

import (
	"fmt"
)

type Pos int

var NoPos Pos = 0

// Start at 1 so the zero value NoPos is distinct from the first character.
var posBase = 1

type FileInfo struct {
	Filename string

	newlines []int // 0-based line index => offset of '\n' in file
}

func (f *FileInfo) AddLine(offset int) {
	f.newlines = append(f.newlines, offset)
}

func (f *FileInfo) Pos(offset int) Pos {
	return Pos(offset + 1)
}

func (f *FileInfo) Position(pos Pos) Position {
	offset := int(pos) - posBase

	lineIndex, lineOffset := 0, 0
	for lineIndex < len(f.newlines) && offset > f.newlines[lineIndex] {
		lineOffset = f.newlines[lineIndex] + 1
		lineIndex++
	}
	return Position{
		Filename: f.Filename,
		Offset:   offset,
		Line:     lineIndex + 1,
		Column:   offset - lineOffset + 1,
	}
}

type Position struct {
	Filename string
	Offset   int
	Line     int
	Column   int
}

func (pos Position) String() string {
	return fmt.Sprintf("%s:%d:%d", pos.Filename, pos.Line, pos.Column)
}
