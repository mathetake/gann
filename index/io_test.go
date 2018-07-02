package index

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

var tcs = []struct {
	dim, nTrees, k, nItem int
}{
	{
		dim:    10,
		nTrees: 5,
		k:      4,
		nItem:  100,
	},
	{
		dim:    100,
		nTrees: 5,
		k:      20,
		nItem:  1000,
	},
}

var path = "tmp.json"

func TestSaveAndLoad(t *testing.T) {
	for _, tc := range tcs {
		rawItems := make([][]float32, 0, tc.nItem)
		rand.Seed(time.Now().UnixNano())

		for i := 0; i < tc.nItem; i++ {
			item := make([]float32, 0, tc.dim)
			for i := 0; i < tc.dim; i++ {
				item = append(item, rand.Float32())
			}
			rawItems = append(rawItems, item)
		}
		gIDx := GetIndex(rawItems, tc.dim, tc.nTrees, tc.k, true)
		gIDx.Build()

		// do search
		q := make([]float32, tc.dim)
		q[0] = 0.1

		expectedANN, _ := gIDx.GetANNbyVector(q, 5, 10)

		if err := gIDx.Save(path); err != nil {
			t.Fatal(err)
		}

		var idx = &Index{}
		if err := idx.Load(path); err != nil {
			t.Fatal(err)
		}

		actualANN, _ := idx.GetANNbyVector(q, 5, 10)
		assert.Equal(t, expectedANN, actualANN)

		if err := os.Remove(path); err != nil {
			t.Fatalf("failed to delete file %s", path)
		}
	}
}
