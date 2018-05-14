package index

import "github.com/mathetake/gann/item"

// Build ... build index forest.
func (idx *Index) Build() {}

// GetANNbyItem ... get ANNs by a item.Item
func (idx *Index) GetANNbyItem (id item.ID, num int, searchBucket int) (ann []int32, err error){
	return ann, nil
}

// GetANNbyVector ... get ANNs by a vector
func (idx *Index) GetANNbyVector (v []float32, num int, searchBucket int) (ann []int32, err error){
	return ann, nil
}