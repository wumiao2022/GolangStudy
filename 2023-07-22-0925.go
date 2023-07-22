package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")
	
	// 练习2：计算两个整数的和
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)
	
	// 练习3：判断一个数是否是偶数
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is an even number")
	} else {
		fmt.Println(num, "is an odd number")
	}
	
	// 练习4：遍历切片，并打印每个元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}
	
	// 练习5：计算一个列表中所有元素的和
	numbers := []int{1, 2, 3, 4, 5}
	result := 0
	for _, num := range numbers {
		result += num
	}
	fmt.Println("Sum:", result)
	
	// 练习6：定义一个函数，实现两个数的交换
	swap := func(a, b *int) {
		temp := *a
		*a = *b
		*b = temp
	}
	num1 := 10
	num2 := 20
	fmt.Println("Before swap:", num1, num2)
	swap(&num1, &num2)
	fmt.Println("After swap:", num1, num2)
	
	// 练习7：实现一个简单的计算器，接收两个数和一个操作符，返回计算结果
	calculator := func(num1, num2 int, operator string) int {
		result := 0
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			result = num1 / num2
		}
		return result
	}
	num1 := 10
	num2 := 5
	operator := "*"
	fmt.Println("Result:", calculator(num1, num2, operator))
}