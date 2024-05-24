package caravan_test

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/maladroitthief/caravan"
)

func Test_priority_queue_Dequeue(t *testing.T) {
	type item struct {
		value    string
		priority int
	}
	tests := []struct {
		name    string
		reverse bool
		items   []item
		err     error
		wants   []string
	}{
		{
			name:    "single insert",
			reverse: false,
			items: []item{
				{value: "test", priority: 1},
			},
			err: nil,
			wants: []string{
				"test",
			},
		},
		{
			name:    "multiple insert",
			reverse: false,
			items: []item{
				{value: "forth", priority: 4},
				{value: "sixth", priority: 6},
				{value: "fifth", priority: 5},
				{value: "seventh", priority: 7},
				{value: "tenth", priority: 10},
				{value: "second", priority: 2},
				{value: "third", priority: 3},
				{value: "eighth", priority: 8},
				{value: "ninth", priority: 9},
				{value: "first", priority: 1},
			},
			err: nil,
			wants: []string{
				"tenth",
				"ninth",
				"eighth",
				"seventh",
				"sixth",
				"fifth",
				"forth",
				"third",
				"second",
				"first",
			},
		},
		{
			name:    "no elements",
			reverse: false,
			items:   []item{},
			err:     caravan.ErrPriorityQueueEmpty,
			wants: []string{
				"",
			},
		},
		{
			name:    "multiple insert",
			reverse: true,
			items: []item{
				{value: "forth", priority: 4},
				{value: "sixth", priority: 6},
				{value: "fifth", priority: 5},
				{value: "seventh", priority: 7},
				{value: "tenth", priority: 10},
				{value: "second", priority: 2},
				{value: "third", priority: 3},
				{value: "eighth", priority: 8},
				{value: "ninth", priority: 9},
				{value: "first", priority: 1},
			},
			err: nil,
			wants: []string{
				"first",
				"second",
				"third",
				"forth",
				"fifth",
				"sixth",
				"seventh",
				"eighth",
				"ninth",
				"tenth",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := caravan.NewPriorityQueue[string](tt.reverse)
			for _, item := range tt.items {
				pq.Enqueue(item.value, item.priority)
			}

			for _, want := range tt.wants {
				got, err := pq.Dequeue()
				if err != nil && !errors.Is(err, tt.err) {
					t.Errorf("PriorityQueue.Dequeue() error: want %+v, got: %+v\n", tt.err, err)
				}

				if want != got {
					t.Errorf("PriorityQueue.Dequeue(): want %+v, got: %+v\n", want, got)
				}
			}
		})
	}
}

func BenchmarkPriorityQueueEnqueue(b *testing.B) {
	type entity struct {
		id int
	}

	pq := caravan.NewPriorityQueue[entity](false)
	for n := 0; n < b.N; n++ {
		pq.Enqueue(entity{id: rand.Int()}, rand.Int())
	}
}

func BenchmarkReversePriorityQueueEnqueue(b *testing.B) {
	type entity struct {
		id int
	}

	pq := caravan.NewPriorityQueue[entity](true)
	for n := 0; n < b.N; n++ {
		pq.Enqueue(entity{id: rand.Int()}, rand.Int())
	}
}

func BenchmarkPriorityQueueDequeue(b *testing.B) {
	type entity struct {
		id int
	}

	pq := caravan.NewPriorityQueue[entity](false)
	for i := 0; i < ContainerSize; i++ {
		pq.Enqueue(entity{id: rand.Int()}, rand.Int())
	}

	for n := 0; n < b.N; n++ {
		if n < ContainerSize {
			_, err := pq.Dequeue()
			if err != nil {
				b.Errorf("PriorityQueue.Dequeue(false) error: %v", err)
			}
		}
	}
}

func BenchmarkReversePriorityQueueDequeue(b *testing.B) {
	type entity struct {
		id int
	}

	pq := caravan.NewPriorityQueue[entity](true)
	for i := 0; i < ContainerSize; i++ {
		pq.Enqueue(entity{id: rand.Int()}, rand.Int())
	}

	for n := 0; n < b.N; n++ {
		if n < ContainerSize {
			_, err := pq.Dequeue()
			if err != nil {
				b.Errorf("PriorityQueue.Dequeue(false) error: %v", err)
			}
		}
	}
}
