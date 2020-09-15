package gann

import (
	"sync"

	"github.com/google/uuid"
	"github.com/mathetake/gann/metric"
)

// Index is the interface of gann's search index. GetANNbyItemID and GetANNbyVector are different in the form of query.
// GetANNbyItemID can be executed by passing a certain item's id contained in the list of items used in the index building phase.
// GetANNbyVector allows us to pass any vector of proper dimension.
//
// searchNum is the number of requested approximated nearest neighbors, and bucketScale can be tuned to make balance between
// the search result's accuracy and computational complexity in the search phase.
//
// see README.md for more details.
type Index interface {
	// GetANNbyItemID ... search approximate nearest neighbors by a given itemID
	GetANNbyItemID(id int64, searchNum int, bucketScale float64) (ann []int64, err error)

	// GetANNbyVector ... search approximate nearest neighbors by a given query vector
	GetANNbyVector(v []float64, searchNum int, bucketScale float64) (ann []int64, err error)
}

type index struct {
	metric metric.Metric

	// dim ... dimension of the target space
	dim int

	// k ... maximum # of items in a single leaf node
	k int

	// itemIDToItem ... ItemIDToItem
	itemIDToItem map[itemId]*item

	// nodeIDToNode ... NodeIDToNode
	nodeIDToNode map[nodeId]*node

	// roots ... roots of the trees
	roots []*node

	mux *sync.Mutex
}

// CreateNewIndex build a new search index for given vectors. rawItems should consist of search target vectors and
// its slice index corresponds to the first argument id of GetANNbyItemID. For example, if we want to search approximate
// nearest neighbors of rawItems[3], it can simply achieved by calling index.GetANNbyItemID(3, ...).
//
// dim is the dimension of target spaces. nTree and k are tunable parameters which affects performances of
// the index (see README.md for details.)
//
// The last argument m is type of metric.Metric and represents the metric of the target search space.
// See https://godoc.org/github.com/mathetake/gann/metric for details.
func CreateNewIndex(rawItems [][]float64, dim, nTree, k int, m metric.Metric) (Index, error) {
	// verify that given items have same dimension
	for _, it := range rawItems {
		if len(it) != dim {
			return nil, errDimensionMismatch
		}
	}

	if len(rawItems) < 2 {
		return nil, errNotEnoughItems
	}

	its := make([]*item, len(rawItems))
	idToItem := make(map[itemId]*item, len(rawItems))
	for i, v := range rawItems {
		it := &item{
			id:     itemId(i),
			vector: v,
		}
		its[i] = it
		idToItem[it.id] = it
	}

	idx := &index{
		metric:       m,
		dim:          dim,
		k:            k,
		itemIDToItem: idToItem,
		roots:        make([]*node, nTree),
		nodeIDToNode: map[nodeId]*node{},
		mux:          &sync.Mutex{},
	}

	// build
	idx.build(its, nTree)
	return idx, nil
}

func (idx *index) build(items []*item, nTree int) {
	vs := make([][]float64, len(idx.itemIDToItem))
	for i, it := range items {
		vs[i] = it.vector
	}

	for i := 0; i < nTree; i++ {
		nv := idx.metric.GetSplittingVector(vs)
		rn := &node{
			id:       nodeId(uuid.New().String()),
			vec:      nv,
			idxPtr:   idx,
			children: map[direction]*node{},
		}
		idx.roots[i] = rn
		idx.nodeIDToNode[rn.id] = rn
	}

	var wg sync.WaitGroup
	wg.Add(nTree)
	for _, rn := range idx.roots {
		rn := rn
		go func() {
			defer wg.Done()
			rn.build(items)
		}()
	}
	wg.Wait()
}
