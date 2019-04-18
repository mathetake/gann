package gann

import (
	"sync"

	"github.com/google/uuid"
)

type nodeId string
type direction string

const (
	left  direction = "left"
	right direction = "right"
)

var directions = []direction{left, right}

type node struct {
	idxPtr *index

	id nodeId

	// the normal vector of the hyper plane which splits the space, represented by the node
	vec []float64

	// children of node. If len equals 0, then it is leaf node.
	children map[direction]*node

	// In our setting, a `leaf` is a kind of node with len(Leaf field) greater than zero
	leaf []itemId
}

func (n *node) build(its []*item) {
	if len(its) < n.idxPtr.k {
		ids := make([]itemId, len(its))
		for i, it := range its {
			ids[i] = it.id
		}
		n.leaf = append(n.leaf, ids...)
		return
	}
	n.buildChildren(its)
}

// build child nodes
func (n *node) buildChildren(its []*item) {
	// split descendants
	dItems := map[direction][]*item{}
	dVectors := map[direction][][]float64{}
	for _, it := range its {
		if n.idxPtr.metrics.GetDirectionPriority(n.vec, it.vector) > 0 {
			dItems[left] = append(dItems[left], it)
			dVectors[left] = append(dVectors[left], it.vector)
		} else {
			dItems[right] = append(dItems[right], it)
			dVectors[right] = append(dVectors[right], it.vector)
		}
	}

	for _, s := range directions {
		if len(dItems[s]) == 0 {
			ids := make([]itemId, len(its))
			for i, it := range its {
				ids[i] = it.id
			}
			n.leaf = append(n.leaf, ids...)
			return
		}
	}

	var wg sync.WaitGroup
	wg.Add(len(directions))
	for _, s := range directions {
		s := s
		go func() {
			defer wg.Done()
			n := &node{
				vec:    n.idxPtr.metrics.GetNormalVectorOfSplittingHyperPlane(dVectors[s]),
				id:     nodeId(uuid.New().String()),
				idxPtr: n.idxPtr,
			}

			n.build(dItems[s])

			// append child for the search phase
			n.children[s] = n

			// append child to global map for the search phase
			n.idxPtr.mux.Lock()
			n.idxPtr.nodeIDToNode[n.id] = n
			n.idxPtr.mux.Unlock()
		}()
	}
	wg.Wait()
	return
}
