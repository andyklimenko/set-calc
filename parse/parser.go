package parse

import (
	"errors"
	"fmt"
	"strings"
)

const (
	OperationUnary = "unary"
)

var (
	ErrInvalidStringFormat = errors.New("invalid string format")
	ErrNoOpeningBrace      = errors.New("can't find opening brace but closing brace was found")
	ErrNoClosingBrace      = errors.New("can't find closing brace but opening brace was found")
)

type Node struct {
	operation string
	files     []string
	children  []*Node
}

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

func Parse(str string) (*Node, error) {
	if str == "" {
		return nil, nil
	}

	s, err := omitBraces(str)
	if err != nil {
		return nil, err
	}

	if strings.Index(s, " ") == -1 {
		return &Node{ // no spaces -> no operation marker -> unary operation
			operation: OperationUnary,
			files:     []string{s},
		}, nil
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

	n := Node{
		operation: string(op),
	}
	args := s[opLastIdx+1:]
	parsedArgs := splitSubExpressions(args)
	for _, parsedArg := range parsedArgs {
		if parsedArg == "" {
			continue
		}

		if strings.Index(parsedArg, "[") == -1 {
			n.files = strings.Split(parsedArg, " ")
			continue
		}

		child, err := Parse(parsedArg)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}
		n.children = append(n.children, child)
	}

	return &n, nil
}
