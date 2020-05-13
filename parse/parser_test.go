package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	assert.Nil(t, Parse(""))

	n := Parse("[DIF a.txt b.txt c.txt]")
	assert.Equal(t, "DIF", n.operation)
	assert.Equal(t, []string{"a.txt", "b.txt", "c.txt"}, n.files)
	assert.Nil(t, n.children)
}

func TestComplexParse(t *testing.T) {
	n := Parse("[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]")
	assert.Equal(t, "SUM", n.operation)
	require.Len(t, n.children, 2)

	assert.Equal(t, "DIF", n.children[0].operation)
	assert.Equal(t, []string{"a.txt", "b.txt", "c.txt"}, n.children[0].files)
	assert.Equal(t, "INT", n.children[1].operation)
	assert.Equal(t, []string{"b.txt", "c.txt"}, n.children[1].files)
}
