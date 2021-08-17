package main

// by Yousuf
// ascii code 33 to 126
// youtube video use 02-02 or 02-03 and 02-04
import (
	"github.com/01-edu/z01"
)

func main() {
	var i int32 = 97
	// z01.PrintRune('a')
	// z01.PrintRune('\n')

	for i <= 122 {

		z01.PrintRune(i)
		i++
	}
	z01.PrintRune('\n')
}
