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

package stack

type Slice[T any] struct {
	elements []T
}

// NewStackSlice returns a new stack
func NewStackSlice[T any]() *Slice[T] {
	return &Slice[T]{}
}

// NewStackSliceWithSize returns a new stack with size
func NewStackSliceWithSize[T any](size int) *Slice[T] {
	return &Slice[T]{
		elements: make([]T, 0, size),
	}
}

// Push pushes an element onto the top of the stack
func (s *Slice[T]) Push(e T) {
	s.elements = append(s.elements, e)
}

// Pop removes the element at the top of the stack and returns that element
func (s *Slice[T]) Pop() (t T) {
	if !s.IsEmpty() {
		lastIdx := len(s.elements) - 1
		t = s.elements[lastIdx]
		s.elements = s.elements[:lastIdx]
	}
	return t
}

// Peek returns the element at the top of the stack
func (s *Slice[T]) Peek() (t T) {
	if !s.IsEmpty() {
		lastIdx := len(s.elements) - 1
		t = s.elements[lastIdx]
	}
	return t
}

// IsEmpty checks whether the stack is empty
func (s *Slice[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size returns the size of the stack
func (s *Slice[T]) Size() int {
	return len(s.elements)
}
