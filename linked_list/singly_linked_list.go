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

package linked_list

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
func (s *SinglyLinkedList[T]) Add(elements ...T) {
	if len(elements) > 0 {
		for _, e := range elements {
			node := &singlyNode[T]{Val: e}
			if s.IsEmpty() {
				s.head, s.tail = node, node
			} else {
				s.tail.Next = node
				s.tail = node
			}
			s.size++
		}
	}
}

// Append appends the specified elements to the end of the list.(same as Add)
func (s *SinglyLinkedList[T]) Append(elements ...T) {
	s.Add(elements...)
}

// Prepend prepends the specified elements to the beginning of the list.
func (s *SinglyLinkedList[T]) Prepend(elements ...T) {
	// reverse the elements. i.e. original elements: [2, 3], prepend elements: [0, 1], result: [0, 1, 2, 3]
	for i := len(elements) - 1; i >= 0; i-- {
		node := &singlyNode[T]{Val: elements[i], Next: s.head}
		s.head = node
		if s.size == 0 {
			s.tail = node
		}
		s.size++
	}
}

// GetFirst returns the first element in the list.
// If the list is empty, b return false
func (s *SinglyLinkedList[T]) GetFirst() (t T, b bool) {
	if s.IsEmpty() {
		return
	}
	return s.head.Val, true
}

// GetLast returns the last element in the list.
// If the list is empty, b return false
func (s *SinglyLinkedList[T]) GetLast() (t T, b bool) {
	if s.IsEmpty() {
		return
	}
	return s.tail.Val, true
}

// Get returns the element at the specified position in the list.
// If the index is invalid, b return false
func (s *SinglyLinkedList[T]) Get(index int) (t T, b bool) {
	if s.isInvalidIndex(index) {
		return
	}
	node := s.head
	for i := 0; i < index; i, node = i+1, node.Next {
	}
	return node.Val, true
}

// Set sets the element at the specified position in the list.
// If the index is invalid, b return false
func (s *SinglyLinkedList[T]) Set(index int, e T) bool {
	if index < 0 || index > s.Size()-1 {
		return false
	}
	node := s.head
	for i := 0; i < index; i, node = i+1, node.Next {
	}
	node.Val = e
	return true
}

// Insert inserts the specified elements at the specified position in the list.
func (s *SinglyLinkedList[T]) Insert(index int, e ...T) bool {
	if s.isInvalidIndex(index) {
		if index == 0 {
			s.Add(e...)
			return true
		}
		return false
	}

	if index == 0 {
		s.Prepend(e...)
		return true
	}
	if index == s.Size()-1 {
		s.Append(e...)
		return true
	}
	prev := s.head
	for i := 0; i < index-1; i, prev = i+1, prev.Next {
	}
	oldNext := prev.Next
	for _, val := range e {
		node := &singlyNode[T]{Val: val}
		prev.Next = node
		prev = node
		s.size++
	}
	prev.Next = oldNext
	return true
}

// RemoveFirst removes the first element from the list.
// If the list is empty, b return false
func (s *SinglyLinkedList[T]) RemoveFirst() (t T, b bool) {
	if s.IsEmpty() {
		return
	}
	node := s.head
	s.head = node.Next
	s.size--
	if s.IsEmpty() {
		s.tail = nil
	}
	return node.Val, true
}

func (s *SinglyLinkedList[T]) RemoveLast() (t T, b bool) {
	if s.IsEmpty() {
		return
	}
	if s.Size() == 1 {
		return s.RemoveFirst()
	}
	prev := s.head
	for prev.Next != s.tail {
		prev = prev.Next
	}
	node := prev.Next
	prev.Next = nil
	s.tail = prev
	s.size--
	return node.Val, true
}

// Remove removes the element at the specified position in the list.
// If the index is invalid, b return false
func (s *SinglyLinkedList[T]) Remove(index int) (t T, b bool) {
	if s.isInvalidIndex(index) {
		return
	}
	if s.Size() == 1 {
		return s.RemoveFirst()
	}
	if index == s.Size()-1 {
		return s.RemoveLast()
	}

	prev := s.head
	// find the previous node of the node to be deleted
	for i := 0; i < index-1; i, prev = i+1, prev.Next {
	}

	node := prev.Next
	prev.Next = node.Next
	s.size--
	return node.Val, true
}

// IsEmpty checks whether the list is empty
func (s *SinglyLinkedList[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Size returns the size of the list
func (s *SinglyLinkedList[T]) Size() int {
	return s.size
}

// isInvalidIndex checks whether the index is invalid
func (s *SinglyLinkedList[T]) isInvalidIndex(index int) bool {
	return index < 0 || index > s.Size()-1
}

// Clear removes all the elements from the list
func (s *SinglyLinkedList[T]) Clear() {
	s.head = nil
	s.tail = nil
	s.size = 0
}

func (s *SinglyLinkedList[T]) ToSlice() []T {
	elements := make([]T, 0, s.Size())
	node := s.head
	for node != nil {
		elements = append(elements, node.Val)
		node = node.Next
	}
	return elements
}
