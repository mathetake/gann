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
	bucketScale float32
	searchNum   int
}

func BenchmarkGetANNByVector1(b *testing.B) {
	tmpl := benchTemplate{
		dim:         300,
		nItem:       100000,
		nTree:       20,
		k:           4,
		bucketScale: 10,
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
	gIDx := GetIndex(its, tmpl.dim, tmpl.nTree, tmpl.k, true)

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

func BenchmarkBuildIndex1(b *testing.B) {
	tmpl := benchTemplate{
		dim:   10,
		nItem: 1000000,
		nTree: 20,
		k:     40,
	}

	its := _getItems(tmpl.dim, tmpl.nItem)

	b.Logf("dim: %d, # of items: %d, # of trees: %d, leaf size: %d", tmpl.dim, tmpl.nItem, tmpl.nTree, tmpl.k)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// create index
		gIDx := GetIndex(its, tmpl.dim, tmpl.nTree, tmpl.k, true)

		// build index
		gIDx.Build()
	}
}

func BenchmarkBuildIndex2(b *testing.B) {
	tmpl := benchTemplate{
		dim:   300,
		nItem: 1000000,
		nTree: 20,
		k:     40,
	}

	its := _getItems(tmpl.dim, tmpl.nItem)

	b.Logf("dim: %d, # of items: %d, # of trees: %d, leaf size: %d", tmpl.dim, tmpl.nItem, tmpl.nTree, tmpl.k)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// create index
		gIDx := GetIndex(its, tmpl.dim, tmpl.nTree, tmpl.k, true)

		// build index
		gIDx.Build()
	}
}

func BenchmarkBuildIndex3(b *testing.B) {
	tmpl := benchTemplate{
		dim:   1000,
		nItem: 1000000,
		nTree: 20,
		k:     40,
	}

	its := _getItems(tmpl.dim, tmpl.nItem)

	b.Logf("dim: %d, # of items: %d, # of trees: %d, leaf size: %d", tmpl.dim, tmpl.nItem, tmpl.nTree, tmpl.k)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// create index
		gIDx := GetIndex(its, tmpl.dim, tmpl.nTree, tmpl.k, true)

		// build index
		gIDx.Build()
	}
}
