package tree

import (
	"cmp"
	"fmt"
)

type TreeNode[K cmp.Ordered, V any] struct {
	key   K
	value V
	left  *TreeNode[K, V]
	right *TreeNode[K, V]
}

func NewBSTArray[K cmp.Ordered, V any](keys []K, values []V) *TreeNode[K, V] {
	if len(keys) == 0 {
		return nil
	}
	if len(values) == 0 {
		return nil
	}
	root := NewBSTRoot(keys[0], values[0])
	for i := 1; i < len(keys); i++ {
		root.Insert(keys[i], values[i])
	}
	return root
}

func NewBSTRoot[K cmp.Ordered, V any](key K, value V) *TreeNode[K, V] {
	return &TreeNode[K, V]{
		key:   key,
		value: value,
	}
}

func (tree *TreeNode[K, V]) Insert(key K, value V) {
	newTree := &TreeNode[K, V]{
		key:   key,
		value: value,
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
		tree.left.Insert(key, value)
	} else {
		if tree.right == nil {
			tree.right = newTree
			return
		}
		tree.right.Insert(key, value)
	}
}

func (tree *TreeNode[K, V]) DebugInorderTraversalAsList() {
	fmt.Printf("Inorder Traversal:\n")
	list := tree.InorderTraversal()
	fmt.Printf("%v", list)
	fmt.Printf("\n")
}

func (tree *TreeNode[K, V]) InorderTraversal() []TreeNode[K, V] {
	if tree == nil {
		return []TreeNode[K, V]{}
	}

	left := tree.left.InorderTraversal()
	right := tree.right.InorderTraversal()

	return append(left, append([]TreeNode[K, V]{*tree}, right...)...)
}
