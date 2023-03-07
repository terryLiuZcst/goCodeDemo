package src

import "fmt"

var src = [4]int{1, 2, 3, 4}

var S1 = src[0:2]
var S2 = src[1:4]

func SliceDemo() {
	fmt.Println(S1)
	fmt.Println(S2)
	fmt.Println(S1[0] == src[0])
}
