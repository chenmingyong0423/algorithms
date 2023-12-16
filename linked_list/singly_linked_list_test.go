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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList_Add(t *testing.T) {
	testCases := []struct {
		name     string
		list     *SinglyLinkedList[int]
		elements []int

		wantListElements []int
	}{
		{
			name:             "add a element to empty list",
			list:             NewSinglyLinkedList[int](),
			elements:         []int{1},
			wantListElements: []int{1},
		},
		{
			name:             "add a element to non-empty list",
			list:             NewSinglyLinkedList[int](1),
			elements:         []int{2},
			wantListElements: []int{1, 2},
		},
		{
			name:             "add multiple elements to empty list",
			list:             NewSinglyLinkedList[int](),
			elements:         []int{1, 2, 3},
			wantListElements: []int{1, 2, 3},
		},
		{
			name:             "add multiple elements to non-empty list",
			list:             NewSinglyLinkedList[int](1, 2, 3),
			elements:         []int{4, 5},
			wantListElements: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.Add(tc.elements...)
			assert.Equal(t, tc.wantListElements, tc.list.ToSlice())
		})
	}
}

func TestSinglyLinkedList_Append(t *testing.T) {
	testCases := []struct {
		name     string
		list     *SinglyLinkedList[int]
		elements []int

		wantListElements []int
	}{
		{
			name:             "append a element to empty list",
			list:             NewSinglyLinkedList[int](),
			elements:         []int{1},
			wantListElements: []int{1},
		},
		{
			name:             "append a element to non-empty list",
			list:             NewSinglyLinkedList[int](1),
			elements:         []int{2},
			wantListElements: []int{1, 2},
		},
		{
			name:             "append multiple elements to empty list",
			list:             NewSinglyLinkedList[int](),
			elements:         []int{1, 2, 3},
			wantListElements: []int{1, 2, 3},
		},
		{
			name:             "append multiple elements to non-empty list",
			list:             NewSinglyLinkedList[int](1, 2, 3),
			elements:         []int{4, 5},
			wantListElements: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.Append(tc.elements...)
			assert.Equal(t, tc.wantListElements, tc.list.ToSlice())
		})
	}
}

func TestSinglyLinkedList_Prepend(t *testing.T) {
	testCases := []struct {
		name     string
		list     *SinglyLinkedList[int]
		elements []int

		wantListElements []int
	}{
		{
			name:             "prepend a element to empty list",
			list:             NewSinglyLinkedList[int](),
			elements:         []int{1},
			wantListElements: []int{1},
		},
		{
			name:             "prepend a element to non-empty list",
			list:             NewSinglyLinkedList[int](2),
			elements:         []int{1},
			wantListElements: []int{1, 2},
		},
		{
			name:             "prepend multiple elements to empty list",
			list:             NewSinglyLinkedList[int](),
			elements:         []int{1, 2, 3},
			wantListElements: []int{1, 2, 3},
		},
		{
			name:             "prepend multiple elements to non-empty list",
			list:             NewSinglyLinkedList[int](4, 5, 6),
			elements:         []int{1, 2, 3},
			wantListElements: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.Prepend(tc.elements...)
			assert.Equal(t, tc.wantListElements, tc.list.ToSlice())
		})
	}
}

func TestSinglyLinkedList_GetFirst(t *testing.T) {
	testCases := []struct {
		name string
		list *SinglyLinkedList[int]

		wantValue int
		wantBool  bool
	}{
		{
			name:      "get first element from empty list",
			list:      NewSinglyLinkedList[int](),
			wantValue: 0,
			wantBool:  false,
		},
		{
			name:      "get first element from non-empty list",
			list:      NewSinglyLinkedList[int](1, 2, 3),
			wantValue: 1,
			wantBool:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			element, b := tc.list.GetFirst()
			assert.Equal(t, tc.wantValue, element)
			assert.Equal(t, tc.wantBool, b)
		})
	}
}

func TestSinglyLinkedList_GetLast(t *testing.T) {
	testCases := []struct {
		name string
		list *SinglyLinkedList[int]

		wantValue int
		wantBool  bool
	}{
		{
			name:      "get last element from empty list",
			list:      NewSinglyLinkedList[int](),
			wantValue: 0,
			wantBool:  false,
		},
		{
			name:      "get last element from non-empty list",
			list:      NewSinglyLinkedList[int](1, 2, 3),
			wantValue: 3,
			wantBool:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			element, b := tc.list.GetLast()
			assert.Equal(t, tc.wantValue, element)
			assert.Equal(t, tc.wantBool, b)
		})
	}
}

func TestSinglyLinkedList_Get(t *testing.T) {
	testCases := []struct {
		name  string
		list  *SinglyLinkedList[int]
		index int

		wantValue int
		wantBool  bool
	}{
		{
			name:      "index is less than zero",
			list:      NewSinglyLinkedList[int](1, 2, 3),
			index:     -1,
			wantValue: 0,
			wantBool:  false,
		},
		{
			name:      "index is greater than size of list",
			list:      NewSinglyLinkedList[int](1, 2, 3),
			index:     3,
			wantValue: 0,
			wantBool:  false,
		},
		{
			name:      "index is equal to size of list with one element",
			list:      NewSinglyLinkedList[int](1),
			index:     0,
			wantValue: 1,
			wantBool:  true,
		},
		{
			name:      "index is equal to size of list with more than one element",
			list:      NewSinglyLinkedList[int](1, 2, 3),
			index:     2,
			wantValue: 3,
			wantBool:  true,
		},
		{
			name:      "index is greater than zero and less than size of list",
			list:      NewSinglyLinkedList[int](1, 2, 3),
			index:     1,
			wantValue: 2,
			wantBool:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			element, b := tc.list.Get(tc.index)
			assert.Equal(t, tc.wantValue, element)
			assert.Equal(t, tc.wantBool, b)
		})
	}
}

func TestSinglyLinkedList_Set(t *testing.T) {
	testCases := []struct {
		name  string
		list  *SinglyLinkedList[int]
		index int
		value int

		want bool
	}{
		{
			name:  "index is less than zero",
			list:  NewSinglyLinkedList[int](1, 2, 3),
			index: -1,
			value: 4,
			want:  false,
		},
		{
			name:  "index is greater than size of list",
			list:  NewSinglyLinkedList[int](1, 2, 3),
			index: 3,
			value: 4,
			want:  false,
		},
		{
			name:  "index is equal to size of list with one element",
			list:  NewSinglyLinkedList[int](1),
			index: 0,
			value: 4,
			want:  true,
		},
		{
			name:  "index is equal to size of list with more than one element",
			list:  NewSinglyLinkedList[int](1, 2, 3),
			index: 2,
			value: 4,
			want:  true,
		},
		{
			name:  "index is greater than zero and less than size of list",
			list:  NewSinglyLinkedList[int](1, 2, 3),
			index: 1,
			value: 4,
			want:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.Set(tc.index, tc.value)
			assert.Equal(t, tc.want, got)
			if got {
				get, b := tc.list.Get(tc.index)
				assert.Equal(t, tc.value, get)
				assert.True(t, b)
			}
		})
	}
}

func TestSinglyLinkedList_RemoveFirst(t *testing.T) {
	testCases := []struct {
		name string
		list *SinglyLinkedList[int]

		wantValue int
		wantBool  bool
	}{
		{
			name:      "remove first element from empty list",
			list:      NewSinglyLinkedList[int](),
			wantValue: 0,
			wantBool:  false,
		},
		{
			name:      "remove first element from the list with one element",
			list:      NewSinglyLinkedList[int](1),
			wantValue: 1,
			wantBool:  true,
		},
		{
			name:      "remove first element from the list with more than one element",
			list:      NewSinglyLinkedList[int](1, 2, 3),
			wantValue: 1,
			wantBool:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			element, b := tc.list.RemoveFirst()
			assert.Equal(t, tc.wantValue, element)
			assert.Equal(t, tc.wantBool, b)
		})
	}
}

func TestSinglyLinkedList_RemoveLast(t *testing.T) {
	testCases := []struct {
		name string
		list *SinglyLinkedList[int]

		wantValue int
		wantBool  bool
	}{
		{
			name:      "remove last element from empty list",
			list:      NewSinglyLinkedList[int](),
			wantValue: 0,
			wantBool:  false,
		},
		{
			name:      "remove last element from the list with one element",
			list:      NewSinglyLinkedList[int](1),
			wantValue: 1,
			wantBool:  true,
		},
		{
			name:      "remove last element from the list with more than one element",
			list:      NewSinglyLinkedList[int](1, 2, 3),
			wantValue: 3,
			wantBool:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			element, b := tc.list.RemoveLast()
			assert.Equal(t, tc.wantValue, element)
			assert.Equal(t, tc.wantBool, b)
		})
	}
}

func TestSinglyLinkedList_Remove(t *testing.T) {
	testCases := []struct {
		name  string
		list  *SinglyLinkedList[int]
		index int

		wantValue int
		wantBool  bool
	}{
		{
			name:      "index is less than zero",
			list:      NewSinglyLinkedList[int](1, 2, 3),
			index:     -1,
			wantValue: 0,
			wantBool:  false,
		},
		{
			name:      "index is greater than size of list",
			list:      NewSinglyLinkedList[int](1, 2, 3),
			index:     3,
			wantValue: 0,
			wantBool:  false,
		},
		{
			name:      "index is equal to size of list with one element",
			list:      NewSinglyLinkedList[int](1),
			index:     0,
			wantValue: 1,
			wantBool:  true,
		},
		{
			name:      "index is equal to size of list with more than one element",
			list:      NewSinglyLinkedList[int](1, 2, 3),
			index:     2,
			wantValue: 3,
			wantBool:  true,
		},
		{
			name:      "index is greater than zero and less than size of list",
			list:      NewSinglyLinkedList[int](1, 2, 3, 4, 5),
			index:     2,
			wantValue: 3,
			wantBool:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			element, b := tc.list.Remove(tc.index)
			assert.Equal(t, tc.wantValue, element)
			assert.Equal(t, tc.wantBool, b)
		})
	}
}

func TestSinglyLinkedList_IsEmpty(t *testing.T) {
	testCases := []struct {
		name string
		list *SinglyLinkedList[int]
		want bool
	}{
		{
			name: "check empty list",
			list: NewSinglyLinkedList[int](),
			want: true,
		},
		{
			name: "check non-empty list",
			list: NewSinglyLinkedList[int](1),
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.IsEmpty()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSinglyLinkedList_Size(t *testing.T) {
	testCases := []struct {
		name string
		list *SinglyLinkedList[int]
		want int
	}{
		{
			name: "empty list",
			list: NewSinglyLinkedList[int](),
			want: 0,
		},
		{
			name: "non-empty list",
			list: NewSinglyLinkedList[int](1),
			want: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.Size()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSinglyLinkedList_Clear(t *testing.T) {
	testCases := []struct {
		name string
		list *SinglyLinkedList[int]
	}{
		{
			name: "empty list",
			list: NewSinglyLinkedList[int](),
		},
		{
			name: "non-empty list",
			list: NewSinglyLinkedList[int](1),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.Clear()
			assert.Equal(t, 0, tc.list.Size())
		})
	}
}

func TestSinglyLinkedList_ToSlice(t *testing.T) {
	testCases := []struct {
		name string
		list *SinglyLinkedList[int]

		want []int
	}{
		{
			name: "empty list",
			list: NewSinglyLinkedList[int](),
			want: []int{},
		},
		{
			name: "non-empty list",
			list: NewSinglyLinkedList[int](1, 2, 3),
			want: []int{1, 2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.ToSlice()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSinglyLinkedList_Insert(t *testing.T) {
	testCases := []struct {
		name     string
		list     *SinglyLinkedList[int]
		index    int
		elements []int

		wantBool         bool
		wantListElements []int
	}{
		{
			name:     "index is less than zero",
			list:     NewSinglyLinkedList[int](1, 2, 5),
			index:    -1,
			elements: []int{3, 4},
			wantBool: false,
		},
		{
			name:     "index is greater than size of list",
			list:     NewSinglyLinkedList[int](1, 2, 3),
			index:    3,
			elements: []int{4, 5},
			wantBool: false,
		},
		{
			name:             "insert one element at the beginning of empty list",
			list:             NewSinglyLinkedList[int](),
			index:            0,
			elements:         []int{1},
			wantBool:         true,
			wantListElements: []int{1},
		},
		{
			name:             "insert one element at the beginning of the list",
			list:             NewSinglyLinkedList[int](2, 3),
			index:            0,
			elements:         []int{1},
			wantBool:         true,
			wantListElements: []int{1, 2, 3},
		},
		{
			name:             "insert multiple elements at the beginning of the list",
			list:             NewSinglyLinkedList[int](3, 4),
			index:            0,
			elements:         []int{1, 2},
			wantBool:         true,
			wantListElements: []int{1, 2, 3, 4},
		},
		{
			name:             "insert one element at the end of the list",
			list:             NewSinglyLinkedList[int](1, 2, 3),
			index:            2,
			elements:         []int{4},
			wantBool:         true,
			wantListElements: []int{1, 2, 3, 4},
		},
		{
			name:             "insert multiple elements at the end of the list",
			list:             NewSinglyLinkedList[int](1, 2, 3),
			index:            2,
			elements:         []int{4, 5},
			wantBool:         true,
			wantListElements: []int{1, 2, 3, 4, 5},
		},
		{
			name:             "insert one element in the middle of the list",
			list:             NewSinglyLinkedList[int](1, 2, 4, 5),
			index:            2,
			elements:         []int{3},
			wantBool:         true,
			wantListElements: []int{1, 2, 3, 4, 5},
		},
		{
			name:             "insert multiple elements in the middle of the list",
			list:             NewSinglyLinkedList[int](1, 2, 5, 6),
			index:            2,
			elements:         []int{3, 4},
			wantBool:         true,
			wantListElements: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.list.Insert(tc.index, tc.elements...)
			assert.Equal(t, tc.wantBool, got)
			if got {
				assert.Equal(t, tc.wantListElements, tc.list.ToSlice())
			}
		})
	}
}
