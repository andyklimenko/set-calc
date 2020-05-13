package int

import (
	"github.com/andyklimenko/set-calc/operation"
)

type Intersection struct {
	args []operation.Resolvable
}

type uniqueElement struct {
	index  int
	unique bool
}

func (i *Intersection) Resolve() []int {
	if len(i.args) == 0 {
		return []int{}
	}
	if len(i.args) == 1 {
		return i.args[0].Resolve()
	}

	var args [][]int
	var length int
	for _, arg := range i.args {
		nums := arg.Resolve()
		length += len(nums)
		args = append(args, nums)
	}

	uniqueElements := map[int]bool{}
	for i, arg := range args {
		if i == 0 {
			for _, a := range arg {
				uniqueElements[a] = true
			}
			continue
		}

		m := map[int]uniqueElement{}
		for k, _ := range uniqueElements {
			m[k] = uniqueElement{index: i, unique: true}
		}
		for _, a := range arg {
			_, contains := m[a]
			if !contains {
				continue
			}

			m[a] = uniqueElement{index: i, unique: false}
		}

		for k, v := range m {
			if v.unique {
				delete(uniqueElements, k)
			}
		}
	}

	res := make([]int, 0, len(uniqueElements))
	for k, _ := range uniqueElements {
		res = append(res, k)
	}
	return res
}

func New(args []operation.Resolvable) *Intersection {
	return &Intersection{args: args}
}
