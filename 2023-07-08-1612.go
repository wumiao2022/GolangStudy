package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// 随机生成一个整数
	num := rand.Intn(100)
	fmt.Println("生成的随机数是:", num)

	// 判断这个数字是奇数还是偶数
	if num%2 == 0 {
		fmt.Println("这个数字是偶数")
	} else {
		fmt.Println("这个数字是奇数")
	}

	// 使用循环输出100以内的所有偶数
	fmt.Println("100以内的所有偶数:")
	for i := 0; i <= 100; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 使用循环求1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("1到100的和是:", sum)
}
