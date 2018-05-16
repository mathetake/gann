package index

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/mathetake/gann/item"
)

func TestInitRootNodes(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	k := 6
	dim := 2
	num := 20
	nTree := 10
	var rawItems [][]float32

	for i := 0; i < num; i++ {
		v := make([]float32, 2)
		for d := 0; d < dim; d++ {
			v[d] = rand.Float32()
		}
		rawItems = append(rawItems, v)
	}

	idx, err := Initialize(rawItems, dim, nTree, k, true)
	if err != nil {
		panic(err)
	}

	idx.initRootNodes()
	assert.Equal(t, nTree, len(idx.roots))
	assert.Equal(t, nTree, len(idx.nodes))

}

func TestBuild(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for _, k := range []int{2, 10, 100} {
		dim := 2
		num := 2000
		nTree := 10
		var rawItems [][]float32

		for i := 0; i < num; i++ {
			v := make([]float32, 2)
			for d := 0; d < dim; d++ {
				if rand.Int()%2 == 0 {
					v[d] = rand.Float32()
				} else {
					v[d] = -rand.Float32()
				}

			}
			rawItems = append(rawItems, v)
		}

		rawItems[0] = rawItems[1]
		rawItems[1] = rawItems[2]

		idx, err := Initialize(rawItems, dim, nTree, k, true)
		if err != nil {
			panic(err)
		}

		idx.Build()
		assert.Equal(t, nTree, len(idx.roots))
		assert.Equal(t, true, len(idx.nodes) > nTree)
		assert.Equal(t, true, len(idx.nodeIDToNode) > nTree)

		for _, n := range idx.nodes {
			if !n.IsLeaf() {
				assert.Equal(t, true, len(n.Leaf) == 0)
			}
		}
	}
}

func TestGetANNByVector(t *testing.T) {
	k := 10
	dim := 20
	num := 10000
	nTree := 10

	var rawItems [][]float32
	for i := 0; i < num; i++ {
		v := make([]float32, dim)
		for d := 0; d < dim; d++ {
			if rand.Int()%2 == 0 {
				v[d] = rand.Float32()
			} else {
				v[d] = -rand.Float32()
			}

		}
		item.Normalize(v)
		rawItems = append(rawItems, v)
	}

	// build index
	idx, err := Initialize(rawItems, dim, nTree, k, false)
	if err != nil {
		panic(err)
	}
	idx.Build()
	for _, n := range idx.nodes {
		if n.IsLeaf() {
			assert.Equal(t, true, len(n.Leaf) < k)
		} else {
			assert.Equal(t, true, len(n.Leaf) == 0)
		}
	}

	query := make([]float32, dim)
	query[0] = 0.1

	// exact neighbors
	aSims := map[int64]float32{}
	ids := make([]int64, len(rawItems))
	for i, v := range rawItems {
		ids[i] = int64(i)
		aSims[int64(i)] = item.DotProduct(v, query)
	}
	sort.Slice(ids, func(i, j int) bool {
		return -aSims[ids[i]] < -aSims[ids[j]]
	})

	searchNum := 10
	for i, id := range ids[:searchNum] {
		t.Logf("%d-th true neighbor: %d. The similarity: %f", i, id, aSims[id])
	}

	ass, err := idx.getANNbyVector(query, searchNum, 100.0)
	if err != nil {
		panic(err)
	}
	for i, id := range ass {
		t.Logf("%d-th approximated neighbor: %d. The similarity: %f", i, id, item.DotProduct(rawItems[id], query))
	}
}
