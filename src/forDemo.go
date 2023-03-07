package src

import "fmt"

func ForDemo(c rune) { // print a-z
	for c <= 122 {
		fmt.Printf("%c  ", c)
		c++
	}
	fmt.Printf("\n")
}

func Demo_Print_Z_A(c rune) {
	for c >= 65 {
		fmt.Printf("%c  ", c)
		c--
	}
}

func Deme_Range(s string) {
	for i, x := range s {
		fmt.Printf("%d", i)
		fmt.Printf("%c ", x)
	}
}
