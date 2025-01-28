package tree

import (
	"cmp"
	"fmt"

	"github.com/dukenmarga/gollection/deque"
)

type AVLTree[K cmp.Ordered, V any] struct {
	*TreeNode[K, V]
	left  *AVLTree[K, V]
	right *AVLTree[K, V]
}

func NewAVLTArray[K cmp.Ordered, V any](keys []K, values []V) *AVLTree[K, V] {
	if len(keys) == 0 {
		return nil
	}
	if len(values) == 0 {
		return nil
	}
	root := NewAVLTRoot(keys[0], values[0])
	for i := 1; i < len(keys); i++ {
		root.Add(keys[i], values[i])
	}
	return root
}

func NewAVLTRoot[K cmp.Ordered, V any](key K, value V) *AVLTree[K, V] {
	return &AVLTree[K, V]{
		TreeNode: &TreeNode[K, V]{
			key:   key,
			value: value,
		},
	}
}

// Add a node to the tree. Usually this function is called
// by root node, but it can be called by any node.
// If called by non-root node, it will add the new node under
// the parent node and probably will not create a balance
// tree overall.
func (tree *AVLTree[K, V]) Add(key K, value V) error {
	newNode := &AVLTree[K, V]{
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

func (tree *AVLTree[K, V]) AddNode(node *AVLTree[K, V]) error {
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

	// Balancing
	if tree.GetBalance() > 1 {
		if tree.left.GetBalance() < 0 {
			tree.left.RotateLeft()
		}
		tree.RotateRight()
	}

	if tree.GetBalance() < -1 {
		if tree.right.GetBalance() > 0 {
			tree.right.RotateRight()
		}
		tree.RotateLeft()
	}
	return nil
}

func (tree *AVLTree[K, V]) Clear() *AVLTree[K, V] {
	if tree == nil {
		return nil
	}
	tree.left = tree.left.Clear()
	tree.right = tree.right.Clear()
	tree = nil
	return tree
}

func (tree *AVLTree[K, V]) DebugInorderTraversalAsList() {
	fmt.Printf("Inorder Traversal:\n")
	list := tree.InorderTraversal()
	for _, node := range list {
		fmt.Printf("%+v", node.TreeNode)
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (tree *AVLTree[K, V]) DebugLevelOrderTraversalAsList() {
	fmt.Printf("Level Order Traversal:\n")
	list := tree.LevelOrderTraversal()
	for _, node := range list {
		fmt.Printf("%+v", node.TreeNode)
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// Delete a node from the tree by key.
// After deleting the node, the parent node
// will re-add the node's children.
func (tree *AVLTree[K, V]) Delete(key K) error {
	var err error
	_, err = delete(tree, key)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (tree *AVLTree[K, V]) Find(key K) (*AVLTree[K, V], error) {
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

// GetBalance calculate the difference of the left
// height and right height.
func (tree *AVLTree[K, V]) GetBalance() int {
	return tree.left.Height() - tree.right.Height()
}

func (tree *AVLTree[K, V]) Height() int {
	if tree == nil {
		return -1
	}
	// fmt.Printf("tree: %+v\n", tree.value)
	return 1 + max(tree.left.Height(), tree.right.Height())
}

func (tree *AVLTree[K, V]) InorderTraversal() []*AVLTree[K, V] {
	if tree == nil {
		return []*AVLTree[K, V]{}
	}

	left := tree.left.InorderTraversal()
	right := tree.right.InorderTraversal()

	return append(left, append([]*AVLTree[K, V]{tree}, right...)...)
}

func (tree *AVLTree[K, V]) LevelOrderTraversal() []*AVLTree[K, V] {
	if tree == nil {
		return []*AVLTree[K, V]{}
	}

	var results []*AVLTree[K, V]
	deq := deque.NewDequeue([]*AVLTree[K, V]{tree})
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

func (tree *AVLTree[K, V]) RotateLeft() {
	// Theoritically, RotateLeft is called
	// only when tree.right != nil.
	// This is just to make sure to avoid
	// panic/nil dereference
	if tree.right == nil {
		return
	}

	// Get the new root temporarily and the old root
	newRoot := *tree.right
	oldRoot := *tree

	// Set the old root's right to the new root's left
	// Set nil if the new root's left is nil
	if newRoot.left != nil {
		*oldRoot.right = *newRoot.left
	} else {
		oldRoot.right = nil
	}

	// Set the new root's left to the old root.
	if newRoot.left == nil {
		newRoot.left = new(AVLTree[K, V])
	}
	*newRoot.left = oldRoot

	// Set the current tree to the new root
	*tree = newRoot
}

func (tree *AVLTree[K, V]) RotateRight() {
	// Theoritically, RotateRight is called
	// only when tree.left != nil.
	// This is just to make sure to avoid
	// panic/nil dereference
	if tree.left == nil {
		return
	}

	// Get the new root temporarily and the old root
	newRoot := *tree.left
	oldRoot := *tree

	// Set the old root's left to the new root's right.
	// Set nil if the new root's right is nil
	if newRoot.right != nil {
		*oldRoot.left = *newRoot.right
	} else {
		oldRoot.left = nil
	}

	// Set the new root's right to the old root.
	if newRoot.right == nil {
		newRoot.right = new(AVLTree[K, V])
	}
	*newRoot.right = oldRoot

	// Set the current tree to the new root
	*tree = newRoot
}

func (tree *AVLTree[K, V]) Update(key K, value V) error {
	node, err := tree.Find(key)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	*node = AVLTree[K, V]{
		TreeNode: &TreeNode[K, V]{
			key:   key,
			value: value,
		},
		left:  node.left,
		right: node.right,
	}
	return nil
}

func delete[K cmp.Ordered, V any](tree *AVLTree[K, V], key K) (*AVLTree[K, V], error) {
	var err error
	// If the node is not found, return an error
	if tree == nil {
		return nil, fmt.Errorf("key not found")
	}

	if key < tree.key {
		tree.left, err = delete(tree.left, key)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
	} else if key > tree.key {
		tree.right, err = delete(tree.right, key)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
	} else if key == tree.key {
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
			return nil, nil
		}
	}
	if tree.GetBalance() > 1 {
		if tree.left.GetBalance() < 0 {
			tree.left.RotateLeft()
		}
		tree.RotateRight()
	}

	if tree.GetBalance() < -1 {
		if tree.right.GetBalance() > 0 {
			tree.right.RotateRight()
		}
		tree.RotateLeft()
	}
	return tree, nil
}
