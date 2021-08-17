package piscine

func StrRev(s string) string {
	/*
		dash := ""
		for _, ma := range s {
			dash = string(ma) + dash
		}
		return dash
		I don't understand how this works so I won't use it
	*/
	r := []rune(s)
	for x, y := 0, len(r)-1; x < len(r)/2; x, y = x+1, y-1 {
		r[x], r[y] = r[y], r[x]
	}
	return string(r)
}
