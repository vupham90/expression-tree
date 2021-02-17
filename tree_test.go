package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfixToTree_SimpleCase(t *testing.T) {
	node, err := infixToTree([]string{"Country", "=", "\"VN\"", "AND", "age", ">", "12"})

	assert.Nil(t, err)
	assert.NotNil(t, node)
	assert.Equal(t, "((Country = \"VN\") AND (age > 12))", node.ToQuery())
}

func TestInfixToTree_SimpleCaseWithNOT(t *testing.T) {
	node, err := infixToTree([]string{"Country", "=", "\"VN\"", "AND", "NOT", "(", "age", ">", "12", ")"})

	assert.Nil(t, err)
	assert.NotNil(t, node)
	assert.Equal(t, "((Country = \"VN\") AND NOT ((age > 12)))", node.ToQuery())
}

func TestInfixToTree_SegmentWithNOT(t *testing.T) {
	node, err := infixToTree([]string{"Country", "=", "\"VN\"", "AND", "NOT", "(", "female", ")"})

	assert.Nil(t, err)
	assert.NotNil(t, node)
	assert.Equal(t, "((Country = \"VN\") AND NOT (female))", node.ToQuery())
}

func TestInfixToTree_ComplexCase(t *testing.T) {
	node, err := infixToTree([]string{"(", "Country", "=", "\"VN\"", "AND", "age", ">", "12", ")", "OR", "Female", "=", "true"})

	assert.Nil(t, err)
	assert.NotNil(t, node)
	assert.Equal(t, "(((Country = \"VN\") AND (age > 12)) OR (Female = true))", node.ToQuery())
}

func TestInfixToTree_ComplexCaseWithNOTAndSegment(t *testing.T) {
	node, err := infixToTree([]string{"Region", "=", "\"HCM\"", "AND", "(", "(", "Country", "=", "\"VN\"", "AND", "age", ">", "12", ")", "OR", "Female", "=", "true", ")", "AND", "(", "NOT", "(", "NewUser", ")", "AND", "balance", ">", "100", ")"})

	assert.Nil(t, err)
	assert.NotNil(t, node)
	assert.Equal(t, "((Region = \"HCM\") AND ((((Country = \"VN\") AND (age > 12)) OR (Female = true)) AND (NOT (NewUser) AND (balance > 100))))", node.ToQuery())
}

func TestInfixToTree_InvalidOperator(t *testing.T) {
	node, err := infixToTree([]string{"Region", "=", "\"HCM\"", "AND", "AND"})

	assert.EqualError(t, err, "invalid expression for AND")
	assert.Nil(t, node)
}

func TestInfixToTree_UnclosedParenthesis(t *testing.T) {
	node, err := infixToTree([]string{"Region", "=", "\"HCM\"", "AND", "(", "(", "name", "=", `"Vu"`, ")"})

	assert.EqualError(t, err, "unbalance parenthesis")
	assert.Nil(t, node)
}

func TestInfixToTree_UnbalanceParenthesis(t *testing.T) {
	node, err := infixToTree([]string{"Region", "=", "\"HCM\"", "AND", "(", "name", "=", `"Vu"`, ")", ")"})

	assert.EqualError(t, err, "unbalance parenthesis")
	assert.Nil(t, node)
}

func TestInfixToTree_InvalidValue(t *testing.T) {
	node, err := infixToTree([]string{"Region", "=", "\"HCM\"", "&&", "(", "name", "=", `"Vu"`, ")"})

	assert.EqualError(t, err, "invalid expression")
	assert.Nil(t, node)
}