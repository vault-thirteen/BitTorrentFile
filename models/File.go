package models

import "github.com/vault-thirteen/BitTorrentFile/models/hash"

// File is information about a single file mentioned in the BitTorrent file.
type File struct {
	// Path elements of the file path.
	// The last element is the file's name.
	Path []string

	// File size in bytes.
	Size int

	// Hash sums.
	HashSum hash.FileHash
}
