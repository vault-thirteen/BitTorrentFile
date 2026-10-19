package models

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_File_SanitiseFilePath(t *testing.T) {
	var aTest = tester.New(t)
	f := File{
		Path: []string{"a", "b", ".", "c", "d", "..", "e", "f"},
	}

	f.SanitiseFilePath()
	aTest.MustBeEqual(f.Path, []string{"a", "b", "c", "d", "e", "f"})
}
