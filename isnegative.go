package piscine

// package main

import "github.com/01-edu/z01"

// by Yousuf isnegative.go Quest 02
// Write a function that prints
// 'T' (true) on a single line if the int passed as
// parameter is negative, otherwise it prints 'F' (false).

func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
	// return nb
}

/*
func main() {
	IsNegative(-157637090634009893)
	IsNegative(0)
	IsNegative(-1)
}
*/
