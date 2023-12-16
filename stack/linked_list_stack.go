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

var _ Stack[any] = (*LinkedListStack[any])(nil)

type node[T any] struct {
	Val  T
	Next *node[T]
}

type LinkedListStack[T any] struct {
	top  *node[T]
	size int
}

func NewLinkedListStack[T any]() *LinkedListStack[T] {
	return &LinkedListStack[T]{}
}

func (l *LinkedListStack[T]) Push(e T) {
	n := &node[T]{Val: e}
	n.Next = l.top
	l.top = n
	l.size++
}

func (l *LinkedListStack[T]) Pop() (t T, b bool) {
	if !l.IsEmpty() {
		t, b = l.top.Val, true
		l.top = l.top.Next
		l.size--
	}
	return
}

func (l *LinkedListStack[T]) Peek() (t T, b bool) {
	if !l.IsEmpty() {
		t, b = l.top.Val, true
	}
	return
}

func (l *LinkedListStack[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedListStack[T]) Size() int {
	return l.size
}

func (l *LinkedListStack[T]) toSlice() (elements []T) {
	cur := l.top
	for cur != nil {
		elements = append(elements, cur.Val)
		cur = cur.Next
	}
	return
}
