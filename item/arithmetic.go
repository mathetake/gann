package item

func DotProduct(v1, v2 Vector) (ret float32) {
	if len(v1) != len(v2) {
		panic("Dimension mismatch.")
	}
	for i := 0; i < len(v1); i++ {
		ret += v1[i] * v2[i]
	}
	return ret
}

// get normal vector which is perpendicular to the splitting hyperplane.
// We chose the vector so that it is the average vectro of a given set of data points.
func GetNormalVectorOfSplittingHyperPlane(vs []Vector, dim int) (nv Vector, err error) {
	for _, v := range vs {
		for i:= 0; i < dim; i++ {
			nv[i] += v[i]
		}
	}
	for i:= 0; i < dim; i++ {
		nv[i] /= float32(len(vs))
	}
	return nv, err
}
