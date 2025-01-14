package tree

import (
	"cmp"
	"fmt"
)

type BinarySearchTree[K cmp.Ordered, V any] struct {
	*NodeTree[K, V]
}

func NewBSTArray[K cmp.Ordered, V any](keys []K, values []V) *BinarySearchTree[K, V] {
	if len(keys) == 0 {
		return nil
	}
	if len(values) == 0 {
		return nil
	}
	root := NewBSTRoot(keys[0], values[0])
	for i := 1; i < len(keys); i++ {
		root.Add(keys[i], values[i])
	}
	return root
}

func NewBSTRoot[K cmp.Ordered, V any](key K, value V) *BinarySearchTree[K, V] {
	return &BinarySearchTree[K, V]{
		NodeTree: &NodeTree[K, V]{
			key:   key,
			value: value,
		},
	}
}

// Add a node to the tree. Usually this function is called
// by root node and will add the node to the appropriate place
// in the tree.
func (tree *BinarySearchTree[K, V]) Add(key K, value V) {
	newTree := &BinarySearchTree[K, V]{
		NodeTree: &NodeTree[K, V]{
			key:   key,
			value: value,
		},
	}
	// If this is the root
	if tree == nil {
		tree = newTree
		fmt.Printf("Root is: %+v", tree)
		return
	}

	if key <= tree.key {
		if tree.left == nil {
			tree.left = newTree
			return
		}
		tree.left.Add(key, value)
	} else {
		if tree.right == nil {
			tree.right = newTree
			return
		}
		tree.right.Add(key, value)
	}
}

func (tree *BinarySearchTree[K, V]) DebugInorderTraversalAsList() {
	fmt.Printf("Inorder Traversal:\n")
	list := tree.InorderTraversal()
	fmt.Printf("%v", list)
	fmt.Printf("\n")
}

func (tree *BinarySearchTree[K, V]) InorderTraversal() []BinarySearchTree[K, V] {
	if tree == nil {
		return []BinarySearchTree[K, V]{}
	}

	left := tree.left.InorderTraversal()
	right := tree.right.InorderTraversal()

	return append(left, append([]BinarySearchTree[K, V]{*tree}, right...)...)
}
