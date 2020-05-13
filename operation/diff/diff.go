package diff

import (
	"github.com/andyklimenko/set-calc/operation"
)

type XOR struct {
	args []operation.Resolvable
}

func (x *XOR) Resolve() []int {
	if len(x.args) == 0 {
		return []int{}
	}

	if len(x.args) == 1 {
		return x.args[0].Resolve()
	}

	var args [][]int
	for _, arg := range x.args {
		nums := arg.Resolve()
		args = append(args, nums)
	}

	uniqueElements := map[int]bool{}
	for _, arg := range args {
		if len(arg) == 0 {
			continue
		}
		if len(uniqueElements) == 0 {
			for _, a := range arg {
				uniqueElements[a] = true
			}
			continue
		}

		m := map[int]bool{}
		for k := range uniqueElements {
			m[k] = true
		}
		for _, a := range arg {
			_, contains := m[a]
			if contains {
				delete(uniqueElements, a)
				continue
			}

			uniqueElements[a] = true
		}
	}

	res := make([]int, 0, len(uniqueElements))
	for k := range uniqueElements {
		res = append(res, k)
	}
	return res
}

func New(args []operation.Resolvable) *XOR {
	return &XOR{args: args}
}
