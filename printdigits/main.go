package main

// by Yousuf
// Write a program that prints the decimal digits in ascending order (from 0 to 9) on a single line.

import (
	"github.com/01-edu/z01"
)

func main() {
	var i int32 = 48
	// z01.PrintRune('a')
	// z01.PrintRune('\n')

	for i <= 57 {

		z01.PrintRune(i)
		i++
	}
	z01.PrintRune('\n')
}
