package item

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	minIteration = 200
)

func Normalize(v1 Vector) {
	n := norm(v1)
	if n == 0 {
		panic("zero vector given.")
	}
	for i := 0; i < len(v1); i++ {
		v1[i] = v1[i] / n
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
func GetNormalVectorOfSplittingHyperPlane(vs []Vector, dim int) Vector {
	lvs := len(vs)
	maxIteration := lvs / 20
	threshold := float32(0.8)

	if maxIteration < minIteration {
		maxIteration = minIteration
	}

	// init centroids
	rand.Seed(time.Now().UnixNano())
	k := rand.Intn(lvs)
	l := rand.Intn(lvs - 1)
	if k == l {
		l++
	}
	c0 := vs[k]
	c1 := vs[l]

	ret := make([]float32, dim)
	for i := 0; ; i++ {
		clusterToVecs := map[int][]Vector{}
		for _, v := range vs {
			ip0 := DotProduct(c0, v)
			ip1 := DotProduct(c1, v)
			if ip0 < ip1 {
				clusterToVecs[0] = append(clusterToVecs[0], v)
			} else {
				clusterToVecs[1] = append(clusterToVecs[1], v)
			}
		}

		lc0 := len(clusterToVecs[0])
		lc1 := len(clusterToVecs[1])

		if (float32(lc0)/float32(lvs) <= threshold) && (float32(lc1)/float32(lvs) <= threshold) {
			break
		} else if i > maxIteration {
			break
		}

		// update centroids
		c0 = make([]float32, dim)
		c1 = make([]float32, dim)
		for _, v := range clusterToVecs[0] {
			for d := 0; d < dim; d++ {
				c0[d] += v[d] / float32(lc0)
			}
		}
		for _, v := range clusterToVecs[1] {
			for d := 0; d < dim; d++ {
				c1[d] += v[d] / float32(lc1)
			}
		}
	}

	for d := 0; d < dim; d++ {
		ret[d] += c0[d] - c1[d]
	}

	// normalize
	Normalize(ret)
	return ret
}

func norm(v1 Vector) float32 {
	var n32 float32
	for _, v := range v1 {
		n32 += v * v
	}
	n64 := math.Sqrt(float64(n32))
	return float32(n64)
}
