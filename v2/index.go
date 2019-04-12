package v2

import (
	"github.com/mathetake/gann/v2/metrics"
)

type GannIndex interface {
	// GetANNbyItemID ... search ANNs by a given itemID
	GetANNbyItemID(id int64, num int, bucketScale float64) (ann []int64, err error)

	// GetANNbyVector ... search ANNs by a given query vector
	GetANNbyVector(v []float64, num int, bucketScale float64) (ann []int64, err error)
}

var _ GannIndex = &index{}

func BuildIndex(items [][]float64, d int, nT int, k int, mt metrics.Type) (GannIndex, error) {
	// TODO:
	return nil, ErrDimensionMismatch
}

// Index ... a core struct in gann
type index struct {
	metrics metrics.Metrics

	// Dim ... dimension of the target space
	dim int

	// NTree ... # of trees
	nTree int

	// K ... minimum of descendants which every node contains.
	k int

	// ItemIDToItem ... ItemIDToItem
	itemIDToItem map[itemId]*item
	items        []*item // items ... only used in building steps.

	// NodeIDToNode ... NodeIDToNode
	nodeIDToNode map[string]*node
	nodes        []*node // only used in building steps.

	// Roots ... roots of the trees
	roots []*node
}

func (idx *index) GetANNbyItemID(id int64, num int, bucketScale float64) (ann []int64, err error) {
	return nil, nil
}

func (idx *index) GetANNbyVector(v []float64, num int, bucketScale float64) (ann []int64, err error) {
	return nil, nil
}
