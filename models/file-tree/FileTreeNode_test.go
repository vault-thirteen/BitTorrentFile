package ft

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_FileTreeNode_AppendChild(t *testing.T) {
	aTest := tester.New(t)

	ftn1 := &FileTreeNode{}
	ftn2 := &FileTreeNode{}
	ftn1.AppendChild(ftn2)
	aTest.MustBeEqual(ftn1.Children, []*FileTreeNode{ftn2})
}
