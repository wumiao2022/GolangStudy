package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：整数加法
	a := 5
	b := 10
	sum := a + b
	fmt.Println("Sum:", sum)

	// 练习3：字符串拼接
	str1 := "Hello"
	str2 := "Golang"
	result := str1 + " " + str2
	fmt.Println("Result:", result)

	// 练习4：判断奇偶
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习5：循环输出
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}