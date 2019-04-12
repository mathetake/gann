package v2

import (
	"github.com/mathetake/gann/v2/metrics"
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

type node struct {
	id string

	// the normal vector of the hyper plane which splits the space, represented by the node
	vec []float64

	// # of descendants items
	NDescendants int

	// children of node. If len equals 0, then it is leaf node.
	Children []*node

	// In our setting, a `leaf` is a kind of node with len(Leaf field) greater than zero
	Leaf []itemId
}

// IsLeaf ... check if it is a leaf node
func (n *node) IsLeaf() bool {
	return len(n.Leaf) > 0
}

// Build ... build
func (n *node) Build(its []item, k int, d int, m *sync.Map) {
	if n.NDescendants < k {
		ids := make([]itemId, len(its))
		for i, it := range its {
			ids[i] = it.id
		}
		n.Leaf = append(n.Leaf, ids...)
		return
	}
	n.buildChildren(its, k, d, m)
}

// build child nodes
func (n *node) buildChildren(its []item, k int, d int, sm *sync.Map, m metrics.Metrics) {
	rand.Seed(time.Now().UnixNano())

	cMap := map[string]*node{
		left: {
			vec:  make([]float64, d),
			Leaf: []itemId{},
		},
		right: {
			vec:  make([]float64, d),
			Leaf: []itemId{},
		},
	}

	// split descendants
	dItems := map[string][]item{}
	dVectors := map[string][]item{}
	for _, it := range its {
		ip := item.DotProduct(n.vec, it.Vec)
		if {
			dItems[left] = append(dItems[left], it)
			dVectors[left] = append(dVectors[left], it.Vec)
		} else if ip <= 0 {
			dItems[right] = append(dItems[right], it)
			dVectors[right] = append(dVectors[right], it.Vec)
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
			cMap[s].vec = m.GetNormalVectorOfSplittingHyperPlane(dVectors[s])
		}
		cMap[s].id = uuid.New().String()
		cMap[s].NDescendants = len(dItems[s])
		// build children nodes recursively
		cMap[s].Build(dItems[s], k, d, )

		// append children.
		n.Children = append(n.Children, cMap[s])
		sm.Store(cMap[s], true)
	}
	return
}
