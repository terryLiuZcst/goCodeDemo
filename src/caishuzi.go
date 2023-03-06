package src

import (
	"fmt"
	"math/rand"
)

func CaiShuZi() {
	a := rand.Intn(100)
	var b int
	for i := 7; i >= 1; i-- {
		fmt.Println("请输入你猜的数字(1到一百)，您还有", i, "次机会")
		fmt.Scan(&b)
		if a == b {
			fmt.Println("cxk,ji ni tai mei")
			return
		} else if b > a {
			fmt.Println("大了")
		} else {
			fmt.Println("小了")
		}
	}
	fmt.Println("机会已用完，谢谢参与，正确答案是:", a)

}
