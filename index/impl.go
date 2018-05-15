package index

import (
	"container/heap"
	"github.com/google/uuid"
	"github.com/mathetake/gann/item"
	"github.com/mathetake/gann/node"
	"github.com/pkg/errors"
	"math"
	"sort"
)

// GetANNbyItem ... get ANNs by a item.Item
func (idx *Index) GetANNbyItemID(id int64, num int, searchBucket float64) (ann []int64, err error) {
	it, ok := idx.itemIDToItem[id]
	if !ok {
		return ann, errors.Errorf("Item not found for %v", id)
	}
	return idx.getANNbyVector(it.Vec, num, searchBucket)
}

// GetANNbyVector ... get ANNs by a vector
func (idx *Index) GetANNbyVector(v []float32, num int, bucketScale float64) (ann []int64, err error) {
	return idx.getANNbyVector(v, num, bucketScale)
}

func (idx *Index) getANNbyVector(v []float32, num int, bucketScale float64) ([]int64, error) {
	/*
		1. insert root nodes into the priority queue
		2. search all trees until len(`ann`) is enough.
		3. calculate actual distances to each elements in ann from v.
		4. sort `ann` by distances.
		5. Return the top `num` ones.
	*/

	if len(idx.roots) == 0 {
		return []int64{}, errors.Errorf("Please build Index before searching.")
	}

	annMap := make(map[int64]interface{}, int(float64(num)*bucketScale))

	pq := node.PriorityQueue{}

	// 1.
	for i, r := range idx.roots {
		n := &node.QueueItem{
			ID:       r.ID,
			Index:    i,
			Priority: float32(math.Inf(1)),
		}
		pq = append(pq, n)
	}

	heap.Init(&pq)

	// 2.
	i := idx.nTree + 1
	for {
		q := pq.Pop().(*node.QueueItem)
		d := q.Priority
		n, ok := idx.nodeIDToNode[q.ID]
		if !ok {
			panic("wrong item set in priority queue")
		}

		if n.NDescendants < idx.k || n.IsLeaf() {
			for _, id := range n.Leaf {
				annMap[id] = true
			}
		} else {
			ip := item.DotProduct(n.Vec, v)
			heap.Push(&pq, &node.QueueItem{
				ID:       n.Children[0].ID,
				Index:    i,
				Priority: min(d, ip),
			})
			i++
			heap.Push(&pq, &node.QueueItem{
				ID:       n.Children[1].ID,
				Index:    i,
				Priority: min(d, -ip),
			})
			i++
		}

		if len(annMap) >= num || len(pq) == 0 {
			break
		}
	}

	// 3.
	idToDist := make(map[int64]float32, len(annMap))
	ann := make([]int64, 0, len(annMap))
	for id := range annMap {
		ann = append(ann, id)
		idToDist[id] = item.DotProduct(idx.itemIDToItem[id].Vec, v)
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
	err := idx.buildRootNodes()
	if err != nil {
		return errors.Wrapf(err, "buildRootNodes failed.")
	}
	for _, rn := range idx.roots {
		err := rn.Build(idx.items, idx.k, idx.dim)
		if err != nil {
			return errors.Wrapf(err, "Build failed.")
		}
	}

	if len(idx.Nodes) == 0 {
		panic("# of nodes is zero.")
	}

	// build nodeIDToNode map
	for _, n := range idx.Nodes {
		idx.nodeIDToNode[n.ID] = n
	}
	return nil
}

func (idx *Index) buildRootNodes() error {
	vecs := make([]item.Vector, len(idx.itemIDToItem))
	for i, it := range idx.items {
		vecs[i] = it.Vec
	}
	for i := 0; i < idx.nTree; i++ {
		nv := item.GetNormalVectorOfSplittingHyperPlane(vecs, idx.dim)
		r := &node.Node{
			ID:           uuid.New().String(),
			Vec:          nv,
			NDescendants: len(idx.items),
			Forest:       &idx.Nodes,
		}
		idx.roots = append(idx.roots, r)
		idx.Nodes = append(idx.Nodes, r)
	}
	return nil
}

// for float32
func min(v1, v2 float32) float32 {
	if v1 > v2 {
		return v2
	}
	return v1
}
