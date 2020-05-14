package operation

type Resolvable interface {
	Resolve() []int
}

type Operation string

const (
	Sum Operation = "SUM"
	Int Operation = "INT"
	Dif Operation = "DIF"
)
