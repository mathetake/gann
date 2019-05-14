// Package gann can be used for approximate nearest neighbor search.
//
// By calling gann.CreateNewIndex function, we can obtain a search index.
// Its interface is defined in gann.Index:
//
//	type Index interface {
// 		GetANNbyItemID(id int64, searchNum int, bucketScale float64) (ann []int64, err error)
//		GetANNbyVector(v []float64, searchNum int, bucketScale float64) (ann []int64, err error)
//	}
//
// GetANNbyItemID allows us to pass id of specific item for search execution
// and instead GetANNbyVector allows us to pass a vector.
//
// See README.md for more details.
package gann
