package index

import (
	"github.com/mathetake/gann/item"
	"github.com/mathetake/gann/node"
)

// Index ... a core struct in gann
// TODO: modify for search with other distances (like L2, Hamming etc.)
type Index struct {
	// Dim ... dimension of the target space
	Dim int `json:"dim"`

	// NTree ... # of trees
	NTree int `json:"n_tree"`

	// K ... minimum of descendants which every node contains.
	K int `json:"k"`

	// Items ... items
	Items []item.Item `json:"items"`

	// ItemIDToItem ... ItemIDToItem
	ItemIDToItem map[int64]item.Item `json:"item_id_to_item"`

	// Nodes ... nodes
	Nodes []*node.Node `json:"nodes"`

	// NodeIDToNode ... NodeIDToNode
	NodeIDToNode map[string]*node.Node `json:"node_id_to_node"`

	// Roots ... roots of the trees
	Roots []*node.Node `json:"roots"`
}

// Initialize ... initialize Index struct.
func Initialize(rawItems [][]float32, d int, nTree int, k int, normalize bool) *Index {
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
		Dim:          d,
		K:            k,
		NTree:        nTree,
		Items:        its,
		ItemIDToItem: idToItem,
		NodeIDToNode: map[string]*node.Node{},
		Roots:        []*node.Node{},
	}
}

// GetIndex ... get index (composed of trees, nodes, etc.)
func GetIndex(items [][]float32, d int, nT int, k int, normalize bool) *Index {
	return Initialize(items, d, nT, k, normalize)
}
