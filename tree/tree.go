package tree

import "cmp"

type TreeNode[K cmp.Ordered, V any] struct {
	key   K
	value V
	left  *BinarySearchTree[K, V]
	right *BinarySearchTree[K, V]
}
