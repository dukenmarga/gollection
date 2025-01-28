package tree

import (
	"cmp"
	"fmt"

	"github.com/dukenmarga/gollection/deque"
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
func (tree *BinarySearchTree[K, V]) Add(key K, value V) error {
	newNode := &BinarySearchTree[K, V]{
		TreeNode: &TreeNode[K, V]{
			key:   key,
			value: value,
		},
	}

	err := tree.AddNode(newNode)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

func (tree *BinarySearchTree[K, V]) AddNode(node *BinarySearchTree[K, V]) error {
	if node == nil {
		return nil
	}

	if node.key < tree.key {
		if tree.left == nil {
			tree.left = node
			return nil
		}
		tree.left.AddNode(node)
	} else if node.key > tree.key {
		if tree.right == nil {
			tree.right = node
			return nil
		}
		tree.right.AddNode(node)
	} else if node.key == tree.key {
		return fmt.Errorf("key already exists")
	} else {
		tree.value = node.value
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

func (tree *BinarySearchTree[K, V]) DebugLevelOrderTraversalAsList() {
	fmt.Printf("Level Order Traversal:\n")
	list := tree.LevelOrderTraversal()
	for _, node := range list {
		fmt.Printf("%+v", node.TreeNode)
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (tree *BinarySearchTree[K, V]) Delete(key K) error {
	var err error
	_, err = deleteBSTNode(tree, key)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

// Delete a node from the tree by key.
// After deleting the node, the parent node
// will re-add the node's children.
func deleteBSTNode[K cmp.Ordered, V any](tree *BinarySearchTree[K, V], key K) (*BinarySearchTree[K, V], error) {
	var err error
	// If the node is not found, return an error
	if tree == nil {
		return nil, fmt.Errorf("key not found")
	}

	if key < tree.key {
		tree.left, err = deleteBSTNode(tree.left, key)
		if err != nil {
			return tree, fmt.Errorf("%w", err)
		}
	} else if key > tree.key {
		tree.right, err = deleteBSTNode(tree.right, key)
		if err != nil {
			return tree, fmt.Errorf("%w", err)
		}
	} else if key == tree.key {
		// Set the tree.left as the new node
		// replacing the deleted node.
		// Then add the right child of the deleted node
		// by calling AddNode, so it can re-determine
		// the position of the node under the new parent.
		if tree.left != nil {
			right := tree.right
			*tree = *tree.left
			_ = tree.AddNode(right)
		} else {
			if tree.right != nil {
				*tree = *tree.right
			}
		}
		if tree.left == nil && tree.right == nil {
			return nil, nil
		}
	}

	return tree, nil
}

func (tree *BinarySearchTree[K, V]) Find(key K) (*BinarySearchTree[K, V], error) {
	if tree == nil {
		return nil, fmt.Errorf("key not found")
	}

	if tree.key == key {
		return tree, nil
	}
	if key <= tree.key {
		return tree.left.Find(key)
	} else {
		return tree.right.Find(key)
	}
}

func (tree *BinarySearchTree[K, V]) InorderTraversal() []*BinarySearchTree[K, V] {
	if tree == nil {
		return []*BinarySearchTree[K, V]{}
	}

	left := tree.left.InorderTraversal()
	right := tree.right.InorderTraversal()

	return append(left, append([]*BinarySearchTree[K, V]{tree}, right...)...)
}

func (tree *BinarySearchTree[K, V]) LevelOrderTraversal() []*BinarySearchTree[K, V] {
	if tree == nil {
		return []*BinarySearchTree[K, V]{}
	}

	var results []*BinarySearchTree[K, V]
	deq := deque.NewDequeue([]*BinarySearchTree[K, V]{tree})
	for !deq.IsEmpty() {
		for deq.Length() > 0 {
			node, _ := deq.PopLeft()
			results = append(results, node)
			if node.left != nil {
				deq.PushRight(node.left)
			}
			if node.right != nil {
				deq.PushRight(node.right)
			}
		}
	}

	return results
}

func (tree *BinarySearchTree[K, V]) Update(key K, value V) error {
	node, err := tree.Find(key)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	*node = BinarySearchTree[K, V]{
		TreeNode: &TreeNode[K, V]{
			key:   key,
			value: value,
		},
		left:  node.left,
		right: node.right,
	}
	return nil
}
