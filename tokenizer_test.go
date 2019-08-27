package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenize_Simple(t *testing.T) {
	str := `Country = "VN" AND age > 12`
	tokens := tokenize(str)
	assert.Equal(t, []string{"Country", "=", "\"VN\"", "AND", "age", ">", "12"}, tokens)
}

func TestTokenize_Complex(t *testing.T) {
	str := `(Country = "VN" AND age > 12) OR Female=true`
	tokens := tokenize(str)
	assert.Equal(t, []string{"(", "Country", "=", "\"VN\"", "AND", "age", ">", "12", ")", "OR", "Female", "=", "true"}, tokens)
}