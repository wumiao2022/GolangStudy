package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 打印当前时间
	fmt.Println(time.Now())

	// 2. 计算1到100的和
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 3. 判断一个数是否是偶数
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 4. 输出斐波那契数列（前10个数）
	n := 10
	fmt.Print("Fibonacci series: ")
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}

	// 5. 将字符串反转
	str := "Hello, Golang!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("\nReversed string:", reversedStr)
}