package main

import (
	"fmt"
)

func addTen(b int) int {
	b = b + 10
	return b
}

func main() {
	a := 10
	a = addTen(a)
	fmt.Println(a)
}
