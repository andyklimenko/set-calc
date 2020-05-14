package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSimpleExp(t *testing.T) {
	n, err := Parse("")
	require.NoError(t, err)
	assert.Nil(t, n)

	root, err := Parse("[DIF a.txt b.txt c.txt]")
	require.NoError(t, err)
	assert.Equal(t, "DIF", root.operation)
	assert.Equal(t, []string{"a.txt", "b.txt", "c.txt"}, root.files)
	assert.Nil(t, root.children)
}

func TestParseComplexExpr(t *testing.T) {
	root, err := Parse("[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]")
	require.NoError(t, err)
	assert.Equal(t, "SUM", root.operation)
	require.Len(t, root.children, 2)

	assert.Equal(t, "DIF", root.children[0].operation)
	assert.Equal(t, []string{"a.txt", "b.txt", "c.txt"}, root.children[0].files)
	assert.Equal(t, "INT", root.children[1].operation)
	assert.Equal(t, []string{"b.txt", "c.txt"}, root.children[1].files)
}

func TestOmitBraces(t *testing.T) {
	type testCase struct {
		name           string
		str            string
		expectedResult string
		expectedError  error
	}

	t.Parallel()
	testCases := []testCase{
		{name: "empty line", str: "", expectedResult: ""},
		{name: "invalid format", str: "]invalid format[", expectedError: ErrInvalidStringFormat},
		{name: "missing opening brace", str: "no opening brace]", expectedError: ErrNoOpeningBrace},
		{name: "missing closing brace", str: "[no closing brace", expectedError: ErrNoClosingBrace},
		{name: "no braces", str: "no braces", expectedResult: "no braces"},
		{name: "space at start", str: "[ space at start]", expectedResult: "space at start"},
		{name: "space at the end", str: "[space at the end ]", expectedResult: "space at the end"},
		{name: "spaces everywhere", str: "[   many spaces everywhere       ]", expectedResult: "many spaces everywhere"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := omitBraces(tc.str)
			if tc.expectedError != nil {
				assert.EqualError(t, err, tc.expectedError.Error())
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, s)
		})
	}
}

func TestSplit(t *testing.T) {
	s := splitSubExpressions("[ab[c]][[c]d]")
	require.Len(t, s, 2)
	assert.Equal(t, []string{"[ab[c]]", "[[c]d]"}, s)

	noBraces1, err := omitBraces(s[0])
	require.NoError(t, err)
	assert.Equal(t, []string{"ab", "[c]"}, splitSubExpressions(noBraces1))

	noBraces2, err := omitBraces(s[1])
	require.NoError(t, err)
	assert.Equal(t, []string{"[c]", "d"}, splitSubExpressions(noBraces2))
}

func TestParseNestedExpr(t *testing.T) {
	root, err := Parse("[ SUM [ DIF a.txt [ SUM b.txt c.txt ] ] [ INT d.txt e.txt ] ]")
	assert.NoError(t, err)
	assert.Equal(t, "SUM", root.operation)
	require.Len(t, root.children, 2)
	require.Equal(t, "DIF", root.children[0].operation)
	require.Len(t, root.children[0].children, 1)

	require.Equal(t, []string{"a.txt"}, root.children[0].files)

	require.Equal(t, "INT", root.children[1].operation)
	require.Equal(t, []string{"d.txt", "e.txt"}, root.children[1].files)
}
