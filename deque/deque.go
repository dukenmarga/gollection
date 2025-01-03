package deque

import "fmt"

// Node contains the value and the next and previous node.
// The value can be anything, from int to string to struct or even map.
type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

// DequeueList contains the head and tail nodes.
// The head points to the first node and the tail points to the last node.
// The head and tail can be nil if the list is empty.
// For the sake of consistency of definition of head/tail terms,
// the head is always on the left side, while tail is always on the right side.
type DequeueList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

// PushLeft adds a new node to the left of the list.
// The new added node becomes the new head of the list.
func (list *DequeueList[T]) PushLeft(value T) {
	// create a new node
	newNode := &Node[T]{
		value: value,
	}

	// check if the list is empty (no head and no tail).
	// if so, set the head and tail to the new node
	if list.tail == nil {
		list.tail = newNode
	}
	if list.head == nil {
		list.head = newNode
	}

	// If the list is not empty (head is not nil)
	if list.head != nil {
		// get the current head
		currHead := list.head

		// set the new node's on the left side to the current head
		currHead.prev = newNode

		// set the current head's on the right side to the new node
		newNode.next = currHead

		// set the new node as the new head
		list.head = newNode
	}
}

// PushRight adds a new node to the right of the list.
// The new added node becomes the new tail of the list.
func (list *DequeueList[T]) PushRight(value T) {
	// create a new node
	newNode := &Node[T]{
		value: value,
	}

	// check if the list is empty (no head and no tail).
	// if so, set the head and tail to the new node
	if list.tail == nil {
		list.tail = newNode
	}
	if list.head == nil {
		list.head = newNode
	}

	// If the list is not empty (tail is not nil)
	if list.tail != nil {
		// get the current tail
		currTail := list.tail

		// set the new node's on the right side to the current tail
		currTail.next = newNode

		// set the current tail's on the left side to the new node
		newNode.prev = currTail

		// set the new node as the new tail
		list.tail = newNode
	}
}

// PopLeft removes the head of the list and returns the value of the removed node.
func (list *DequeueList[T]) PopLeft() T {
	// get current head
	popNode := list.head

	// get the next node
	next := list.head.next

	// set the next node as the new head
	list.head = next

	// the new head will has no previous node, so we set it to nil
	if next != nil {
		list.head.prev = nil
	}

	return popNode.value
}

func (list *DequeueList[T]) PopRight() T {
	// get current tail
	popNode := list.tail

	// get the previous node
	prev := list.tail.prev

	// set the previous node as the new tail
	list.tail = prev

	// the new tail will has no next node, so we set it to nil
	if prev != nil {
		list.tail.next = nil
	}

	return popNode.value
}

func (list *DequeueList[T]) DisplayQueue() {
	current := list.head
	for current != nil {
		fmt.Printf("Val: %v\n", current.value)
		current = current.next
	}
}
