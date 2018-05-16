package index

import (
	"math/rand"
	"testing"
	"time"

	"github.com/bmizerany/assert"
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

	k := 6
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

	idx, err := Initialize(rawItems, dim, nTree, k, true)
	if err != nil {
		panic(err)
	}

	idx.Build()
	assert.Equal(t, nTree, len(idx.roots))
	assert.Equal(t, true, len(idx.nodes) > nTree)
	assert.Equal(t, true, len(idx.nodeIDToNode) > nTree)

	for _, n := range idx.nodes {
		if n.IsLeaf() {
			assert.Equal(t, true, len(n.Leaf) < k)
		} else {
			assert.Equal(t, true, len(n.Leaf) == 0)
		}
		assert.Equal(t, idx.nodes, *n.Forest)
	}
}

func TestGetANNByVector(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	k := 2
	dim := 2
	num := 100
	nTree := 1

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

	idx, err := Initialize(rawItems, dim, nTree, k, true)
	if err != nil {
		panic(err)
	}
	idx.Build()

	for _, n := range idx.nodes {
		if n.IsLeaf() {
			assert.Equal(t, 1, len(n.Leaf))
		} else {
			assert.Equal(t, true, len(n.Leaf) == 0)
		}
		assert.Equal(t, idx.nodes, *n.Forest)
	}
}
