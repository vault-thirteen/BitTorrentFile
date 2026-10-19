package ft

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_NodeRoute_NewNodeRoute(t *testing.T) {
	aTest := tester.New(t)

	nr := NewNodeRoute(nil)
	aTest.MustBeEqual([]*FileTreeNode(nr), []*FileTreeNode{nil})
}

func Test_NodeRoute_ConvertToPath(t *testing.T) {
	aTest := tester.New(t)

	ftn1 := &FileTreeNode{Name: "a"}
	ftn2 := &FileTreeNode{Name: "b"}
	ftn3 := &FileTreeNode{Name: "c"}
	nr := NewNodeRoute(ftn1)
	nr.AddNode(ftn2)
	nr.AddNode(ftn3)
	aTest.MustBeEqual(nr.ConvertToPath(), []string{"a", "b", "c"})
}

func Test_NodeRoute_AddNode(t *testing.T) {
	aTest := tester.New(t)

	ftn1 := &FileTreeNode{Name: "a"}
	ftn2 := &FileTreeNode{Name: "b"}
	nr := NewNodeRoute(ftn1)
	nr.AddNode(ftn2)
	aTest.MustBeEqual([]*FileTreeNode(nr), []*FileTreeNode{ftn1, ftn2})
}

func Test_NodeRoute_RemoveNode(t *testing.T) {
	aTest := tester.New(t)

	ftn1 := &FileTreeNode{Name: "a"}
	ftn2 := &FileTreeNode{Name: "b"}
	ftn3 := &FileTreeNode{Name: "c"}
	nr := NewNodeRoute(ftn1)
	nr.AddNode(ftn2)
	nr.AddNode(ftn3)
	nr.RemoveNode()
	aTest.MustBeEqual([]*FileTreeNode(nr), []*FileTreeNode{ftn1, ftn2})
}
