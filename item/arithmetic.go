package item

import (
	"github.com/pkg/errors"
)

func DotProduct(v1, v2 Vector) (ret float32) {
	if len(v1) != len(v2) {
		panic("Dimension mismatch.")
	}
	for i := 0; i < len(v1); i++ {
		ret += v1[i] * v2[i]
	}
	return ret
}

func GetNormalVectorOfSplittingHyperPlane(its []Item) (nv Vector, err error) {
	cs, err := twoMeans(its)
	if err != nil {
		return nv, errors.Wrap(err, "TwoMeans failed.")
	}
	nv = subtract(cs[0], cs[1])
	return nv, err
}

// Given a set of vectors, do 2-means algorithm and returns its centroids.
func twoMeans(its []Item) (cs map[int]Vector, err error) {
	return cs, nil
}

func subtract(v1, v2 Vector) Vector {
	if len(v1) != len(v2) {
		panic("dimension mimatch")
	}
	v := make([]float32, len(v1))
	for i := 0; i < len(v1); i++ {
		v[i] = v1[i] - v2[i]
	}
	return Vector(v)
}
