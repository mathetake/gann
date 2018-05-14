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
	itemIDToItem map[item.ID]item.Item

	// nodes
	nodes        []*node.Node
	nodeIDToNode map[node.ID]*node.Node

	// roots of trees
	roots []*node.Node
}

func (idx *Index) getNItems() int {
	return len(idx.items)
}

// Initialize ... initialize Index struct.
func Initialize(rawItems [][]float32, d int, nTree int, k int) (*Index, error) {
	if k >= len(rawItems) {
		panic("k must be smaller than len(rawItems).")
	}

	its := make([]item.Item, len(rawItems))
	idToItem := make(map[item.ID]item.Item, len(rawItems))
	for i, v := range rawItems {
		it := item.Item{
			ID:  item.ID(i),
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
	}, nil
}
