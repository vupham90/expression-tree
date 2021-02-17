package tree

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFindStatement_SimpleCase(t *testing.T) {
	exp, err := InfixToExp(`Country = "VN" AND age > 12 OR (NOT(Female) AND NewUser)`)

	assert.Nil(t, err)
	assert.NotNil(t, exp)

	var n *Node
	n = exp.FindStatement("Country")
	assert.Equal(t, "(Country = \"VN\")", n.ToQuery())

	n = exp.FindStatement("age")
	assert.Equal(t, "(age > 12)", n.ToQuery())

	n = exp.FindStatement("Female")
	assert.Equal(t, "NOT (Female)", n.ToQuery())

	n = exp.FindStatement("NewUser")
	assert.Equal(t, "NewUser", n.ToQuery())

	n = exp.FindStatement("NahNah")
	assert.Nil(t, n)
}


/**
$ go test -bench=. -benchmem -cpuprofile profile.out -test.benchtime=10s
goos: darwin
goarch: amd64
pkg: github.com/vupham90/expression-tree
BenchmarkInfixToExp-4             500000             24958 ns/op           11960 B/op        192 allocs/op
PASS
ok      github.com/vupham90/expression-tree     29.063s
*/
func BenchmarkInfixToExp(b *testing.B) {
	for i:=0; i< b.N; i++ {
		exp, err := InfixToExp(`Country = "VN" AND age > 12 OR (NOT(Female) AND productID = 1000 OR paymentMethod = "Dash" AND productID = 1000 OR paymentMethod = "Dash" AND productID = 1000 OR paymentMethod = "Dash" AND productID = 1000 OR paymentMethod = "Dash" OR paymentMethod = "Dash" OR paymentMethod = "Dash" OR paymentMethod = "Dash") OR (NOT(Female) AND productID = 1000 OR paymentMethod = "Dash" AND productID = 1000 OR paymentMethod = "Dash" AND productID = 1000 OR paymentMethod = "Dash" AND productID = 1000 OR paymentMethod = "Dash" OR paymentMethod = "Dash" OR paymentMethod = "Dash" OR paymentMethod = "Dash")`)
		_, _ = exp, err
	}
}