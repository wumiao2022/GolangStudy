package main

import "fmt"

func main() {
	// 实例1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 实例2：计算两个数的和并输出
	a := 5
	b := 7
	sum := a + b
	fmt.Println("The sum of", a, "and", b, "is", sum)

	// 实例3：判断一个数是否为奇数
	num := 13
	if num%2 == 1 {
		fmt.Println(num, "is an odd number")
	} else {
		fmt.Println(num, "is an even number")
	}

	// 实例4：循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 实例5：使用递归计算阶乘
	num = 6
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", num, "is", factorial)
}