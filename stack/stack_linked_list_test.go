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
	stackLinkedList := NewStackLinkedList[int]()
	assert.NotNil(t, stackLinkedList)
}

func TestLinkedList_Push(t *testing.T) {
	testCases := []struct {
		name          string
		stack         *LinkedList[int]
		element       int
		stackElements []int
	}{
		{
			name:    "push element to empty stack",
			stack:   NewStackLinkedList[int](),
			element: 1,
			stackElements: []int{
				1,
			},
		},
		{
			name: "push element to non-empty stack",
			stack: func() *LinkedList[int] {
				s := NewStackLinkedList[int]()
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
		name            string
		stack           *LinkedList[int]
		stackElements   []int
		expectedElement int
	}{
		{
			name: "pop element from empty stack",
			stack: func() *LinkedList[int] {
				s := NewStackLinkedList[int]()
				return s
			}(),
			stackElements:   []int{},
			expectedElement: 0,
		},
		{
			name: "pop element from the stack with one element",
			stack: func() *LinkedList[int] {
				s := NewStackLinkedList[int]()
				s.Push(1)
				return s
			}(),
			stackElements:   []int{},
			expectedElement: 1,
		},
		{
			name: "pop element from stack with more than one element",
			stack: func() *LinkedList[int] {
				s := NewStackLinkedList[int]()
				s.Push(1)
				s.Push(2)
				return s
			}(),
			stackElements: []int{
				1,
			},
			expectedElement: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			element := tc.stack.Pop()
			assert.ElementsMatch(t, tc.stackElements, tc.stack.toSlice())
			assert.Equal(t, tc.expectedElement, element)
		})
	}
}

func TestLinkedList_Peek(t *testing.T) {
	testCases := []struct {
		name            string
		stack           *LinkedList[int]
		expectedElement int
	}{
		{
			name: "peek element from empty stack",
			stack: func() *LinkedList[int] {
				s := NewStackLinkedList[int]()
				return s
			}(),
			expectedElement: 0,
		},
		{
			name: "peek element from the stack with one element",
			stack: func() *LinkedList[int] {
				s := NewStackLinkedList[int]()
				s.Push(1)
				return s
			}(),
			expectedElement: 1,
		},
		{
			name: "peek element from stack with more than one element",
			stack: func() *LinkedList[int] {
				s := NewStackLinkedList[int]()
				s.Push(1)
				s.Push(2)
				return s
			}(),
			expectedElement: 2,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			element := tc.stack.Peek()
			assert.Equal(t, tc.expectedElement, element)
		})
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	testCases := []struct {
		name  string
		stack *LinkedList[int]

		wantBool bool
	}{
		{
			name:     "empty stack",
			stack:    NewStackLinkedList[int](),
			wantBool: true,
		},
		{
			name: "non-empty stack",
			stack: func() *LinkedList[int] {
				s := NewStackLinkedList[int]()
				s.Push(1)
				return s
			}(),

			wantBool: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.wantBool, tc.stack.IsEmpty())
		})
	}
}

func TestLinkedList_Size(t *testing.T) {
	testCases := []struct {
		name  string
		stack *LinkedList[int]

		wantSize int
	}{
		{
			name:     "empty stack",
			stack:    NewStackLinkedList[int](),
			wantSize: 0,
		},
		{
			name: "non-empty stack",
			stack: func() *LinkedList[int] {
				s := NewStackLinkedList[int]()
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
