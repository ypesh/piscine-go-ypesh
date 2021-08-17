package piscine

import "github.com/01-edu/z01"

func BasicAtoi(s string) int {
	// return len([]rune(s))
	//r _, c := range s {
	//	z01.PrintRune(c)
	// return len([]rune(s))
	slice := []rune(s)
	for _, character := range slice {
		z01.PrintRune(character)
	}
	return 0

}
