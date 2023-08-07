package main

import "fmt"

func main() {
	// 练习1: 输出 Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个数的和并输出结果
	num1 := 10
	num2 := 5
	sum := num1 + num2
	fmt.Println("The sum is:", sum)

	// 练习3: 判断一个数是否为偶数，并输出结果
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is an even number.")
	} else {
		fmt.Println(num, "is an odd number.")
	}

	// 练习4: 使用循环输出 1 到 10
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// 练习5: 定义一个切片，存储数值，并输出切片内容
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println("Numbers in the slice:", numbers)

	// 练习6: 使用函数计算两个数的乘积
	product := multiply(3, 4)
	fmt.Println("The product is:", product)
}

func multiply(a, b int) int {
	return a * b
}