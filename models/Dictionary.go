package models

import (
	"encoding/hex"
	"errors"
	"reflect"

	e "github.com/vault-thirteen/BitTorrentFile/models/error"
	ft "github.com/vault-thirteen/BitTorrentFile/models/file-tree"
	iface "github.com/vault-thirteen/BitTorrentFile/models/interface"
	"github.com/vault-thirteen/auxie/hash"
	b "github.com/vault-thirteen/bencode"
)

// Dictionary is a "bencoded" dictionary.
type Dictionary []b.DictionaryItem

// InterfaceAsDictionary tries to interpret an interface as a dictionary.
func InterfaceAsDictionary(x any) (d Dictionary, err error) {
	// Direct type cast into 'Dictionary' type is not allowed in Go language.
	// So, we are doing it in two steps:
	//	1. Cast into an array of 'DictionaryItem's;
	//	2. Cast into the 'Dictionary' type.
	dictionaryItems, ok := x.([]b.DictionaryItem)
	if !ok {
		return nil, errors.New(e.ErrTypeAssertion)
	}

	return dictionaryItems, nil
}

// FindDictionaryItem tries to search for a dictionary entry (item) specified
// by its key (name). On success, the entry is returned.
func (d *Dictionary) FindDictionaryItem(name string) (item *b.DictionaryItem, err error) {
	for _, x := range *d {
		if string(x.Key) == name {
			return &x, nil
		}
	}

	return nil, errors.New(e.ErrItemIsNotFound)
}

// IsFieldPresent tries to search for a dictionary entry (item) specified
// by its key (name). On success, the presence flag is returned.
func (d *Dictionary) IsFieldPresent(fieldName string) (isFieldPresent bool) {
	_, err := d.FindDictionaryItem(fieldName)
	if err == nil {
		return true
	}

	return false
}

// GetFieldValue returns a dictionary's entry specified by its key (name).
// The difference between this method and 'FindDictionaryItem' is that the
// latter is a more "low-level" function and is used by other functions.
func (d *Dictionary) GetFieldValue(fieldName string) (fv any, err error) {
	var field *b.DictionaryItem
	field, err = d.FindDictionaryItem(fieldName)
	if err != nil {
		if err.Error() == e.ErrItemIsNotFound {
			return nil, errors.New(e.ErrFieldDoesNotExist)
		} else {
			return nil, err
		}
	}

	return field.Value, nil
}

// GetFieldValueAsInt returns a dictionary's entry specified by its key (name)
// as int.
func (d *Dictionary) GetFieldValueAsInt(fieldName string) (fv int, err error) {
	var tmp any
	tmp, err = d.GetFieldValue(fieldName)
	if err != nil {
		return 0, err
	}

	return iface.InterfaceAsInt(tmp)
}

// GetFieldValueAsString returns a dictionary's entry specified by its key (name)
// as string.
func (d *Dictionary) GetFieldValueAsString(fieldName string) (fv string, err error) {
	var tmp any
	tmp, err = d.GetFieldValue(fieldName)
	if err != nil {
		return "", err
	}

	return iface.InterfaceAsString(tmp)
}

// GetFieldValueAsStringArray returns a dictionary's entry specified by its key
// (name) as a string array.
func (d *Dictionary) GetFieldValueAsStringArray(fieldName string) (fv []string, err error) {
	var tmp any
	tmp, err = d.GetFieldValue(fieldName)
	if err != nil {
		return nil, err
	}

	return iface.InterfaceAsStringArray(tmp)
}

// GuessVersion tries to guess the version of the dictionary. This method is
// used only for the 'info' section dictionary.
func (d *Dictionary) GuessVersion() (version Version, err error) {
	var isVersionFieldPresent = d.IsFieldPresent(FieldMetaVersion)

	// We can only make assumptions about the version when it is not set.
	// The BitTorrent file format is crazy.
	if !isVersionFieldPresent {
		return Version_One, nil
	}

	var versionNumber int
	versionNumber, err = d.GetFieldValueAsInt(FieldMetaVersion)
	if err != nil {
		return Version_Unknown, err
	}

	switch versionNumber {
	case 2:
		// The BitTorrent Protocol Specification v2
		// http://bittorrent.org/beps/bep_0052.html
		return Version_Two, nil

	default:
		return Version_Unknown, nil
	}
}

// GuessFormat tries to guess the format of the dictionary. This method is
// used only for the 'info' section dictionary and when the BitTorrent file is
// of the original first version.
func (d *Dictionary) GuessFormat() (format InfoSectionFormat) {
	var isLengthFieldPresent = d.IsFieldPresent(FieldLength)
	var isFilesFieldPresent = d.IsFieldPresent(FieldFiles)

	// Exactly one field should be present.
	if isLengthFieldPresent && isFilesFieldPresent {
		return InfoSectionFormat_Unknown
	}
	if (!isLengthFieldPresent) && (!isFilesFieldPresent) {
		return InfoSectionFormat_Unknown
	}

	if isLengthFieldPresent {
		return InfoSectionFormat_SingleFile
	}

	return InfoSectionFormat_MultiFile
}

// ReadFileSize reads the 'length' field of the dictionary.
func (d *Dictionary) ReadFileSize() (fs int, err error) {
	fs, err = d.GetFieldValueAsInt(FieldLength)
	if err != nil {
		return 0, err
	}

	return fs, nil
}

// ReadOptionalFileCrc32 reads the optional field of the dictionary, CRC32
// check sum field.
func (d *Dictionary) ReadOptionalFileCrc32() (isCrc32Set bool, crc32 *hash.Crc32Sum, err error) {
	var buf1 string
	buf1, err = d.GetFieldValueAsString(FieldCrc32Sum)
	if err != nil {
		if err.Error() == e.ErrFieldDoesNotExist {
			return false, nil, nil
		} else {
			return false, nil, err
		}
	}

	var buf2 []byte
	buf2, err = hex.DecodeString(buf1)
	if err != nil {
		return true, nil, err
	}

	var buf3 hash.Crc32Sum
	if len(buf2) != int(reflect.TypeOf(buf3).Size()) {
		return true, nil, errors.New(e.ErrFieldValueSizeMismatch)
	}

	buf3 = hash.Crc32Sum(buf2)
	crc32 = &buf3
	return true, crc32, nil
}

// ReadOptionalFileMd5 reads the optional field of the dictionary, MD5
// check sum field.
func (d *Dictionary) ReadOptionalFileMd5() (isMd5Set bool, md5 *hash.Md5Sum, err error) {
	var buf1 string
	buf1, err = d.GetFieldValueAsString(FieldMd5Sum)
	if err != nil {
		if err.Error() == e.ErrFieldDoesNotExist {
			return false, nil, nil
		} else {
			return false, nil, err
		}
	}

	var buf2 []byte
	buf2, err = hex.DecodeString(buf1)
	if err != nil {
		return true, nil, err
	}

	var buf3 hash.Md5Sum
	if len(buf2) != int(reflect.TypeOf(buf3).Size()) {
		return true, nil, errors.New(e.ErrFieldValueSizeMismatch)
	}

	buf3 = hash.Md5Sum(buf2)
	md5 = &buf3
	return true, md5, nil
}

// ReadOptionalFileSha1 reads the optional field of the dictionary, SHA-1
// check sum field.
func (d *Dictionary) ReadOptionalFileSha1() (isSha1Set bool, sha1 *hash.Sha1Sum, err error) {
	var buf1 string
	buf1, err = d.GetFieldValueAsString(FieldSha1Sum)
	if err != nil {
		if err.Error() == e.ErrFieldDoesNotExist {
			return false, nil, nil
		} else {
			return false, nil, err
		}
	}

	var buf2 []byte
	buf2, err = hex.DecodeString(buf1)
	if err != nil {
		return true, nil, err
	}

	var buf3 hash.Sha1Sum
	if len(buf2) != int(reflect.TypeOf(buf3).Size()) {
		return true, nil, errors.New(e.ErrFieldValueSizeMismatch)
	}

	buf3 = hash.Sha1Sum(buf2)
	sha1 = &buf3
	return true, sha1, nil
}

// ReadOptionalFileSha256 reads the optional field of the dictionary, SHA-256
// check sum field.
func (d *Dictionary) ReadOptionalFileSha256() (isSha256Set bool, sha256 *hash.Sha256Sum, err error) {
	var buf1 string
	buf1, err = d.GetFieldValueAsString(FieldSha256Sum)
	if err != nil {
		if err.Error() == e.ErrFieldDoesNotExist {
			return false, nil, nil
		} else {
			return false, nil, err
		}
	}

	var buf2 []byte
	buf2, err = hex.DecodeString(buf1)
	if err != nil {
		return true, nil, err
	}

	var buf3 hash.Sha256Sum
	if len(buf2) != int(reflect.TypeOf(buf3).Size()) {
		return true, nil, errors.New(e.ErrFieldValueSizeMismatch)
	}

	buf3 = hash.Sha256Sum(buf2)
	sha256 = &buf3
	return true, sha256, nil
}

// ReadFilePath reads the file path.
// Beware that depending on the format of the 'info' section, this may be
// either the full path or the path without the root path. If you are looking
// for a full path, you should prepend this path with the root path when it is
// needed. Unfortunately, BitTorrent file format is crazy.
func (d *Dictionary) ReadFilePath(isf InfoSectionFormat) (filePath []string, err error) {
	switch isf {
	case InfoSectionFormat_SingleFile:
		{
			var fileName string
			fileName, err = d.GetFieldValueAsString(FieldName)
			if err != nil {
				return nil, err
			}

			return []string{fileName}, nil
		}

	case InfoSectionFormat_MultiFile:
		{
			var filePathWithoutRootFolder []string
			filePathWithoutRootFolder, err = d.GetFieldValueAsStringArray(FieldPath)
			if err != nil {
				return nil, err
			}

			return filePathWithoutRootFolder, nil
		}

	default:
		return nil, errors.New(e.ErrInfoSectionFormatIsUnknown)
	}
}

// IsFileParametersNodeV2 checks whether the dictionary is a special dictionary
// used for storing file parameters. This approach proves that BitTorrent file
// format is a nightmare.
func (d *Dictionary) IsFileParametersNodeV2() bool {
	dic := *d

	if len(dic) != 1 {
		return false
	}

	dicItemName := dic[0].Key
	if len(dicItemName) == 0 {
		return true
	}

	return false
}

// FillFileParameters fills (writes) file parameters into the specified file
// tree node pointer. The 'outNode' parameter must be a valid pointer. Current
// dictionary is a dictionary which has an item containing the special
// dictionary with file parameters.
func (d *Dictionary) FillFileParameters(outNode *ft.FileTreeNode) (err error) {
	var paramsDic Dictionary
	paramsDic, err = InterfaceAsDictionary((*d)[0].Value)
	if err != nil {
		return err
	}

	outNode.IsFile = true

	// 1. File size.
	outNode.Size, err = paramsDic.GetFieldValueAsInt(FieldLength)
	if err != nil {
		return err
	}

	// 2. Optional check sums.
	_, outNode.HashSum.Crc32, err = paramsDic.ReadOptionalFileCrc32()
	if err != nil {
		return err
	}

	_, outNode.HashSum.Md5, err = paramsDic.ReadOptionalFileMd5()
	if err != nil {
		return err
	}

	_, outNode.HashSum.Sha1, err = paramsDic.ReadOptionalFileSha1()
	if err != nil {
		return err
	}

	_, outNode.HashSum.Sha256, err = paramsDic.ReadOptionalFileSha256()
	if err != nil {
		return err
	}

	return nil
}
