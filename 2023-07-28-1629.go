package main

import "fmt"

func main() {
	// 练习1：输出Hello World
	fmt.Println("Hello World")

	// 练习2：求两个数之和并输出结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是奇数还是偶数
	num := 7
	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习4：计算1到100的和并输出结果
	total := 0
	for i := 1; i <= 100; i++ {
		total += i
	}
	fmt.Println("Total:", total)
	
	// 练习5：实现一个阶乘函数并输出结果
	num := 5
	factorial := 1
	for i := 1; i <= num; i++ {
		factorial *= i
	}
	fmt.Println("Factorial:", factorial)
}
