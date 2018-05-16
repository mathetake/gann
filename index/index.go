package index

import (
	"github.com/mathetake/gann/item"
	"github.com/mathetake/gann/node"
)

// TODO: modify for search with other distances (like L2, Hamming etc.)
type Index struct {
	// dimension
	dim int

	// # of trees
	nTree int

	// minimum num of descendants which any node contains.
	k int

	// items
	items        []item.Item
	itemIDToItem map[int64]item.Item

	// nodes
	nodes        []*node.Node
	nodeIDToNode map[string]*node.Node

	// roots of trees
	roots []*node.Node
}

func (idx *Index) getNItems() int {
	return len(idx.items)
}

// Initialize ... initialize Index struct.
func Initialize(rawItems [][]float32, d int, nTree int, k int, normalize bool) (*Index, error) {
	if k >= len(rawItems) {
		panic("k must be smaller than len(rawItems).")
	}

	if normalize {
		for i := 0; i < len(rawItems); i++ {
			item.Normalize(rawItems[i])
		}
	}

	its := make([]item.Item, len(rawItems))
	idToItem := make(map[int64]item.Item, len(rawItems))
	for i, v := range rawItems {
		it := item.Item{
			ID:  int64(i),
			Vec: v,
		}
		its[i] = it
		idToItem[it.ID] = it
	}
	return &Index{
		dim:          d,
		k:            k,
		nTree:        nTree,
		items:        its,
		itemIDToItem: idToItem,
		nodeIDToNode: map[string]*node.Node{},
		roots:        []*node.Node{},
	}, nil
}
