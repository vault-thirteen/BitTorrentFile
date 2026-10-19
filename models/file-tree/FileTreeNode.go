package ft

import (
	"github.com/vault-thirteen/BitTorrentFile/models/hash"
)

// FileTreeNode is a file tree node used in BitTorrent file of Version 2.
type FileTreeNode struct {
	// Base parameters of files and folders.
	Name    string
	Size    int
	HashSum hash.FileHash

	// Flags.
	IsRoot      bool
	IsDirectory bool
	IsFile      bool

	// Links to other nodes.
	Parent   *FileTreeNode
	Children []*FileTreeNode
}

// AppendChild adds a child node to the parent node.
func (ftn *FileTreeNode) AppendChild(childNode *FileTreeNode) {
	ftn.Children = append(ftn.Children, childNode)
}
