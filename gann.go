package gann

import (
	"github.com/mathetake/gann/index"
)

type GannIndex interface {
	Build() error // build search trees.
	GetANNbyItem(id int64, num int, bucketScale float64) (ann []int64, err error)
	GetANNbyVector(v []float32, num int, bucketScale float64) (ann []int64, err error)
}

// GetIndex ... get index (composed of trees, nodes, etc.)
func GetIndex(items [][]float32, d int, nT int, k int) (GannIndex, error) {
	return index.Initialize(items, d, nT, k)
}
