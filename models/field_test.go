package models

import (
	"testing"

	"github.com/vault-thirteen/auxie/tester"
)

func Test_field(t *testing.T) {
	var aTest = tester.New(t)

	aTest.MustBeEqual(Field_Crc32Sum, "crc32sum")
	aTest.MustBeEqual(Field_FileTree, "file tree")
	aTest.MustBeEqual(Field_Files, "files")
	aTest.MustBeEqual(Field_Length, "length")
	aTest.MustBeEqual(Field_Md5Sum, "md5sum")
	aTest.MustBeEqual(Field_MetaVersion, "meta version")
	aTest.MustBeEqual(Field_Name, "name")
	aTest.MustBeEqual(Field_Path, "path")
	aTest.MustBeEqual(Field_PieceLayers, "piece layers")
	aTest.MustBeEqual(Field_PieceLength, "piece length")
	aTest.MustBeEqual(Field_PieceRoot, "piece root")
	aTest.MustBeEqual(Field_Pieces, "pieces")
	aTest.MustBeEqual(Field_Private, "private")
	aTest.MustBeEqual(Field_Sha1Sum, "sha1sum")
	aTest.MustBeEqual(Field_Sha256Sum, "sha256sum")
}
