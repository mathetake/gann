package item

func DotProduct (v1, v2 Vector) (ret float32) {
	if len(v1) != len(v2) {
		panic("Dimension mismatch.")
	}
	for i := 0; i < len(v1); i++ {
		ret += v1[i]*v2[i]
	}
	return ret
}
