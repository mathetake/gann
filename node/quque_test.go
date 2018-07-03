package node

import (
	"container/heap"
	"testing"

	"math"

	"github.com/bmizerany/assert"
)

func TestPriorityQueue(t *testing.T) {
	qItems := map[string]float32{
		"a": 1,
		"b": float32(math.Inf(0)),
		"c": 3,
		"d": 100,
	}

	pq := make(PriorityQueue, 4)

	var i = 0
	for v, priprity := range qItems {
		pq[i] = &QueueItem{
			Value:    v,
			Priority: priprity,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	q1 := heap.Pop(&pq).(*QueueItem)
	assert.Equal(t, pq.Len(), 3)
	assert.Equal(t, q1.Priority, float32(math.Inf(0)))
	assert.Equal(t, q1.Value, "b")

	q2 := heap.Pop(&pq).(*QueueItem)
	assert.Equal(t, pq.Len(), 2)
	assert.Equal(t, q2.Priority, float32(100))
	assert.Equal(t, q2.Value, "d")

	q3 := heap.Pop(&pq).(*QueueItem)
	assert.Equal(t, pq.Len(), 1)
	assert.Equal(t, q3.Priority, float32(3))
	assert.Equal(t, q3.Value, "c")

	q4 := heap.Pop(&pq).(*QueueItem)
	assert.Equal(t, pq.Len(), 0)
	assert.Equal(t, q4.Priority, float32(1))
	assert.Equal(t, q4.Value, "a")
}
