package parse

import "strings"

type Node struct {
	operation string
	files     []string
	children  []*Node
}

func Parse(s string) *Node {
	if s == "" {
		return nil
	}
	if s[0] != '[' && s[len(s)-1] != ']' {
		return nil
	}

	s = s[1 : len(s)-1]

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
	subArgsStringsToParse := strings.Split(args, "[ ")
	if len(subArgsStringsToParse) == 1 {
		files := strings.Split(args, " ")
		for _, f := range files {
			if f != "" {
				n.files = append(n.files, f)
			}
		}
	}

	if len(subArgsStringsToParse) > 1 {
		n.children = make([]*Node, 0, len(subArgsStringsToParse))
		for _, s := range subArgsStringsToParse {
			if s == "" {
				continue
			}

			if s[len(s)-1] == ' ' {
				s = s[:len(s)-1]
			}
			child := Parse("[" + s)
			n.children = append(n.children, child)
		}
	}

	return &n
}
