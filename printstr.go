package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	slice := []rune(s) // convert string to runes

	for _, character := range slice {
		z01.PrintRune(character)
	}
}
