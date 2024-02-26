package caravan

import "errors"

var (
	ErrQueueEmpty = errors.New("queue has no elements")
)

type Queue[T any] struct {
	l *List[T]
}

func NewQueue[T any]() *Queue[T] {
	return new(Queue[T]).Init()
}

func (q *Queue[T]) Init() *Queue[T] {
	q.l = NewList[T]()
	return q
}

func (q *Queue[T]) Len() int {
	return q.l.Len()
}

func (q *Queue[T]) Enqueue(v T) {
	q.l.PushBack(v)
}

func (q *Queue[T]) Dequeue() (T, error) {
	var v T
	front := q.l.Front()

	if front == nil {
		return v, ErrQueueEmpty
	}

	v = front.Value
	q.l.Remove(front)

	return v, nil
}
