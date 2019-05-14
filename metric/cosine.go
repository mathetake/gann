package metric

import (
	"math/rand"
	"time"
)

const (
	cosineMetricsMaxIteration      = 200
	cosineMetricsMaxTargetSample   = 100
	cosineMetricsTwoMeansThreshold = 0.7
	cosineMetricsCentroidCalcRatio = 0.0001
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type cosineDistance struct {
	dim int
}

// NewCosineMetric returns cosineDistance.
// NOTE: We assume that the given vectors are already normalize, i.e. the norm equals 1
func NewCosineMetric(dim int) (Metric, error) {
	return &cosineDistance{
		dim: dim,
	}, nil
}

func (c *cosineDistance) CalcDistance(v1, v2 []float64) float64 {
	var ret float64
	for i := range v1 {
		ret += v1[i] * v2[i]
	}
	return -ret
}

func (c *cosineDistance) GetSplittingVector(vs [][]float64) []float64 {
	lvs := len(vs)
	// init centroids
	k := rand.Intn(lvs)
	l := rand.Intn(lvs - 1)
	if k == l {
		l++
	}
	c0 := vs[k]
	c1 := vs[l]

	for i := 0; i < cosineMetricsMaxIteration; i++ {
		clusterToVecs := map[int][][]float64{}

		iter := cosineMetricsMaxTargetSample
		if len(vs) < cosineMetricsMaxTargetSample {
			iter = len(vs)
		}
		for i := 0; i < iter; i++ {
			v := vs[rand.Intn(len(vs))]
			ip0 := c.CalcDistance(c0, v)
			ip1 := c.CalcDistance(c1, v)
			if ip0 > ip1 {
				clusterToVecs[0] = append(clusterToVecs[0], v)
			} else {
				clusterToVecs[1] = append(clusterToVecs[1], v)
			}
		}

		lc0 := len(clusterToVecs[0])
		lc1 := len(clusterToVecs[1])

		if (float64(lc0)/float64(iter) <= cosineMetricsTwoMeansThreshold) &&
			(float64(lc1)/float64(iter) <= cosineMetricsTwoMeansThreshold) {
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

		c0 = make([]float64, c.dim)
		it0 := int(float64(lvs) * cosineMetricsCentroidCalcRatio)
		for i := 0; i < it0; i++ {
			for d := 0; d < c.dim; d++ {
				c0[d] += clusterToVecs[0][rand.Intn(lc0)][d] / float64(it0)
			}
		}

		c1 = make([]float64, c.dim)
		it1 := int(float64(lvs)*cosineMetricsCentroidCalcRatio + 1)
		for i := 0; i < int(float64(lc1)*cosineMetricsCentroidCalcRatio+1); i++ {
			for d := 0; d < c.dim; d++ {
				c1[d] += clusterToVecs[1][rand.Intn(lc1)][d] / float64(it1)
			}
		}
	}

	ret := make([]float64, c.dim)
	for d := 0; d < c.dim; d++ {
		v := c0[d] - c1[d]
		ret[d] += v
	}
	return ret
}

func (c *cosineDistance) CalcDirectionPriority(base, target []float64) float64 {
	var ret float64
	for i := range base {
		ret += base[i] * target[i]
	}
	return ret
}
