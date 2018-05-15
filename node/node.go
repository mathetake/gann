package node

import (
	"github.com/mathetake/gann/item"
	"github.com/pkg/errors"
)

const (
	left  = "left"
	right = "right"
)

// Node ... node for tree
type Node struct {
	ID int

	// the normal vector of the hyper plane which splits the space, represented by the node
	Vec item.Vector

	// # of descendants items
	NDescendants int

	// children of node. If len equals 0, then it is leaf node.
	Children []*Node

	// In our setting, a `leaf` is a kind of node with len(Leaf field) greater than zero
	Leaf []int64

	Forest []*Node
}

func (n *Node) IsLeaf() bool {
	return len(n.Leaf) > 0
}

func (n *Node) Build(its []item.Item, k int) error {
	if n.NDescendants < k {
		ids := make([]int64, len(its))
		for i, it := range its {
			ids[i] = it.ID
		}
		n.Leaf = append(n.Leaf, ids...)
		return nil
	}
	err := n.buildChildren(its, k)
	if err != nil {
		return errors.Wrap(err, "buildChild failed.")
	}
	return nil
}

// build child nodes
func (n *Node) buildChildren(its []item.Item, k int) error {
	var cMap map[string]*Node

	// split descendants
	dItems := map[string][]item.Item{}
	dVectors := map[string][]item.Vector{}
	for _, it := range its {
		if item.DotProduct(n.Vec, it.Vec) > 0 {
			dItems[left] = append(dItems[left], it)
			dVectors[left] = append(dVectors[left], it.Vec)
		} else {
			dItems[right] = append(dItems[right], it)
			dVectors[right] = append(dVectors[right], it.Vec)
		}
	}

	for i, s := range []string{left, right} {
		if len(dItems[s]) >= k {
			nv, err := item.GetNormalVectorOfSplittingHyperPlane(dVectors[s])
			if err != nil {
				return errors.Wrap(err, "GetNormalVectorOfSplittingHyperPlane failed.")
			}
			cMap[s].Vec = nv
		}
		cMap[s].ID = n.ID + i
		cMap[s].NDescendants = len(dItems[s])
		cMap[s].Forest = n.Forest
		// build children nodes recursively
		err := cMap[s].Build(dItems[s], k)
		if err != nil {
			return errors.Wrap(err, "Build failed.")
		}

		// append children.
		n.Children = append(n.Children, cMap[s])
		n.Forest = append(n.Forest, cMap[s])
	}
	return nil
}
