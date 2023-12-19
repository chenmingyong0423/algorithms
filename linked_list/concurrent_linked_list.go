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

package linkedlist

import "sync"

type ConcurrentLinkedList[T any] struct {
	list LinkedList[T]
	lock *sync.RWMutex
}

// NewDefaultConcurrentLinkedList returns a new ConcurrentLinkedList with the default SinglyLinkedList.
func NewDefaultConcurrentLinkedList[T any]() *ConcurrentLinkedList[T] {
	return &ConcurrentLinkedList[T]{
		list: NewSinglyLinkedList[T](),
		lock: &sync.RWMutex{},
	}
}

func NewConcurrentLinkedList[T any](list LinkedList[T]) *ConcurrentLinkedList[T] {
	return &ConcurrentLinkedList[T]{
		list: list,
		lock: &sync.RWMutex{},
	}
}

func (l *ConcurrentLinkedList[T]) Add(elements ...T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.list.Add(elements...)
}

func (l *ConcurrentLinkedList[T]) Append(elements ...T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.list.Append(elements...)
}

func (l *ConcurrentLinkedList[T]) Prepend(elements ...T) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.list.Prepend(elements...)
}

func (l *ConcurrentLinkedList[T]) GetFirst() (T, bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.list.GetFirst()
}

func (l *ConcurrentLinkedList[T]) GetLast() (T, bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.list.GetLast()
}

func (l *ConcurrentLinkedList[T]) Get(index int) (T, bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.list.Get(index)
}

func (l *ConcurrentLinkedList[T]) Set(index int, e T) bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.list.Set(index, e)
}

func (l *ConcurrentLinkedList[T]) Insert(index int, elements ...T) bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.list.Insert(index, elements...)
}

func (l *ConcurrentLinkedList[T]) RemoveFirst() (T, bool) {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.list.RemoveFirst()
}

func (l *ConcurrentLinkedList[T]) RemoveLast() (T, bool) {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.list.RemoveLast()
}

func (l *ConcurrentLinkedList[T]) Remove(index int) (T, bool) {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.list.Remove(index)
}

func (l *ConcurrentLinkedList[T]) IsEmpty() bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.list.IsEmpty()
}

func (l *ConcurrentLinkedList[T]) Size() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.list.Size()
}

func (l *ConcurrentLinkedList[T]) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.list.Clear()
}

func (l *ConcurrentLinkedList[T]) ToSlice() []T {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return l.list.ToSlice()
}

func (l *ConcurrentLinkedList[T]) Reverse() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.list.Reverse()
}
