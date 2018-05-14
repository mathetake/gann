# gann

gann (go-approximate-nearest-neighbor) is a library for approximate nearest neighbor search written in golang.

The implemented algorithm is truly inspired by Annoy (https://github.com/spotify/annoy).

# usage

```golang
import "github.com/mathetake/gann"
    
func main () {
	rawItems := [][]float32{
		{0.1, 0.1 ,0.1},
		{0.2, 0.2 ,0.2},
		{0.3, 0.3 ,0.3},
		{0.4, 0.4 ,0.4},
		{0.5, 0.5 ,0.5},
	}
	
	// create index
	gIDx := gann.GetIndex(rawItems, 3, 1)
	
	// do search
	q := []float32{0.1, 0.02, 0.001}
	ann, err := gIDx.getANNbyItem(q, 1, 10)
}
```
