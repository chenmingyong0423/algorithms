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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStackLinkedList(t *testing.T) {
	stackLinkedList := NewLinkedListStack[int]()
	assert.NotNil(t, stackLinkedList)
}

func TestLinkedList_Push(t *testing.T) {
	testCases := []struct {
		name          string
		stack         *LinkedListStack[int]
		element       int
		stackElements []int
	}{
		{
			name:    "push element to empty stack",
			stack:   NewLinkedListStack[int](),
			element: 1,
			stackElements: []int{
				1,
			},
		},
		{
			name: "push element to non-empty stack",
			stack: func() *LinkedListStack[int] {
				s := NewLinkedListStack[int]()
				s.Push(1)
				return s
			}(),
			element: 1,
			stackElements: []int{
				1,
				1,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.stack.Push(tc.element)
			assert.ElementsMatch(t, tc.stackElements, tc.stack.toSlice())
		})
	}
}

func TestLinkedList_Pop(t *testing.T) {
	testCases := []struct {
		name          string
		stack         *LinkedListStack[int]
		stackElements []int
		wantValue     int
		wantBool      bool
	}{
		{
			name: "pop element from empty stack",
			stack: func() *LinkedListStack[int] {
				s := NewLinkedListStack[int]()
				return s
			}(),
			stackElements: []int{},
			wantValue:     0,
			wantBool:      false,
		},
		{
			name: "pop element from the stack with one element",
			stack: func() *LinkedListStack[int] {
				s := NewLinkedListStack[int]()
				s.Push(1)
				return s
			}(),
			stackElements: []int{},
			wantValue:     1,
			wantBool:      true,
		},
		{
			name: "pop element from stack with more than one element",
			stack: func() *LinkedListStack[int] {
				s := NewLinkedListStack[int]()
				s.Push(1)
				s.Push(2)
				return s
			}(),
			stackElements: []int{
				1,
			},
			wantValue: 2,
			wantBool:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			element, b := tc.stack.Pop()
			assert.ElementsMatch(t, tc.stackElements, tc.stack.toSlice())
			assert.Equal(t, tc.wantValue, element)
			assert.Equal(t, tc.wantBool, b)
		})
	}
}

func TestLinkedList_Peek(t *testing.T) {
	testCases := []struct {
		name      string
		stack     *LinkedListStack[int]
		wantValue int
		wantBool  bool
	}{
		{
			name: "peek element from empty stack",
			stack: func() *LinkedListStack[int] {
				s := NewLinkedListStack[int]()
				return s
			}(),
			wantValue: 0,
			wantBool:  false,
		},
		{
			name: "peek element from the stack with one element",
			stack: func() *LinkedListStack[int] {
				s := NewLinkedListStack[int]()
				s.Push(1)
				return s
			}(),
			wantValue: 1,
			wantBool:  true,
		},
		{
			name: "peek element from stack with more than one element",
			stack: func() *LinkedListStack[int] {
				s := NewLinkedListStack[int]()
				s.Push(1)
				s.Push(2)
				return s
			}(),
			wantValue: 2,
			wantBool:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			element, b := tc.stack.Peek()
			assert.Equal(t, tc.wantValue, element)
			assert.Equal(t, tc.wantBool, b)
		})
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	testCases := []struct {
		name  string
		stack *LinkedListStack[int]

		want bool
	}{
		{
			name:  "check empty stack",
			stack: NewLinkedListStack[int](),
			want:  true,
		},
		{
			name: "check non-empty stack",
			stack: func() *LinkedListStack[int] {
				s := NewLinkedListStack[int]()
				s.Push(1)
				return s
			}(),

			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, tc.stack.IsEmpty())
		})
	}
}

func TestLinkedList_Size(t *testing.T) {
	testCases := []struct {
		name  string
		stack *LinkedListStack[int]

		wantSize int
	}{
		{
			name:     "empty stack",
			stack:    NewLinkedListStack[int](),
			wantSize: 0,
		},
		{
			name: "non-empty stack",
			stack: func() *LinkedListStack[int] {
				s := NewLinkedListStack[int]()
				s.Push(1)
				return s
			}(),

			wantSize: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantSize, tc.stack.Size())
		})
	}
}
