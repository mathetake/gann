# gann
[![CircleCI](https://circleci.com/gh/mathetake/gann.svg?style=shield&circle-token=9a6608c5baa7a400661a700127778a9ff8baeee3)](https://circleci.com/gh/mathetake/gann)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

<img width="600" alt="portfolio_view" src="https://mathetake.github.io/blogs/assets/gann/recursive_build.png">

gann (go-approximate-nearest-neighbor) is a library for approximate nearest neighbor search purely written in golang.

The implemented algorithm is truly inspired by Annoy (https://github.com/spotify/annoy).

## feature
1. purely written in Go: no dependencies out of Go world.
2. easy to tune with a bit of parameters

## usage

```golang
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mathetake/gann"
	"github.com/mathetake/gann/metric"
)

var (
	dim    = 3
	nTrees = 2
	k      = 10
	nItem  = 1000
)

func main() {
	rawItems := make([][]float64, 0, nItem)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < nItem; i++ {
		item := make([]float64, 0, dim)
		for j := 0; j < dim; j++ {
			item = append(item, rand.Float64())
		}
		rawItems = append(rawItems, item)
	}

	m, err := metric.NewCosineMetric(dim)
	if err != nil {
		// err handling
		return
	}

	// create index
	idx, err := gann.CreateNewIndex(rawItems, dim, nTrees, k, m)
	if err != nil {
		// error handling
		return
	}

	// search
	q := []float64{0.1, 0.02, 0.001}
	res, err := idx.GetANNbyVector(q, 5, 10)
	if err != nil {
		// error handling
		return
	}

	fmt.Printf("res: %v\n", res)
}
```

## parameters

See the blog post describing the parameters and algorithms in _gann_  :

https://mathetake.github.io/blogs/gann.html

## references

- https://github.com/spotify/annoy
- https://en.wikipedia.org/wiki/Nearest_neighbor_search#Approximate_nearest_neighbor

## License

MIT
