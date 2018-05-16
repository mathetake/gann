package node

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/mathetake/gann/item"
	"github.com/pkg/errors"
)

const (
	left  = "left"
	right = "right"
)

// Node ... node for tree
type Node struct {
	ID string

	// the normal vector of the hyper plane which splits the space, represented by the node
	Vec item.Vector

	// # of descendants items
	NDescendants int

	// children of node. If len equals 0, then it is leaf node.
	Children []*Node

	// In our setting, a `leaf` is a kind of node with len(Leaf field) greater than zero
	Leaf []int64

	Forest *[]*Node
}

func (n *Node) IsLeaf() bool {
	return len(n.Leaf) > 0
}

func (n *Node) Build(its []item.Item, k int, d int) error {
	if n.NDescendants < k {
		ids := make([]int64, len(its))
		for i, it := range its {
			ids[i] = it.ID
		}
		n.Leaf = append(n.Leaf, ids...)
		return nil
	}
	err := n.buildChildren(its, k, d)
	if err != nil {
		return errors.Wrap(err, "buildChild failed.")
	}
	return nil
}

// build child nodes
func (n *Node) buildChildren(its []item.Item, k int, d int) error {
	rand.Seed(time.Now().UnixNano())

	cMap := map[string]*Node{
		left: {
			Vec:  make([]float32, d),
			Leaf: []int64{},
		},
		right: {
			Vec:  make([]float32, d),
			Leaf: []int64{},
		},
	}

	// split descendants
	dItems := map[string][]item.Item{}
	dVectors := map[string][]item.Vector{}
	for _, it := range its {
		ip := item.DotProduct(n.Vec, it.Vec)
		if ip > 0 {
			dItems[left] = append(dItems[left], it)
			dVectors[left] = append(dVectors[left], it.Vec)
		} else if ip < 0 {
			dItems[right] = append(dItems[right], it)
			dVectors[right] = append(dVectors[right], it.Vec)
		} else {
			// if ip == 0, we assign the item randomly. Just in case.
			if rand.Int()%2 == 0 {
				dItems[left] = append(dItems[left], it)
				dVectors[left] = append(dVectors[left], it.Vec)
			} else {
				dItems[right] = append(dItems[right], it)
				dVectors[right] = append(dVectors[right], it.Vec)
			}
		}
	}

	for _, s := range []string{left, right} {
		if len(dItems[s]) == 0 {
			vs := make([]item.Vector, len(its))
			for i, it := range its {
				vs[i] = it.Vec
			}
			n.Vec = item.GetNormalVectorOfSplittingHyperPlane(vs, d)
			return n.buildChildren(its, k, d)
		}
	}

	for _, s := range []string{left, right} {
		if len(dItems[s]) >= k {
			nv := item.GetNormalVectorOfSplittingHyperPlane(dVectors[s], d)
			copy(cMap[s].Vec, nv)
		}
		cMap[s].ID = uuid.New().String()
		cMap[s].NDescendants = len(dItems[s])
		cMap[s].Forest = n.Forest
		// build children nodes recursively
		err := cMap[s].Build(dItems[s], k, d)
		if err != nil {
			return errors.Wrap(err, "Build failed.")
		}

		// append children.
		n.Children = append(n.Children, cMap[s])
		*n.Forest = append(*n.Forest, cMap[s])
	}
	return nil
}
