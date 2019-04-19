package gann

import (
	"fmt"
	"math"
	"math/rand"
	"testing"

	"github.com/mathetake/gann/metrics"
)

func TestIndex_GetANNbyItemID(t *testing.T) {
	for i, c := range []struct {
		dim, num, nTree, k int
	}{
		{dim: 2, num: 1000, nTree: 10, k: 2},
		{dim: 10, num: 100, nTree: 5, k: 10},
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

			idx, err := CreateNewIndex(rawItems, c.dim, c.nTree,
				c.k, metrics.TypeCosineDistance)
			if err != nil {
				t.Fatal(err)
			}

			ann, err := idx.GetANNbyItemID(0, 10, 2)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(ann)
		})
	}
}

func TestIndex_GetANNbyVector(t *testing.T) {
	for i, c := range []struct {
		dim, num, nTree, k int
	}{
		{dim: 2, num: 1000, nTree: 10, k: 2},
		{dim: 10, num: 100, nTree: 5, k: 10},
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

			idx, err := CreateNewIndex(rawItems, c.dim, c.nTree,
				c.k, metrics.TypeCosineDistance)
			if err != nil {
				t.Fatal(err)
			}

			key := make([]float64, c.dim)
			for i := range key {
				key[i] = rand.Float64() - 0.5
			}

			ann, err := idx.GetANNbyVector(key, 10, 2)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(ann)
		})
	}
}
