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

type LinkedList[T any] interface {
	Add(elements ...T)
	Append(elements ...T)
	Prepend(elements ...T)
	Get(index int) (T, bool)
	GetFirst() (T, bool)
	GetLast() (T, bool)
	Set(index int, e T) bool
	Insert(index int, e ...T) bool
	RemoveFirst() (T, bool)
	RemoveLast() (T, bool)
	Remove(index int) (T, bool)
	IsEmpty() bool
	Size() int
	Clear()
	ToSlice() []T
}
