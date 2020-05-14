package sum

import (
	"sort"
	"testing"

	"github.com/andyklimenko/set-calc/operation"
	"github.com/andyklimenko/set-calc/operation/unary"
	"github.com/stretchr/testify/assert"
)

func TestUnions(t *testing.T) {
	type testCase struct {
		name string
		args []operation.Resolvable
		want []int
	}

	t.Parallel()
	empty := unary.New(nil)
	testCases := []testCase{
		{name: "nil", args: nil, want: []int{}},
		{name: "[]", args: []operation.Resolvable{empty}, want: []int{}},
		{name: "[1, 2, 3]", args: []operation.Resolvable{unary.New([]int{1, 2, 3})}, want: []int{1, 2, 3}},
		{name: "[1, 2, 3]SUM[]", args: []operation.Resolvable{unary.New([]int{1, 2, 3}), empty}, want: []int{1, 2, 3}},
		{name: "[]SUM[1, 2, 3]", args: []operation.Resolvable{empty, unary.New([]int{1, 2, 3})}, want: []int{1, 2, 3}},
		{name: "[1, 2, 3]SUM[4, 5, 6]", args: []operation.Resolvable{unary.New([]int{1, 2, 3}), unary.New([]int{4, 5, 6})}, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "[1, 2, 3]SUM[2, 3, 4]", args: []operation.Resolvable{unary.New([]int{1, 2, 3}), unary.New([]int{2, 3, 4})}, want: []int{1, 2, 3, 4}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			op := New(tc.args)
			got := op.Resolve()
			sort.Ints(got)
			assert.Equal(t, tc.want, got)
		})
	}
}
