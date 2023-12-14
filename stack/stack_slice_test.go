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

func TestSlice_Push(t *testing.T) {
	testCases := []struct {
		name          string
		stack         *Slice[int]
		element       int
		stackElements []int
	}{
		{
			name:    "push element to empty stack",
			stack:   NewStackSlice[int](),
			element: 1,
			stackElements: []int{
				1,
			},
		},
		{
			name: "push element to non-empty stack",
			stack: func() *Slice[int] {
				s := NewStackSlice[int]()
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
			assert.ElementsMatch(t, tc.stackElements, tc.stack.elements)
		})
	}
}

func TestSlice_Pop(t *testing.T) {
	testCases := []struct {
		name          string
		stack         *Slice[int]
		stackElements []int
		want          int
	}{
		{
			name: "pop element from empty stack",
			stack: func() *Slice[int] {
				s := NewStackSlice[int]()
				return s
			}(),
			stackElements: []int{},
			want:          0,
		},
		{
			name: "pop element from non-empty stack",
			stack: func() *Slice[int] {
				s := NewStackSlice[int]()
				s.Push(1)
				return s
			}(),
			stackElements: []int{},
			want:          1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.stack.Pop()
			assert.ElementsMatch(t, tc.stackElements, tc.stack.elements)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSlice_Peek(t *testing.T) {
	testCases := []struct {
		name  string
		stack *Slice[int]
		want  int
	}{
		{
			name: "peek element from non-empty stack",
			stack: func() *Slice[int] {
				s := NewStackSlice[int]()
				s.Push(1)
				return s
			}(),
			want: 1,
		},
		{
			name: "peek element from empty stack",
			stack: func() *Slice[int] {
				s := NewStackSlice[int]()
				return s
			}(),
			want: 0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.stack.Peek()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSlice_IsEmpty(t *testing.T) {
	testCases := []struct {
		name  string
		stack *Slice[int]
		want  bool
	}{
		{
			name: "check empty stack",
			stack: func() *Slice[int] {
				s := NewStackSlice[int]()
				return s
			}(),
			want: true,
		},
		{
			name: "check non-empty stack",
			stack: func() *Slice[int] {
				s := NewStackSlice[int]()
				s.Push(1)
				return s
			}(),
			want: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.stack.IsEmpty()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSlice_Size(t *testing.T) {
	testCases := []struct {
		name  string
		stack *Slice[int]
		want  int
	}{
		{
			name: "empty stack",
			stack: func() *Slice[int] {
				s := NewStackSlice[int]()
				return s
			}(),
			want: 0,
		},
		{
			name: "non-empty stack",
			stack: func() *Slice[int] {
				s := NewStackSlice[int]()
				s.Push(1)
				return s
			}(),
			want: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.stack.Size()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestNewStackSliceWithSize(t *testing.T) {
	s := NewStackSliceWithSize[int](10)
	assert.Equal(t, 10, cap(s.elements))
}
