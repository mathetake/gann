package index

import (
	"container/heap"
	"math"
	"sort"
	"sync"

	"github.com/google/uuid"
	"github.com/mathetake/gann/item"
	"github.com/mathetake/gann/node"
	"github.com/pkg/errors"
)

// GetANNbyItemID ... get ANNs by a item.Item
func (idx *Index) GetANNbyItemID(id int64, num int, searchBucket float32) (ann []int64, err error) {
	it, ok := idx.ItemIDToItem[id]
	if !ok {
		return ann, errors.Errorf("Item not found for %v", id)
	}
	return idx.getANNbyVector(it.Vec, num, searchBucket)
}

// GetANNbyVector ... get ANNs by a vector
func (idx *Index) GetANNbyVector(v []float32, num int, bucketScale float32) (ann []int64, err error) {
	return idx.getANNbyVector(v, num, bucketScale)
}

func (idx *Index) getANNbyVector(v []float32, num int, bucketScale float32) ([]int64, error) {
	/*
		1. insert root nodes into the priority queue
		2. search all trees until len(`ann`) is enough.
		3. calculate actual distances to each elements in ann from v.
		4. sort `ann` by distances.
		5. Return the top `num` ones.
	*/

	if len(idx.Roots) == 0 {
		return []int64{}, errors.Errorf("Please build Index before searching.")
	}

	bucketSize := int(float32(num) * bucketScale)
	annMap := make(map[int64]interface{}, bucketSize)

	pq := node.PriorityQueue{}

	// 1.
	for i, r := range idx.Roots {
		n := &node.QueueItem{
			Value:    r.ID,
			Index:    i,
			Priority: float32(math.Inf(1)),
		}
		pq = append(pq, n)
	}

	heap.Init(&pq)

	// 2.
	for {
		q := heap.Pop(&pq).(*node.QueueItem)
		d := q.Priority
		n, ok := idx.NodeIDToNode[q.Value]
		if !ok {
			panic("wrong item set in priority queue")
		}

		if n.IsLeaf() {
			for _, id := range n.Leaf {
				annMap[id] = true
			}
		} else {
			ip := item.DotProduct(n.Vec, v)
			heap.Push(&pq, &node.QueueItem{
				Value:    n.Children[0].ID,
				Priority: min(d, ip),
			})
			heap.Push(&pq, &node.QueueItem{
				Value:    n.Children[1].ID,
				Priority: min(d, -ip),
			})
		}

		if len(annMap) >= bucketSize || len(pq) == 0 {
			break
		}
	}

	// 3.
	idToDist := make(map[int64]float32, len(annMap))
	ann := make([]int64, 0, len(annMap))
	for id := range annMap {
		ann = append(ann, id)
		idToDist[id] = item.DotProduct(idx.ItemIDToItem[id].Vec, v)
	}

	// 4.
	sort.Slice(ann, func(i, j int) bool {
		return -idToDist[ann[i]] < -idToDist[ann[j]]
	})

	// 5.
	if len(ann) > num {
		ann = ann[:num]
	}
	return ann, nil
}

// Build ... build index forest.
func (idx *Index) Build() error {
	if idx.isLoadedIndex {
		return errors.Errorf("This index is loaded from disk. Rebuild on such a index is not allowed because `nodes` and `items` fields is nil.")
	}

	idx.initRootNodes()

	var wg sync.WaitGroup
	var m sync.Map
	for i := range idx.Roots {
		wg.Add(1)
		ii := i
		go func() {
			idx.Roots[ii].Build(idx.items, idx.K, idx.Dim, &m)
			wg.Done()
		}()
	}
	wg.Wait()

	m.Range(func(key, _ interface{}) bool {
		n := key.(*node.Node)
		idx.nodes = append(idx.nodes, n)
		return true
	})

	if len(idx.nodes) == 0 {
		panic("# of nodes is zero.")
	}

	// build nodeIDToNode map
	for _, n := range idx.nodes {
		idx.NodeIDToNode[n.ID] = n
	}
	return nil
}

func (idx *Index) initRootNodes() {
	vecs := make([]item.Vector, len(idx.ItemIDToItem))
	for i, it := range idx.items {
		vecs[i] = it.Vec
	}
	for i := 0; i < idx.NTree; i++ {
		nv := item.GetNormalVectorOfSplittingHyperPlane(vecs, idx.Dim)
		r := &node.Node{
			ID:           uuid.New().String(),
			Vec:          nv,
			NDescendants: len(idx.items),
		}
		idx.Roots = append(idx.Roots, r)
		idx.nodes = append(idx.nodes, r)
	}
}

// for float32
func min(v1, v2 float32) float32 {
	if v1 > v2 {
		return v2
	}
	return v1
}
