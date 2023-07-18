package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")
	
	// 练习2：输出0到9的数字
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	
	// 练习3：计算两个数的和
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)
	
	// 练习4：判断一个数是否为偶数
	num := 6
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
	
	// 练习5：输出斐波那契数列前10项
	n := 10
	a, b := 0, 1
	fmt.Print("Fibonacci sequence: ")
	for i := 1; i <= n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
	
	// 练习6：计算n的阶乘
	n = 5
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println(n, "factorial:", factorial)
	
	// 练习7：输出一个5行5列的星号图案
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}