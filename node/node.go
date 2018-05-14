package node

import (
	"github.com/mathetake/gann/item"
	"github.com/pkg/errors"
)

const (
	left  = "left"
	right = "right"
)

// ID ... identifier for nodes
type ID int32

// Node ... node for tree
type Node struct {
	ID ID

	// the normal vector of the hyper plane which splits the space, represented by the node
	Vec item.Vector

	// # of descendants items
	nDescendants int

	// children of node. If len equals 0, then it is leaf node.
	Children []*Node

	// In our setting, a `leaf` is a kind of node with len(Leaf field) greater than zero
	Leaf []item.ID
}

func (n *Node) IsLeaf() bool {
	return len(n.Leaf) > 0
}

func (n *Node) Build(its []item.Item, k int) error {
	// NOTE: len(its) must equal nDescendants
	if len(its) < k {
		ids := make([]item.ID, len(its))
		for i, it := range its {
			ids[i] = it.ID
		}
		n.Leaf = append(n.Leaf, ids...)
		return nil
	}
	err := n.buildChild(its, k)
	if err != nil {
		return errors.Wrap(err, "buildChild failed.")
	}
	return nil
}

// build child nodes
func (n *Node) buildChild(its []item.Item, k int) error {
	var cMap map[string]*Node

	// split descendants
	ds := map[string][]item.Item{}
	for _, it := range its {
		if item.DotProduct(n.Vec, it.Vec) > 0 {
			ds[left] = append(ds[left], it)
		} else {
			ds[right] = append(ds[right], it)
		}
	}

	for i, s := range []string{left, right} {
		if len(ds[s]) >= k {
			nv1, err := item.GetNormalVectorOfSplittingHyperPlane(ds[s])
			if err != nil {
				return errors.Wrap(err, "GetNormalVectorOfSplittingHyperPlane failed.")
			}
			cMap[s].Vec = nv1
			cMap[s].ID = n.ID + ID(i)
			cMap[s].nDescendants = len(ds[s])

			// build children nodes recursively
			err = cMap[s].Build(ds[s], k)
			if err != nil {
				return errors.Wrap(err, "Build failed.")
			}
			// append children.
			n.Children = append(n.Children, cMap[s])
		} else {
			err := cMap[s].Build(ds[s], k)
			if err != nil {
				return errors.Wrap(err, "Build failed.")
			}
			n.Children = append(n.Children, cMap[s])
		}
	}
	return nil
}
