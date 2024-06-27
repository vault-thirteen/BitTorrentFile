package models

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_field(t *testing.T) {
	var aTest = tester.New(t)

	aTest.MustBeEqual(FieldCrc32Sum, "crc32sum")
	aTest.MustBeEqual(FieldFileTree, "file tree")
	aTest.MustBeEqual(FieldFiles, "files")
	aTest.MustBeEqual(FieldLength, "length")
	aTest.MustBeEqual(FieldMd5Sum, "md5sum")
	aTest.MustBeEqual(FieldMetaVersion, "meta version")
	aTest.MustBeEqual(FieldName, "name")
	aTest.MustBeEqual(FieldPath, "path")
	aTest.MustBeEqual(FieldPieceLayers, "piece layers")
	aTest.MustBeEqual(FieldPieceLength, "piece length")
	aTest.MustBeEqual(FieldPieceRoot, "piece root")
	aTest.MustBeEqual(FieldPieces, "pieces")
	aTest.MustBeEqual(FieldPrivate, "private")
	aTest.MustBeEqual(FieldSha1Sum, "sha1sum")
	aTest.MustBeEqual(FieldSha256Sum, "sha256sum")
}
