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

// Index ... an interface for gann's index in `index` package (only used for interface declaration on its methods)
type Index interface {
	// Build ... build gann's index
	Build() error

	// GetANNbyItemID ... search ANNs by a given itemID
	GetANNbyItemID(id int64, num int, bucketScale float32) (ann []int64, err error)

	// GetANNbyVector ... search ANNs by a given query vector
	GetANNbyVector(v []float32, num int, bucketScale float32) (ann []int64, err error)

	// Load ... load index from disk
	Load(path string) error

	// Save ... save index to disk
	Save(path string) error
}

var _ Index = &index.Index{}
