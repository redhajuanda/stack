package stack

import "testing"

import "github.com/stretchr/testify/assert"

func TestPush(t *testing.T) {
	var stack Stack
	stack.Push(1)
	assert.Equal(t, len(stack), 1)
}

func TestPop(t *testing.T) {
	var stack Stack
	stack.Pop()
	assert.Equal(t, len(stack), 0)
	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	assert.Equal(t, len(stack), 1)
	assert.Equal(t, stack[0], 1)
	stack.Pop()
	assert.Equal(t, len(stack), 0)
}

func TestTop(t *testing.T) {
	var stack Stack
	stack.Push(1)
	stack.Push(2)
	top := stack.Top()
	assert.Equal(t, top, 2)
	stack.Pop()
	top = stack.Top()
	assert.Equal(t, top, 1)
}

func TestEmpty(t *testing.T) {
	var stack Stack
	empty := stack.Empty()
	assert.Equal(t, empty, true)
	stack.Push(1)
	empty = stack.Empty()
	assert.Equal(t, empty, false)
}
