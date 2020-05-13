package unary

type Unary struct {
	nums []int
}

func (u *Unary) Resolve() []int {
	return u.nums
}

func New(nums []int) *Unary {
	return &Unary{
		nums: nums,
	}
}
