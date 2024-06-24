package caravan

import (
	"container/heap"
	"errors"
)

type (
	PriorityQueue[T any] struct {
		heap    *pqHeap[T]
		reverse bool
	}

	pqHeapItem[T any] struct {
		reverse  bool
		value    T
		index    int
		priority int
	}

	pqHeap[T any] []*pqHeapItem[T]
)

var (
	ErrPriorityQueueEmpty = errors.New("priority queue has no elements")
)

func NewPriorityQueue[T any](reverse bool) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		reverse: reverse,
		heap:    &pqHeap[T]{},
	}

	return pq
}

func (pq *PriorityQueue[T]) Enqueue(value T, priority int) {
	heap.Push(pq.heap, &pqHeapItem[T]{
		reverse:  pq.reverse,
		value:    value,
		priority: priority,
	})
}

func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	if pq.heap.Len() <= 0 {
		var v T
		return v, ErrPriorityQueueEmpty
	}

	item := heap.Pop(pq.heap).(*pqHeapItem[T])
	return item.value, nil
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.heap.Len()
}

func (pqh pqHeap[T]) Len() int {
	return len(pqh)
}

func (pqh pqHeap[T]) Less(i, j int) bool {
	if pqh[i].reverse {
		return pqh[i].priority < pqh[j].priority
	}

	return pqh[i].priority > pqh[j].priority
}

func (pqh pqHeap[T]) Swap(i, j int) {
	pqh[i], pqh[j] = pqh[j], pqh[i]
	pqh[i].index = i
	pqh[j].index = j
}

func (pqh *pqHeap[T]) Push(x any) {
	n := len(*pqh)
	item := x.(*pqHeapItem[T])
	item.index = n
	*pqh = append(*pqh, item)
}

func (pqh *pqHeap[T]) Pop() any {
	old := *pqh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pqh = old[0 : n-1]

	return item
}