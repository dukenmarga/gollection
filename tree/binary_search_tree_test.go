package tree

import (
	"cmp"
	"testing"
)

type testTree[K cmp.Ordered, V any] struct {
	name        string
	inputKeys   []K
	inputVals   []V
	wantTreeVal []V
}

func TestNewBinarySearchTreeFromArray(t *testing.T) {
	tests := []testTree[int, int]{
		{
			name: "Test new tree int values (1 node)",
			inputKeys: []int{
				10,
			},
			inputVals: []int{
				10,
			},
			wantTreeVal: []int{
				10,
			},
		},
		{
			name: "Test new tree int values (2 nodes incremental)",
			inputKeys: []int{
				10,
				20,
			},
			inputVals: []int{
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
			inputKeys: []int{
				20,
				10,
			},
			inputVals: []int{
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
			inputKeys: []int{
				0,
				0,
			},
			inputVals: []int{
				0,
				0,
			},
			wantTreeVal: []int{
				0,
			},
		},
		{
			name:        "Test new dequeue list from empty list",
			inputKeys:   []int{},
			wantTreeVal: []int{},
		},
		{
			name: "Test new tree int values many nodes",
			inputKeys: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			inputVals: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			wantTreeVal: []int{
				1, 2, 3, 5, 6, 9, 10, 12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := NewBSTArray[int](tt.inputKeys, tt.inputVals)

			got := root.InorderTraversal()
			for i, wantVal := range tt.wantTreeVal {
				if wantVal != got[i].value {
					t.Errorf("actual = %v, want %v", got[i], wantVal)
				}
			}
		})
	}
}

func TestNewBinarySearchTreeRoot(t *testing.T) {
	tests := []testTree[int, int]{
		{
			name: "Test new tree int values (1 node)",
			inputKeys: []int{
				10,
			},
			inputVals: []int{
				10,
			},
			wantTreeVal: []int{
				10,
			},
		},
		{
			name: "Test new tree int values (2 nodes incremental)",
			inputKeys: []int{
				10,
				20,
			},
			inputVals: []int{
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
			inputKeys: []int{
				20,
				10,
			},
			inputVals: []int{
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
			inputKeys: []int{
				0,
				0,
			},
			inputVals: []int{
				0,
				0,
			},
			wantTreeVal: []int{
				0,
			},
		},
		{
			name: "Test new tree int values many nodes",
			inputKeys: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			inputVals: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			wantTreeVal: []int{
				1, 2, 3, 5, 6, 9, 10, 12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add root
			root := NewBSTRoot[int](tt.inputKeys[0], tt.inputVals[0])

			// Add the remaining nodes
			for i := 1; i < len(tt.inputKeys); i++ {
				root.Add(tt.inputKeys[i], tt.inputVals[i])
			}

			got := root.InorderTraversal()
			for i, wantVal := range tt.wantTreeVal {
				if wantVal != got[i].value {
					t.Errorf("actual = %v, want %v", got[i], wantVal)
				}
			}
		})
	}
}

type testBSTSearching[K cmp.Ordered, V any] struct {
	name           string
	inputKeys      []K
	inputVals      []V
	inputSearchVal V
	wantError      bool
	wantSearchVal  V
}

func TestSearchBinaryTreeSearching(t *testing.T) {
	tests := []testBSTSearching[int, int]{
		{
			name: "Test search binary tree: search from several items",
			inputKeys: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			inputVals: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			inputSearchVal: 1,
			wantError:      false,
			wantSearchVal:  1,
		},
		{
			name: "Test search binary tree: search but item not found",
			inputKeys: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			inputVals: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			inputSearchVal: 99,
			wantError:      true,
			wantSearchVal:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add root
			root := NewBSTArray[int](tt.inputKeys, tt.inputVals)
			got, err := root.Search(tt.inputSearchVal)

			if tt.wantError {
				if err == nil {
					t.Errorf("actual = %v, want %v", err == nil, tt.wantError)
				}
			} else {
				if got.value != tt.wantSearchVal {
					t.Errorf("actual = %v, want %v", got.value, tt.wantSearchVal)
				}
			}

		})
	}
}

type testTreeDelete[K cmp.Ordered, V any] struct {
	name           string
	inputKeys      []K
	inputVals      []V
	inputDeleteKey K
	wantError      bool
	wantTreeVal    []V
}

func TestBinarySearchTreeDeleteNode(t *testing.T) {
	tests := []testTreeDelete[int, int]{
		{
			name: "Test delete binary tree: delete from several items",
			inputKeys: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			inputVals: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			inputDeleteKey: 6,
			wantError:      false,
			wantTreeVal: []int{
				1, 2, 3, 5, 9, 10, 12,
			},
		},
		{
			name: "Test delete binary tree: delete but item not found",
			inputKeys: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			inputVals: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			inputDeleteKey: 99,
			wantError:      true,
			wantTreeVal: []int{
				1, 2, 3, 5, 6, 9, 10, 12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add root
			root := NewBSTArray[int](tt.inputKeys, tt.inputVals)
			err := root.Delete(tt.inputDeleteKey)

			if tt.wantError {
				if err == nil {
					t.Errorf("actual = %v, want %v", err == nil, tt.wantError)
				}
			} else {
				got := root.InorderTraversal()
				for i, wantVal := range tt.wantTreeVal {
					if wantVal != got[i].value {
						t.Errorf("actual = %v, want %v", got[i].value, wantVal)
					}
				}
			}

		})
	}
}
