package main

import (
	"fmt"
)

func main() {
	// 示例1：计算两个整数的和，并输出结果
	num1 := 10
	num2 := 20
	sum := num1 + num2
	fmt.Println("Sum:", sum)

	// 示例2：判断一个数是否为偶数，并输出结果
	num := 15
	if num%2 == 0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// 示例3：遍历一个数组，并输出每个元素
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// 示例4：计算斐波那契数列，并输出前10个数
	fib := []int{0, 1}
	for i := 2; i < 10; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	fmt.Println(fib)
}