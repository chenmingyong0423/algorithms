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

func TestStackSlice(t *testing.T) {
	stackSlice := NewStackSlice[int]()
	assert.True(t, stackSlice.IsEmpty())

	stackSlice.Push(1)
	assert.Equal(t, 1, stackSlice.Size())

	assert.Equal(t, 1, stackSlice.Peek())
	assert.Equal(t, 1, stackSlice.Pop())

	// test pop empty stack
	assert.Equal(t, 0, stackSlice.Pop())

	stackSliceWithSize := NewStackSliceWithSize[int](16)
	assert.Equal(t, 0, stackSliceWithSize.Size())
}
