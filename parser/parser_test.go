package parser

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Minnozz/gospl/ast"
	"github.com/Minnozz/gospl/token"
)

func TestParserValid(t *testing.T) {
	tests, err := ioutil.ReadDir("../testdata/valid")
	if err != nil {
		t.Fatalf("Error reading test directory: %v", err)
	}

	for _, test := range tests {
		name := test.Name()
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			parseTestFile(t, name)
		})
	}
}

func parseTestFile(t *testing.T, name string) {
	file, err := os.Open("../testdata/valid/" + name)
	if err != nil {
		t.Fatalf("Error opening test %s: %v", name, err)
	}
	defer file.Close()

	src, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Error reading test %s: %v", name, err)
	}

	fileInfo := &token.FileInfo{
		Filename: name,
	}

	p := &Parser{}
	p.Init(fileInfo, src)

	fileNode := p.Parse()
	for _, err := range p.Errors {
		t.Error(err)
	}

	t.Logf("AST:\n%s\n", ast.Print(fileNode, fileInfo))
	t.Logf("Reconstructed source from AST:\n%s\n", ast.PrintSource(fileNode))

	// Check position informtion in AST
	coverage := make([]bool, len(src))
	ast.WalkFunc(fileNode, func(n ast.Node) {
		if n.Pos() == token.NoPos {
			t.Errorf("Pos() empty in AST node %T", n)
		}
		if n.End() == token.NoPos {
			t.Errorf("End() empty in AST node %T", n)
		}
		if n.Pos() != token.NoPos && n.End() != token.NoPos {
			for pos := n.Pos(); pos < n.End(); pos++ {
				position := fileInfo.Position(pos)
				if position.Offset < 0 || position.Offset >= len(src) {
					t.Errorf("Position %v of AST node %+v outside of source file", position, n)
				} else {
					coverage[position.Offset] = true
				}
			}
		}
	})
	for offset, covered := range coverage {
		if ch := src[offset]; ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			// Whitespace is skipped by the scanner
		} else if !covered {
			t.Errorf("Position %v in source file (%q) is not claimed by any AST node", fileInfo.Position(fileInfo.Pos(offset)), ch)
		}
	}
}
