package btf

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_format(t *testing.T) {
	var aTest = tester.New(t)

	aTest.MustBeEqual(SectionInfo, "info")
	aTest.MustBeEqual(SectionAnnounce, "announce")
	aTest.MustBeEqual(SectionAnnounceList, "announce-list")
	aTest.MustBeEqual(SectionCreationDate, "creation date")
	aTest.MustBeEqual(SectionComment, "comment")
	aTest.MustBeEqual(SectionCreatedBy, "created by")
	aTest.MustBeEqual(SectionEncoding, "encoding")
	aTest.MustBeEqual(SectionPieceLayers, "piece layers")

	aTest.MustBeEqual(FieldFileTree, "file tree")
	aTest.MustBeEqual(FieldFiles, "files")
	aTest.MustBeEqual(FieldLength, "length")
	aTest.MustBeEqual(FieldMD5Sum, "md5sum")
	aTest.MustBeEqual(FieldMetaVersion, "meta version")
	aTest.MustBeEqual(FieldName, "name")
	aTest.MustBeEqual(FieldPath, "path")
	aTest.MustBeEqual(FieldPieceLayers, "piece layers")
	aTest.MustBeEqual(FieldPieceLength, "piece length")
	aTest.MustBeEqual(FieldPieceRoot, "piece root")
	aTest.MustBeEqual(FieldPieces, "pieces")
	aTest.MustBeEqual(FieldPrivate, "private")
}
