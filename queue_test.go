package gann

import (
	"container/heap"
	"fmt"
	"math"
	"testing"

	"github.com/bmizerany/assert"
)

func TestPriorityQueue(t *testing.T) {
	for i, c := range []struct {
		valueToPriority map[nodeId]float64
		expValues       []nodeId
	}{
		{
			valueToPriority: map[nodeId]float64{
				"a": 1,
				"b": math.Inf(0),
				"c": 3,
				"d": 100,
			},
		},
	} {
		c := c
		t.Run(fmt.Sprintf("%d-th case", i), func(t *testing.T) {
			var i int
			pq := make(priorityQueue, len(c.valueToPriority))
			for v, pr := range c.valueToPriority {
				pq[i] = &queueItem{
					value:    v,
					priority: pr,
					index:    i,
				}
				i++
			}
			heap.Init(&pq)

			for _, v := range c.expValues {
				qi := heap.Pop(&pq).(queueItem)
				assert.Equal(t, qi.value, v)
			}
		})
	}
}
