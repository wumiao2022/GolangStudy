package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 打印 1-10
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	// 打印 10-1
	for i := 10; i >= 1; i-- {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	// 打印 0-100之间所有偶数
	for i := 0; i <= 100; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	// 打印一个 5 * 5 的正方形
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	// 生成并打印一个随机数
	randomNum := rand.Intn(100)
	fmt.Printf("随机数是：%d\n", randomNum)
}