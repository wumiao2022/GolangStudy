package main

import (
	"fmt"
)

func main() {
	// 练习1：输出"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数的和，并输出结果
	a := 10
	b := 5
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：判断一个整数是否为偶数，并输出结果
	num := 16
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：计算一个浮点数的平方，并输出结果
	floatNum := 3.14
	square := floatNum * floatNum
	fmt.Println("Square:", square)

	// 练习5：将一个字符串翻转，并输出结果
	str := "Hello, World!"
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	fmt.Println("Reversed String:", reversedStr)
}
