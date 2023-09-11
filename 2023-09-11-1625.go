package main

import (
	"fmt"
)

func main() {
	// 练习1：打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：计算1到100的和并打印结果
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum from 1 to 100:", sum)

	// 练习3：编写一个函数，接收两个整数参数，返回它们的和
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("Addition of 5 and 3:", add(5, 3))

	// 练习4：将字符串反转并打印结果
	str := "Hello, Golang!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed String:", reversedStr)
}