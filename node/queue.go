// Implementation of priority queue for nodes
// https://golang.org/pkg/container/heap/#example__priorityQueue
package node

import (
	"container/heap"
)

type nodeQueueItem struct {
	// ID ... node ID
	ID int

	// The index of the item in the heap.
	index int

	priority float32
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*nodeQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*nodeQueueItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *nodeQueueItem, id int, priority float32) {
	item.ID = id
	item.priority = priority
	heap.Fix(pq, item.index)
}
