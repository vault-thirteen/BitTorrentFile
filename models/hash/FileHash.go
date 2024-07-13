package hash

import "github.com/vault-thirteen/auxie/hash"

// FileHash is a collection of file's hash sums.
// All these sums are optional.
type FileHash struct {
	// CRC32 check sum of the file.
	// This field is optional.
	// This is an un-official extension.
	Crc32 *hash.Crc32Sum

	// MD5 check sum of the file.
	// This field is optional.
	// Source: Bittorrent Protocol Specification v1.0
	// https://wiki.theory.org/BitTorrentSpecification
	Md5 *hash.Md5Sum

	// SHA-1 check sum of the file.
	// This field is optional.
	// This is an un-official extension.
	Sha1 *hash.Sha1Sum

	// SHA-256 check sum of the file.
	// This field is optional.
	// This is an un-official extension.
	Sha256 *hash.Sha256Sum
}
