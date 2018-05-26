package item

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"github.com/pkg/errors"
)

const (
	maxIteration      = 200
	maxTargetSample   = 100
	twoMeansThreshold = float32(0.7)
	centroidCalcRatio = float32(0.0001)
)

func Normalize(v Vector) error {
	n := norm(v)
	if n == 0 {
		return errors.Errorf("zero vector is given.")
	}
	for i := 0; i < len(v); i++ {
		v[i] = v[i] / n
	}
	return nil
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
	for i := 0; i < maxIteration; i++ {
		clusterToVecs := map[int][]Vector{}

		iter := maxTargetSample
		if len(vs) < maxTargetSample {
			iter = len(vs)
		}
		for i := 0; i < iter; i++ {
			v := vs[rand.Intn(len(vs))]
			ip0 := DotProduct(c0, v)
			ip1 := DotProduct(c1, v)
			if ip0 > ip1 {
				clusterToVecs[0] = append(clusterToVecs[0], v)
			} else {
				clusterToVecs[1] = append(clusterToVecs[1], v)
			}
		}

		lc0 := len(clusterToVecs[0])
		lc1 := len(clusterToVecs[1])

		if (float32(lc0)/float32(iter) <= twoMeansThreshold) && (float32(lc1)/float32(iter) <= twoMeansThreshold) {
			break
		}

		// update centroids
		if lc0 == 0 || lc1 == 0 {
			k := rand.Intn(lvs)
			l := rand.Intn(lvs - 1)
			if k == l {
				l++
			}
			c0 = vs[k]
			c1 = vs[l]
			continue
		}

		c0 = make([]float32, dim)
		it0 := int(float32(lvs) * centroidCalcRatio)
		for i := 0; i < it0; i++ {
			for d := 0; d < dim; d++ {
				c0[d] += clusterToVecs[0][rand.Intn(lc0)][d] / float32(it0)
			}
		}

		c1 = make([]float32, dim)
		it1 := int(float32(lvs)*centroidCalcRatio + 1)
		for i := 0; i < int(float32(lc1)*centroidCalcRatio+1); i++ {
			for d := 0; d < dim; d++ {
				c1[d] += clusterToVecs[1][rand.Intn(lc1)][d] / float32(it1)
			}
		}
	}

	for d := 0; d < dim; d++ {
		v := c0[d] - c1[d]
		ret[d] += v
	}

	// normalize
	err := Normalize(ret)
	if err != nil {
		d := rand.Intn(dim)
		ret[d] = 1
	}
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
