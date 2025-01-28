package tree

import (
	"cmp"
	"testing"
)

type testBSTCreation[K cmp.Ordered, V any] struct {
	name        string
	inputKeys   []K
	inputVals   []V
	wantTreeVal []V
}

func TestNewBinarySearchTreeFromArray(t *testing.T) {
	tests := []testBSTCreation[int, int]{
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
	tests := []testBSTCreation[int, int]{
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
				_ = root.Add(tt.inputKeys[i], tt.inputVals[i])
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
			got, err := root.Find(tt.inputSearchVal)

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

type testBSTDelete[K cmp.Ordered, V any] struct {
	name           string
	inputKeys      []K
	inputVals      []V
	inputDeleteKey K
	wantError      bool
	wantTreeVal    []V
}

func TestBinarySearchTreeDeleteNode(t *testing.T) {
	tests := []testBSTDelete[int, int]{
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
		{
			name: "Test delete tree: delete the root node",
			inputKeys: []int{
				4, 3, 1, 2, 6, 5, 7,
			},
			inputVals: []int{
				4, 3, 1, 2, 6, 5, 7,
			},
			inputDeleteKey: 4,
			wantError:      false,
			wantTreeVal: []int{
				1, 2, 3, 5, 6, 7,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add root
			root := NewBSTArray[int](tt.inputKeys, tt.inputVals)
			err := root.Delete(tt.inputDeleteKey)

			got := root.InorderTraversal()
			if len(got) != len(tt.wantTreeVal) {
				t.Errorf("actual length = %v, want length %v", len(got), len(tt.wantTreeVal))
			}
			if tt.wantError {
				if (err != nil) != tt.wantError {
					t.Errorf("actual = %v, want %v", (err != nil) == tt.wantError, tt.wantError)
				}
			}
			for i, wantVal := range tt.wantTreeVal {
				if wantVal != got[i].value {
					t.Errorf("actual = %v, want %v", got[i].value, wantVal)
				}
			}
		})
	}
}

type testBSTUpdate[K cmp.Ordered, V any] struct {
	name        string
	inputKeys   []K
	inputVals   []V
	updateKey   K
	updateVal   V
	wantError   bool
	wantTreeVal []V
}

func TestBSTreeUpdate(t *testing.T) {
	tests := []testBSTUpdate[int, int]{
		{
			name: "Test update 1 node tree ",
			inputKeys: []int{
				2, 1, 4, 3, 5,
			},
			inputVals: []int{
				2, 1, 4, 3, 5,
			},
			updateKey: 1,
			updateVal: 99,
			wantError: false,
			wantTreeVal: []int{
				2, 99, 4, 3, 5,
			},
		},
		{
			name: "Test update the root ",
			inputKeys: []int{
				2, 1, 4, 3, 5,
			},
			inputVals: []int{
				2, 1, 4, 3, 5,
			},
			updateKey: 2,
			updateVal: 99,
			wantError: false,
			wantTreeVal: []int{
				99, 1, 4, 3, 5,
			},
		},
		{
			name: "Test update tree: key not found",
			inputKeys: []int{
				2, 1, 4, 3, 5,
			},
			inputVals: []int{
				2, 1, 4, 3, 5,
			},
			updateKey: 99,
			updateVal: 99,
			wantError: true,
			wantTreeVal: []int{
				2, 1, 4, 3, 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add root
			root := NewBSTArray[int](tt.inputKeys, tt.inputVals)
			err := root.Update(tt.updateKey, tt.updateVal)

			got := root.LevelOrderTraversal()
			if len(got) != len(tt.wantTreeVal) {
				t.Errorf("actual length = %v, want length %v", len(got), len(tt.wantTreeVal))
			}
			if tt.wantError {
				if (err != nil) != tt.wantError {
					t.Errorf("actual = %v, want %v", (err != nil) == tt.wantError, tt.wantError)
				}
			}
			for i, wantTreeVal := range tt.wantTreeVal {
				if got[i].value != wantTreeVal {
					t.Errorf("actual = %v, want %v", got[i].value, wantTreeVal)
				}
			}
		})
	}
}

type testBSTAdd[K, V any] struct {
	name        string
	inputKeys   []K
	inputVals   []V
	wantError   []bool
	wantTreeVal []V
}

func TestBSTreeAdd(t *testing.T) {
	tests := []testBSTAdd[int, int]{
		{
			name: "Test add node int values (1 node)",
			inputKeys: []int{
				10,
			},
			inputVals: []int{
				10,
			},
			wantError: []bool{false},
			wantTreeVal: []int{
				10,
			},
		},
		{
			name: "Test add node int values (2 nodes incremental)",
			inputKeys: []int{
				10, 20,
			},
			inputVals: []int{
				10, 20,
			},
			wantError: []bool{false, false},
			wantTreeVal: []int{
				10, 20,
			},
		},
		{
			name: "Test new tree int values (2 nodes decremental)",
			inputKeys: []int{
				20, 10,
			},
			inputVals: []int{
				20, 10,
			},
			wantError: []bool{false, false},
			wantTreeVal: []int{
				10, 20,
			},
		},
		{
			name: "Test new tree from several zero int",
			inputKeys: []int{
				0, 0,
			},
			inputVals: []int{
				0, 0,
			},
			wantError: []bool{false, true},
			wantTreeVal: []int{
				0,
			},
		},
		{
			name: "Test new tree int values many nodes",
			inputKeys: []int{
				5, 6, 2, 10, 12, 5, 1, 9,
			},
			inputVals: []int{
				5, 6, 2, 10, 12, 5, 1, 9,
			},
			wantError: []bool{false, false, false, false, false, true, false, false},
			wantTreeVal: []int{
				1, 2, 5, 6, 9, 10, 12,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add root
			root := NewBSTRoot[int](tt.inputKeys[0], tt.inputVals[0])

			// Add the remaining nodes
			for i := 1; i < len(tt.inputKeys); i++ {
				err := root.Add(tt.inputKeys[i], tt.inputVals[i])
				if tt.wantError[i] {
					if (err != nil) != tt.wantError[i] {
						t.Errorf("actual = %v, want %v", (err != nil) == tt.wantError[i], tt.wantError)
					}
				}
			}

			got := root.InorderTraversal()
			if len(got) != len(tt.wantTreeVal) {
				t.Errorf("actual length = %v, want length %v", len(got), len(tt.wantTreeVal))
			}
			for i, wantVal := range tt.wantTreeVal {
				if got[i].value != wantVal {
					t.Errorf("actual = %v, want %v", got[i].value, wantVal)
				}
			}
		})
	}
}

type testBSTreeClear[K cmp.Ordered, V any] struct {
	name        string
	inputKeys   []K
	inputVals   []V
	wantTreeVal []V
}

func TestBSTreeClear(t *testing.T) {
	tests := []testBSTreeClear[int, int]{
		{
			name: "Test clear the tree 1",
			inputKeys: []int{
				2, 1, 4, 3, 5,
			},
			inputVals: []int{
				2, 1, 4, 3, 5,
			},
			wantTreeVal: []int{},
		},
		{
			name: "Test clear the tree 2",
			inputKeys: []int{
				6, 2, 9, 1, 4, 8, 11, 3, 5, 7, 10, 12, 13,
			},
			inputVals: []int{
				6, 2, 9, 1, 4, 8, 11, 3, 5, 7, 10, 12, 13,
			},
			wantTreeVal: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Add root
			root := NewBSTArray[int](tt.inputKeys, tt.inputVals)
			root = root.Clear()

			got := root.LevelOrderTraversal()
			if len(got) != len(tt.wantTreeVal) {
				t.Errorf("actual length = %v, want length %v", len(got), len(tt.wantTreeVal))
			}
			for i, wantVal := range tt.wantTreeVal {
				if got[i].value != wantVal {
					t.Errorf("actual = %v, want %v", got[i].value, wantVal)
				}
			}
		})
	}
}
