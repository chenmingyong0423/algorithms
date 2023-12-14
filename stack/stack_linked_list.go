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

type node[T any] struct {
	Val  T
	Next *node[T]
}

type LinkedList[T any] struct {
	top  *node[T]
	size int
}

func NewStackLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Push(e T) {
	n := &node[T]{Val: e}
	n.Next = l.top
	l.top = n
	l.size++
}

func (l *LinkedList[T]) Pop() (t T) {
	if !l.IsEmpty() {
		t = l.top.Val
		l.top = l.top.Next
		l.size--
	}
	return t
}

func (l *LinkedList[T]) Peek() (t T) {
	if !l.IsEmpty() {
		t = l.top.Val
	}
	return t
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) toSlice() (elements []T) {
	cur := l.top
	for cur != nil {
		elements = append(elements, cur.Val)
		cur = cur.Next
	}
	return
}
