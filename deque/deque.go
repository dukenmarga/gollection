package deque

import "fmt"

// Node contains the value and the next and previous node.
// The value can be anything, from int to string to struct or even map.
type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

// DequeueList contains the head and tail nodes.
// The head points to the first node and the tail points to the last node.
// The head and tail can be nil if the list is empty.
// For the sake of consistency of definition of head/tail terms,
// the head is always on the left side, while tail is always on the right side.
type DequeueList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length uint
}

func (list *DequeueList[T]) At(index uint) (T, error) {
	current := list.head

	var pos uint = 0

	for current != nil {
		if pos == index {
			return current.value, nil
		}
		current = current.next
		pos++
	}
	var empty T
	return empty, fmt.Errorf("index out of range: %v (total length: %v)", index, list.length)
}

// Clear removes all nodes from the list
func (list *DequeueList[T]) Clear() {
	current := list.head

	for current != nil {
		next := current.next
		current.next = nil
		current = next
	}

	list.head = nil
	list.tail = nil
	list.length = 0
}

func NewDequeue[T any](list []T) *DequeueList[T] {
	deque := &DequeueList[T]{}
	for _, value := range list {
		deque.PushRight(value)
	}
	return deque
}

// PushLeft adds a new node to the left of the list.
// The new added node becomes the new head of the list.
func (list *DequeueList[T]) PushLeft(value T) {
	// create a new node
	newNode := &Node[T]{
		value: value,
	}

	// If the list is empty (no tail),
	// set the tail to the new node
	if list.tail == nil {
		list.tail = newNode
	}

	// get the current head
	currHead := list.head

	// If the list is not empty
	if list.head != nil {
		// set the new node's on the left side to the current head
		currHead.prev = newNode

		// set the current head's on the right side to the new node
		newNode.next = currHead
	}

	// set new node as the new head
	list.head = newNode

	// update length
	list.length++

}

// PushRight adds a new node to the right of the list.
// The new added node becomes the new tail of the list.
func (list *DequeueList[T]) PushRight(value T) {
	// create a new node
	newNode := &Node[T]{
		value: value,
	}

	// If the list is empty (no head),
	// set the head to the new node
	if list.head == nil {
		list.head = newNode
	}

	// get the current tail
	currTail := list.tail

	// If the list is not empty
	if list.tail != nil {
		// set the new node's on the right side to the current tail
		currTail.next = newNode

		// set the current tail's on the left side to the new node
		newNode.prev = currTail
	}

	// set new node as the new tail
	list.tail = newNode

	// update length
	list.length++
}

// PopLeft removes the head of the list and returns the value of the removed node.
func (list *DequeueList[T]) PopLeft() (T, error) {
	// return empty value if the list is empty
	if list.head == nil {
		var empty T
		return empty, fmt.Errorf("list is empty")
	}

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

	// update length
	list.length--

	return popNode.value, nil
}

func (list *DequeueList[T]) PopRight() (T, error) {
	if list.tail == nil {
		var empty T
		return empty, fmt.Errorf("list is empty")
	}

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

	// update length
	list.length--

	return popNode.value, nil
}

// Debug prints the value of each node in the list.
// It contains basic information about the list.
func (list *DequeueList[T]) Debug() {
	current := list.head
	count := 1
	for current != nil {
		fmt.Printf("No.: %v\n", count)
		if current.prev == nil {
			fmt.Print("nil <- ")
		} else {
			fmt.Printf("%v <- ", current.prev.value)
		}

		fmt.Printf("%v", current.value)

		if current.next == nil {
			fmt.Print(" -> nil\n")
		} else {
			fmt.Printf(" -> %v\n\n", current.next.value)
		}

		current = current.next
		count++
	}
}

func (list DequeueList[T]) IsEmpty() bool {
	return list.length == 0
}

func (list DequeueList[T]) Length() uint {
	return list.length
}
