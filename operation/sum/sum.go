package sum

import (
	"github.com/andyklimenko/set-calc/operation"
)

type Union struct {
	args []operation.Resolvable
}

func (u *Union) Resolve() []int {
	if len(u.args) == 0 {
		return []int{}
	}

	m := map[int]bool{}
	for _, arg := range u.args {
		nums := arg.Resolve()
		for _, n := range nums {
			m[n] = true
		}
	}

	uniqueElements := make([]int, 0, len(m))
	for k, _ := range m {
		uniqueElements = append(uniqueElements, k)
	}

	return uniqueElements
}

func New(args []operation.Resolvable) *Union {
	return &Union{args: args}
}
