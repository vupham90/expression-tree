package tree

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindStatement_SimpleCase(t *testing.T) {
	exp, err := InfixToExp(`Country = "VN" AND age > 12 OR (NOT(Female) AND GrabPayUser)`)

	assert.Nil(t, err)
	assert.NotNil(t, exp)

	var n *Node
	n = exp.FindStatement("Country")
	assert.Equal(t, "(Country = \"VN\")", n.ToQuery())

	n = exp.FindStatement("age")
	assert.Equal(t, "(age > 12)", n.ToQuery())

	n = exp.FindStatement("Female")
	assert.Equal(t, "NOT (Female)", n.ToQuery())

	n = exp.FindStatement("GrabPayUser")
	assert.Equal(t, "GrabPayUser", n.ToQuery())

	n = exp.FindStatement("NahNah")
	assert.Nil(t, n)
}