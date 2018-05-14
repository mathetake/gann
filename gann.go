package gann

import (
	"github.com/mathetake/gann/index"
	"github.com/mathetake/gann/item"
)

type GannIndex interface {
	Build() error // build search trees.
	GetANNbyItem(id item.ID, num int, searchBucket int) (ann []int32, err error)
	GetANNbyVector(v []float32, num int, searchBucket int) (ann []int32, err error)
}

// GetIndex ... get index (composed of trees, nodes, etc.)
func GetIndex(items [][]float32, d int, nT int, k int) (GannIndex, error) {
	return index.Initialize(items, d, nT, k)
}
