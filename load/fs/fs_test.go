package fs

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	f, err := ioutil.TempFile("", "")
	require.NoError(t, err)

	defer os.Remove(f.Name())

	_, err = f.WriteString("123\n")
	require.NoError(t, err)
	_, err = f.WriteString("100500\n")
	require.NoError(t, err)
	_, err = f.WriteString("987654321\n")
	require.NoError(t, err)

	require.NoError(t, f.Sync())

	loader := New(f.Name())
	numbers, err := loader.Load()
	require.NoError(t, err)
	assert.Equal(t, []int{123, 100500, 987654321}, numbers)
}
