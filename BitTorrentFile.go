package btf

import (
	"errors"
	"time"

	"github.com/vault-thirteen/BitTorrentFile/models"
	e "github.com/vault-thirteen/BitTorrentFile/models/error"
	"github.com/vault-thirteen/BitTorrentFile/models/generic"
	"github.com/vault-thirteen/BitTorrentFile/models/hash"
	iface "github.com/vault-thirteen/BitTorrentFile/models/interface"
	b "github.com/vault-thirteen/bencode"
)

type BitTorrentFile struct {
	// Source is a file which is being parsed.
	Source *b.File

	// Version of the BitTorrent File.
	// It can be numeric or textual while there are some crazy things such as
	// the 'Hybrid' file format.
	Version string

	// BitTorrent Info Hash.
	// The original first version of info hash.
	BTIH hash.BtihData

	// New BitTorrent Info Hash.
	// Second version of info hash.
	BTIH2 hash.BtihData2

	// Decoded raw data.
	RawData *b.DecodedObject

	// List of files described in the BitTorrent File.
	Files []models.File

	// List of announce URLs of trackers described in the BitTorrent File.
	AnnounceUrlMain models.AnnounceAddress
	AnnounceUrlsAux [][]models.AnnounceAddress

	// Creation time of the BitTorrent File.
	CreationTime time.Time

	// Comment of the BitTorrent File.
	Comment string

	// Creator of the BitTorrent File.
	Creator string

	// String encoding format used to generate the 'pieces' field of the 'info'
	// dictionary of the BitTorrent File.
	Encoding string
}

// NewBitTorrentFile is a constructor of the BitTorrentFile object.
func NewBitTorrentFile(filePath string) (tf *BitTorrentFile) {
	tf = new(BitTorrentFile)
	tf.Source = b.NewFile(filePath)
	return tf
}

// Open opens an existing BitTorrent file and parses it.
func (tf *BitTorrentFile) Open() (err error) {
	tf.RawData, err = tf.Source.Parse(true)
	if err != nil {
		return err
	}

	err = tf.calculateBtih()
	if err != nil {
		return err
	}

	err = tf.readAnnounceUrls()
	if err != nil {
		return err
	}

	err = tf.readCreationTime()
	if err != nil {
		return err
	}

	err = tf.readComment()
	if err != nil {
		return err
	}

	err = tf.readCreator()
	if err != nil {
		return err
	}

	err = tf.readEncoding()
	if err != nil {
		return err
	}

	err = tf.readFiles()
	if err != nil {
		return err
	}

	return nil
}

// GetSection gets a section specified by its name from the object.
func (tf *BitTorrentFile) GetSection(sectionName string) (result any, err error) {
	if tf.RawData == nil {
		return nil, errors.New(e.ErrFileIsNotOpened)
	}

	// Get the dictionary.
	var dictionary []b.DictionaryItem
	var ok bool
	dictionary, ok = tf.RawData.RawObject.([]b.DictionaryItem)
	if !ok {
		return nil, errors.New(e.ErrTypeAssertion)
	}

	// Get the section from the decoded object.
	var dictItem b.DictionaryItem
	for _, dictItem = range dictionary {
		if string(dictItem.Key) == sectionName {
			return dictItem.Value, nil
		}
	}

	return nil, errors.New(e.ErrSectionDoesNotExist)
}

// GetInfoSection gets an 'info' section from the object.
func (tf *BitTorrentFile) GetInfoSection() (is models.Dictionary, err error) {
	var tmp any
	tmp, err = tf.GetSection(models.SectionInfo)
	if err != nil {
		return nil, err
	}

	var ok bool
	var tmp2 []b.DictionaryItem
	tmp2, ok = tmp.([]b.DictionaryItem)
	if !ok {
		return nil, errors.New(e.ErrTypeAssertion)
	}

	is = (models.Dictionary)(tmp2)
	return is, nil
}

// GetSectionValueAsInt reads section value and returns it as int.
func (tf *BitTorrentFile) GetSectionValueAsInt(sectionName string) (i int, err error) {
	var section any
	section, err = tf.GetSection(sectionName)
	if err != nil {
		return 0, err
	}

	return iface.InterfaceAsInt(section)
}

// GetSectionValueAsString reads section value and returns it as string.
func (tf *BitTorrentFile) GetSectionValueAsString(sectionName string) (sv string, err error) {
	var section any
	section, err = tf.GetSection(sectionName)
	if err != nil {
		return "", err
	}

	return iface.InterfaceAsString(section)
}

// GetSectionValueAsStringArray reads section value and returns it as string
// array.
func (tf *BitTorrentFile) GetSectionValueAsStringArray(sectionName string) (sa []string, err error) {
	var section any
	section, err = tf.GetSection(sectionName)
	if err != nil {
		return nil, err
	}

	return iface.InterfaceAsStringArray(section)
}

// GetSectionValueAsArrayOfStringArrays reads section value and returns it as
// array of string arrays.
func (tf *BitTorrentFile) GetSectionValueAsArrayOfStringArrays(sectionName string) (asa [][]string, err error) {
	var section any
	section, err = tf.GetSection(sectionName)
	if err != nil {
		return nil, err
	}

	return iface.InterfaceAsArrayOfStringArrays(section)
}

// calculateBtih calculates the BitTorrent Info Hash (BTIH) check sums.
func (tf *BitTorrentFile) calculateBtih() (err error) {

	// Get the 'info' section from the decoded object.
	var infoSection models.Dictionary
	infoSection, err = tf.GetInfoSection()
	if err != nil {
		return err
	}

	// Encode the 'info' section.
	var infoSectionBA []byte
	infoSectionBA, err = b.NewEncoder().EncodeAnInterface(([]b.DictionaryItem)(infoSection))
	if err != nil {
		return err
	}

	// Calculate the BTIH check sums.
	tf.BTIH.Bytes, tf.BTIH.Text = hash.CalculateSha1(infoSectionBA)
	tf.BTIH2.Bytes, tf.BTIH2.Text = hash.CalculateSha256(infoSectionBA)

	return nil
}

// readAnnounceUrls reads announce URLs.
func (tf *BitTorrentFile) readAnnounceUrls() (err error) {

	// 1. Get the 'announce' section from the decoded object.
	var buf1 string
	buf1, err = tf.GetSectionValueAsString(models.SectionAnnounce)
	if err != nil {
		return err
	}

	var mainAnnounceUrl *models.AnnounceAddress
	mainAnnounceUrl, err = models.NewAnnounceAddressFromString(buf1)
	if err != nil {
		return err
	}

	tf.AnnounceUrlMain = *mainAnnounceUrl

	// 2. Get the optional 'announce-list' section from the decoded object.
	var buf2 []string
	var buf3 [][]string
	buf3, err = tf.GetSectionValueAsArrayOfStringArrays(models.SectionAnnounceList)
	if err != nil {
		if err.Error() == e.ErrSectionDoesNotExist {
			return nil
		}
		return err
	}

	tf.AnnounceUrlsAux = make([][]models.AnnounceAddress, 0, len(buf3))

	var aa []models.AnnounceAddress
	for _, buf2 = range buf3 {
		aa, err = models.NewAnnounceAddressListFromStringArray(buf2)
		if err != nil {
			return err
		}

		aa = generic.RemoveDuplicatesFromList[models.AnnounceAddress](aa)

		tf.AnnounceUrlsAux = append(tf.AnnounceUrlsAux, aa)
	}

	return nil
}

// readCreationTime reads creation time.
func (tf *BitTorrentFile) readCreationTime() (err error) {
	var i int
	i, err = tf.GetSectionValueAsInt(models.SectionCreationDate)
	if err != nil {
		if err.Error() == e.ErrSectionDoesNotExist {
			return nil
		}
		return err
	}

	tf.CreationTime = time.Unix(int64(i), 0)

	return nil
}

// readComment reads comment.
func (tf *BitTorrentFile) readComment() (err error) {
	tf.Comment, err = tf.GetSectionValueAsString(models.SectionComment)
	if err != nil {
		if err.Error() == e.ErrSectionDoesNotExist {
			return nil
		}
		return err
	}

	return nil
}

// readCreator reads creator.
func (tf *BitTorrentFile) readCreator() (err error) {
	tf.Creator, err = tf.GetSectionValueAsString(models.SectionCreatedBy)
	if err != nil {
		if err.Error() == e.ErrSectionDoesNotExist {
			return nil
		}
		return err
	}

	return nil
}

// readEncoding reads encoding.
func (tf *BitTorrentFile) readEncoding() (err error) {
	tf.Encoding, err = tf.GetSectionValueAsString(models.SectionEncoding)
	if err != nil {
		if err.Error() == e.ErrSectionDoesNotExist {
			return nil
		}
		return err
	}

	return nil
}

// readFiles reads the list of files.
func (tf *BitTorrentFile) readFiles() (err error) {
	//TODO: V1
	//TODO: V2 (?)

	// Guess the format of 'info' section.

	var infoSection models.Dictionary
	infoSection, err = tf.GetInfoSection()
	if err != nil {
		return err
	}

	var infoSectionFormat models.InfoSectionFormat
	infoSectionFormat = infoSection.GuessFormat()

	switch infoSectionFormat {
	case models.InfoSectionFormat_SingleFile:
		tf.Files, err = tf.readSingleFile(infoSection)
		if err != nil {
			return err
		}

	case models.InfoSectionFormat_MultiFile:
		tf.Files, err = tf.readMultipleFiles(infoSection)
		if err != nil {
			return err
		}

	default:
		return errors.New(e.ErrInfoSectionFormatIsUnknown)
	}

	return nil
}

func (tf *BitTorrentFile) readSingleFile(infoSection models.Dictionary) (files []models.File, err error) {
	var f models.File

	// 1. Read the file size.
	f.Size, err = infoSection.ReadFileSize()
	if err != nil {
		return nil, err
	}

	// 2. Read optional check sums.
	_, f.HashSum.Crc32, err = infoSection.ReadOptionalFileCrc32()
	if err != nil {
		return nil, err
	}

	_, f.HashSum.Md5, err = infoSection.ReadOptionalFileMd5()
	if err != nil {
		return nil, err
	}

	_, f.HashSum.Sha1, err = infoSection.ReadOptionalFileSha1()
	if err != nil {
		return nil, err
	}

	_, f.HashSum.Sha256, err = infoSection.ReadOptionalFileSha256()
	if err != nil {
		return nil, err
	}

	// 3. Read the file name.
	f.Path, err = infoSection.ReadFilePath(models.InfoSectionFormat_SingleFile)
	if err != nil {
		return nil, err
	}

	// Save the result.
	files = []models.File{f}
	return files, err
}

func (tf *BitTorrentFile) readMultipleFiles(infoSection models.Dictionary) (files []models.File, err error) {
	var rootFolderName string
	rootFolderName, err = infoSection.GetFieldValueAsString(models.FieldName)
	if err != nil {
		return nil, err
	}

	var buf1 any
	buf1, err = infoSection.GetFieldValue(models.FieldFiles)
	if err != nil {
		return nil, err
	}

	var buf2 []any
	var ok bool
	buf2, ok = buf1.([]any)
	if !ok {
		return nil, errors.New(e.ErrTypeAssertion)
	}

	files = make([]models.File, 0, len(buf2))
	var buf3 []b.DictionaryItem
	var filesDictionary models.Dictionary
	var f models.File

	for _, x := range buf2 {
		buf3, ok = x.([]b.DictionaryItem)
		if !ok {
			return nil, errors.New(e.ErrTypeAssertion)
		}

		filesDictionary = buf3

		// 1. Read the file size.
		f.Size, err = filesDictionary.ReadFileSize()
		if err != nil {
			return nil, err
		}

		// 2. Read optional check sums.
		_, f.HashSum.Crc32, err = filesDictionary.ReadOptionalFileCrc32()
		if err != nil {
			return nil, err
		}

		_, f.HashSum.Md5, err = filesDictionary.ReadOptionalFileMd5()
		if err != nil {
			return nil, err
		}

		_, f.HashSum.Sha1, err = filesDictionary.ReadOptionalFileSha1()
		if err != nil {
			return nil, err
		}

		_, f.HashSum.Sha256, err = filesDictionary.ReadOptionalFileSha256()
		if err != nil {
			return nil, err
		}

		// 3. Read the file path.
		var filePathWithoutRootFolder []string
		filePathWithoutRootFolder, err = filesDictionary.ReadFilePath(models.InfoSectionFormat_MultiFile)
		if err != nil {
			return nil, err
		}

		f.Path = []string{rootFolderName}
		f.Path = append(f.Path, filePathWithoutRootFolder...)

		// Save the result.
		files = append(files, f)
	}

	return files, nil
}
