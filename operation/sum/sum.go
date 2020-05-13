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

	var args [][]int
	var length int
	for _, arg := range u.args {
		nums := arg.Resolve()
		length += len(nums)
		args = append(args, nums)
	}

	uniqueElements := make([]int, 0, length)
	for i, arg := range args {
		if i == 0 {
			uniqueElements = append(uniqueElements, arg...)
			continue
		}

		m := map[int]bool{}
		for _, u := range uniqueElements {
			m[u] = true
		}
		for _, a := range arg {
			_, contains := m[a]
			if !contains {
				uniqueElements = append(uniqueElements, a)
			}
		}
	}

	return uniqueElements
}

func New(args []operation.Resolvable) *Union {
	return &Union{args: args}
}
