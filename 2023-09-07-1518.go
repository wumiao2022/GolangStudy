package main

import "fmt"

func main() {
	// 练习1：打印出Hello World
	fmt.Println("Hello World")

	// 练习2：计算两个数的和并打印结果
	num1 := 5
	num2 := 8
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：判断一个数是否为偶数并打印结果
	number := 9
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	// 练习4：使用for循环打印数字1到10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5：使用while循环打印数字10到1
	j := 10
	for j >= 1 {
		fmt.Println(j)
		j--
	}
}