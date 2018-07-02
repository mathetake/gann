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

	idx := Initialize(rawItems, dim, nTree, k, true)

	idx.initRootNodes()
	assert.Equal(t, nTree, len(idx.Roots))
	assert.Equal(t, nTree, len(idx.Nodes))

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

		idx := Initialize(rawItems, dim, nTree, k, true)

		idx.Build()
		assert.Equal(t, nTree, len(idx.Roots))
		assert.Equal(t, true, len(idx.Nodes) > nTree)
		assert.Equal(t, true, len(idx.NodeIDToNode) > nTree)

		for _, n := range idx.Nodes {
			if !n.IsLeaf() {
				assert.Equal(t, true, len(n.Leaf) == 0)
			}
		}
	}
}

// This unit test is made to verify if our algorithm can correctly find
// the `exact` neighbors. That is done by checking the ration of exact
// neighbors in the result returned by `getANNbyVector` to searchResult.
// is less than the given threshold.
func TestGetANNByVector(t *testing.T) {
	k := 2
	dim := 20
	num := 10000
	nTree := 20
	threshold := 0.5
	searchNum := 20
	query := make([]float32, dim)
	query[0] = 0.1

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
	idx := Initialize(rawItems, dim, nTree, k, false)

	idx.Build()
	for _, n := range idx.Nodes {
		if n.IsLeaf() {
			assert.Equal(t, true, len(n.Leaf) < k)
		} else {
			assert.Equal(t, true, len(n.Leaf) == 0)
		}
	}

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

	expectedIDsMap := make(map[int64]interface{}, searchNum)
	for _, id := range ids[:searchNum] {
		expectedIDsMap[int64(id)] = true
	}

	ass, err := idx.getANNbyVector(query, searchNum, 100.0)
	if err != nil {
		panic(err)
	}

	var count int
	for _, id := range ass {
		if _, ok := expectedIDsMap[id]; ok {
			count++
		}
	}

	if float64(count)/float64(searchNum) < threshold {
		t.Errorf("Too few exact neighbor in approximated result.")
	}
}
