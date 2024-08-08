package models

import (
	"testing"

	e "github.com/vault-thirteen/BitTorrentFile/models/error"
	ft "github.com/vault-thirteen/BitTorrentFile/models/file-tree"
	"github.com/vault-thirteen/auxie/hash"
	"github.com/vault-thirteen/auxie/tester"
	b "github.com/vault-thirteen/bencode"
)

func Test_InterfaceAsDictionary(t *testing.T) {
	aTest := tester.New(t)
	var result Dictionary
	var err error

	// Test #1.
	result, err = InterfaceAsDictionary(0)
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, Dictionary(nil))

	// Test #2.
	result, err = InterfaceAsDictionary([]b.DictionaryItem{{KeyStr: "A"}})
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, Dictionary([]b.DictionaryItem{{KeyStr: "A"}}))
}

func Test_FindDictionaryItem(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var result *b.DictionaryItem
	var err error

	diA := b.DictionaryItem{Key: []byte("A")}
	diB := b.DictionaryItem{Key: []byte("B")}

	d, err = InterfaceAsDictionary([]b.DictionaryItem{diA, diB})
	aTest.MustBeNoError(err)

	// Test #1.
	result, err = d.FindDictionaryItem("C")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, (*b.DictionaryItem)(nil))

	// Test #2.
	result, err = d.FindDictionaryItem("B")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, &diB)
}

func Test_IsFieldPresent(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var err error

	diA := b.DictionaryItem{Key: []byte("A")}
	diB := b.DictionaryItem{Key: []byte("B")}

	d, err = InterfaceAsDictionary([]b.DictionaryItem{diA, diB})
	aTest.MustBeNoError(err)

	// Tests.
	aTest.MustBeEqual(d.IsFieldPresent("A"), true)
	aTest.MustBeEqual(d.IsFieldPresent("B"), true)
	aTest.MustBeEqual(d.IsFieldPresent("C"), false)
	aTest.MustBeEqual(d.IsFieldPresent("D"), false)
}

func Test_GetFieldValue(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var result any
	var err error

	diA := b.DictionaryItem{Key: []byte("A"), Value: 123}
	diB := b.DictionaryItem{Key: []byte("B")}

	d, err = InterfaceAsDictionary([]b.DictionaryItem{diA, diB})
	aTest.MustBeNoError(err)

	// Test #1. Key is not found.
	result, err = d.GetFieldValue("C")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(err.Error(), e.ErrFieldDoesNotExist)
	aTest.MustBeEqual(result, nil)

	// Test #2. Other error.
	// This is not implemented.

	// Test #3. Key is found.
	result, err = d.GetFieldValue("A")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, 123)
}

func Test_GetFieldValueAsInt(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var result int
	var err error

	diA := b.DictionaryItem{Key: []byte("A"), Value: int8(1)}
	diB := b.DictionaryItem{Key: []byte("B"), Value: "HAHA"}

	d, err = InterfaceAsDictionary([]b.DictionaryItem{diA, diB})
	aTest.MustBeNoError(err)

	// Test #1. Key is not found.
	result, err = d.GetFieldValueAsInt("C")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, 0)

	// Test #2. Key is found, but value is bad.
	result, err = d.GetFieldValueAsInt("B")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, 0)

	// Test #3. Key is found, and value is good.
	result, err = d.GetFieldValueAsInt("A")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, 1)
}

func Test_GetFieldValueAsString(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var result string
	var err error

	diA := b.DictionaryItem{Key: []byte("A"), Value: int8(1)}
	diB := b.DictionaryItem{Key: []byte("B"), Value: "Test"}

	d, err = InterfaceAsDictionary([]b.DictionaryItem{diA, diB})
	aTest.MustBeNoError(err)

	// Test #1. Key is not found.
	result, err = d.GetFieldValueAsString("C")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, "")

	// Test #2. Key is found, but value is bad.
	result, err = d.GetFieldValueAsString("A")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, "")

	// Test #3. Key is found, and value is good.
	result, err = d.GetFieldValueAsString("B")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, "Test")
}

func Test_GetFieldValueAsStringArray(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var result []string
	var err error

	diA := b.DictionaryItem{Key: []byte("A"), Value: int8(1)}
	diB := b.DictionaryItem{Key: []byte("B"), Value: []any{"Q", "W"}}

	d, err = InterfaceAsDictionary([]b.DictionaryItem{diA, diB})
	aTest.MustBeNoError(err)

	// Test #1. Key is not found.
	result, err = d.GetFieldValueAsStringArray("C")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, []string(nil))

	// Test #2. Key is found, but value is bad.
	result, err = d.GetFieldValueAsStringArray("A")
	aTest.MustBeAnError(err)
	aTest.MustBeEqual(result, []string(nil))

	// Test #3. Key is found, and value is good.
	result, err = d.GetFieldValueAsStringArray("B")
	aTest.MustBeNoError(err)
	aTest.MustBeEqual(result, []string{"Q", "W"})
}

func Test_GuessVersion(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var result Version
	var err error

	// Test #1.
	{
		d, err = InterfaceAsDictionary([]b.DictionaryItem{})
		aTest.MustBeNoError(err)

		result, err = d.GuessVersion()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, Version_One)
	}

	// Test #2.
	{
		diA := b.DictionaryItem{Key: []byte("meta version"), Value: 2}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diA})
		aTest.MustBeNoError(err)

		result, err = d.GuessVersion()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, Version_Two)
	}

	// Test #3.
	{
		diA := b.DictionaryItem{Key: []byte("meta version"), Value: 99}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diA})
		aTest.MustBeNoError(err)

		result, err = d.GuessVersion()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, Version_Unknown)
	}

	// Test #4.
	{
		diA := b.DictionaryItem{Key: []byte("meta version"), Value: "Boo"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diA})
		aTest.MustBeNoError(err)

		result, err = d.GuessVersion()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(result, Version_Unknown)
	}
}

func Test_GuessFormat(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var result InfoSectionFormat
	var err error

	// Test #1.
	{
		diLength := b.DictionaryItem{Key: []byte("length"), Value: 123}
		diFiles := b.DictionaryItem{Key: []byte("files"), Value: "abc"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diLength, diFiles})
		aTest.MustBeNoError(err)

		result = d.GuessFormat()
		aTest.MustBeEqual(result, InfoSectionFormat_Unknown)
	}

	// Test #2.
	{
		d, err = InterfaceAsDictionary([]b.DictionaryItem{})
		aTest.MustBeNoError(err)

		result = d.GuessFormat()
		aTest.MustBeEqual(result, InfoSectionFormat_Unknown)
	}

	// Test #3.
	{
		diLength := b.DictionaryItem{Key: []byte("length"), Value: 123}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diLength})
		aTest.MustBeNoError(err)

		result = d.GuessFormat()
		aTest.MustBeEqual(result, InfoSectionFormat_SingleFile)
	}

	// Test #4.
	{
		diLength := b.DictionaryItem{Key: []byte("files"), Value: "Q"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diLength})
		aTest.MustBeNoError(err)

		result = d.GuessFormat()
		aTest.MustBeEqual(result, InfoSectionFormat_MultiFile)
	}
}

func Test_ReadFileSize(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var result int
	var err error

	// Test #1.
	{
		diLength := b.DictionaryItem{Key: []byte("length"), Value: 123}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diLength})
		aTest.MustBeNoError(err)

		result, err = d.ReadFileSize()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, 123)
	}

	// Test #2.
	{
		diLength := b.DictionaryItem{Key: []byte("length"), Value: "Q"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diLength})
		aTest.MustBeNoError(err)

		result, err = d.ReadFileSize()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(result, 0)
	}

	// Test #3.
	{
		d, err = InterfaceAsDictionary([]b.DictionaryItem{})
		aTest.MustBeNoError(err)

		result, err = d.ReadFileSize()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(result, 0)
	}
}

func Test_ReadOptionalFileCrc32(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var isSumSet bool
	var result *hash.Crc32Sum
	var err error

	// Test #1. Sum filed is not present.
	{
		d, err = InterfaceAsDictionary([]b.DictionaryItem{})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileCrc32()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(isSumSet, false)
		aTest.MustBeEqual(result, (*hash.Crc32Sum)(nil))
	}

	// Test #2. Sum filed is not a hexadecimal number string.
	{
		diSum := b.DictionaryItem{Key: []byte("crc32sum"), Value: "boo"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diSum})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileCrc32()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(isSumSet, true)
		aTest.MustBeEqual(result, (*hash.Crc32Sum)(nil))
	}

	// Test #3. Length of the sum filed is wrong.
	{
		diSum := b.DictionaryItem{Key: []byte("crc32sum"), Value: "123456"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diSum})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileCrc32()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(isSumSet, true)
		aTest.MustBeEqual(result, (*hash.Crc32Sum)(nil))
	}

	// Test #4. Normal data.
	{
		diSum := b.DictionaryItem{Key: []byte("crc32sum"), Value: "FF00BB22"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diSum})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileCrc32()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(isSumSet, true)
		aTest.MustBeEqual(*result, hash.Crc32Sum([4]byte{0xFF, 0x00, 0xBB, 0x22}))
	}
}

func Test_ReadOptionalFileMd5(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var isSumSet bool
	var result *hash.Md5Sum
	var err error

	// Test #1. Sum filed is not present.
	{
		d, err = InterfaceAsDictionary([]b.DictionaryItem{})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileMd5()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(isSumSet, false)
		aTest.MustBeEqual(result, (*hash.Md5Sum)(nil))
	}

	// Test #2. Sum filed is not a hexadecimal number string.
	{
		diSum := b.DictionaryItem{Key: []byte("md5sum"), Value: "boo"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diSum})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileMd5()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(isSumSet, true)
		aTest.MustBeEqual(result, (*hash.Md5Sum)(nil))
	}

	// Test #3. Length of the sum filed is wrong.
	{
		diSum := b.DictionaryItem{Key: []byte("md5sum"), Value: "123456"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diSum})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileMd5()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(isSumSet, true)
		aTest.MustBeEqual(result, (*hash.Md5Sum)(nil))
	}

	// Test #4. Normal data.
	{
		diSum := b.DictionaryItem{Key: []byte("md5sum"), Value: "0102030405060708090A0B0C0D0E0F10"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diSum})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileMd5()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(isSumSet, true)
		aTest.MustBeEqual(*result, hash.Md5Sum([16]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10}))
	}
}

func Test_ReadOptionalFileSha1(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var isSumSet bool
	var result *hash.Sha1Sum
	var err error

	// Test #1. Sum filed is not present.
	{
		d, err = InterfaceAsDictionary([]b.DictionaryItem{})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileSha1()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(isSumSet, false)
		aTest.MustBeEqual(result, (*hash.Sha1Sum)(nil))
	}

	// Test #2. Sum filed is not a hexadecimal number string.
	{
		diSum := b.DictionaryItem{Key: []byte("sha1sum"), Value: "boo"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diSum})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileSha1()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(isSumSet, true)
		aTest.MustBeEqual(result, (*hash.Sha1Sum)(nil))
	}

	// Test #3. Length of the sum filed is wrong.
	{
		diSum := b.DictionaryItem{Key: []byte("sha1sum"), Value: "123456"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diSum})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileSha1()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(isSumSet, true)
		aTest.MustBeEqual(result, (*hash.Sha1Sum)(nil))
	}

	// Test #4. Normal data.
	{
		diSum := b.DictionaryItem{Key: []byte("sha1sum"), Value: "0102030405060708090A0B0C0D0E0F1011121314"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diSum})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileSha1()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(isSumSet, true)
		aTest.MustBeEqual(*result, hash.Sha1Sum([20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x11, 0x12, 0x13, 0x14}))
	}
}

func Test_ReadOptionalFileSha256(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var isSumSet bool
	var result *hash.Sha256Sum
	var err error

	// Test #1. Sum filed is not present.
	{
		d, err = InterfaceAsDictionary([]b.DictionaryItem{})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileSha256()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(isSumSet, false)
		aTest.MustBeEqual(result, (*hash.Sha256Sum)(nil))
	}

	// Test #2. Sum filed is not a hexadecimal number string.
	{
		diSum := b.DictionaryItem{Key: []byte("sha256sum"), Value: "boo"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diSum})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileSha256()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(isSumSet, true)
		aTest.MustBeEqual(result, (*hash.Sha256Sum)(nil))
	}

	// Test #3. Length of the sum filed is wrong.
	{
		diSum := b.DictionaryItem{Key: []byte("sha256sum"), Value: "123456"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diSum})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileSha256()
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(isSumSet, true)
		aTest.MustBeEqual(result, (*hash.Sha256Sum)(nil))
	}

	// Test #4. Normal data.
	{
		diSum := b.DictionaryItem{Key: []byte("sha256sum"), Value: "0102030405060708090A0B0C0D0E0F100102030405060708090A0B0C0D0E0F10"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diSum})
		aTest.MustBeNoError(err)

		isSumSet, result, err = d.ReadOptionalFileSha256()
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(isSumSet, true)
		aTest.MustBeEqual(*result, hash.Sha256Sum([32]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10}))
	}
}

func Test_ReadFilePath(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var result []string
	var err error

	// Test #1.1. Single file format. Field is set.
	{
		diName := b.DictionaryItem{Key: []byte("name"), Value: "TheFile.txt"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diName})
		aTest.MustBeNoError(err)

		result, err = d.ReadFilePath(InfoSectionFormat_SingleFile)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, []string{"TheFile.txt"})
	}

	// Test #1.2. Single file format. Field is not set.
	{
		diX := b.DictionaryItem{Key: []byte("x"), Value: "x"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diX})
		aTest.MustBeNoError(err)

		result, err = d.ReadFilePath(InfoSectionFormat_SingleFile)
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(result, []string(nil))
	}

	// Test #2.1. Multi file format. Field is set.
	{
		diPath := b.DictionaryItem{Key: []byte("path"), Value: []any{"File1.txt", "File2.txt"}}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diPath})
		aTest.MustBeNoError(err)

		result, err = d.ReadFilePath(InfoSectionFormat_MultiFile)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(result, []string{"File1.txt", "File2.txt"})
	}

	// Test #2.2. Multi file format. Field is not set.
	{
		diX := b.DictionaryItem{Key: []byte("x"), Value: "x"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diX})
		aTest.MustBeNoError(err)

		result, err = d.ReadFilePath(InfoSectionFormat_MultiFile)
		aTest.MustBeAnError(err)
		aTest.MustBeEqual(result, []string(nil))
	}
}

func Test_IsFileParametersNodeV2(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var err error

	// Test #1. No entries.
	{
		d, err = InterfaceAsDictionary([]b.DictionaryItem{})
		aTest.MustBeNoError(err)

		aTest.MustBeEqual(d.IsFileParametersNodeV2(), false)
	}

	// Test #2. Two entries.
	{
		diX := b.DictionaryItem{Key: []byte("x"), Value: "x"}
		diY := b.DictionaryItem{Key: []byte("y"), Value: "y"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diX, diY})
		aTest.MustBeNoError(err)

		aTest.MustBeEqual(d.IsFileParametersNodeV2(), false)
	}

	// Test #3. Single entry, but it is not empty.
	{
		diEntry := b.DictionaryItem{Key: []byte("non-empty"), Value: "non-empty"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diEntry})
		aTest.MustBeNoError(err)

		aTest.MustBeEqual(d.IsFileParametersNodeV2(), false)
	}

	// Test #4. Single entry, and it is empty.
	{
		diEntry := b.DictionaryItem{Key: []byte(""), Value: "BitTorrent is horse shit"}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diEntry})
		aTest.MustBeNoError(err)

		aTest.MustBeEqual(d.IsFileParametersNodeV2(), true)
	}
}

func Test_FillFileParameters(t *testing.T) {
	aTest := tester.New(t)
	var d Dictionary
	var fileDic []b.DictionaryItem
	var err error
	var outNode *ft.FileTreeNode

	// Test.
	{
		fileDic = []b.DictionaryItem{
			{Key: []byte("length"), Value: 123},
			{Key: []byte("crc32sum"), Value: "FF00BB22"},
			{Key: []byte("md5sum"), Value: "0102030405060708090A0B0C0D0E0F10"},
			{Key: []byte("sha1sum"), Value: "0102030405060708090A0B0C0D0E0F1011121314"},
			{Key: []byte("sha256sum"), Value: "0102030405060708090A0B0C0D0E0F100102030405060708090A0B0C0D0E0F10"},
		}

		diParamsEntry := b.DictionaryItem{Key: []byte(""), Value: fileDic}

		d, err = InterfaceAsDictionary([]b.DictionaryItem{diParamsEntry})
		aTest.MustBeNoError(err)

		outNode = new(ft.FileTreeNode)
		err = d.FillFileParameters(outNode)
		aTest.MustBeNoError(err)
		aTest.MustBeEqual(outNode.IsFile, true)
		aTest.MustBeEqual(outNode.Size, 123)
		aTest.MustBeEqual(*outNode.HashSum.Crc32, hash.Crc32Sum([4]byte{0xFF, 0x00, 0xBB, 0x22}))
		aTest.MustBeEqual(*outNode.HashSum.Md5, hash.Md5Sum([16]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10}))
		aTest.MustBeEqual(*outNode.HashSum.Sha1, hash.Sha1Sum([20]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x11, 0x12, 0x13, 0x14}))
		aTest.MustBeEqual(*outNode.HashSum.Sha256, hash.Sha256Sum([32]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10}))
	}
}
