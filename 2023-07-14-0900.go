package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 生成一个随机数并打印
	randomNum := rand.Intn(100)
	fmt.Println("Random number:", randomNum)

	// 打印1到100之间的偶数
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	// 打印斐波那契数列的前20个数字
	fmt.Println("Fibonacci sequence:")
	fibonacci := make([]int, 20)
	fibonacci[0] = 0
	fibonacci[1] = 1
	for i := 2; i < 20; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}
	fmt.Println(fibonacci)

	// 打印字符串反转
	str := "Hello, world!"
	reversedStr := ""
	for _, ch := range str {
		reversedStr = string(ch) + reversedStr
	}
	fmt.Println("Reversed string:", reversedStr)
}