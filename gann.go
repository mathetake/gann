package gann

import (
	"github.com/mathetake/gann/item"
	"github.com/mathetake/gann/index"
)


type GannIndex interface {
	build() // build search trees.
	getANNbyItem (id item.ID, num int, searchBucket int) (ann []int32, err error)
	getANNbyVector (v []float32, num int, searchBucket int) (ann []int32, err error)
}

// GetIndex ... get index (composed of trees, nodes, etc.)
func GetIndex (items [][]float32, d int, nT int) (GannIndex, error) {
	return index.Initialize(items, d, nT)
}
