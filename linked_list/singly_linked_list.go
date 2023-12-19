// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use the file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linkedlist

var _ LinkedList[any] = (*SinglyLinkedList[any])(nil)

type singlyNode[T any] struct {
	Val  T
	Next *singlyNode[T]
}

type SinglyLinkedList[T any] struct {
	head *singlyNode[T]
	tail *singlyNode[T]
	size int
}

// NewSinglyLinkedList returns a new singly linked list.
// If the elements is not empty, add the elements to the list.
func NewSinglyLinkedList[T any](elements ...T) *SinglyLinkedList[T] {
	list := &SinglyLinkedList[T]{}
	if len(elements) > 0 {
		list.Add(elements...)
	}
	return list
}

// Add appends the specified elements to the end of the list.(same as Append)
func (l *SinglyLinkedList[T]) Add(elements ...T) {
	if len(elements) > 0 {
		for _, e := range elements {
			node := &singlyNode[T]{Val: e}
			if l.IsEmpty() {
				l.head, l.tail = node, node
			} else {
				l.tail.Next = node
				l.tail = node
			}
			l.size++
		}
	}
}

// Append appends the specified elements to the end of the list.(same as Add)
func (l *SinglyLinkedList[T]) Append(elements ...T) {
	l.Add(elements...)
}

// Prepend prepends the specified elements to the beginning of the list.
func (l *SinglyLinkedList[T]) Prepend(elements ...T) {
	// reverse the elements. i.e. original elements: [2, 3], prepend elements: [0, 1], result: [0, 1, 2, 3]
	for i := len(elements) - 1; i >= 0; i-- {
		node := &singlyNode[T]{Val: elements[i], Next: l.head}
		l.head = node
		if l.size == 0 {
			l.tail = node
		}
		l.size++
	}
}

// GetFirst returns the first element in the list.
// If the list is empty, b return false
func (l *SinglyLinkedList[T]) GetFirst() (t T, b bool) {
	if l.IsEmpty() {
		return
	}
	return l.head.Val, true
}

// GetLast returns the last element in the list.
// If the list is empty, b return false
func (l *SinglyLinkedList[T]) GetLast() (t T, b bool) {
	if l.IsEmpty() {
		return
	}
	return l.tail.Val, true
}

// Get returns the element at the specified position in the list.
// If the index is invalid, b return false
func (l *SinglyLinkedList[T]) Get(index int) (t T, b bool) {
	if l.isInvalidIndex(index) {
		return
	}
	node := l.head
	for i := 0; i < index; i, node = i+1, node.Next {
	}
	return node.Val, true
}

// Set sets the element at the specified position in the list.
// If the index is invalid, b return false
func (l *SinglyLinkedList[T]) Set(index int, e T) bool {
	if index < 0 || index > l.Size()-1 {
		return false
	}
	node := l.head
	for i := 0; i < index; i, node = i+1, node.Next {
	}
	node.Val = e
	return true
}

// Insert inserts the specified elements at the specified position in the list.
func (l *SinglyLinkedList[T]) Insert(index int, elements ...T) bool {
	if l.isInvalidIndex(index) {
		if index == 0 {
			l.Add(elements...)
			return true
		}
		return false
	}

	if index == 0 {
		l.Prepend(elements...)
		return true
	}
	if index == l.Size()-1 {
		l.Append(elements...)
		return true
	}
	prev := l.head
	for i := 0; i < index-1; i, prev = i+1, prev.Next {
	}
	oldNext := prev.Next
	for _, val := range elements {
		node := &singlyNode[T]{Val: val}
		prev.Next = node
		prev = node
		l.size++
	}
	prev.Next = oldNext
	return true
}

// RemoveFirst removes the first element from the list.
// If the list is empty, b return false
func (l *SinglyLinkedList[T]) RemoveFirst() (t T, b bool) {
	if l.IsEmpty() {
		return
	}
	node := l.head
	l.head = node.Next
	l.size--
	if l.IsEmpty() {
		l.tail = nil
	}
	return node.Val, true
}

// RemoveLast removes the last element from the list.
// If the list is empty, b return false
func (l *SinglyLinkedList[T]) RemoveLast() (t T, b bool) {
	if l.IsEmpty() {
		return
	}
	if l.Size() == 1 {
		return l.RemoveFirst()
	}
	prev := l.head
	for prev.Next != l.tail {
		prev = prev.Next
	}
	node := prev.Next
	prev.Next = nil
	l.tail = prev
	l.size--
	return node.Val, true
}

// Remove removes the element at the specified position in the list.
// If the index is invalid, b return false
func (l *SinglyLinkedList[T]) Remove(index int) (t T, b bool) {
	if l.isInvalidIndex(index) {
		return
	}
	if l.Size() == 1 {
		return l.RemoveFirst()
	}
	if index == l.Size()-1 {
		return l.RemoveLast()
	}

	prev := l.head
	// find the previous node of the node to be deleted
	for i := 0; i < index-1; i, prev = i+1, prev.Next {
	}

	node := prev.Next
	prev.Next = node.Next
	l.size--
	t, node = node.Val, nil
	return t, true
}

// IsEmpty checks whether the list is empty
func (l *SinglyLinkedList[T]) IsEmpty() bool {
	return l.Size() == 0
}

// Size returns the size of the list
func (l *SinglyLinkedList[T]) Size() int {
	return l.size
}

// isInvalidIndex checks whether the index is invalid
func (l *SinglyLinkedList[T]) isInvalidIndex(index int) bool {
	return index < 0 || index > l.Size()-1
}

// Clear removes all the elements from the list
func (l *SinglyLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// ToSlice returns a slice containing all the elements in this list.
func (l *SinglyLinkedList[T]) Values() []T {
	elements := make([]T, 0, l.Size())
	node := l.head
	for node != nil {
		elements = append(elements, node.Val)
		node = node.Next
	}
	return elements
}

// Reverse reverses the list
func (l *SinglyLinkedList[T]) Reverse() {
	var prev, next *singlyNode[T]
	cur := l.head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.head = prev
}
