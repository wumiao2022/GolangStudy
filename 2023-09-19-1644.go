package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成随机数
	num1 := rand.Intn(10)
	num2 := rand.Intn(10)

	// 判断两个数之间的关系
	if num1 > num2 {
		fmt.Printf("%d 大于 %d\n", num1, num2)
	} else if num1 < num2 {
		fmt.Printf("%d 小于 %d\n", num1, num2)
	} else {
		fmt.Printf("%d 等于 %d\n", num1, num2)
	}

	// 使用for循环打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d   ", j, i, i*j)
		}
		fmt.Println()
	}
}