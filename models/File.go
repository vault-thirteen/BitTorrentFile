package models

import "github.com/vault-thirteen/BitTorrentFile/models/hash"

// Forbidden path elements.
const (
	PathElement_Dot1 = `.`
	PathElement_Dot2 = `..`
)

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

func (f *File) SanitiseFilePath() {
	buf := make([]string, 0, len(f.Path))

	// Remove special folders used in operating systems.
	for _, element := range f.Path {
		if (element == PathElement_Dot1) ||
			(element == PathElement_Dot2) {
			continue
		}

		buf = append(buf, element)

		f.Path = buf
	}
}
