package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 生成随机数
	randNum := rand.Intn(100) + 1
	fmt.Println("生成的随机数:", randNum)

	// 判断随机数的奇偶性
	if randNum%2 == 0 {
		fmt.Println("随机数是偶数")
	} else {
		fmt.Println("随机数是奇数")
	}

	// 输出10次Hello Golang
	for i := 0; i < 10; i++ {
		fmt.Println("Hello Golang")
	}

	// 遍历打印1到100之间的数字
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}