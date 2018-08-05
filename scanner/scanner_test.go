package scanner

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/Minnozz/gompiler/token"
)

func TestScannerValid(t *testing.T) {
	tests, err := ioutil.ReadDir("../testdata/valid")
	if err != nil {
		t.Fatalf("Error reading test directory: %v", err)
	}

	for _, test := range tests {
		name := test.Name()
		t.Run(name, func(t *testing.T) {
			scanTestFile(t, name)
		})
	}
}

func scanTestFile(t *testing.T, name string) {
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

	var errors ErrorList

	s := &Scanner{}
	s.Init(info, src, func(pos token.Position, msg string) {
		errors.Add(pos, msg)
	})

	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		if tok == token.INVALID {
			t.Errorf("Error: invalid character scanned @ %v: %+q", info.Position(pos), lit)
		} else if len(lit) > 0 {
			t.Logf("%v %+q @ %v", tok, lit, info.Position(pos))
		} else {
			t.Logf("%v @ %v", tok, info.Position(pos))
		}
	}

	for _, err := range errors {
		t.Error(err)
	}
}
