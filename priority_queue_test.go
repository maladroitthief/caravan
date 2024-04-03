package caravan_test

import (
	"testing"

	"github.com/maladroitthief/caravan"
)

func Test_priority_queue_Dequeue(t *testing.T) {
	type item struct {
		value    string
		priority int
	}
	tests := []struct {
		name  string
		items []item
		wants []string
	}{
		{
			name: "single insert",
			items: []item{
				{value: "test", priority: 1},
			},
			wants: []string{
				"test",
			},
		},
		{
			name: "multiple insert",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := caravan.NewPriorityQueue[string]()
			for _, item := range tt.items {
				pq.Enqueue(item.value, item.priority)
			}

			for _, want := range tt.wants {
				got := pq.Dequeue()
				if want != got {
					t.Errorf("PriorityQueue.Dequeue(): want %+v, got: %+v\n", want, got)
				}
			}
		})
	}
}
