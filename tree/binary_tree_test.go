package tree

import (
	"testing"
)

type testTree[T any] struct {
	name        string
	inputRest   []T
	wantTreeVal []T
}

func TestNewTree(t *testing.T) {
	tests := []testTree[int]{
		{
			name: "Test new tree int values (1 node)",
			inputRest: []int{
				10,
			},
			wantTreeVal: []int{
				10,
			},
		},
		{
			name: "Test new tree int values (2 nodes incremental)",
			inputRest: []int{
				10,
				20,
			},
			wantTreeVal: []int{
				10,
				20,
			},
		},
		{
			name: "Test new tree int values (2 nodes decremental)",
			inputRest: []int{
				20,
				10,
			},
			wantTreeVal: []int{
				10,
				20,
			},
		},
		{
			name: "Test new tree from several zero int",
			inputRest: []int{
				0,
				0,
			},
			wantTreeVal: []int{
				0,
				0,
			},
		},
		{
			name:        "Test new dequeue list from empty list",
			inputRest:   []int{},
			wantTreeVal: []int{},
		},
		{
			name: "Test new tree int values many nodes",
			inputRest: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			wantTreeVal: []int{
				1, 2, 3, 5, 6, 9, 10, 12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := NewBinarySearchTree[int](tt.inputRest)

			got := root.InorderTraversal()
			for i, wantVal := range tt.wantTreeVal {
				if wantVal != got[i] {
					t.Errorf("actual = %v, want %v", got[i], wantVal)
				}
			}
		})
	}
}
