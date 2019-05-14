package metric

type Metric interface {
	CalcDirectionPriority(base, target []float64) float64
	CalcDistance(v1, v2 []float64) float64
	GetSplittingVector(vs [][]float64) []float64
}
