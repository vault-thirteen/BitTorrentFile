package generic

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_RemoveDuplicatesFromList(t *testing.T) {
	aTest := tester.New(t)
	aTest.MustBeEqual(RemoveDuplicatesFromList([]int{1, 3, 5, 3, 7, 1}), []int{1, 3, 5, 7})
	aTest.MustBeEqual(RemoveDuplicatesFromList([]int{7, 7, 1}), []int{7, 1})
}
