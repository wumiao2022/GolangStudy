package main

import "fmt"

func main() {
	// 练习1：打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2：计算和打印两个整数的和
	var num1, num2 int = 5, 10
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3：计算并打印从1加到100的和
	total := 0
	for i := 1; i <= 100; i++ {
		total += i
	}
	fmt.Println("Total:", total)
}

// 练习4: 编写一个函数，接受一个字符串作为参数，返回该字符串的反转结果
func reverseString(str string) string {
	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

// 练习5: 编写一个函数，接受一个整数切片作为参数，返回切片元素的总和
func sumSlice(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	return sum
}