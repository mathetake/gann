package index

import "github.com/mathetake/gann/item"

func (idx *Index) build() {}


func (idx *Index) getANNbyItem (id item.ID, num int, searchBucket int) (ann []int32, err error){
	return ann, nil
}

func (idx *Index) getANNbyVector (v []float32, num int, searchBucket int) (ann []int32, err error){
	return ann, nil
}