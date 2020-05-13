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

	var args [][]int
	var length int
	for _, arg := range i.args {
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

		m := map[int]uniqueElement{}
		for i, u := range uniqueElements {
			m[u] = uniqueElement{index: i, unique: true}
		}
		for _, a := range arg {
			_, contains := m[a]
			if !contains {
				continue
			}

			m[a] = uniqueElement{index: i, unique: false}
		}

		newUniqueElements := make([]int, 0, length)
		for k, v := range m {
			if !v.unique {
				newUniqueElements = append(newUniqueElements, k)
			}
		}
		uniqueElements = newUniqueElements
	}
	return uniqueElements
}

func New(args []operation.Resolvable) *Intersection {
	return &Intersection{args: args}
}
