package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	// u deals with going across
	// test values are greater than zero
	//	if x > 0 && y > 0 {
	for u := 1; u <= x; u++ {
		for i := 1; i <= y; i++ {
			if u == 1 {
				if i == 1 {
					z01.PrintRune('o')
				} else if i == x {
					z01.PrintRune('o')
					z01.PrintRune('\n')
				} else {
					z01.PrintRune('-')
				}
			}
		}
	}
	z01.PrintRune('\n')
	z01.PrintRune('x')
	z01.PrintRune('y')
}
