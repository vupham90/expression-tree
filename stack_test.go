package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func TestStringStack(t *testing.T) {
	stk := &stringStack{}

	stk.push("1")
	stk.push("2")

	assert.Equal(t, 2, stk.size())
	assert.Equal(t, "2", stk.pop())

	assert.Equal(t, 1, stk.size())
	assert.Equal(t, "1", stk.pop())
	
	assert.Equal(t, 0, stk.size())
}
