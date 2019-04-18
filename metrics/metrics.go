package metrics

import "github.com/pkg/errors"

type Type string

const (
	TypeCosineDistance Type = "cosineDistance"
)

var (
	ErrInvalidMetricsType = errors.New("invalid metrics type")
)

type Metrics interface {
	CalcDistance(v1, v2 []float64) float64
	GetNormalVectorOfSplittingHyperPlane(vs [][]float64) []float64
	GetDirectionPriority(base, target []float64) float64
}

func NewMetrics(t Type, dim int) (Metrics, error) {
	switch t {
	case TypeCosineDistance:
		return &cosineDistance{}, nil
	default:
		return nil, ErrInvalidMetricsType
	}
}
