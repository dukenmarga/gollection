package tree

import (
	"cmp"
	"fmt"
)

type BinarySearchTree[K cmp.Ordered, V any] struct {
	*TreeNode[K, V]
	left  *BinarySearchTree[K, V]
	right *BinarySearchTree[K, V]
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
		TreeNode: &TreeNode[K, V]{
			key:   key,
			value: value,
		},
	}
}

// Add a node to the tree. Usually this function is called
// by root node, but it can be called by any node.
// If called by non-root node, it will add the new node under
// the parent node and probably will break the BST rule if the
// new node is not in the correct order under the root.
func (tree *BinarySearchTree[K, V]) Add(key K, value V) {
	newNode := &BinarySearchTree[K, V]{
		TreeNode: &TreeNode[K, V]{
			key:   key,
			value: value,
		},
	}

	tree.AddNode(newNode)
}

func (tree *BinarySearchTree[K, V]) AddNode(node *BinarySearchTree[K, V]) {
	if node == nil {
		return
	}

	if node.key < tree.key {
		if tree.left == nil {
			tree.left = node
			return
		}
		tree.left.AddNode(node)
	} else if node.key > tree.key {
		if tree.right == nil {
			tree.right = node
			return
		}
		tree.right.AddNode(node)
	} else {
		tree.value = node.value
	}
}

func (tree *BinarySearchTree[K, V]) Search(key K) (*BinarySearchTree[K, V], error) {
	if tree == nil {
		return nil, fmt.Errorf("key not found")
	}

	if tree.key == key {
		return tree, nil
	}
	if key <= tree.key {
		return tree.left.Search(key)
	} else {
		return tree.right.Search(key)
	}
}

// Delete a node from the tree by key.
// After deleting the node, the parent node
// will re-add the node's children.
func (tree *BinarySearchTree[K, V]) Delete(key K) error {
	// If the node is not found, return an error
	if tree == nil {
		return fmt.Errorf("key not found")
	}

	if key < tree.key {
		err := tree.left.Delete(key)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	if key > tree.key {
		err := tree.right.Delete(key)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
	}

	if key == tree.key {
		// Set the tree.left as the new node
		// replacing the deleted node.
		// Then add the right child of the deleted node
		// by calling AddNode, so it can re-determine
		// the position of the node under the new parent.
		if tree.left != nil {
			*tree = *tree.left
			tree.AddNode(tree.right)
		} else {
			if tree.right != nil {
				*tree = *tree.right
			}
		}
		if tree.left == nil && tree.right == nil {
			return nil
		}
	}

	return nil
}

func (tree *BinarySearchTree[K, V]) DebugInorderTraversalAsList() {
	fmt.Printf("Inorder Traversal:\n")
	list := tree.InorderTraversal()
	for _, node := range list {
		fmt.Printf("%+v", node.TreeNode)
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (tree *BinarySearchTree[K, V]) InorderTraversal() []*BinarySearchTree[K, V] {
	if tree == nil {
		return []*BinarySearchTree[K, V]{}
	}

	left := tree.left.InorderTraversal()
	right := tree.right.InorderTraversal()

	return append(left, append([]*BinarySearchTree[K, V]{tree}, right...)...)
}
