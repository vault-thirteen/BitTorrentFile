package models

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_section(t *testing.T) {
	var aTest = tester.New(t)

	aTest.MustBeEqual(SectionInfo, "info")
	aTest.MustBeEqual(SectionAnnounce, "announce")
	aTest.MustBeEqual(SectionAnnounceList, "announce-list")
	aTest.MustBeEqual(SectionCreationDate, "creation date")
	aTest.MustBeEqual(SectionComment, "comment")
	aTest.MustBeEqual(SectionCreatedBy, "created by")
	aTest.MustBeEqual(SectionEncoding, "encoding")
	aTest.MustBeEqual(SectionPieceLayers, "piece layers")
}
