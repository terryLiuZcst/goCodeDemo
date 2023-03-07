package src

import (
	"fmt"
	"unsafe"
)

func DemoSwitch(i int) {
	fmt.Println("====================switch demo======================")
	//基于类型构建的switch语句
	switch unsafe.Sizeof(i) {
	case 8:
		fmt.Println("this is ", i)
	default:
		fmt.Println("this is no int")
	}
}
