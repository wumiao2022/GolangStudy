package main

import "fmt"

func main() {
	// 1. 输出Hello World!
	fmt.Println("Hello World!")

	// 2. 打印1到10的整数
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 3. 计算两个数的和并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 4. 判断一个数是否为偶数并输出结果
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 5. 调用函数计算两个数的积并输出结果
	product := multiply(5, 6)
	fmt.Println("Product:", product)
}

func multiply(a, b int) int {
	return a * b
}