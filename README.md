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
	var searchNum = 5
	var bucketScale float64 = 10
	q := []float64{0.1, 0.02, 0.001}
	res, err := idx.GetANNbyVector(q, searchNum, bucketScale)
	if err != nil {
		// error handling
		return
	}

	fmt.Printf("res: %v\n", res)
}
```

## parameters

### setup phase parameters

|name|type|description|run-time computational complexity|space complexity|accuracy|
|:---:|:---:|:---:|:---:|:---:|:---:|
|dim|int| dimension of target vectors| the larger, the more expensive | the larger, the more expensive |  N/A |
|nTree|int| # of trees|the larger, the more expensive| the larger, the more expensive | the larger, the more accurate|
|k|int|maximum # of items in a single leaf|the larger, the less expensive| N/A| the larger, the less accurate|

### runtime (search phase) parameters

|name|type|description|computational complexity|accuracy|
|:---:|:---:|:---:|:---:|:---:|
|searchNum|int| # of requested neighbors|the larger, the more expensive|N/A|
|bucketScale|float64| affects the size of `bucket` |the larger, the more expensive|the larger, the more accurate|

`bucketScale` affects the size of `bucket` which consists of items for exact distance calculation. 
The actual size of the bucket is [calculated by](https://github.com/mathetake/gann/blob/357c3abd241bd6455e895a5b392251b06507a8e8/search.go#L30) `int(searchNum * bucketScale)`.

In the search phase, we traverse index trees and continuously put items on reached leaves to the bucket [until the bucket becomes full](https://github.com/mathetake/gann/blob/357c3abd241bd6455e895a5b392251b06507a8e8/search.go#L48).
Then we [calculate the exact distances between a item in the bucket and the query vector](https://github.com/mathetake/gann/blob/357c3abd241bd6455e895a5b392251b06507a8e8/search.go#L74-L81) to get approximate nearest neighbors.

Therefore, the larger `bucketScale` the more computational complexity while the more accurate result to be produced.

## references

- https://github.com/spotify/annoy
- https://en.wikipedia.org/wiki/Nearest_neighbor_search#Approximate_nearest_neighbor

## License

MIT
