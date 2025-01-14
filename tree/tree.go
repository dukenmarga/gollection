package tree

import "cmp"

type NodeTree[K cmp.Ordered, V any] struct {
	key   K
	value V
	left  *BinarySearchTree[K, V]
	right *BinarySearchTree[K, V]
}
