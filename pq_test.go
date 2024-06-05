package caravan_test

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/maladroitthief/caravan"
)

func Test_pq_Dequeue(t *testing.T) {
	type item struct {
		value    string
		priority float64
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
			name:    "reverse multiple insert",
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
			pq := caravan.NewPQ[string](tt.reverse)
			for _, item := range tt.items {
				pq.Enqueue(item.value, item.priority)
			}

			for _, want := range tt.wants {
				got, err := pq.Dequeue()
				if err != nil && !errors.Is(err, tt.err) {
					t.Errorf("pq.Dequeue() error: want %+v, got: %+v\n", tt.err, err)
				}

				if want != got {
					t.Errorf("pq.Dequeue(): want %+v, got: %+v\n", want, got)
				}
			}
		})
	}
}

func BenchmarkPQEnqueue(b *testing.B) {
	type entity struct {
		id int
	}

	pq := caravan.NewPQ[entity](false)
	for n := 0; n < b.N; n++ {
		pq.Enqueue(entity{id: rand.Int()}, rand.Float64())
	}
}

func BenchmarkPQDequeue(b *testing.B) {
	type entity struct {
		id int
	}

	pq := caravan.NewPQ[entity](false)
	for i := 0; i < ContainerSize; i++ {
		pq.Enqueue(entity{id: rand.Int()}, rand.Float64())
	}

	for n := 0; n < b.N; n++ {
		if n < ContainerSize {
			_, err := pq.Dequeue()
			if err != nil {
				b.Errorf("PQ.Dequeue(false) error: %v", err)
			}
		}
	}
}

func BenchmarkPQReverseEnqueue(b *testing.B) {
	type entity struct {
		id int
	}

	pq := caravan.NewPQ[entity](true)
	for n := 0; n < b.N; n++ {
		pq.Enqueue(entity{id: rand.Int()}, rand.Float64())
	}
}

func BenchmarkPQReverseDequeue(b *testing.B) {
	type entity struct {
		id int
	}

	pq := caravan.NewPQ[entity](true)
	for i := 0; i < ContainerSize; i++ {
		pq.Enqueue(entity{id: rand.Int()}, rand.Float64())
	}

	for n := 0; n < b.N; n++ {
		if n < ContainerSize {
			_, err := pq.Dequeue()
			if err != nil {
				b.Errorf("PQ.Dequeue(false) error: %v", err)
			}
		}
	}
}
