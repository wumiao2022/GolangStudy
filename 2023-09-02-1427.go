package main

import (
	"fmt"
)

func main() {
	// 练习1：打印"Hello, World!"
	fmt.Println("Hello, World!")

	// 练习2：计算1+2+3+...+10的结果并打印
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	// 练习3：遍历一个整数切片，打印每个元素的值
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 练习4：计算100以内的所有奇数的和并打印
	sumOdd := 0
	for i := 1; i <= 100; i += 2 {
		sumOdd += i
	}
	fmt.Println("Sum of odd numbers:", sumOdd)
}