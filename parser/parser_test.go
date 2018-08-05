package parser

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Minnozz/gompiler/token"
)

func TestParserValid(t *testing.T) {
	tests, err := ioutil.ReadDir("../testdata/valid")
	if err != nil {
		t.Fatalf("Error reading test directory: %v", err)
	}

	for _, test := range tests {
		t.Run(test.Name(), func(t *testing.T) {
			parseTestFile(t, test.Name())
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

	info := &token.FileInfo{
		Filename: name,
	}

	p := &Parser{}
	p.Init(info, src)

	f := p.Parse()

	for _, err := range p.Errors {
		t.Error(err)
	}

	t.Logf("Reconstructed source from AST:\n%s\n", f.Print())
}
