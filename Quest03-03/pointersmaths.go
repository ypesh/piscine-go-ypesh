package main

import (
	"fmt"
)

func addTenPointer(b *int) {
	*b = *b + 10
}

func main() {
	a := 10
	addTenPointer(&a)
	fmt.Println(a)
}
