# gann
[![CircleCI](https://circleci.com/gh/mathetake/gann.svg?style=shield&circle-token=9a6608c5baa7a400661a700127778a9ff8baeee3)](https://circleci.com/gh/mathetake/gann)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

<img width="600" alt="portfolio_view" src="https://mathetake.github.io/blogs/assets/gann/recursive_build.png">

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
	"math/rand"
	"time"
	
	"github.com/mathetake/gann/index"
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
