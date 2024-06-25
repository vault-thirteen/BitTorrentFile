package btf

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
	"time"

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
	BTIH BtihData

	// New BitTorrent Info Hash.
	// Second version of info hash.
	BTIH2 BtihData2

	// Decoded raw data.
	RawData *b.DecodedObject

	// List of files described in the BitTorrent File.
	Files []File

	// List of announce URLs of trackers described in the BitTorrent File.
	AnnounceAddresses []AnnounceAddress

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

	//TODO: Parse the object.

	return nil
}

// GetSection gets a section specified by its name from the object.
func (tf *BitTorrentFile) GetSection(sectionName string) (result any, err error) {
	if tf.RawData == nil {
		return nil, errors.New(ErrFileIsNotOpened)
	}

	// Get the dictionary.
	var dictionary []b.DictionaryItem
	var ok bool
	dictionary, ok = tf.RawData.RawObject.([]b.DictionaryItem)
	if !ok {
		return nil, errors.New(ErrTypeAssertion)
	}

	// Get the section from the decoded object.
	var dictItem b.DictionaryItem
	for _, dictItem = range dictionary {
		if string(dictItem.Key) == sectionName {
			return dictItem.Value, nil
		}
	}

	return nil, errors.New(ErrSectionDoesNotExist)
}

// GetInfoSection gets an 'info' section from the object.
func (tf *BitTorrentFile) GetInfoSection() (result any, err error) {
	return tf.GetSection(SectionInfo)
}

// calculateBtih calculates the BitTorrent Info Hash (BTIH) check sums.
func (tf *BitTorrentFile) calculateBtih() (err error) {

	// Get the 'info' section from the decoded object.
	var infoSection any
	infoSection, err = tf.GetInfoSection()
	if err != nil {
		return err
	}

	// Encode the 'info' section.
	var infoSectionBA []byte
	infoSectionBA, err = b.NewEncoder().EncodeAnInterface(infoSection)
	if err != nil {
		return err
	}

	// Calculate the BTIH check sums.
	tf.BTIH.Bytes, tf.BTIH.Text = CalculateSha1(infoSectionBA)
	tf.BTIH2.Bytes, tf.BTIH2.Text = CalculateSha256(infoSectionBA)

	return nil
}

// CalculateSha1 calculates the SHA-1 check sum and returns it as a hexadecimal
// text and byte array.
func CalculateSha1(data []byte) (resultAsBytes Sha1Sum, resultAsText string) {
	resultAsBytes = sha1.Sum(data)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText
}

// CalculateSha256 calculates the SHA-256 check sum and returns it as a
// hexadecimal text and byte array.
func CalculateSha256(data []byte) (resultAsBytes Sha256Sum, resultAsText string) {
	resultAsBytes = sha256.Sum256(data)
	resultAsText = strings.ToUpper(hex.EncodeToString(resultAsBytes[:]))
	return resultAsBytes, resultAsText
}
