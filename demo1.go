package main

import (
	"./src"
	"fmt"
	"time"
)

func main() {
	fmt.Println("cxk")
	//var score int
	//fmt.Println("please input your score")
	////fmt.Scanln(&score)
	//fmt.Scanf("%d", &score)
	//switch score { //switch后面有一个布尔值变量默认是true
	//case 90:
	//	fmt.Println("sing")
	//	fallthrough
	//case 70, 80: //这里的case可以case多个值
	//	if score != 70 && score != 80 {
	//		fmt.Println("he can't jump")
	//		break
	//	}
	//	fmt.Println("jump")
	//case 60:
	//	fmt.Println("rap")
	//default:
	//	fmt.Println("basketball")
	//}
	//for i := 1; i <= 9; i++ { //打印99乘法表
	//	for j := 1; j <= i; j++ {
	//		fmt.Print(j, "*", i, "=", i*j, "\t")
	//	}
	//	fmt.Println()
	//}

	//for i := 0; i < 8; i++ {
	//
	//	if i == 5 {
	//		//break
	//		//continue
	//	}
	//	fmt.Println(i)
	//}
	time.Sleep(1000 * time.Millisecond)
	src.DateDemo()

	//src.WeekDayDemo(src.DateDemo())
	//输入一个日期并显示是星期几
	//var time1 time.Time = time.Date(2000, time.August, 12, 0, 0, 0, 0, time.Local)
	//src.WeekDayDemo(time1)
	//var x int = 7
	//src.AddOne(x)
	//fmt.Println("调用非指针函数后还是x", x)
	//src.AddOneDemo(&x)
	//fmt.Println("调用指针函数后还是x?", x)

	//var d src.Demo = src.Demo{Tag: "1234"}
	//fmt.Println(d.Tag)
	//src.SetDemo(&d, "uuuu")
	//fmt.Println(d.Tag)

	//var a, _ = src.Demo1()
	//fmt.Println(a)

	//力扣第一题
	//a1 := [4]int{1, 5, 6, 9}
	//var tag int
	//fmt.Println("please input a int")
	//fmt.Scan(&tag)
	//fmt.Printf("%v", leetcode.TwoNumsSum(a1, tag))

	src.CaiShuZi()
}
