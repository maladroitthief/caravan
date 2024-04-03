package caravan

import "container/heap"

type (
	PriorityQueue[T any] struct {
		heap *pqHeap[T]
	}

	pqHeapItem[T any] struct {
		value    T
		index    int
		priority int
	}

	pqHeap[T any] []*pqHeapItem[T]
)

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{
		heap: &pqHeap[T]{},
	}

	return pq
}

func (pq *PriorityQueue[T]) Enqueue(value T, priority int) {
	heap.Push(pq.heap, &pqHeapItem[T]{
		value:    value,
		priority: priority,
	})
}

func (pq *PriorityQueue[T]) Dequeue() T {
	item := heap.Pop(pq.heap).(*pqHeapItem[T])

	return item.value
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.heap.Len()
}

func (pqh pqHeap[T]) Len() int {
	return len(pqh)
}

func (pqh pqHeap[T]) Less(i, j int) bool {
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
