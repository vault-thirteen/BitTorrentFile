package models

type File struct {
	// Path elements of the file path.
	// The last element is the file's name.
	Path []string

	// File size in bytes.
	Size int

	// Hash sums.
	HashSum FileHash
}
