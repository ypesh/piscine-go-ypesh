package piscine

func UltimateDivMod(a *int, b *int) {
	div := *a / *b
	modx := *a % *b
	*a = div
	*b = modx
}
