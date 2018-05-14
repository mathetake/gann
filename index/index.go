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

	// items
	items []*item.Item
	itemIDToItem map[item.ID]*item.Item

	// nodes
	nodes []*node.Node
	nodeIDToNode map[node.ID]*node.Node

	// roots of trees
	roots []*node.Node
}


// Initialize ... initialize Index struct.
func Initialize(rawItems [][]float32, d int, nTree int) (*Index, error) {
	its := make([]*item.Item, len(rawItems))
	idToItem := make(map[item.ID]*item.Item)
	for i, v := range rawItems {
		it := &item.Item{
			ID: item.ID(i),
			Vec: v,
		}
		its[i] = it
		idToItem[it.ID] = it
	}
	return &Index{
		dim: d,
		nTree: nTree,
		items: its,
		itemIDToItem: idToItem,
	}, nil
}
