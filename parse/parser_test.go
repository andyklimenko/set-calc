package parse

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupFile(nums ...string) (string, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return "", err
	}

	for _, n := range nums {
		if _, err := f.WriteString(fmt.Sprintf("%s\n", n)); err != nil {
			return "", err
		}
	}

	return f.Name(), f.Sync()
}

func TestParseEmpty(t *testing.T) {
	t.Parallel()
	n, err := Parse("")
	require.NoError(t, err)
	assert.Nil(t, n)
}

func TestParseSimpleExp(t *testing.T) {
	t.Parallel()
	a, err := setupFile("1", "2", "3")
	require.NoError(t, err)
	defer os.Remove(a)

	b, err := setupFile("2", "3", "4")
	require.NoError(t, err)
	defer os.Remove(b)

	c, err := setupFile("3", "4", "5")
	require.NoError(t, err)
	defer os.Remove(c)

	rootExpr, err := Parse(fmt.Sprintf("[DIF %s %s %s]", a, b, c))
	require.NoError(t, err)

	res := rootExpr.Resolve()
	sort.Ints(res)
	assert.Equal(t, []int{1, 3, 5}, res)
}

func TestParseComplexExpr(t *testing.T) {
	t.Parallel()
	a, err := setupFile("1", "2", "3")
	require.NoError(t, err)
	defer os.Remove(a)

	b, err := setupFile("2", "3", "4")
	require.NoError(t, err)
	defer os.Remove(b)

	c, err := setupFile("3", "4", "5")
	require.NoError(t, err)
	defer os.Remove(c)

	line := fmt.Sprintf("[ SUM [ DIF %s %s %s ] [ INT %s %s ] ]", a, b, c, b, c)
	rootExpr, err := Parse(line)
	require.NoError(t, err)

	res := rootExpr.Resolve()
	sort.Ints(res)

	assert.Equal(t, []int{1, 3, 4, 5}, res)
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
	t.Parallel()
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
	t.Parallel()
	a, err := setupFile("1", "2", "3")
	require.NoError(t, err)
	defer os.Remove(a)

	b, err := setupFile("2", "3", "4")
	require.NoError(t, err)
	defer os.Remove(b)

	c, err := setupFile("3", "4", "5")
	require.NoError(t, err)
	defer os.Remove(c)

	d, err := setupFile("4", "5", "6")
	require.NoError(t, err)
	defer os.Remove(b)

	e, err := setupFile("5", "6", "7")
	require.NoError(t, err)
	defer os.Remove(c)

	line := fmt.Sprintf("[ SUM [ DIF %s [ SUM %s %s ] ] [ INT %s %s ] ]", a, b, c, d, e)
	rootExpr, err := Parse(line)
	assert.NoError(t, err)

	res := rootExpr.Resolve()
	sort.Ints(res)

	assert.Equal(t, []int{1, 4, 5, 6}, res)

}
