package gann

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/bmizerany/assert"
	"github.com/mathetake/gann/metric"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestCreateNewIndex(t *testing.T) {
	for i, c := range []struct {
		dim, num, nTree, k int
	}{
		{dim: 2, num: 1000, nTree: 10, k: 2},
		{dim: 10, num: 100, nTree: 5, k: 10},
		{dim: 10, num: 100000, nTree: 5, k: 10},
		{dim: 1000, num: 10000, nTree: 5, k: 10},
	} {
		c := c
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			rawItems := make([][]float64, c.num)
			for i := range rawItems {
				v := make([]float64, c.dim)

				var norm float64
				for j := range v {
					cof := rand.Float64() - 0.5
					v[j] = cof
					norm += cof * cof
				}

				norm = math.Sqrt(norm)
				for j := range v {
					v[j] /= norm
				}

				rawItems[i] = v
			}

			m, err := metric.NewCosineMetric(c.dim)
			if err != nil {
				t.Fatal(err)
			}

			idx, err := CreateNewIndex(rawItems, c.dim, c.nTree, c.k, m)
			if err != nil {
				t.Fatal(err)
			}

			rawIdx, ok := idx.(*index)
			if !ok {
				t.Fatal("type assertion failed")
			}

			assert.Equal(t, c.nTree, len(rawIdx.roots))
			assert.Equal(t, true, len(rawIdx.nodeIDToNode) > c.nTree)
		})
	}

}

func TestCreateNewIndexNotEnoughItems(t *testing.T) {
	rawItems := make([][]float64, 1)
	rawItems[0] = []float64{1, 2, 3, 4}

	m, err := metric.NewCosineMetric(4)
	if err != nil {
		t.Fatal(err)
	}

	//1 vector is not enough
	_, err = CreateNewIndex(rawItems, 4, 4, 2, m)
	if err != errNotEnoughItems {
		t.Fatalf("expected error errNotEnoughItems, got %v instead", err)
	}

	rawItems2 := make([][]float64, 2)
	rawItems2[0] = []float64{1, 2, 3, 4}
	rawItems2[1] = []float64{2, 2, 2, 2}

	//2 vectors are ok
	_, err = CreateNewIndex(rawItems2, 4, 4, 2, m)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

}
