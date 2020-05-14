package parse

import (
	"errors"
	"fmt"
	"strings"

	"github.com/andyklimenko/set-calc/load/fs"
	"github.com/andyklimenko/set-calc/operation"
	"github.com/andyklimenko/set-calc/operation/unary"
)

var (
	ErrInvalidStringFormat = errors.New("invalid string format")
	ErrNoOpeningBrace      = errors.New("can't find opening brace but closing brace was found")
	ErrNoClosingBrace      = errors.New("can't find closing brace but opening brace was found")
)

func omitBraces(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	startIdx := strings.Index(s, "[")
	endIdx := strings.LastIndex(s, "]")

	if startIdx == -1 && endIdx == -1 {
		return s, nil
	}
	if startIdx == -1 && endIdx != -1 {
		return "", ErrNoOpeningBrace
	}
	if startIdx != -1 && endIdx == -1 {
		return "", ErrNoClosingBrace
	}
	if startIdx > endIdx {
		return "", ErrInvalidStringFormat
	}

	s = s[startIdx+1 : endIdx]

	for i, c := range s {
		startIdx = i
		if c == ' ' {
			continue
		}

		break
	}

	for i := len(s) - 1; i >= 0; i-- {
		endIdx = i
		if s[i] == ' ' {
			continue
		}

		break
	}

	return s[startIdx : endIdx+1], nil
}

func splitSubExpressions(s string) []string {
	var res []string
	unprocessed := s

	var globalEndIndex int
	for unprocessed != "]" && unprocessed != "" {
		var beginIndicators, endIndicators int
		var beginIndex, endIndex int
		for i, c := range unprocessed {
			if c == '[' {
				if beginIndicators == 0 {
					beginIndex = i
				}

				if beginIndex > 0 {
					endIndex = beginIndex - 1
					for i := endIndex; i > 0; i-- {
						if unprocessed[i] == ' ' {
							endIndex--
						}
					}
					globalEndIndex += endIndex
					beginIndex = 0
					break
				}

				beginIndicators++
			}
			if c == ']' {
				endIndicators++
				if endIndicators == beginIndicators {
					endIndex = i
					globalEndIndex += i
					break
				}
			}
		}

		if beginIndex == 0 && endIndex == 0 {
			res = append(res, unprocessed)
			unprocessed = ""
			continue
		}
		res = append(res, unprocessed[beginIndex:endIndex+1])
		unprocessed = s[globalEndIndex+1:]
	}
	return res
}

func newUnary(fileName string) (*unary.Unary, error) {
	l := fs.New(fileName)
	nums, err := l.Load()
	if err != nil {
		return nil, fmt.Errorf("loading numbers from %s: %w", fileName, err)
	}

	return unary.New(nums), nil
}

func Parse(str string) (operation.Resolvable, error) {
	if str == "" {
		return nil, nil
	}

	s, err := omitBraces(str)
	if err != nil {
		return nil, err
	}

	if strings.Index(s, " ") == -1 {
		return newUnary(s) // no spaces -> no operation marker -> unary operation
	}

	var op []byte
	var opLastIdx int
	for i, c := range s {
		if i != 0 && c == ' ' {
			opLastIdx = i
			break
		}

		if c != ' ' {
			op = append(op, byte(c))
		}
	}

	if !supportOperation(string(op)) {
		return nil, ErrOperationNotSupported
	}

	args := s[opLastIdx+1:]
	parsedArgs := splitSubExpressions(args)
	operationArgs := make([]operation.Resolvable, 0, len(parsedArgs))
	for _, parsedArg := range parsedArgs {
		if strings.Index(parsedArg, "[") == -1 {
			files := strings.Split(parsedArg, " ")
			for _, f := range files {
				a, err := newUnary(f)
				if err != nil {
					return nil, err
				}
				operationArgs = append(operationArgs, a)
			}
			continue
		}

		child, err := Parse(parsedArg)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		operationArgs = append(operationArgs, child)
	}

	return newOperation(string(op), operationArgs...), nil
}
