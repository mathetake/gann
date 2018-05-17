/*

MIT License

Copyright (c) 2018 @mathetake

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
package gann

import (
	"github.com/mathetake/gann/index"
)

type GannIndex interface {
	Build() error // build search trees.
	GetANNbyItemID(id int64, num int, bucketScale float32) (ann []int64, err error)
	GetANNbyVector(v []float32, num int, bucketScale float32) (ann []int64, err error)
}

// GetIndex ... get index (composed of trees, nodes, etc.)
func GetIndex(items [][]float32, d int, nT int, k int, normalize bool) (GannIndex, error) {
	return index.Initialize(items, d, nT, k, normalize)
}
