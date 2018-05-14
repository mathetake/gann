package index

import (
	"github.com/mathetake/gann/item"
	"github.com/mathetake/gann/node"
	"github.com/pkg/errors"
)

// GetANNbyItem ... get ANNs by a item.Item
func (idx *Index) GetANNbyItem(id item.ID, num int, searchBucket int) (ann []int32, err error) {
	it, ok := idx.itemIDToItem[id]
	if !ok {
		errors.Errorf("Item not found for %v", id)
	}
	return idx.getANNbyVector(it.Vec, num, searchBucket)
}

// GetANNbyVector ... get ANNs by a vector
func (idx *Index) GetANNbyVector(v []float32, num int, searchBucket int) (ann []int32, err error) {
	return idx.getANNbyVector(v, num, searchBucket)
}

func (idx *Index) getANNbyVector(v []float32, num int, searchBucket int) (ann []int32, err error) {
	return ann, nil
}

// Build ... build index forest.
func (idx *Index) Build() {
	/*
		1. build root nodes.
		2. build forests by execute `node.Build` method of root nodes.
	*/
}

func (idx *Index) buildRootNodes() {}

// build a forest with a given root node.Node
func (idx *Index) buildForest(root *node.Node) {}
