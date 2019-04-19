package gann

import (
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/mathetake/gann/metrics"
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
	metrics metrics.Metrics

	// Dim ... dimension of the target space
	dim int

	// K ... minimum of descendants which every node contains.
	k int

	// ItemIDToItem ... ItemIDToItem
	itemIDToItem map[itemId]*item
	items        []*item // items ... only used in building steps.

	// NodeIDToNode ... NodeIDToNode
	nodeIDToNode map[nodeId]*node
	nodes        []*node // only used in building steps.

	// Roots ... roots of the trees
	roots []*node

	mux *sync.Mutex
}

func CreateNewIndex(rawItems [][]float64, dim int, nTree int, k int, mt metrics.Type) (Index, error) {
	// verify that given items have same dimension
	for _, it := range rawItems {
		if len(it) != dim {
			return nil, ErrDimensionMismatch
		}
	}

	m, err := metrics.NewMetrics(mt, dim)
	if err != nil {
		return nil, err
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
		metrics:      m,
		dim:          dim,
		k:            k,
		items:        its,
		itemIDToItem: idToItem,
		roots:        make([]*node, nTree),
		nodeIDToNode: map[nodeId]*node{},
		mux:          &sync.Mutex{},
	}

	// build
	idx.build(nTree)
	return idx, nil
}

func (idx *index) build(nTree int) {
	vs := make([][]float64, len(idx.itemIDToItem))
	for i, it := range idx.items {
		vs[i] = it.vector
	}

	for i := 0; i < nTree; i++ {
		nv := idx.metrics.GetSplittingVector(vs)
		rn := &node{
			id:       nodeId(uuid.New().String()),
			vec:      nv,
			idxPtr:   idx,
			children: map[direction]*node{},
		}
		idx.roots[i] = rn
		idx.nodes = append(idx.nodes, rn)
	}

	var wg sync.WaitGroup
	wg.Add(nTree)
	for _, rn := range idx.roots {
		rn := rn
		go func() {
			defer wg.Done()
			rn.build(idx.items)
		}()
	}
	wg.Wait()
}
