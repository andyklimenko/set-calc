package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/andyklimenko/set-calc/parse"
)

func main() {
	flag.Parse()
	args := flag.Args()

	var expr string
	for _, a := range args {
		expr += " " + a
	}
	if expr == "" {
		fmt.Println("can't parse empty expression")
		return
	}

	r, err := parse.Parse(expr)
	if err != nil {
		fmt.Println(err)
		return
	}

	res := r.Resolve()
	sort.Ints(res)
	for _, n := range res {
		fmt.Println(n)
	}
}
