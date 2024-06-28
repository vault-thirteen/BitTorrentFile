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
