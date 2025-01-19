package tree

import "cmp"

type TreeNode[K cmp.Ordered, V any] struct {
	key   K
	value V
}
