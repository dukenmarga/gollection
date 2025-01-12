package tree

import (
	"cmp"
	"fmt"
)

type TreeNode[T cmp.Ordered] struct {
	value T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func NewBinarySearchTree[T cmp.Ordered](list []T) *TreeNode[T] {
	if len(list) == 0 {
		return nil
	}
	root := NewRoot(list[0])
	for i := 1; i < len(list); i++ {
		root.Insert(list[i])
	}
	return root
}

func NewRoot[T cmp.Ordered](value T) *TreeNode[T] {
	return &TreeNode[T]{
		value: value,
	}
}

func (tree *TreeNode[T]) Insert(value T) {
	newTree := &TreeNode[T]{
		value: value,
	}
	// If this is the root
	if tree == nil {
		tree = newTree
		fmt.Printf("Root is: %+v", tree)
		return
	}

	if value <= tree.value {
		if tree.left == nil {
			tree.left = newTree
			return
		}
		tree.left.Insert(value)
	} else {
		if tree.right == nil {
			tree.right = newTree
			return
		}
		tree.right.Insert(value)
	}
}

func (tree *TreeNode[T]) PrintInorderTraversalAsList() []T {
	fmt.Printf("Inorder Traversal:\n")
	list := tree.InorderTraversal()
	fmt.Printf("%v", list)
	fmt.Printf("\n")
	return list
}

func (tree *TreeNode[T]) InorderTraversal() []T {
	if tree == nil {
		return []T{}
	}

	left := tree.left.InorderTraversal()
	value := tree.value
	right := tree.right.InorderTraversal()

	return append(left, append([]T{value}, right...)...)
}
