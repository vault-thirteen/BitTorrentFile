package hash

import (
	"encoding/hex"
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

//TODO

func Test_CalculateSha1(t *testing.T) {
	const (
		Data        string = "Just a Test."
		HashSumText string = "7B708EF0A8EFED41F005C67546A9467BF612A145"
	)

	var aTest = tester.New(t)

	var (
		ba                    []byte
		data                  []byte
		err                   error
		expectedResultAsBytes Sha1Sum
		expectedResultAsText  string
		resultAsBytes         Sha1Sum
		resultAsText          string
	)

	// Test #1.
	{
		data = []byte(Data)
		expectedResultAsText = HashSumText
		ba, err = hex.DecodeString(HashSumText)
		aTest.MustBeNoError(err)
		copy(expectedResultAsBytes[:], ba)
		resultAsBytes, resultAsText = CalculateSha1(data)
		aTest.MustBeEqual(resultAsText, expectedResultAsText)
		aTest.MustBeEqual(resultAsBytes, expectedResultAsBytes)
	}
}
