package main

import (
	"fmt"
)

func main() {
	// 练习1: 打印Hello, World!
	fmt.Println("Hello, World!")

	// 练习2: 计算两个数的和并打印结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 练习3: 判断一个数是否为偶数，并打印结果
	num := 15
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// 练习4: 求一个数组的平均值并打印结果
	numbers := [5]int{1, 2, 3, 4, 5}
	sum = 0
	for _, value := range numbers {
		sum += value
	}
	average := float64(sum) / float64(len(numbers))
	fmt.Println("Average:", average)
}