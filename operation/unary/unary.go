package unary

type Unary struct {
	nums []int
}

func (u *Unary) Resolve() []int {
	if len(u.nums) == 0 {
		return []int{}
	}
	return u.nums
}

func New(nums []int) *Unary {
	return &Unary{
		nums: nums,
	}
}
