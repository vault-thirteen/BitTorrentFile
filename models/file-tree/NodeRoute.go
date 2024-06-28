package ft

// NodeRoute is a routed chain of file tree nodes used in BitTorrent file of
// Version 2. The route starts from the root node and ends with the file node.
type NodeRoute []*FileTreeNode

// NewNodeRoute creates a new route of nodes.
func NewNodeRoute(rootNode *FileTreeNode) (nr NodeRoute) {
	return []*FileTreeNode{rootNode}
}

// ConvertToPath composes an array of path elements from a route.
func (nr *NodeRoute) ConvertToPath() (path []string) {
	path = make([]string, 0, len(*nr))

	for _, node := range *nr {
		// We store BitTorrent name in the root node, but we must ignore it for
		// the file path. Here is the explaination why. The BitTorrent Protocol
		// Specification v2 states that 'name' field stores "a display name for
		// the torrent". This means that BitTorrent name is not part of the
		// file path while it may be any string. Though some BitTorrent clients
		// store folder name in this field, this is not the rule. The 'name'
		// field can be any string, and thus, we can not be sure that this
		// field contains the name of the root folder or anything else. We
		// simply ignore this field for the file path.
		//
		// The BitTorrent Protocol Specification v2 can be found at the
		// following address: http://bittorrent.org/beps/bep_0052.html.
		if node.IsRoot {
			continue
		}

		path = append(path, node.Name)
	}

	return path
}

// AddNode adds a node to the end of the route.
func (nr *NodeRoute) AddNode(node *FileTreeNode) {
	*nr = append(*nr, node)
}

// RemoveNode removes a node from the end of the route.
func (nr *NodeRoute) RemoveNode() {
	*nr = (*nr)[:len(*nr)-1]
}
