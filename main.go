package main

import (
	"fmt"
	"os"
)

func main() {
	d := os.Args[1]
	p := os.Args[2]

	if len(os.Args) < 3 {
		fmt.Println("go run *.go day part")
		fmt.Println("example: go run *.go 1 1")
		return
	}

	switch d {
	case "1":
		dayOne(p)
	}
}
