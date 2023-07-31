package main

import "fmt"

func main() {
	// 练习1：打印Hello World
	fmt.Println("Hello World")

	// 练习2：计算两个整数的和并打印结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数并打印结果
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4：使用for循环打印1到10的数字
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用for循环计算1到100的和并打印结果
	sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of 1 to 100:", sum)
}