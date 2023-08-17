package main

import (
	"fmt"
)

func main() {
	// 练习1：打印 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：求两个数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是正数还是负数
	num := 7
	if num >= 0 {
		fmt.Println(num, "is a positive number")
	} else {
		fmt.Println(num, "is a negative number")
	}

	// 练习4：计算阶乘
	num3 := 5
	fact := 1
	for i := 1; i <= num3; i++ {
		fact *= i
	}
	fmt.Printf("Factorial of %d is %d\n", num3, fact)
}