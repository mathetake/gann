package gann

import (
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/mathetake/gann/metric"
)

type Index interface {
	// GetANNbyItemID ... search ANNs by a given itemID
	GetANNbyItemID(id int64, num int, bucketScale float64) (ann []int64, err error)

	// GetANNbyVector ... search ANNs by a given query vector
	GetANNbyVector(v []float64, num int, bucketScale float64) (ann []int64, err error)
}

var _ Index = &index{}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Index ... a core struct in gann
type index struct {
	metric metric.Metric

	// dim ... dimension of the target space
	dim int

	// k ... minimum of descendants which every node contains.
	k int

	// itemIDToItem ... ItemIDToItem
	itemIDToItem map[itemId]*item

	// nodeIDToNode ... NodeIDToNode
	nodeIDToNode map[nodeId]*node

	// roots ... roots of the trees
	roots []*node

	mux *sync.Mutex
}

func CreateNewIndex(rawItems [][]float64, dim, nTree, k int, m metric.Metric) (Index, error) {
	// verify that given items have same dimension
	for _, it := range rawItems {
		if len(it) != dim {
			return nil, ErrDimensionMismatch
		}
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
