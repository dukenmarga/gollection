package tree

import (
	"cmp"
	"reflect"
	"testing"
)

type testBPlusIsOverflow[K cmp.Ordered, V any] struct {
	name      string
	inputM    int
	inputKeys []K
	wantVal   bool
}

func TestNodeIsOverflow(t *testing.T) {
	tests := []testBPlusIsOverflow[int, int]{
		{
			name:   "Test: A node has 3 keys, m=4, return false",
			inputM: 4,
			inputKeys: []int{
				20, 4, 15,
			},
			wantVal: false,
		},
		{
			name:   "Test: A node has 4 keys, m=4, return true",
			inputM: 4,
			inputKeys: []int{
				20, 4, 15, 16,
			},
			wantVal: true,
		},
		{
			name:   "Test: A node has 5 keys, m=4, return true",
			inputM: 4,
			inputKeys: []int{
				20, 4, 15, 16, 17,
			},
			wantVal: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := BPlusTreeNode[int, int]{
				keys: tt.inputKeys,
			}

			got := node.IsOverflow(tt.inputM)
			if got != tt.wantVal {
				t.Errorf("actual = %v, want %v", got, tt.wantVal)
			}
		})
	}
}

type testBPlusInsertKeyAndChild[K cmp.Ordered, V any] struct {
	name             string
	inputNode        *BPlusTreeNode[K, V]
	insertedKey      K
	insertedChildren []*BPlusTreeNode[K, V]
	wantError        bool
	wantKeys         []K
	wantChildren     []*BPlusTreeNode[K, V]
}

func TestNodeInsertKeyAndChild(t *testing.T) {
	tests := []testBPlusInsertKeyAndChild[int, int]{
		{
			name: "Test: A node has 3 keys & 4 children, added 1 in the middle, now has 4 keys & 5 children",
			inputNode: &BPlusTreeNode[int, int]{
				keys: []int{
					10, 20, 30,
				},
				children: []*BPlusTreeNode[int, int]{
					{
						keys:   []int{8},
						values: []int{8},
					},
					{
						keys:   []int{10, 15},
						values: []int{10, 15},
					},
					{
						keys:   []int{20, 24},
						values: []int{20, 24},
					},
					{
						keys:   []int{30, 34, 35},
						values: []int{30, 34, 35},
					},
				},
			},
			insertedKey: 25,
			insertedChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{25},
					values: []int{25},
				},
			},
			wantError: false,
			wantKeys: []int{
				10, 20, 25, 30,
			},
			wantChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{8},
					values: []int{8},
				},
				{
					keys:   []int{10, 15},
					values: []int{10, 15},
				},
				{
					keys:   []int{20, 24},
					values: []int{20, 24},
				},
				{
					keys:   []int{25},
					values: []int{25},
				},
				{
					keys:   []int{30, 34, 35},
					values: []int{30, 34, 35},
				},
			},
		},
		{
			name: "Test: A node has 3 keys & 4 children, added 1 at the beginning, now has 4 keys & 5 children",
			inputNode: &BPlusTreeNode[int, int]{
				keys: []int{
					10, 20, 30,
				},
				children: []*BPlusTreeNode[int, int]{
					{
						keys:   []int{8},
						values: []int{8},
					},
					{
						keys:   []int{10, 15},
						values: []int{10, 15},
					},
					{
						keys:   []int{20, 24},
						values: []int{20, 24},
					},
					{
						keys:   []int{30, 34, 35},
						values: []int{30, 34, 35},
					},
				},
			},
			insertedKey: 5,
			insertedChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{5},
					values: []int{5},
				},
			},
			wantError: false,
			wantKeys: []int{
				5, 10, 20, 30,
			},
			wantChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{5},
					values: []int{5},
				},
				{
					keys:   []int{8},
					values: []int{8},
				},
				{
					keys:   []int{10, 15},
					values: []int{10, 15},
				},
				{
					keys:   []int{20, 24},
					values: []int{20, 24},
				},
				{
					keys:   []int{30, 34, 35},
					values: []int{30, 34, 35},
				},
			},
		},
		{
			name: "Test: A node has 3 keys & 4 children, added 1 at the end, now has 4 keys & 5 children",
			inputNode: &BPlusTreeNode[int, int]{
				keys: []int{
					10, 20, 30,
				},
				children: []*BPlusTreeNode[int, int]{
					{
						keys:   []int{8},
						values: []int{8},
					},
					{
						keys:   []int{10, 15},
						values: []int{10, 15},
					},
					{
						keys:   []int{20, 24},
						values: []int{20, 24},
					},
					{
						keys:   []int{30, 34, 35},
						values: []int{30, 34, 35},
					},
				},
			},
			insertedKey: 55,
			insertedChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{55},
					values: []int{55},
				},
			},
			wantError: false,
			wantKeys: []int{
				10, 20, 30, 55,
			},
			wantChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{8},
					values: []int{8},
				},
				{
					keys:   []int{10, 15},
					values: []int{10, 15},
				},
				{
					keys:   []int{20, 24},
					values: []int{20, 24},
				},
				{
					keys:   []int{30, 34, 35},
					values: []int{30, 34, 35},
				},
				{
					keys:   []int{55},
					values: []int{55},
				},
			},
		},
		{
			name:        "Test: If a node is nil, will return an error",
			inputNode:   nil,
			insertedKey: 55,
			insertedChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{55},
					values: []int{55},
				},
			},
			wantError: true,
			wantKeys:  []int{},
			wantChildren: []*BPlusTreeNode[int, int]{
				{},
			},
		},
		{
			name: "Test: If a node is empty, but add only 1 key & 1 child, will return an error (it should add 2 nodes)",
			inputNode: &BPlusTreeNode[int, int]{
				keys:     []int{},
				children: []*BPlusTreeNode[int, int]{},
			},
			insertedKey: 55,
			insertedChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{55},
					values: []int{55},
				},
			},
			wantError: true,
			wantKeys:  []int{},
			wantChildren: []*BPlusTreeNode[int, int]{
				{},
			},
		},
		{
			name: "Test: If a node is empty and add 1 keys & 2 children, will not return an error",
			inputNode: &BPlusTreeNode[int, int]{
				keys:     []int{},
				children: []*BPlusTreeNode[int, int]{},
			},
			insertedKey: 55,
			insertedChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{27, 43},
					values: []int{27, 43},
				},
				{
					keys:   []int{55, 64},
					values: []int{55, 64},
				},
			},
			wantError: false,
			wantKeys: []int{
				55,
			},
			wantChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{27, 43},
					values: []int{27, 43},
				},
				{
					keys:   []int{55, 64},
					values: []int{55, 64},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := tt.inputNode
			err := node.InsertKeyAndChild(tt.insertedKey, tt.insertedChildren)

			// assert error
			if tt.wantError {
				if (err != nil) != tt.wantError {
					t.Errorf("actual = %v, want %v", (err != nil) == tt.wantError, tt.wantError)
				}
				return
			}

			// assert length
			if len(node.keys) != len(tt.wantKeys) {
				t.Errorf("actual keys length = %v, want keys length %v", len(node.keys), len(tt.wantKeys))
			}
			if len(node.children) != len(tt.wantChildren) {
				t.Errorf("actual children length = %v, want children length %v", len(node.children), len(tt.wantChildren))
			}

			// assert keys
			for i, key := range node.keys {
				if key != tt.wantKeys[i] {
					t.Errorf("actual = %v, want %v", key, tt.wantKeys[i])
				}
			}

			// assert children
			for i, child := range node.children {
				if !reflect.DeepEqual(child, tt.wantChildren[i]) {
					t.Errorf("actual = %v, want %v", child, tt.wantChildren[i])
				}
			}
		})
	}
}

type testBPlusInsertKeyAndValue[K cmp.Ordered, V any] struct {
	name          string
	inputNode     *BPlusTreeNode[K, V]
	insertedKey   K
	insertedValue V
	wantError     bool
	wantKeys      []K
	wantValues    []V
	wantIsLeaf    bool
}

func TestNodeInsertKeyAndValue(t *testing.T) {
	tests := []testBPlusInsertKeyAndValue[int, int]{
		{
			name: "Test: A node has 4 keys/values, added 1 in the middle, now 5 keys/values",
			inputNode: &BPlusTreeNode[int, int]{
				keys: []int{
					10, 20, 30,
				},
				values: []int{
					10, 20, 30,
				},
			},
			insertedKey:   25,
			insertedValue: 25,
			wantError:     false,
			wantKeys: []int{
				10, 20, 25, 30,
			},
			wantValues: []int{
				10, 20, 25, 30,
			},
			wantIsLeaf: true,
		},
		{
			name: "Test: A node has 4 keys/values, added 1 at the beginning, now 5 keys/values",
			inputNode: &BPlusTreeNode[int, int]{
				keys: []int{
					10, 20, 30,
				},
				values: []int{
					10, 20, 30,
				},
			},
			insertedKey:   5,
			insertedValue: 5,
			wantError:     false,
			wantKeys: []int{
				5, 10, 20, 30,
			},
			wantValues: []int{
				5, 10, 20, 30,
			},
			wantIsLeaf: true,
		},
		{
			name: "Test: A node has 4 keys/values, added 1 at the end, now 5 keys/values",
			inputNode: &BPlusTreeNode[int, int]{
				keys: []int{
					10, 20, 30,
				},
				values: []int{
					10, 20, 30,
				},
			},
			insertedKey:   55,
			insertedValue: 55,
			wantError:     false,
			wantKeys: []int{
				10, 20, 30, 55,
			},
			wantValues: []int{
				10, 20, 30, 55,
			},
			wantIsLeaf: true,
		},
		{
			name: "Test: If a node is empty",
			inputNode: &BPlusTreeNode[int, int]{
				keys:   []int{},
				values: []int{},
			},
			insertedKey:   55,
			insertedValue: 55,
			wantError:     false,
			wantKeys: []int{
				55,
			},
			wantValues: []int{
				55,
			},
			wantIsLeaf: true,
		},
		{
			name:          "Test: If a node is nil",
			inputNode:     nil,
			insertedKey:   55,
			insertedValue: 55,
			wantError:     true,
			wantKeys:      []int{},
			wantValues:    []int{},
			wantIsLeaf:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := tt.inputNode
			err := node.InsertKeyAndValue(tt.insertedKey, tt.insertedValue)

			// assert error
			if tt.wantError {
				if (err != nil) != tt.wantError {
					t.Errorf("actual = %v, want %v", (err != nil) == tt.wantError, tt.wantError)
				}
				return
			}

			// assert length
			if len(node.keys) != len(tt.wantKeys) {
				t.Errorf("actual keys length = %v, want keys length %v", len(node.keys), len(tt.wantKeys))
			}
			if len(node.values) != len(tt.wantValues) {
				t.Errorf("actual values length = %v, want values length %v", len(node.values), len(tt.wantValues))
			}

			// assert keys
			for i, key := range node.keys {
				if key != tt.wantKeys[i] {
					t.Errorf("actual = %v, want %v", key, tt.wantKeys[i])
				}
			}

			// assert values
			for i, value := range node.values {
				if !reflect.DeepEqual(value, tt.wantValues[i]) {
					t.Errorf("actual = %v, want %v", value, tt.wantValues[i])
				}
			}

			// assert isLeaf
			if node.isLeaf != tt.wantIsLeaf {
				t.Errorf("actual = %v, want %v", node.isLeaf, tt.wantIsLeaf)
			}
		})
	}
}

type testBPlus_searchCommonNode[K cmp.Ordered, V any] struct {
	name       string
	inputNode  *BPlusTreeNode[K, V]
	searchKey  K
	wantKeys   []K
	wantValues []V
}

func TestNode_searchCommonNode(t *testing.T) {
	tests := []testBPlus_searchCommonNode[int, int]{
		{
			name: "Test: Search item at the start, from 1 level node",
			inputNode: &BPlusTreeNode[int, int]{
				keys: []int{
					10, 20, 30,
				},
				children: []*BPlusTreeNode[int, int]{
					{
						keys:   []int{8},
						values: []int{8},
						isLeaf: true,
					},
					{
						keys:   []int{10, 15},
						values: []int{10, 15},
						isLeaf: true,
					},
					{
						keys:   []int{20, 24},
						values: []int{20, 24},
						isLeaf: true,
					},
					{
						keys:   []int{30, 34, 35},
						values: []int{30, 34, 35},
						isLeaf: true,
					},
				},
			},
			searchKey: 8,
			wantKeys: []int{
				8,
			},
			wantValues: []int{
				8,
			},
		},
		{
			name: "Test: Search item in the middle, from 1 level node",
			inputNode: &BPlusTreeNode[int, int]{
				keys: []int{
					10, 20, 30,
				},
				children: []*BPlusTreeNode[int, int]{
					{
						keys:   []int{8},
						values: []int{8},
						isLeaf: true,
					},
					{
						keys:   []int{10, 15},
						values: []int{10, 15},
						isLeaf: true,
					},
					{
						keys:   []int{20, 24},
						values: []int{20, 24},
						isLeaf: true,
					},
					{
						keys:   []int{30, 34, 35},
						values: []int{30, 34, 35},
						isLeaf: true,
					},
				},
			},
			searchKey: 10,
			wantKeys: []int{
				10, 15,
			},
			wantValues: []int{
				10, 15,
			},
		},
		{
			name: "Test: Search item at the end, from 1 level node",
			inputNode: &BPlusTreeNode[int, int]{
				keys: []int{
					10, 20, 30,
				},
				children: []*BPlusTreeNode[int, int]{
					{
						keys:   []int{8},
						values: []int{8},
						isLeaf: true,
					},
					{
						keys:   []int{10, 15},
						values: []int{10, 15},
						isLeaf: true,
					},
					{
						keys:   []int{20, 24},
						values: []int{20, 24},
						isLeaf: true,
					},
					{
						keys:   []int{30, 34, 35},
						values: []int{30, 34, 35},
						isLeaf: true,
					},
				},
			},
			searchKey: 35,
			wantKeys: []int{
				30, 34, 35,
			},
			wantValues: []int{
				30, 34, 35,
			},
		},
		{
			name: "Test: Search item that doesn't exist, but it will still return the node, from 1 level node",
			inputNode: &BPlusTreeNode[int, int]{
				keys: []int{
					10, 20, 30,
				},
				children: []*BPlusTreeNode[int, int]{
					{
						keys:   []int{8},
						values: []int{8},
						isLeaf: true,
					},
					{
						keys:   []int{10, 15},
						values: []int{10, 15},
						isLeaf: true,
					},
					{
						keys:   []int{20, 24},
						values: []int{20, 24},
						isLeaf: true,
					},
					{
						keys:   []int{30, 34, 35},
						values: []int{30, 34, 35},
						isLeaf: true,
					},
				},
			},
			searchKey: 40,
			wantKeys: []int{
				30, 34, 35,
			},
			wantValues: []int{
				30, 34, 35,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := tt.inputNode
			searchNode := searchCommonNode(node, tt.searchKey)

			// assert keys
			for i, key := range searchNode.keys {
				if key != tt.wantKeys[i] {
					t.Errorf("actual = %v, want %v", key, tt.wantKeys[i])
				}
			}

			// assert values
			for i, value := range searchNode.values {
				if !reflect.DeepEqual(value, tt.wantValues[i]) {
					t.Errorf("actual = %v, want %v", value, tt.wantValues[i])
				}
			}
		})
	}
}

type testBPlusAdd[K cmp.Ordered, V any] struct {
	name              string
	inputM            int
	inputKeys         []K
	inputValues       []V
	wantError         bool
	wantRootKeys      []K
	wantChildren      []*BPlusTreeNode[K, V]
	wantGrandchildren []*BPlusTreeNode[K, V]
}

func TestNodeAdd(t *testing.T) {
	tests := []testBPlusAdd[int, int]{
		{
			name:   "Test: Add 4 nodes to an order-4 tree, will create a parent node",
			inputM: 4,
			inputKeys: []int{
				10, 20, 30, 40,
			},
			inputValues: []int{
				10, 20, 30, 40,
			},
			wantError: false,
			wantRootKeys: []int{
				30,
			},
			wantChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{10, 20},
					values: []int{10, 20},
					isLeaf: true,
				},
				{
					keys:   []int{30, 40},
					values: []int{30, 40},
					isLeaf: true,
				},
			},
		},
		{
			name:   "Test: Add 5 nodes to an order-4 tree, will create a parent node",
			inputM: 4,
			inputKeys: []int{
				10, 20, 30, 40, 50,
			},
			inputValues: []int{
				10, 20, 30, 40, 50,
			},
			wantError: false,
			wantRootKeys: []int{
				30,
			},
			wantChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{10, 20},
					values: []int{10, 20},
					isLeaf: true,
				},
				{
					keys:   []int{30, 40, 50},
					values: []int{30, 40, 50},
					isLeaf: true,
				},
			},
		},
		{
			name:   "Test: Add 6 nodes to an order-4 tree, will create a parent node",
			inputM: 4,
			inputKeys: []int{
				10, 20, 30, 40, 50, 60,
			},
			inputValues: []int{
				10, 20, 30, 40, 50, 60,
			},
			wantError: false,
			wantRootKeys: []int{
				30, 50,
			},
			wantChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{10, 20},
					values: []int{10, 20},
					isLeaf: true,
				},
				{
					keys:   []int{30, 40},
					values: []int{30, 40},
					isLeaf: true,
				},
				{
					keys:   []int{50, 60},
					values: []int{50, 60},
					isLeaf: true,
				},
			},
		},
		{
			name:   "Test: Add 8 nodes to an order-4 tree, will create a parent node",
			inputM: 4,
			inputKeys: []int{
				10, 20, 30, 40, 50, 60, 70, 80,
			},
			inputValues: []int{
				10, 20, 30, 40, 50, 60, 70, 80,
			},
			wantError: false,
			wantRootKeys: []int{
				30, 50, 70,
			},
			wantChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{10, 20},
					values: []int{10, 20},
					isLeaf: true,
				},
				{
					keys:   []int{30, 40},
					values: []int{30, 40},
					isLeaf: true,
				},
				{
					keys:   []int{50, 60},
					values: []int{50, 60},
					isLeaf: true,
				},
				{
					keys:   []int{70, 80},
					values: []int{70, 80},
					isLeaf: true,
				},
			},
		},
		{
			name:   "Test: Add 10 nodes to an order-4 tree, will create two parent nodes and a grandparent node",
			inputM: 4,
			inputKeys: []int{
				10, 20, 30, 40, 50, 60, 70, 80, 31, 32,
			},
			inputValues: []int{
				10, 20, 30, 40, 50, 60, 70, 80, 31, 32,
			},
			wantError: false,
			wantRootKeys: []int{
				50,
			},
			wantChildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{30, 32},
					values: []int{},
					isLeaf: false,
				},
				{
					keys:   []int{50, 70},
					values: []int{},
					isLeaf: false,
				},
			},
			wantGrandchildren: []*BPlusTreeNode[int, int]{
				{
					keys:   []int{10, 20},
					values: []int{10, 20},
					isLeaf: true,
				},
				{
					keys:   []int{30, 31},
					values: []int{30, 31},
					isLeaf: true,
				},
				{
					keys:   []int{32, 40},
					values: []int{32, 40},
					isLeaf: true,
				},
				{
					keys:   []int{50, 60},
					values: []int{50, 60},
					isLeaf: true,
				},
				{
					keys:   []int{70, 80},
					values: []int{70, 80},
					isLeaf: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree := NewBPlusTree[int, int](tt.inputM)
			for i, key := range tt.inputKeys {
				tree.Add(key, tt.inputValues[i])
			}

			root := tree.root

			// fmt.Printf("root: %+v\n", root)

			// assert length
			if len(root.keys) != len(tt.wantRootKeys) {
				t.Errorf("actual keys length = %v, want keys length %v", len(root.keys), len(tt.wantRootKeys))
			}

			// assert root keys
			for i, key := range root.keys {
				if key != tt.wantRootKeys[i] {
					t.Errorf("actual = %v, want %v", key, tt.wantRootKeys[i])
				}
			}

			// assert child keys
			for i, child := range root.children {
				if !reflect.DeepEqual(child.keys, tt.wantChildren[i].keys) {
					t.Errorf("actual = %v, want %v", child.keys, tt.wantChildren[i].keys)
				}
			}

			// assert child values
			for i, child := range root.children {
				if !reflect.DeepEqual(child.values, tt.wantChildren[i].values) {
					t.Errorf("actual = %v, want %v", child.values, tt.wantChildren[i].values)
				}
			}

			// assert child isLeaf
			for i, child := range root.children {
				if child.isLeaf != tt.wantChildren[i].isLeaf {
					t.Errorf("actual = %v, want %v", child.isLeaf, tt.wantChildren[i].isLeaf)
				}
			}

			// assert grandchild keys
			for i, grandchild := range root.children[0].children {
				if !reflect.DeepEqual(grandchild.keys, tt.wantGrandchildren[i].keys) {
					t.Errorf("actual = %v, want %v", grandchild.keys, tt.wantGrandchildren[i].keys)
				}
			}

			// assert parent of first child
			if root != root.children[0].parent {
				t.Errorf("actual = %v, want %v", root, root.children[0].parent)
			}

			// assert parent of second child
			if root != root.children[1].parent {
				t.Errorf("actual = %v, want %v", root, root.children[1].parent)
			}
		})
	}
}
