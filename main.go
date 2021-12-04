package main

import (
	"aoc21/pkg/solution"
	"fmt"
	"os"
	"reflect"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("go run *.go day part fileName")
		fmt.Println("example: go run *.go 1 1 f1_1")
		return
	}

	d := os.Args[1]
	p := os.Args[2]
	fn := os.Args[3]

	s := &solution.Solution{}
	v := reflect.ValueOf(s)
	m := v.MethodByName(fmt.Sprintf("Day%sPart%s", d, p))

	in := []reflect.Value{reflect.ValueOf(fn)}
	out := m.Call(in)
	for _, v := range out {
		fmt.Println(v)
	}
}
