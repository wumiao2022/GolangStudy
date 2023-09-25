package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：求两个数的和
	num1 := 15
	num2 := 10
	sum := num1 + num2
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)

	// 练习3：判断一个数是否为奇数
	num := 7
	if num%2 == 0 {
		fmt.Printf("%d is an even number\n", num)
	} else {
		fmt.Printf("%d is an odd number\n", num)
	}

	// 练习4：计算n的阶乘
	n := 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Printf("The factorial of %d is %d\n", n, factorial)
}