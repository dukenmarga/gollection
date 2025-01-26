package tree

import (
	"cmp"
	"testing"
)

type testAVLT[K cmp.Ordered, V any] struct {
	name        string
	inputKeys   []K
	inputVals   []V
	wantTreeVal []V
}

func TestNewAVLTreeFromArray(t *testing.T) {
	// Test case is taken from:
	// https://stackoverflow.com/questions/3955680/how-to-check-if-my-avl-tree-implementation-is-correct
	tests := []testAVLT[int, int]{
		{
			name: "Test new tree: insert 15 for tree 1 (1A)",
			inputKeys: []int{
				20,
				4,
				15,
			},
			inputVals: []int{
				20,
				4,
				15,
			},
			wantTreeVal: []int{
				15,
				4,
				20,
			},
		},
		{
			name: "Test new tree: insert 15 for tree 2 (2A)",
			inputKeys: []int{
				20,
				4,
				26,
				3,
				9,
				15,
			},
			inputVals: []int{
				20,
				4,
				26,
				3,
				9,
				15,
			},
			wantTreeVal: []int{
				9,
				4,
				20,
				3,
				15,
				26,
			},
		},
		{
			name: "Test new tree: insert 15 for tree 3 (3A)",
			inputKeys: []int{
				20,
				4,
				26,
				3,
				9,
				21,
				30,
				2,
				7,
				11,
				15,
			},
			inputVals: []int{
				20,
				4,
				26,
				3,
				9,
				21,
				30,
				2,
				7,
				11,
				15,
			},
			wantTreeVal: []int{
				9,
				4,
				20,
				3,
				7,
				11,
				26,
				2,
				15,
				21,
				30,
			},
		},
		{
			name: "Test new tree: insert 8 for tree 1 (1B)",
			inputKeys: []int{
				20,
				4,
				8,
			},
			inputVals: []int{
				20,
				4,
				8,
			},
			wantTreeVal: []int{
				8,
				4,
				20,
			},
		},
		{
			name: "Test new tree: insert 8 for tree 2 (2B)",
			inputKeys: []int{
				20,
				4,
				26,
				3,
				9,
				8,
			},
			inputVals: []int{
				20,
				4,
				26,
				3,
				9,
				8,
			},
			wantTreeVal: []int{
				9,
				4,
				20,
				3,
				8,
				26,
			},
		},
		{
			name: "Test new tree: insert 8 for tree 3 (3B)",
			inputKeys: []int{
				20,
				4,
				26,
				3,
				9,
				21,
				30,
				2,
				7,
				11,
				8,
			},
			inputVals: []int{
				20,
				4,
				26,
				3,
				9,
				21,
				30,
				2,
				7,
				11,
				8,
			},
			wantTreeVal: []int{
				9,
				4,
				20,
				3,
				7,
				11,
				26,
				2,
				8,
				21,
				30,
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
			name:        "Test new tree from empty list",
			inputKeys:   []int{},
			wantTreeVal: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := NewAVLTArray[int](tt.inputKeys, tt.inputVals)

			got := root.LevelOrderTraversal()
			for i, wantVal := range tt.wantTreeVal {
				if wantVal != got[i].value {
					t.Errorf("actual = %v, want %v", got[i].value, wantVal)
				}
			}
		})
	}
}

func TestNewAVLTreeRoot(t *testing.T) {
	// These test below is only checking the creation of the
	// root. Actual test to check balancing see TestAVLTreeFromArray.
	tests := []testAVLT[int, int]{
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
			root := NewAVLTRoot[int](tt.inputKeys[0], tt.inputVals[0])

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

type testAVLTSearching[K cmp.Ordered, V any] struct {
	name           string
	inputKeys      []K
	inputVals      []V
	inputSearchVal V
	wantError      bool
	wantSearchVal  V
}

func TestAVLTreeSearching(t *testing.T) {
	tests := []testAVLTSearching[int, int]{
		{
			name: "Test search tree: search from several items",
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
			name: "Test search tree: search but item not found",
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
			root := NewAVLTArray[int](tt.inputKeys, tt.inputVals)
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

type testAVLTreeDelete[K cmp.Ordered, V any] struct {
	name           string
	inputKeys      []K
	inputVals      []V
	inputDeleteKey K
	wantError      bool
	wantTreeVal    []V
}

func TestAVLTreeDeleteNode(t *testing.T) {
	// Test case is taken from:
	// https://stackoverflow.com/questions/3955680/how-to-check-if-my-avl-tree-implementation-is-correct
	tests := []testAVLTreeDelete[int, int]{
		{
			name: "Test delete tree: delete item 1 from tree",
			inputKeys: []int{
				2, 1, 4, 3, 5,
			},
			inputVals: []int{
				2, 1, 4, 3, 5,
			},
			inputDeleteKey: 1,
			wantError:      false,
			wantTreeVal: []int{
				4, 2, 5, 3,
			},
		},
		{
			name: "Test delete tree: delete item 1 from tree",
			inputKeys: []int{
				6, 2, 9, 1, 4, 8, 11, 3, 5, 7, 10, 12, 13,
			},
			inputVals: []int{
				6, 2, 9, 1, 4, 8, 11, 3, 5, 7, 10, 12, 13,
			},
			inputDeleteKey: 1,
			wantError:      false,
			wantTreeVal: []int{
				6, 4, 9, 2, 5, 8, 11, 3, 7, 10, 12, 13,
			},
		},
		{
			name: "Test delete tree: delete item 1 from tree",
			inputKeys: []int{
				5, 2, 8, 1, 3, 7, 10, 4, 6, 9, 11, 12,
			},
			inputVals: []int{
				5, 2, 8, 1, 3, 7, 10, 4, 6, 9, 11, 12,
			},
			inputDeleteKey: 1,
			wantError:      false,
			wantTreeVal: []int{
				8, 5, 10, 3, 7, 9, 11, 2, 4, 6, 12,
			},
		},
		{
			name: "Test delete tree: delete but item not found",
			inputKeys: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			inputVals: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
			inputDeleteKey: 99,
			wantError:      true,
			wantTreeVal: []int{
				5, 6, 2, 10, 12, 3, 1, 9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add root
			root := NewAVLTArray[int](tt.inputKeys, tt.inputVals)
			err := root.Delete(tt.inputDeleteKey)

			if tt.wantError {
				if (err != nil) != tt.wantError {
					t.Errorf("actual = %v, want %v", (err != nil) == tt.wantError, tt.wantError)
				}
			} else {
				got := root.LevelOrderTraversal()
				for i, wantVal := range tt.wantTreeVal {
					if wantVal != got[i].value {
						t.Errorf("actual = %v, want %v", got[i].value, wantVal)
					}
				}
			}
		})
	}
}
