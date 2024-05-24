package caravan

type (
	PQ[T any] struct {
		heap pqh[T]
	}

	pqhi[T any] struct {
		value    T
		index    int
		priority int
	}

	pqh[T any] []pqhi[T]
)

func NewPQ[T any]() *PQ[T] {
	pq := &PQ[T]{
		heap: pqh[T]{},
	}

	return pq
}

func (pq *PQ[T]) Enqueue(value T, priority int) {
	pq.heap = pq.heap.Push(pqhi[T]{
		value:    value,
		priority: priority,
	})
}

func (pq *PQ[T]) Dequeue() (T, error) {
	if pq.heap.Len() <= 0 {
		var v T
		return v, ErrPriorityQueueEmpty
	}
	var value T
	value, pq.heap = pq.heap.Pop()

	return value, nil
}

func (pqh pqh[T]) Len() int {
	return len(pqh)
}

func (pqh pqh[T]) Less(i, j int) bool {
	return pqh[i].priority > pqh[j].priority
}

func (pqh pqh[T]) Swap(i, j int) pqh[T] {
	pqh[i], pqh[j] = pqh[j], pqh[i]
	pqh[i].index = i
	pqh[j].index = j

	return pqh
}

func (pqh pqh[T]) Push(item pqhi[T]) pqh[T] {
	item.index = len(pqh)
	pqh = append(pqh, item)
	return pqh.up(pqh.Len() - 1)
}

func (pqh pqh[T]) Pop() (T, pqh[T]) {
	n := pqh.Len() - 1
	pqh = pqh.Swap(0, n)
	_, pqh = pqh.down(0, n)

	old := pqh
	n = len(old)
	item := old[n-1]
	// old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	pqh = old[0 : n-1]

	return item.value, pqh
}

func (pqh pqh[T]) heapify() pqh[T] {
	n := pqh.Len()
	for i := n/2 - 1; i >= 0; i-- {
		_, pqh = pqh.down(i, n)
	}

	return pqh
}

func (pqh pqh[T]) fix(i int) pqh[T] {
	down, pqh := pqh.down(i, pqh.Len())
	if !down {
		pqh = pqh.up(i)
	}

	return pqh
}

func (pqh pqh[T]) up(j int) pqh[T] {
	for {
		i := (j - 1) / 2
		if i == j || !pqh.Less(j, i) {
			break
		}
		pqh = pqh.Swap(i, j)
		j = i
	}

	return pqh
}

func (pqh pqh[T]) down(i0, n int) (bool, pqh[T]) {
	i := i0

	for {
		j1 := 2*i + 1

		if j1 >= n || j1 < 0 {
			break
		}

		j := j1
		j2 := j1 + 1
		if j2 < n && pqh.Less(j2, j1) {
			j = j2
		}

		if !pqh.Less(j, i) {
			break
		}

		pqh = pqh.Swap(i, j)
		i = j
	}

	return i > i0, pqh
}
