package item

import (
	"math"
	"math/rand"
	"time"
	"fmt"
)

const (
	minIteration = 20
)

func Normalize(v1 Vector) {
	var n32 float32
	for _, v := range v1 {
		n32 += v * v
	}
	n64 := math.Sqrt(float64(n32))
	n32 = float32(n64)

	for i := 0; i < len(v1); i++ {
		v1[i] = v1[i] / n32
	}
}

func DotProduct(v1, v2 Vector) (ret float32) {
	if len(v1) != len(v2) {
		panic(fmt.Sprintf("Dimension mismatch: %d != %d", len(v1), len(v2)))
	}
	for i := 0; i < len(v1); i++ {
		ret += v1[i] * v2[i]
	}
	return ret
}

// get normal vector which is perpendicular to the splitting hyperplane.
// We chose the vector so that it is the average vector of a given set of data points.
func GetNormalVectorOfSplittingHyperPlane(vs []Vector, dim int) Vector {
	lvs := len(vs)
	iter := lvs / 20
	if iter < minIteration {
		iter = minIteration
	}

	rand.Seed(time.Now().UnixNano())

	nvs := make([]Vector, iter)
	for i := 0; i < iter; i++ {
		k := rand.Intn(lvs)
		l := rand.Intn(lvs - 1)
		if k == l {
			l++
		}
		diff := make([]float32, dim)
		for m := 0; m < dim; m++ {
			diff[m] = vs[k][m] - vs[l][m]
		}
		nvs[i] = diff
	}

	ret := make([]float32, dim)
	for i := 0; i < dim; i++ {
		for _, v := range nvs {
			ret[i] += v[i] / float32(len(nvs))
		}
	}

	return ret
}
