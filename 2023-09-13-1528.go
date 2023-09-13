package main

import (
	"fmt"
)

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算两个整数之和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, sum)

	// 练习3：判断一个整数是否为偶数
	num := 7
	if num%2 == 0 {
		fmt.Printf("%d is an even number\n", num)
	} else {
		fmt.Printf("%d is an odd number\n", num)
	}

	// 练习4：输出1到10的所有奇数
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}

	// 练习5：计算阶乘
	num = 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Printf("The factorial of %d is: %d\n", num, factorial)
}
