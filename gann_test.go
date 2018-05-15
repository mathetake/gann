package gann

import (
	"math/rand"
	"testing"
	"time"
)

type benchTemplate struct {
	dim         int
	nItem       int
	nTree       int
	k           int
	bucketScale float64
	searchNum   int
}

func BenchmarkGetANNByVector1(b *testing.B) {
	tmpl := benchTemplate{
		dim:         300,
		nItem:       100000,
		nTree:       20,
		k:           4,
		bucketScale: 2,
		searchNum:   50,
	}
	gIDx := _getTestIndex(&tmpl)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q := _getRandomVector(tmpl.dim)
		_, err := gIDx.GetANNbyVector(q, tmpl.searchNum, tmpl.bucketScale)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGetANNByVector2(b *testing.B) {
	tmpl := benchTemplate{
		dim:         300,
		nItem:       1000000,
		nTree:       20,
		k:           4,
		bucketScale: 2,
		searchNum:   500,
	}
	gIDx := _getTestIndex(&tmpl)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q := _getRandomVector(tmpl.dim)
		_, err := gIDx.GetANNbyVector(q, tmpl.searchNum, tmpl.bucketScale)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGetANNByVector3(b *testing.B) {
	tmpl := benchTemplate{
		dim:         2000,
		nItem:       100000,
		nTree:       20,
		k:           40,
		bucketScale: 2,
		searchNum:   500,
	}

	gIDx := _getTestIndex(&tmpl)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q := _getRandomVector(tmpl.dim)
		_, err := gIDx.GetANNbyVector(q, tmpl.searchNum, tmpl.bucketScale)
		if err != nil {
			panic(err)
		}
	}
}

func _getTestIndex(tmpl *benchTemplate) GannIndex {
	its := _getItems(tmpl.dim, tmpl.nItem)

	// create index
	gIDx, err := GetIndex(its, tmpl.dim, tmpl.nTree, tmpl.k, true)
	if err != nil {
		panic(err)
	}
	// build index
	gIDx.Build()
	return gIDx
}

func _getItems(dim int, l int) [][]float32 {
	data := [][]float32{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 1000; i++ {
		data = append(data, _getRandomVector(dim))
	}
	return data
}

func _getRandomVector(dim int) []float32 {
	rand.Seed(time.Now().UnixNano())
	v := make([]float32, dim)
	for j := 0; j < dim; j++ {
		v[j] = rand.Float32()
	}
	return v
}
