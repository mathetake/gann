package node

import (
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/mathetake/gann/item"
)

const (
	left  = "left"
	right = "right"
)

// Node ... node for tree
type Node struct {
	ID string `json:"id"`

	// the normal vector of the hyper plane which splits the space, represented by the node
	Vec item.Vector `json:"vec"`

	// # of descendants items
	NDescendants int `json:"n_descendants"`

	// children of node. If len equals 0, then it is leaf node.
	Children []*Node `json:"children"`

	// In our setting, a `leaf` is a kind of node with len(Leaf field) greater than zero
	Leaf []int64 `json:"leaf"`
}

// IsLeaf ... check if it is a leaf node
func (n *Node) IsLeaf() bool {
	return len(n.Leaf) > 0
}

// Build ... build
func (n *Node) Build(its []item.Item, k int, d int, m *sync.Map) {
	if n.NDescendants < k {
		ids := make([]int64, len(its))
		for i, it := range its {
			ids[i] = it.ID
		}
		n.Leaf = append(n.Leaf, ids...)
		return
	}
	n.buildChildren(its, k, d, m)
}

// build child nodes
func (n *Node) buildChildren(its []item.Item, k int, d int, m *sync.Map) {
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
			ids := make([]int64, len(its))
			for i, it := range its {
				ids[i] = it.ID
			}
			n.Leaf = append(n.Leaf, ids...)
			return
		}
	}

	for _, s := range []string{left, right} {
		if len(dItems[s]) >= k {
			cMap[s].Vec = item.GetNormalVectorOfSplittingHyperPlane(dVectors[s], d)
		}
		cMap[s].ID = uuid.New().String()
		cMap[s].NDescendants = len(dItems[s])
		// build children nodes recursively
		cMap[s].Build(dItems[s], k, d, m)

		// append children.
		n.Children = append(n.Children, cMap[s])
		m.Store(cMap[s], true)
	}
	return
}
