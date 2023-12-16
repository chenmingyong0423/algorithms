// Copyright 2023 chenmingyong0423

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package linked_list

var _ LinkedList[any] = (*DoublyLinkedList[any])(nil)

type doublyNode[T any] struct {
	Val  T
	Prev *doublyNode[T]
	Next *doublyNode[T]
}

type DoublyLinkedList[T any] struct {
	head *doublyNode[T]
	tail *doublyNode[T]
	size int
}

// NewDoublyLinkedList returns a new doubly linked list.
// If the elements is not empty, add the elements to the list.
func NewDoublyLinkedList[T any](elements ...T) *DoublyLinkedList[T] {
	list := &DoublyLinkedList[T]{}
	if len(elements) > 0 {
		list.Add(elements...)
	}
	return list
}

// Add appends the specified elements to the end of the list.(same as Append)
func (l *DoublyLinkedList[T]) Add(elements ...T) {
	if len(elements) > 0 {
		for _, e := range elements {
			node := &doublyNode[T]{Val: e, Prev: l.tail}
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
func (l *DoublyLinkedList[T]) Append(elements ...T) {
	l.Add(elements...)
}

// Prepend prepends the specified elements to the beginning of the list.
func (l *DoublyLinkedList[T]) Prepend(elements ...T) {
	for i := len(elements) - 1; i >= 0; i-- {
		node := &doublyNode[T]{Val: elements[i], Next: l.head}
		if l.size == 0 {
			l.tail = node
		} else {
			l.head.Prev = node
		}
		l.head = node
		l.size++
	}
}

// GetFirst returns the first element in the list.
// If the list is empty, b return false
func (l *DoublyLinkedList[T]) GetFirst() (t T, b bool) {
	if l.IsEmpty() {
		return
	}
	return l.head.Val, true
}

// GetLast returns the last element in the list.
// If the list is empty, b return false
func (l *DoublyLinkedList[T]) GetLast() (t T, b bool) {
	if l.IsEmpty() {
		return
	}
	return l.tail.Val, true
}

// Get returns the element at the specified position in the list.
// If the index is invalid, b return false
func (l *DoublyLinkedList[T]) Get(index int) (t T, b bool) {
	if l.isInvalidIndex(index) {
		return
	}
	if index == 0 {
		return l.head.Val, true
	}
	if index == l.size-1 {
		return l.tail.Val, true
	}
	node := l.head
	for i := 0; i < index; i, node = i+1, node.Next {
	}
	return node.Val, true
}

// Set sets the element at the specified position in the list.
// If the index is invalid, b return false
func (l *DoublyLinkedList[T]) Set(index int, e T) bool {
	if l.isInvalidIndex(index) {
		return false
	}
	node := l.head
	for i := 0; i < index; i, node = i+1, node.Next {
	}
	node.Val = e
	return true
}

// Insert inserts the specified elements at the specified position in the list.
func (l *DoublyLinkedList[T]) Insert(index int, elements ...T) bool {
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
	for _, e := range elements {
		node := &doublyNode[T]{Val: e, Prev: prev}
		prev.Next = node
		prev = node
		l.size++
	}
	prev.Next = oldNext
	return true
}

// RemoveFirst removes the first element from the list.
// If the list is empty, b return false
func (l *DoublyLinkedList[T]) RemoveFirst() (t T, b bool) {
	if l.IsEmpty() {
		return
	}
	head := l.head
	l.head = head.Next
	l.size--
	if l.IsEmpty() {
		l.tail = nil
	} else {
		l.head.Prev = nil
	}
	return head.Val, true
}

// RemoveLast removes the last element from the list.
// If the list is empty, b return false
func (l *DoublyLinkedList[T]) RemoveLast() (t T, b bool) {
	if l.IsEmpty() {
		return
	}
	if l.Size() == 1 {
		return l.RemoveFirst()
	}
	tail := l.tail
	l.tail = tail.Prev
	l.tail.Next = nil
	l.size--
	return tail.Val, true
}

// Remove removes the element at the specified position in the list.
// If the index is invalid, b return false
func (l *DoublyLinkedList[T]) Remove(index int) (t T, b bool) {
	if l.isInvalidIndex(index) {
		return
	}
	if index == 0 {
		return l.RemoveFirst()
	}
	if index == l.Size()-1 {
		return l.RemoveLast()
	}
	prev := l.head
	for i := 0; i < index-1; i, prev = i+1, prev.Next {
	}
	dst := prev.Next
	prev.Next = dst.Next
	dst.Next.Prev = prev
	l.size--
	t, dst = dst.Val, nil
	return t, true
}

// IsEmpty checks whether the list is empty
func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

// Size returns the size of the list
func (l *DoublyLinkedList[T]) Size() int {
	return l.size
}

// Clear removes all the elements from the list
func (l *DoublyLinkedList[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

// ToSlice returns a slice containing all the elements in this list.
func (l *DoublyLinkedList[T]) ToSlice() []T {
	elements := make([]T, 0, l.Size())
	node := l.head
	for node != nil {
		elements = append(elements, node.Val)
		node = node.Next
	}
	return elements
}

// isInvalidIndex checks whether the index is invalid
func (l *DoublyLinkedList[T]) isInvalidIndex(index int) bool {
	return index < 0 || index > l.Size()-1
}
