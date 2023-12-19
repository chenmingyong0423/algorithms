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

import linkedlist "github.com/chenmingyong0423/algorithms/linked_list"

var _ Stack[any] = (*LinkedListStack[any])(nil)

type LinkedListStack[T any] struct {
	list linkedlist.LinkedList[T]
}

func NewLinkedListStack[T any]() *LinkedListStack[T] {
	return &LinkedListStack[T]{
		list: linkedlist.NewSinglyLinkedList[T](),
	}
}

func (l *LinkedListStack[T]) Push(e T) {
	l.list.Add(e)
}

func (l *LinkedListStack[T]) Pop() (t T, b bool) {
	return l.list.RemoveLast()
}

func (l *LinkedListStack[T]) Peek() (t T, b bool) {
	return l.list.GetLast()
}

func (l *LinkedListStack[T]) IsEmpty() bool {
	return l.list.IsEmpty()
}

func (l *LinkedListStack[T]) Size() int {
	return l.list.Size()
}

func (l *LinkedListStack[T]) toSlice() (elements []T) {
	return l.list.Values()
}
