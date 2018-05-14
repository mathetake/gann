package node

import (
	"github.com/mathetake/gann/item"
)

// ID ... identifier for nodes
type ID int32

// Node ... node for tree
type Node struct {
 	id ID
	// the normal vector of the hyper plane which splits the space, represented by the node
	vec item.Vector
	// children of node. If len ==0 => it is leaf node.
	children []ID

	// descendants items
	descendants []item.Item

	// # of descendants
	nDescendants int
}

func (n *Node) isLeaf () bool {
	return len(n.children) == 0
}
