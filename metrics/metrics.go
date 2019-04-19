package metrics

import "github.com/pkg/errors"

type Type int

const (
	TypeCosineDistance Type = iota
)

var (
	ErrInvalidMetricsType = errors.New("invalid metrics type")
)

type Metrics interface {
	CalcDirectionPriority(base, target []float64) float64
	CalcDistance(v1, v2 []float64) float64
	GetSplittingVector(vs [][]float64) []float64
}

func NewMetrics(t Type, dim int) (Metrics, error) {
	switch t {
	case TypeCosineDistance:
		return &cosineDistance{
			dim: dim,
		}, nil
	default:
		return nil, ErrInvalidMetricsType
	}
}
