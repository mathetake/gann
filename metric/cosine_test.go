package metric

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/bmizerany/assert"
)

func TestCosineDistance_CalcDirectionPriority(t *testing.T) {
	for i, c := range []struct {
		v1, v2 []float64
		exp    float64
		dim    int
	}{
		{
			v1:  []float64{1.2, 0.1},
			v2:  []float64{-1.2, 0.2},
			dim: 2,
			exp: -1.42,
		},
		{
			v1:  []float64{1.2, 0.1, 0, 0, 0, 0, 0, 0, 0, 0},
			v2:  []float64{-1.2, 0.2, 0, 0, 0, 0, 0, 0, 0, 0},
			dim: 10,
			exp: -1.42,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			cosine := &cosineDistance{dim: c.dim}
			actual := cosine.CalcDirectionPriority(c.v1, c.v2)
			assert.Equal(t, c.exp, actual)
		})
	}
}

func TestCosineDistance_GetSplittingVector(t *testing.T) {
	for i, c := range []struct {
		dim, num int
	}{
		{
			dim: 5, num: 100,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			cosine := &cosineDistance{dim: c.dim}
			vs := make([][]float64, c.num)
			for i := 0; i < c.num; i++ {
				v := make([]float64, c.dim)
				for d := 0; d < c.dim; d++ {
					v[d] = rand.Float64()
				}
				vs[i] = v
			}

			cosine.GetSplittingVector(vs)
		})
	}
}

func TestCosineDistance_CalcDistance(t *testing.T) {
	for i, c := range []struct {
		v1, v2 []float64
		exp    float64
		dim    int
	}{
		{
			v1:  []float64{1.2, 0.1},
			v2:  []float64{-1.2, 0.2},
			dim: 2,
			exp: 1.42,
		},
		{
			v1:  []float64{1.2, 0.1, 0, 0, 0, 0, 0, 0, 0, 0},
			v2:  []float64{-1.2, 0.2, 0, 0, 0, 0, 0, 0, 0, 0},
			dim: 10,
			exp: 1.42,
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			cosine := &cosineDistance{dim: c.dim}
			actual := cosine.CalcDistance(c.v1, c.v2)
			assert.Equal(t, c.exp, actual)
		})
	}
}
