package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：打印1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习3：计算两个数的和并打印结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习4：判断一个整数是否为偶数并打印结果
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习5：计算一个整数的阶乘并打印结果
	factorial := 1
	n := 5
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	fmt.Println("Factorial of", n, "is", factorial)
}

注意：以上代码仅为示例练习，并没有实际运行效果。请将代码复制到你的本地编译器或Go Playground中运行。