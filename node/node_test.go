package node

import (
	"math/rand"
	"testing"
	"time"

	"sync"

	"github.com/bmizerany/assert"
	"github.com/mathetake/gann/item"
)

func TestIsLeaf(t *testing.T) {
	n := Node{
		Leaf: []int64{1, 1},
	}
	assert.Equal(t, true, n.IsLeaf())
	n.Leaf = []int64{}
	assert.Equal(t, false, n.IsLeaf())
}

func TestBuild(t *testing.T) {
	n := Node{
		NDescendants: 2,
	}
	its := []item.Item{
		{ID: 1, Vec: []float32{0.1, 0.1}},
		{ID: 2, Vec: []float32{0.1, 0.1}},
	}
	n.Build(its, 10, 0, &sync.Map{})
	assert.Equal(t, []int64{1, 2}, n.Leaf)
}

// both child are supposed to be leaf nodes
func TestBuildChildren1(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var k = 6
	var dim = 2
	var num = 5
	var its []item.Item

	// positive side
	for i := 0; i < num; i++ {
		v := make([]float32, 2)
		for d := 0; d < dim; d++ {
			v[d] = rand.Float32()
		}
		item.Normalize(v)
		its = append(its, item.Item{
			ID:  int64(i),
			Vec: v,
		})
	}

	// negative side
	for i := 0; i < num; i++ {
		v := make([]float32, 2)
		for d := 0; d < dim; d++ {
			v[d] = -rand.Float32()
		}
		item.Normalize(v)
		its = append(its, item.Item{
			ID:  int64(num + i),
			Vec: v,
		})
	}

	n := Node{
		Vec: []float32{
			1, 0,
		},
	}

	n.buildChildren(its, k, 2, &sync.Map{})

	leftChild := n.Children[0]
	assert.Equal(t, true, leftChild.IsLeaf())
	assert.Equal(t, []int64{0, 1, 2, 3, 4}, leftChild.Leaf)

	rightChild := n.Children[1]
	assert.Equal(t, true, rightChild.IsLeaf())
	assert.Equal(t, []int64{5, 6, 7, 8, 9}, rightChild.Leaf)
}

// Only one of children is supposed to be leaf node
func TestBuildChildren2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	var k = 6
	var dim = 2
	var pNum = 5
	var nNum = 10
	var its []item.Item

	// positive side
	for i := 0; i < pNum; i++ {
		v := make([]float32, 2)
		for d := 0; d < dim; d++ {
			v[d] = rand.Float32()
		}
		item.Normalize(v)
		its = append(its, item.Item{
			ID:  int64(i),
			Vec: v,
		})
	}

	// negative side
	for i := 0; i < nNum; i++ {
		v := make([]float32, 2)
		for d := 0; d < dim; d++ {
			v[d] = -rand.Float32()
		}
		item.Normalize(v)
		its = append(its, item.Item{
			ID:  int64(i + pNum),
			Vec: v,
		})
	}

	n := Node{
		Vec: []float32{
			1, 0,
		},
	}

	n.buildChildren(its, k, 2, &sync.Map{})

	leftChild := n.Children[0]
	assert.Equal(t, true, leftChild.IsLeaf())
	assert.Equal(t, []int64{0, 1, 2, 3, 4}, leftChild.Leaf)

	rightChild := n.Children[1]
	assert.Equal(t, false, rightChild.IsLeaf())
	assert.Equal(t, 2, len(rightChild.Children))
}
