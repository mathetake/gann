package index

import (
	"github.com/mathetake/gann/item"
	"github.com/mathetake/gann/node"
	"github.com/pkg/errors"
)

// GetANNbyItem ... get ANNs by a item.Item
func (idx *Index) GetANNbyItem(id item.ID, num int, searchBucket int) (ann []int32, err error) {
	it, ok := idx.itemIDToItem[id]
	if !ok {
		errors.Errorf("Item not found for %v", id)
	}
	return idx.getANNbyVector(it.Vec, num, searchBucket)
}

// GetANNbyVector ... get ANNs by a vector
func (idx *Index) GetANNbyVector(v []float32, num int, searchBucket int) (ann []int32, err error) {
	return idx.getANNbyVector(v, num, searchBucket)
}

func (idx *Index) getANNbyVector(v []float32, num int, searchBucket int) (ann []int32, err error) {
	return ann, nil
}

// Build ... build index forest.
func (idx *Index) Build() error {
	err := idx.buildRootNodes()
	if err != nil {
		return errors.Wrapf(err, "buildRootNodes failed.")
	}
	for _, rn := range idx.roots {
		err := rn.Build(idx.items, idx.k)
		if err != nil {
			return errors.Wrapf(err, "Build failed.")
		}
	}
	return nil
}

func (idx *Index) buildRootNodes() error {
	n := idx.getNItems()
	for i := 0; i < idx.nTree; i++ {
		nv, err := item.GetNormalVectorOfSplittingHyperPlane(idx.items)
		if err != nil {
			return errors.Wrapf(err, "GetNormalVectorOfSplittingHyperPlane failed.")
		}
		r := &node.Node{
			ID:           node.ID(i + n*i),
			Vec:          nv,
			NDescendants: len(idx.items),
		}
		idx.roots = append(idx.roots, r)
	}
	return nil
}
