package gann

import (
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

	// In our setting, a `leaf` is a kind of node with len(leaf) > 0
	leaf []itemId
}

func (n *node) build(its []*item) {
	if len(its) <= n.idxPtr.k {
		n.leaf = make([]itemId, len(its))
		for i, it := range its {
			n.leaf[i] = it.id
		}
		return
	}
	n.buildChildren(its)
}

func (n *node) buildChildren(its []*item) {
	dItems := map[direction][]*item{}
	dVectors := map[direction][][]float64{}
	for _, it := range its {
		if n.idxPtr.metrics.CalcDirectionPriority(n.vec, it.vector) < 0 {
			dItems[left] = append(dItems[left], it)
			dVectors[left] = append(dVectors[left], it.vector)
		} else {
			dItems[right] = append(dItems[right], it)
			dVectors[right] = append(dVectors[right], it.vector)
		}
	}

	var shouldMerge = false
	for _, s := range directions {
		if len(dItems[s]) <= n.idxPtr.k {
			shouldMerge = true
		}
	}

	if shouldMerge {
		n.leaf = make([]itemId, len(its))
		for i, it := range its {
			n.leaf[i] = it.id
		}
		return
	}

	for _, s := range directions {
		// build child
		c := &node{
			vec:      n.idxPtr.metrics.GetSplittingVector(dVectors[s]),
			id:       nodeId(uuid.New().String()),
			idxPtr:   n.idxPtr,
			children: make(map[direction]*node, len(directions)),
		}

		c.build(dItems[s])

		// append child for the search phase
		n.children[s] = c

		// append child to global map for the search phase
		n.idxPtr.mux.Lock()
		n.idxPtr.nodeIDToNode[c.id] = c
		n.idxPtr.mux.Unlock()
	}
	return
}
