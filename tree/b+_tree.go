package tree

import (
	"cmp"
	"fmt"
)

// B+ Tree rules
// - Order (m): max number of children that an internal node can have before splitting.
//   m should be even or 2n
// - Root: initially will be a leaf node, but later will be an internal node
//   Root can have min 2 children if it is an internal node
// - Internal nodes: has (min m/2-1, max m-1) keys and (min m/2, max m) nodes as children (pointer)
// - Leaf nodes: has (min m/2, max m-1) keys and (min m/2, max m-1) values

type BPlusTreeNode[K cmp.Ordered, V any] struct {
	isLeaf   bool
	keys     []K
	values   []V
	parent   *BPlusTreeNode[K, V]
	children []*BPlusTreeNode[K, V]
	next     *BPlusTreeNode[K, V]
}

type BPlusTree[K cmp.Ordered, V any] struct {
	root *BPlusTreeNode[K, V]
	m    int // order: max chilren
}

func NewBPlusTree[K cmp.Ordered, V any](m int) *BPlusTree[K, V] {
	return &BPlusTree[K, V]{
		root: &BPlusTreeNode[K, V]{
			isLeaf: true,
		},
		m: m,
	}
}

func (tree *BPlusTree[K, V]) Add(key K, value V) {
	// Search the node which the new record will be inserted
	node := tree.whichNodeToInsert(key)

	// Insert the key to the designated node
	node.InsertKeyAndValue(key, value)

	// We run Split operation and get the updated root
	// (the new root will be created if current root is splitting)
	tree.root = node.Split(tree.m)
}

// Find will return a node where the key is located
func (tree *BPlusTree[K, V]) Find(key K) (*BPlusTreeNode[K, V], error) {
	root := tree.root

	// node will contain a series of keys and values
	// that we need to compare to the search key.
	// If the key is in the node, then return the node
	node := searchCommonNode(root, key)
	for _, keyInNode := range node.keys {
		if key == keyInNode {
			return node, nil
		}
	}
	return nil, fmt.Errorf("key not found")
}

// Check whether a node contains too many
// key/value. This function is recommended
// to be called after adding a new key/value.
// E.g if order m=4, then maximum key/value is 3.
// Overflow means it needs to be split
// to maintain the order.
func (node *BPlusTreeNode[K, V]) IsOverflow(m int) bool {
	if m < 3 {
		panic("order m should be greater or equal than 3")
	}
	return len(node.keys) > m-1
}

// InsertKeyAndChild will insert the key/child to an internal node
func (node *BPlusTreeNode[K, V]) InsertKeyAndChild(key K, children []*BPlusTreeNode[K, V]) error {
	// If this node is empty, return error
	if node == nil {
		return fmt.Errorf("node is nil, please initialize it first")
	}

	// If the node is empty, make sure that the inserted children is 2
	if len(node.children) == 0 {
		if len(children) != 2 {
			return fmt.Errorf("node is empty, then children should be 2")
		}
	}

	// If the node is empty, just insert key and children to the node
	if len(node.children) == 0 {
		node.children = children
		node.keys = []K{key}
		return nil
	}

	// Otherwise, we need to append key and children according to the order of the key
	for _, child := range children {
		finishAppendChildren := false
		for i, nodeKey := range node.keys {
			// Find index to put the new key/value.
			// For example we will insert key 5 to existing
			// keys [1,3,4,7,8,9]. The index will be 3 (where 7 > 5).
			// Then we will insert it by appending [1,3,4] + [5] + [7,8,9]
			if nodeKey > key {
				// index selector for left and right is determined
				// by the position of the key.
				// we have m-1 key, but m children, that's why
				// we have special condition for the first item
				j := i
				if i != 0 {
					j = i + 1
				}
				node.keys = append(node.keys[:i], append([]K{key}, node.keys[i:]...)...)
				node.children = append(node.children[0:j], append([]*BPlusTreeNode[K, V]{child}, node.children[j:]...)...)
				finishAppendChildren = true
				break
			}
		}
		if finishAppendChildren {
			continue
		}

		// This is if the node is empty or if the inserted key is the largest compared to the existing keys
		node.keys = append(node.keys, key)
		node.children = append(node.children, child)
	}

	return nil
}

// InsertKeyAndValue will insert the key/value into a leaf node
func (node *BPlusTreeNode[K, V]) InsertKeyAndValue(key K, value V) error {
	// If this node is empty, return error
	if node == nil {
		return fmt.Errorf("node is nil, please initialize it first")
	}

	// set that this node is a leaf
	node.isLeaf = true

	for i, nodeKey := range node.keys {
		// Find index to put the new key/value.
		// For example we will insert key 5 to existing
		// keys [1,3,4,7,8,9]. The index will be 3 (where 7 > 5).
		// Then we will insert it by appending [1,3,4] + [5] + [7,8,9]
		if nodeKey > key {
			node.keys = append(node.keys[:i], append([]K{key}, node.keys[i:]...)...)
			node.values = append(node.values[0:i], append([]V{value}, node.values[i:]...)...)
			return nil
		}
	}

	// This is if the node is empty or if the inserted key is the largest compared to the existing keys
	node.keys = append(node.keys, key)
	node.values = append(node.values, value)

	return nil
}

// Split will split a node into 2 nodes if it is overflow.
// It will be called recursively start from the leaf until
// it reaches the root, then it will return the root.
// Split will handle leaf and non-leaf nodes.
func (node *BPlusTreeNode[K, V]) Split(m int) *BPlusTreeNode[K, V] {
	// Split is different for leaf and non-leaf node.
	// For leaf node:
	// - we need to split the keys and values
	// - splitted nodes hold all keys/values
	// For non-leaf node:
	// - we need to split the keys
	// - splitted nodes hold all keys, except the middle key
	//   that will be promoted to the parent

	// - If the node keys is overflow and it is the root (node.parent is nil)
	//   return the node as root
	// - If the node keys is overflow, but it is not the root,
	//   we keep calling Split() until we find the root node.
	if !node.IsOverflow(m) {
		if node.parent == nil {
			return node
		}
		return node.parent.Split(m)
	}

	// Proceed to the split (split 1 node into 2 nodes)

	// keys are used because it is available for both leaf and non-leaf node
	length := len(node.keys)
	mid := length / 2

	// Determine whether the node is leaf or not
	isLeaf := false
	if node.children == nil {
		isLeaf = true
	}

	// Split the first node and second node keys
	// First node refers to the current node
	// that is being split and use the first half of the keys
	// Second node refers to the new node that is created
	// and use the second half of the keys
	firstNodeKeys := node.keys[0:mid]
	var secondNodeKeys []K
	if isLeaf {
		secondNodeKeys = node.keys[mid:length]
	} else {
		secondNodeKeys = node.keys[mid+1 : length]
	}
	midKey := node.keys[mid]

	// If it is a leaf node, set the values
	firstNodeValues := []V{}
	secondNodeValues := []V{}
	if isLeaf {
		firstNodeValues = node.values[0:mid]
		secondNodeValues = node.values[mid:length]
	}

	// Update the current node and create the second node
	*node = BPlusTreeNode[K, V]{
		keys:   firstNodeKeys,
		values: firstNodeValues,
		parent: node.parent,
		isLeaf: isLeaf,
	}
	secondNode := BPlusTreeNode[K, V]{
		keys:   secondNodeKeys,
		values: secondNodeValues,
		isLeaf: isLeaf,
	}

	// Prepare the parent and its pointer to children
	// There are 2 cases:
	// - If the parent is nil, it means:
	//   a. The current first node is the root
	//   b. Then we need to register both nodes to the new parent
	// - Else, means:
	//   a. The first node has been registered on the parent before splitting
	//   b. Then we only need to register the second node, while
	//      the first node is automatically changed by re-assigning the value
	children := []*BPlusTreeNode[K, V]{}
	if node.parent == nil {
		// Initialize the parent (from nil to a node)
		node.parent = &BPlusTreeNode[K, V]{}

		// Insert the first key/child to the parent
		children = append(children, node)
	}

	if secondNode.parent == nil {
		// First and second node have the same parent
		secondNode.parent = node.parent

		// Insert the second key/child to the parent
		children = append(children, &secondNode)
	}

	// Promote the middle key to the parent
	// (middle key before split is the first key of the second node)
	err := node.parent.InsertKeyAndChild(midKey, children)
	if err != nil {
		panic(err)
	}

	// Recursively check the parent if it needs to split
	return node.parent.Split(m)
}

// searchCommonNode() will return a node that contain
// all the values that are associated with the search key.
// It doesn't necessarily mean that the key exists in the node.
// For example below we have a parent node that has pointer: 60, 85, and 98.
// If we want to searchCommonNode key 51 and the parent pointer that close is 60,
// then searchCommonNode() will return a node that has values
// 38 and 51 (item < 60).
//
//	    60       |     85        |      98
//	38, 51 | 67, 68, 82 | 87, 88, 90, 93 | 98, 100
//
// Using the same example above, if we want to searchCommonNode key 70
// (event if it doesn't exist), then searchCommonNode will return
// a node that contains values 67, 68, 82 (60 <= item < 85)
func searchCommonNode[K cmp.Ordered, V any](node *BPlusTreeNode[K, V], key K) *BPlusTreeNode[K, V] {
	// Operation if leaf
	if node.isLeaf {
		return node
	}

	// Operation if not leaf
	for i, pointerKey := range node.keys {
		// If the key is less than the pointer key,
		// then search the child node (with the same index)
		if key < pointerKey {
			return searchCommonNode(node.children[i], key)
		}
	}
	// If the key is greater than the last pointer key,
	// it means that the key is probably on the last child
	return searchCommonNode(node.children[len(node.children)-1], key)
}

// whichNodeToInsert will return a selected node where a key/value can be inserted
// according to its cluster order.
func (tree *BPlusTree[K, V]) whichNodeToInsert(key K) *BPlusTreeNode[K, V] {
	root := tree.root

	// node will contain a series of keys and values
	// that we need to compare to the search key.
	// If the key is in the node, then return the node
	node := searchCommonNode(root, key)
	return node
}
