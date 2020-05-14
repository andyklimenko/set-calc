package parse

import (
	"errors"
	"strings"

	"github.com/andyklimenko/set-calc/operation"
	"github.com/andyklimenko/set-calc/operation/diff"
	"github.com/andyklimenko/set-calc/operation/intersection"
	"github.com/andyklimenko/set-calc/operation/sum"
)

var (
	ErrOperationNotSupported = errors.New("operation not supported")
)

func supportOperation(s string) bool {
	o := operation.Operation(strings.ToUpper(s))
	return o == operation.Sum || o == operation.Int || o == operation.Dif
}

func newOperation(op string, args ...operation.Resolvable) operation.Resolvable {
	switch operation.Operation(strings.ToUpper(op)) {
	case operation.Sum:
		return sum.New(args)
	case operation.Dif:
		return diff.New(args)
	case operation.Int:
		return intersection.New(args)
	}
	return nil
}
