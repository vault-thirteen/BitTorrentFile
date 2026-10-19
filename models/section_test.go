package models

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_section(t *testing.T) {
	var aTest = tester.New(t)

	aTest.MustBeEqual(Section_Info, "info")
	aTest.MustBeEqual(Section_Announce, "announce")
	aTest.MustBeEqual(Section_AnnounceList, "announce-list")
	aTest.MustBeEqual(Section_CreationDate, "creation date")
	aTest.MustBeEqual(Section_Comment, "comment")
	aTest.MustBeEqual(Section_CreatedBy, "created by")
	aTest.MustBeEqual(Section_Encoding, "encoding")
	aTest.MustBeEqual(Section_PieceLayers, "piece layers")
}
