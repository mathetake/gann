package gann

import (
	"fmt"
	"sync"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/mathetake/gann/metric"
)

func TestNodeBuild(t *testing.T) {
	for i, c := range []struct {
		vec     []float64
		items   []*item
		dim, k  int
		expLeaf bool
	}{
		{
			vec: []float64{0.0, 1.0},
			items: []*item{
				{id: 0, vector: []float64{0.0, 1.0}},
				{id: 1, vector: []float64{0.0, -1.0}},
			},
			k:       2,
			dim:     2,
			expLeaf: true,
		},
		{
			vec: []float64{0.0, 1.0},
			items: []*item{
				{id: 0, vector: []float64{0.0, 1.0}},
				{id: 0, vector: []float64{0.0, 1.1}},
				{id: 0, vector: []float64{0.0, 1.2}},
				{id: 1, vector: []float64{0.0, -1.0}},
				{id: 2, vector: []float64{0.0, -1.1}},
				{id: 2, vector: []float64{0.0, -1.2}},
			},
			k:       2,
			dim:     2,
			expLeaf: false,
		},
	} {
		c := c
		i := i
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			m, err := metric.NewCosineMetric(c.dim)
			if err != nil {
				t.Fatal(err)
			}

			idxPtr := &index{
				k:            1,
				mux:          &sync.Mutex{},
				metric:       m,
				nodeIDToNode: map[nodeId]*node{},
			}

			n := &node{
				id:       nodeId(fmt.Sprintf("%d", i)),
				vec:      c.vec,
				idxPtr:   idxPtr,
				children: make(map[direction]*node, len(directions)),
			}
			n.build(c.items)

			if c.expLeaf {
				assert.Equal(t, true, len(n.leaf) > 0)
			} else {
				assert.Equal(t, true, len(n.leaf) == 0)
				assert.Equal(t, true, len(n.children) > 0)
			}
		})
	}
}
