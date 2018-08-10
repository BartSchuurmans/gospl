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
}
