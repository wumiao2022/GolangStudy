package main

import "fmt"

func main() {
	// 实例1： Hello, World!
	fmt.Println("Hello, World!")

	// 实例2： 输出整数运算结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	sub := num1 - num2
	mul := num1 * num2
	div := num1 / num2
	fmt.Println("Sum:", sum)
	fmt.Println("Subtraction:", sub)
	fmt.Println("Multiplication:", mul)
	fmt.Println("Division:", div)

	// 实例3： 判断奇偶数
	number := 13
	if number%2 == 0 {
		fmt.Println(number, "is even.")
	} else {
		fmt.Println(number, "is odd.")
	}

	// 实例4： 打印1到10的乘法表
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// 实例5： 求两个数的最大公约数
	num3 := 12
	num4 := 20
	var gcd int
	for i := 1; i <= num3 && i <= num4; i++ {
		if num3%i == 0 && num4%i == 0 {
			gcd = i
		}
	}
	fmt.Println("GCD of", num3, "and", num4, "is", gcd)
}
