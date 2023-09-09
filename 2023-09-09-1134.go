package main

import (
	"fmt"
	"time"
)

func main() {
	// 练习1：输出Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算1到10的和
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为素数
	num := 17
	isPrime := true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Printf("%d is prime: %v\n", num, isPrime)

	// 练习4：获取当前时间
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)

	// 练习5：生成随机数
	randomNum := 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		randomNum = rand.Intn(100)
		fmt.Println("Random number:", randomNum)
	}
}