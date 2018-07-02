package index

import (
	"testing"

	"github.com/bmizerany/assert"
	"github.com/mathetake/gann/item"
)

func TestInitializeWithNormalize(t *testing.T) {
	rawItems := [][]float32{
		{2, 0},
		{0, 2},
	}
	d := 2
	idx := Initialize(rawItems, d, 1, 1, true)

	assert.Equal(t, 2, len(idx.ItemIDToItem))
	assert.Equal(t, item.Item{
		ID:  0,
		Vec: []float32{1, 0},
	}, idx.ItemIDToItem[0])
	assert.Equal(t, item.Item{
		ID:  1,
		Vec: []float32{0, 1},
	}, idx.ItemIDToItem[1])
}

func TestInitializeWithoutNormalize(t *testing.T) {
	rawItems := [][]float32{
		{2, 0},
		{0, 2},
	}
	d := 2
	idx := Initialize(rawItems, d, 1, 1, false)

	assert.Equal(t, 2, len(idx.ItemIDToItem))
	assert.Equal(t, item.Item{
		ID:  0,
		Vec: []float32{2, 0},
	}, idx.ItemIDToItem[0])
	assert.Equal(t, item.Item{
		ID:  1,
		Vec: []float32{0, 2},
	}, idx.ItemIDToItem[1])
}
