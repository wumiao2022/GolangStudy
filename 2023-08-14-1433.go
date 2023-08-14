package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Golang每日多练")

	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：计算1+2+...+100并打印结果
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Printf("1+2+...+100 = %d\n", sum)

	// 练习3：输出当前时间
	currentTime := time.Now()
	fmt.Println("当前时间：", currentTime)

	// 练习4：判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Printf("%d是偶数\n", num)
	} else {
		fmt.Printf("%d是奇数\n", num)
	}

	// 练习5：遍历切片并打印元素
	slice := []int{1, 2, 3, 4, 5}
	for _, num := range slice {
		fmt.Println(num)
	}
}
