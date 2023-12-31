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

type Stack[T any] interface {
	Push(e T)
	// Pop removes the element at the top of the stack and returns that element
	// If the stack is empty, b return false
	Pop() (T, bool)
	// Peek returns the element at the top of the stack
	// If the stack is empty, b return false
	Peek() (T, bool)
	IsEmpty() bool
	Size() int
}
