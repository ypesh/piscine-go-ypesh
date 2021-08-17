package main

import (
	"fmt"
	"piscine"
)

func main() {
	l := piscine.StrLen("Hello World!")
	fmt.Println(l)
}

/*
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Hello, 世界", len("世界"), utf8.RuneCountInString("世界"))
}
*/
