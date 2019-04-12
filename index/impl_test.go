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

	var k = 6
	var dim = 2
	var num = 20
	var nTree = 10
	var rawItems [][]float32

	for i := 0; i < num; i++ {
		v := make([]float32, dim)
		for d := 0; d < dim; d++ {
			v[d] = rand.Float32()
		}
		rawItems = append(rawItems, v)
	}

	idx := Initialize(rawItems, dim, nTree, k, true)

	idx.initRootNodes()
	assert.Equal(t, nTree, len(idx.Roots))
	assert.Equal(t, nTree, len(idx.nodes))

}

func TestBuild(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	for _, k := range []int{2, 10, 100} {
		var dim = 2
		var num = 2000
		var nTree = 10
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
			rawItems = append(rawItems, v)
		}

		rawItems[0] = rawItems[1]
		rawItems[1] = rawItems[2]

		idx := Initialize(rawItems, dim, nTree, k, true)

		idx.Build()
		assert.Equal(t, nTree, len(idx.Roots))
		assert.Equal(t, true, len(idx.nodes) > nTree)
		assert.Equal(t, true, len(idx.NodeIDToNode) > nTree)

		for _, n := range idx.nodes {
			if !n.IsLeaf() {
				assert.Equal(t, true, len(n.Leaf) == 0)
			}
		}
	}
}

func TestBuildOnLoadedIndex(t *testing.T) {
	var idx = &Index{}
	err := idx.Build()
	if err == nil {
		t.Fatal("build method on loaded indices should return error.")
	}
}

// This unit test is made to verify if our algorithm can correctly find
// the `exact` neighbors. That is done by checking the ration of exact
// neighbors in the result returned by `getANNbyVector` to searchResult.
// is less than the given threshold.
func TestGetANNByVector(t *testing.T) {
	var k = 2
	var dim = 20
	var num = 10000
	var nTree = 20
	var threshold = 0.01
	var searchNum = 20
	var query = make([]float32, dim)
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
	for _, n := range idx.nodes {
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
