# gann
[![CircleCI](https://circleci.com/gh/mathetake/gann.svg?style=shield&circle-token=9a6608c5baa7a400661a700127778a9ff8baeee3)](https://circleci.com/gh/mathetake/gann)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

gann (go-approximate-nearest-neighbor) is a library for approximate nearest neighbor search purely written in golang.

The implemented algorithm is truly inspired by Annoy (https://github.com/spotify/annoy).

# feature
1. __ONLY__ written in golang, no dependencies out of go world.
2. easy to tune with a bit of parameters
3. __ONLY support for cosine similarity search.__ (issue: https://github.com/mathetake/gann/issues/12)

# usage

```golang
import (
	"fmt"
	"github.com/mathetake/gann/index"
	"math/rand"
	"time"
)

func main() {
	var dim = 3
	var nTrees = 2
	var k = 10
	var nItem = 1000

	rawItems := make([][]float32, 0, nItem)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nItem; i++ {
		item := make([]float32, 0, dim)
		for j := 0; j < dim; j++ {
			item = append(item, rand.Float32())
		}
		rawItems = append(rawItems, item)
	}

	// create index
	gIDx := index.GetIndex(rawItems, dim, nTrees, k, true)
	gIDx.Build()

	// do search
	q := []float32{0.1, 0.02, 0.001}
	ann, _ := gIDx.GetANNbyVector(q, 5, 10)
	fmt.Println("result:", ann)
}
```


You can also save and load your index to/from disk:

```golang
gIDx := index.GetIndex(rawItems, dim, nTrees, k, true)
gIDx.Build()

var path = "foo.gann"

err := gIDx.Save(path)
if err != nil {
    panic(err)
}

var idx = &index.Index{}
err := idx.Load(path)
if err != nil {
    panic(err)
}
```

# parameters

See the blog post describing the parameters and algorithms in _gann_  :

https://mathetake.github.io/blogs/gann.html

# references

- https://github.com/spotify/annoy
- https://en.wikipedia.org/wiki/Nearest_neighbor_search#Approximate_nearest_neighbor

# License

MIT
