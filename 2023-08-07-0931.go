package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 练习1：生成5个随机数并打印出来
	for i := 0; i < 5; i++ {
		num := rand.Intn(100)
		fmt.Println(num)
	}

	// 练习2：实现一个阶乘函数
	num := 5
	fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
