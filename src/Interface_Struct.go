package src

import "fmt"

//type Lock interface {
//	Lock() uint16
//	UnLock(id uint16)
//}
//
//type RepreLocker struct {
//}
//
//func (r *RepreLocker) Lock() uint16 {
//	fmt.Println("资源被上锁")
//	return 1
//}
//
//func (r *RepreLocker) UnLock() uint16 {
//	fmt.Println("资源未被上锁")
//	return 1
//}

type Animal interface {
	SaySomething()
}

type Cat struct {
}

func (c Cat) SaySomething() {
	fmt.Println("miao miao miao")
}

type Dog struct {
}

func (d Dog) SaySomething() {
	fmt.Println("wang wang wang")
}

type Kun struct {
}

func (k Kun) SaySomething() {
	fmt.Println("只因 只因 只因")
}
