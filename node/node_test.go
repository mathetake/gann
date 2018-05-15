package node

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/mathetake/gann/item"
)

func TestIsLeaf(t *testing.T) {
	n := Node{
		Leaf: []int64{1,1},
	}
	assert.Equal(t, true, n.IsLeaf())
	n.Leaf = []int64{}
	assert.Equal(t, false, n.IsLeaf())
}


func TestBuild(t *testing.T) {
	n := Node{
		NDescendants:2,
	}
	its := []item.Item{
		{ID: 1, Vec: []float32{0.1,0.1},},
		{ID: 2, Vec: []float32{0.1,0.1},},
	}
	err := n.Build(its, 10, 0)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, []int64{1,2}, n.Leaf)
}

func TestChildren(t *testing.T) {
	k := 6
	dim := 2
	var its []item.Item

	// positive side
	for i:= 0; i < 5; i++ {
		v := make([]float32, 2)
		for d := 0; d < dim; d++ {
			v[d] = 0.1 * float32(i+1)
		}
		its = append(its, item.Item{
			ID: int64(i),
			Vec: v,
		})
	}

	// negative side
	for i:= 0; i < 5; i++ {
		v := make([]float32, 2)
		for d := 0; d < dim; d++ {
			v[d] = - 0.1 * float32(i+1)
		}
		its = append(its, item.Item{
			ID: int64(i + 5),
			Vec: v,
		})
	}

	n := Node{
		Vec: []float32{
			1, 1,
		},
		Forest: &[]*Node{},
	}

	err := n.buildChildren(its, k, 2)
	if err != nil {
		panic(err)
	}

	leftChild := n.Children[0]
	assert.Equal(t, true, leftChild.IsLeaf())
	assert.Equal(t, []int64{0,1,2,3,4,}, leftChild.Leaf)


	rightChild := n.Children[1]
	assert.Equal(t, true, rightChild.IsLeaf())
	assert.Equal(t, []int64{5,6,7,8,9,}, rightChild.Leaf)
}