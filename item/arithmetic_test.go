package item

import (
	"math/rand"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestNormalize(t *testing.T) {
	v1 := []float32{2, 0}
	err := Normalize(v1)
	assert.Equal(t, nil, err)
	assert.Equal(t, []float32{1, 0}, v1)
}

func TestDotProduct(t *testing.T) {
	v1 := []float32{1.2, 0.1}
	v2 := []float32{-1.2, 0.2}
	expected := int(float32(-1.42) * 100)
	actual := int(DotProduct(v1, v2) * 100)
	assert.Equal(t, expected, actual)
}

func TestGetNormalVectorOfSplittingHyperPlane(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	dim := 2
	num := 5
	var vs []Vector

	for i := 0; i < num; i++ {
		v := make([]float32, dim)
		for d := 0; d < dim; d++ {
			v[d] = rand.Float32()
		}
		Normalize(v)
		vs = append(vs, v)
	}

	GetNormalVectorOfSplittingHyperPlane(vs, dim)
}
