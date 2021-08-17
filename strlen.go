package piscine

func StrLen(s string) int {
	// slice := []rune(s) // convert string to runes
	/*sum := 0
	for index, _ := range slice {
		sum += int(index)
	}
	*/
	return len([]rune(s))
}
