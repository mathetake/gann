package metrics

import "github.com/pkg/errors"

type Type string

const (
	TypeCosineDistance = iota
)

var (
	ErrInvalidMetricsType = errors.New("invalid metrics type")
)

type Metrics interface {
	CalcDistance(v1, v2 []float64) float64
	// GetNormalVectorOfSplittingHyperPlane ... get normal vector which is perpendicular to the splitting hyperplane.
	GetNormalVectorOfSplittingHyperPlane(vs [][]float64) []float64
	ShouldGoLeft(base, target []float64) bool
}

func NewMetrics(t Type, dim int) (Metrics, error) {
	switch t {
	case TypeCosineDistance:
		return &cosineDistance{}, nil
	default:
		return nil, ErrInvalidMetricsType
	}
}
